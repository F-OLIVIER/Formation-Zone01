// ██████  ██       █████  ██    ██ ███████ ██████
// ██   ██ ██      ██   ██  ██  ██  ██      ██   ██
// ██████  ██      ███████   ████   █████   ██████
// ██      ██      ██   ██    ██    ██      ██   ██
// ██      ███████ ██   ██    ██    ███████ ██   ██

use super::audio::*;
use super::environment::*;
use super::moteur::*;
use crate::bevy_client::{struct_manager::*, struct_menu::*};
use crate::interaction_server::config::*;
use crate::interaction_server::udpsocket::*;
use bevy::time::Timer;
use bevy::utils::hashbrown::HashMap;
use bevy_rapier3d::prelude::*;
use renet::transport::NetcodeClientTransport;
use renet::RenetClient;

// ╭───────────────────────────────────────────────╮
// │                                               │
// │                   STRUCTURE                   │
// │                                               │
// ╰───────────────────────────────────────────────╯
#[derive(Debug, Component)]
pub struct Player;

#[derive(Debug, Component)]
pub struct Projectile {
    pub id: u128,
}

#[derive(Debug, Resource)]
pub struct ProjectileNumber {
    pub id: u64,
}

// Composant pour la durée de vie des entités temporaires
#[derive(Component)]
pub struct Lifetime {
    pub timer: Timer,
}

// ╭───────────────────────────────────────────────╮
// │                                               │
// │                    FONCTION                   │
// │                                               │
// ╰───────────────────────────────────────────────╯

// ██████  ██████   ██████       ██ ███████  ██████ ████████ ██ ██      ███████     ██████  ██    ██          ██  ██████  ██    ██ ███████ ██    ██ ██████
// ██   ██ ██   ██ ██    ██      ██ ██      ██         ██    ██ ██      ██          ██   ██ ██    ██          ██ ██    ██ ██    ██ ██      ██    ██ ██   ██
// ██████  ██████  ██    ██      ██ █████   ██         ██    ██ ██      █████       ██   ██ ██    ██          ██ ██    ██ ██    ██ █████   ██    ██ ██████
// ██      ██   ██ ██    ██ ██   ██ ██      ██         ██    ██ ██      ██          ██   ██ ██    ██     ██   ██ ██    ██ ██    ██ ██      ██    ██ ██   ██
// ██      ██   ██  ██████   █████  ███████  ██████    ██    ██ ███████ ███████     ██████   ██████       █████   ██████   ██████  ███████  ██████  ██   ██

pub fn shoot(
    mut commands: Commands,
    mut client: ResMut<RenetClient>,
    mut transport: ResMut<NetcodeClientTransport>,
    mut exit: EventWriter<AppExit>,
    keyboard_input: Res<ButtonInput<KeyCode>>,
    mouse_input: Res<ButtonInput<MouseButton>>,
    resources: Res<MyResources>,
    mut projectile_number: ResMut<ProjectileNumber>,
    time: Res<Time>,
    volume: Res<AudioVolume>,
    mut players: ResMut<Players>,
    player_query: Query<&Transform, With<Player>>,
    mut last_shot_time: Local<f32>,
    _score_table: Query<&mut Text, With<ScoreTable>>,
) {
    // Quitter la partie
    if keyboard_input.pressed(KeyCode::Escape) {
        exit_system(&mut client, &mut transport, &mut players, &mut exit);
    }

    // Tir du joueur
    *last_shot_time += time.delta_seconds();
    if mouse_input.pressed(MouseButton::Left)
        && *last_shot_time >= SHOOTING_COOLDOWN
        && players.player_private.munition > 0
    {
        if let Ok(arm_transform) = player_query.get_single() {
            let arm_rotation = arm_transform.rotation;
            let arm_position = Vec3::new(
                arm_transform.translation.x,
                arm_transform.translation.y,
                arm_transform.translation.z,
            );
            let forward = arm_rotation * -Vec3::Z;
            let right = arm_rotation * Vec3::X;
            let up = arm_rotation * Vec3::Y;
            let spawn_position = arm_position + forward + up * -0.18 + right * 0.19;
            let final_rotation = arm_rotation * Quat::from_rotation_x(1.7);

            // Récupérer puis incrémente l'ID et l'utilise pour le projectile
            let current_id = format!(
                "{}{}",
                players.player_public.info_player.id, projectile_number.id
            )
            .parse::<u128>()
            .expect("Failed to parse concatenated ID");
            projectile_number.id += 1;

            // Spawn du projectile du joueur
            commands.spawn((
                Lifetime {
                    timer: Timer::from_seconds(10.0, TimerMode::Once),
                },
                PbrBundle {
                    mesh: resources.projectile.0.clone(),
                    material: resources.projectile.1.clone(),
                    transform: Transform::from_translation(spawn_position)
                        .with_rotation(final_rotation),
                    ..default()
                },
                Projectile { id: current_id }, // Assignation de l'ID du projectile
                Collider::ball(0.1),
                Ccd { enabled: true },
            ));

            *last_shot_time = 0.0;
            players.player_private.munition -= 1;

            // Ajout du projectile à la HashMap
            players
                .player_public
                .info_player_in_game
                .projectile
                .insert(current_id, (spawn_position.into(), final_rotation));

            // Son de tir
            play_once_audio(commands, resources.audio.shoot.clone(), volume);
        }
    }

    if keyboard_input.just_pressed(KeyCode::KeyR) {
        players.player_private.munition = MAX_MUNITION;
    }

    // Pour récupérer des position de spawn
    // if keyboard_input.just_pressed(KeyCode::KeyC) {
    //     players.player_public.info_player_in_game.score += 1;
    //     println!(
    //         "Position de depart : {:?}",
    //         players.player_public.info_player_in_game.position,
    //     );
    // }
}

