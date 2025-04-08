// ███████ ███    ██ ███    ██ ███████ ███    ███ ██    ██ ███████
// ██      ████   ██ ████   ██ ██      ████  ████  ██  ██  ██
// █████   ██ ██  ██ ██ ██  ██ █████   ██ ████ ██   ████   ███████
// ██      ██  ██ ██ ██  ██ ██ ██      ██  ██  ██    ██         ██
// ███████ ██   ████ ██   ████ ███████ ██      ██    ██    ███████

use super::struct_manager::*;
use crate::interaction_server::config::{Commands, PlayerPublic};
use bevy::animation::AnimationPlayer;
use bevy::prelude::*;
use bevy_rapier3d::prelude::{
    ActiveEvents, AdditionalMassProperties, Ccd, Collider, GravityScale, RigidBody, Sleeping,
    Velocity,
};
use std::time::Duration;

// ╭───────────────────────────────────────────────╮
// │                                               │
// │                   STRUCTURE                   │
// │                                               │
// ╰───────────────────────────────────────────────╯

// ╭───────────────────────────────────────────────╮
// │                                               │
// │                    FONCTION                   │
// │                                               │
// ╰───────────────────────────────────────────────╯

// pub fn setup_ennemy(mut commands: Commands, ressource: Res<MyResources>, id: String) {
pub fn setup_ennemy(id: u64, mut commands: Commands, ressource: &Res<MyResources>) {
    commands
        .spawn(SceneBundle {
            scene: ressource.ennemy.0.clone(),
            ..default()
        })
        .insert(EnemyTag { id })
        .insert(AnimationPlayer::default())
        .insert(Collider::cylinder(3.0, 0.5))
        .insert(RigidBody::KinematicPositionBased)
        .insert(Ccd { enabled: true })
        .insert(GravityScale(0.0))
        .insert(Sleeping::disabled())
        .insert(Velocity::zero())
        .insert(ActiveEvents::COLLISION_EVENTS)
        .insert(AdditionalMassProperties::Mass(1.0));
}

pub fn load_ennemy(
    asset_server: Res<AssetServer>,
    mut graphs: ResMut<Assets<AnimationGraph>>,
) -> (Handle<Scene>, Animations) {
    let path = "textures/enemy/soldier.glb";
    let mut graph = AnimationGraph::new();
    let animation_clip =
        asset_server.load(GltfAssetLabel::Animation(0).from_asset(format!("{}#Animation0", path)));
    let animations = graph.add_clips([animation_clip], 1.0, graph.root).collect();

    // Insert a resource with the current scene information
    let graph = graphs.add(graph);
    let handle_scene: Handle<Scene> = asset_server.load(format!("{}#Scene0", path));

    let animation = Animations { animations, graph };

    (handle_scene, animation)
}

pub fn play_animation(
    mut commands: Commands,
    mut ennemys: Query<(Entity, &mut AnimationPlayer), Added<AnimationPlayer>>,
    ressource: Res<MyResources>,
) {
    for (entity, mut ennemy) in &mut ennemys {
        let mut transitions = AnimationTransitions::new();
        transitions
            .play(
                &mut ennemy,
                ressource.ennemy.1.animations[0].clone(),
                Duration::ZERO,
            )
            .repeat();

        commands
            .entity(entity)
            .insert(ressource.ennemy.1.graph.clone())
            .insert(transitions);
    }
}

pub fn keyboard_animation_control(
    keyboard_input: Res<ButtonInput<KeyCode>>,
    mut animation_ennemys: Query<&mut AnimationPlayer>,
) {
    for mut ennemy in &mut animation_ennemys {
        let Some((&playing_animation_index, _)) = ennemy.playing_animations().next() else {
            continue;
        };
        let playing_animation = ennemy.animation_mut(playing_animation_index).unwrap();

        if keyboard_input.pressed(KeyCode::KeyU) {
            playing_animation.resume();
        } else {
            playing_animation.pause();
        }
    }
}

// Mise à jour ou création d'un ennemi.
pub fn update_or_create_enemy(
    player: &PlayerPublic,
    query_enemy: &mut Query<(Entity, &EnemyTag, &mut Transform)>,
    commands: &mut Commands,
    ressource: &Res<MyResources>,
) {
    if let Some((_entity, _, mut transform)) = query_enemy
        .iter_mut()
        .find(|(_, tag, _)| tag.id == player.info_player.id)
    {
        // Mise à jour de la position et de la rotation de l'ennemi existant
        transform.translation = Vec3::new(
            player.info_player_in_game.position.0,
            player.info_player_in_game.position.1 + 1.0,
            player.info_player_in_game.position.2,
        );
        transform.rotation = Quat::from_euler(
            EulerRot::YXZ,
            player.info_player_in_game.rotation.1,
            player.info_player_in_game.rotation.0,
            0.0,
        );
    } else {
        // Création d'un nouvel ennemi si inexistant
        setup_ennemy(player.info_player.id, commands.reborrow(), ressource);
    }
}