// Déplacement des projectile du joueur
pub fn move_projectiles(
    time: Res<Time>,
    mut commands: Commands,
    mut client: ResMut<RenetClient>,
    mut query: ParamSet<(
        Query<(Entity, &mut Transform, &Projectile, &Collider)>, // Query pour les projectiles
        Query<(Entity, &mut EnemyTag)>,                          // Query pour les ennemis
    )>,
    physics_context: Res<RapierContext>,
    mut players: ResMut<Players>,
) {
    let max_distance = 50.0;
    let mut collisions = Vec::new();
    let mut list_projectile: HashMap<u128, ((f32, f32, f32), Quat)> = HashMap::new(); // : Vec<((f32, f32, f32), (f32, f32))> = Vec::new();

    // Phase 1 : Déplacement des projectiles et détection des collisions
    for (entity, mut transform, projectile, collider) in query.p0().iter_mut() {
        let forward = transform.rotation * -Vec3::Y;
        let downward = -Vec3::Y * 0.15;
        let rightward = transform.rotation * Vec3::X * 0.08;
        let combined_direction = (forward + downward + rightward).normalize();
        let velocity = 10.0;

        transform.translation += combined_direction
            * (velocity + players.player_private.velocity)
            * time.delta_seconds();

        list_projectile.insert(
            projectile.id,
            (
                (
                    transform.translation.x,
                    transform.translation.y,
                    transform.translation.z,
                ),
                Quat::from_euler(
                    EulerRot::YXZ,
                    transform.rotation.y,
                    transform.rotation.x,
                    0.0,
                ),
            ),
        );

        // Détection des collisions
        if let Some((hit_entity, toi)) = physics_context.cast_shape(
            transform.translation,
            transform.rotation,
            combined_direction
                * (velocity + players.player_private.velocity)
                * time.delta_seconds(),
            &collider,
            ShapeCastOptions::with_max_time_of_impact(0.5),
            QueryFilter::default().exclude_collider(entity),
        ) {
            if toi.time_of_impact <= max_distance {
                // Stockez la collision détectée
                collisions.push((entity, hit_entity));
            }
        }
    }

    // mise a jour de la liste des projectiles
    players.player_public.info_player_in_game.projectile = list_projectile;

    // Phase 2 : Gérer les collisions
    for (projectile_entity, hit_entity) in collisions {
        if let Ok((_, enemy_tag)) = query.p1().get_mut(hit_entity) {
            send_event(
                &mut client,
                EventMessage {
                    types: "player_dead".to_string(),
                    id: players.player_public.info_player.id,
                    username: players.player_public.info_player.username.clone(),
                    information: enemy_tag.id.to_string(),
                },
            );
            players.player_public.info_player_in_game.score += 1;
        }

        // Supprimer le projectile
        commands.entity(projectile_entity).despawn();
    }
}

// Système pour supprimer les entités dont le timer est expiré (projectile du joueur)
pub fn remove_expired_projectiles(
    mut commands: Commands,
    time: Res<Time>,
    mut players: ResMut<Players>,
    mut query: Query<(Entity, &mut Lifetime, &Projectile)>,
) {
    let mut expired_entities = Vec::new();

    // Mise a jour du lifetime + deletion de la Hashmap si besoin
    for (entity, mut lifetime, projectile) in query.iter_mut() {
        lifetime.timer.tick(time.delta());

        if lifetime.timer.finished() {
            // sauvegarde pour suppression ulterieur afin d'éviter l'erreur B0003
            expired_entities.push(entity);

            players
                .player_public
                .info_player_in_game
                .projectile
                .remove(&projectile.id);
        }
    }

    // Suppression des projectiles avec vérification si deja remove ou non (projectile despawn via les colissions)
    for entity in expired_entities {
        // Vérification si l'entité est toujours présente dans le monde avant de la supprimer
        if query.get(entity).is_ok() {
            commands.entity(entity).despawn();
        }
    }
}

// ██████  ██████   ██████       ██ ███████  ██████ ████████ ██ ██      ███████      █████  ██    ██ ████████ ██████  ███████ ███████          ██  ██████  ██    ██ ███████ ██    ██ ██████
// ██   ██ ██   ██ ██    ██      ██ ██      ██         ██    ██ ██      ██          ██   ██ ██    ██    ██    ██   ██ ██      ██               ██ ██    ██ ██    ██ ██      ██    ██ ██   ██
// ██████  ██████  ██    ██      ██ █████   ██         ██    ██ ██      █████       ███████ ██    ██    ██    ██████  █████   ███████          ██ ██    ██ ██    ██ █████   ██    ██ ██████
// ██      ██   ██ ██    ██ ██   ██ ██      ██         ██    ██ ██      ██          ██   ██ ██    ██    ██    ██   ██ ██           ██     ██   ██ ██    ██ ██    ██ ██      ██    ██ ██   ██
// ██      ██   ██  ██████   █████  ███████  ██████    ██    ██ ███████ ███████     ██   ██  ██████     ██    ██   ██ ███████ ███████      █████   ██████   ██████  ███████  ██████  ██   ██

// Mise à jour ou création des projectiles pour un joueur donné.
pub fn update_or_create_projectiles(
    player_ennemy: &PlayerPublic,
    query_projectiles: &mut Query<(Entity, &ProjectileInfo, &mut Transform)>,
    commands: &mut Commands,
    resource: &Res<MyResources>,
) {
    for (id_projectile, (position, rotation)) in &player_ennemy.info_player_in_game.projectile {
        // Cherche le projectile concerné
        if let Some((_, _, mut transform)) = query_projectiles
            .iter_mut()
            .find(|(_, projectile_info, _)| projectile_info.id_projectile == *id_projectile)
        {
            // Si le projectile est trouvé, mettre à jour sa position et sa rotation
            transform.translation = Vec3::new(position.0, position.1, position.2);
            transform.rotation = *rotation;
        } else {
            // Si le projectile n'existe pas sur la scéne, on le crée
            spawn_projectile(
                commands,
                resource,
                ProjectileInfo {
                    id_projectile: *id_projectile,
                    position: *position,
                    rotation: (rotation.x, rotation.y),
                },
            );
        }
    }
}

// Spawn des projectile des autres joueurs
fn spawn_projectile(
    commands: &mut Commands,
    resources: &Res<MyResources>,
    projectile_info: ProjectileInfo,
) {
    commands
        .spawn((
            Lifetime {
                timer: Timer::from_seconds(10.0, TimerMode::Once),
            },
            PbrBundle {
                mesh: resources.projectile.0.clone(),
                material: resources.projectile.1.clone(),
                transform: Transform::from_translation(projectile_info.position.into())
                    .with_rotation(Quat::from_euler(
                        EulerRot::YXZ,
                        projectile_info.rotation.1,
                        projectile_info.rotation.0,
                        0.0,
                    )),
                ..default()
            },
            Collider::ball(0.1),
            Ccd { enabled: true },
        ))
        .insert(projectile_info);
}

// Système pour supprimer les projectile des autres joueurs
pub fn remove_orphan_projectiles(
    player: &PlayerPublic,
    query_projectiles: &mut Query<(Entity, &ProjectileInfo, &mut Transform)>,
    commands: &mut Commands,
) {
    // Collecte des ID de projectiles présents dans la HashMap du joueur
    let existing_projectiles: Vec<u128> = player
        .info_player_in_game
        .projectile
        .keys()
        .cloned()
        .collect();

    // Parcours des projectiles dans la query
    for (entity, projectile_info, _) in query_projectiles.iter_mut() {
        // Si l'ID du projectile n'est pas dans la HashMap du joueur, Suppression du projectile
        if !existing_projectiles.contains(&projectile_info.id_projectile) {
            commands.entity(entity).despawn();
        }
    }
}

pub fn respawn(mut query: Query<(&mut Transform, &mut Velocity)>) {
    for (mut transform, mut velocity) in &mut query {
        if transform.translation.y > -50.0 {
            continue;
        }
        velocity.linvel = Vec3::ZERO;
        transform.translation = respawn_point().into();
    }
}
