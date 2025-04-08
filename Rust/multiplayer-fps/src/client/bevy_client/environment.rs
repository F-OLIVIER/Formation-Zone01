// ███████ ███    ██ ██    ██ ██ ██████   ██████  ███    ██ ███    ██ ███████ ███    ███ ███████ ███    ██ ████████
// ██      ████   ██ ██    ██ ██ ██   ██ ██    ██ ████   ██ ████   ██ ██      ████  ████ ██      ████   ██    ██
// █████   ██ ██  ██ ██    ██ ██ ██████  ██    ██ ██ ██  ██ ██ ██  ██ █████   ██ ████ ██ █████   ██ ██  ██    ██
// ██      ██  ██ ██  ██  ██  ██ ██   ██ ██    ██ ██  ██ ██ ██  ██ ██ ██      ██  ██  ██ ██      ██  ██ ██    ██
// ███████ ██   ████   ████   ██ ██   ██  ██████  ██   ████ ██   ████ ███████ ██      ██ ███████ ██   ████    ██

use super::{audio::Audio, map::*, setup_ennemy::*, struct_manager::*};
use crate::interaction_server::config::*;
use bevy::prelude::*;

// ╭───────────────────────────────────────────────╮
// │                                               │
// │                     ENUM                      │
// │                                               │
// ╰───────────────────────────────────────────────╯
#[derive(Debug)]
pub enum BarType {
    Text,
    Green,
    Red,
}

// ╭───────────────────────────────────────────────╮
// │                                               │
// │                   STRUCTURE                   │
// │                                               │
// ╰───────────────────────────────────────────────╯
#[derive(Debug, Component)]
pub struct MunitionText;

#[derive(Component)]
pub struct TimerText;

#[derive(Debug, Component)]
pub struct ScoreTable;

#[derive(Debug, Component)]
pub struct ScoreText;

#[derive(Debug, Component)]
pub struct LifeBar {
    pub bar_type: BarType,
}

// ╭───────────────────────────────────────────────╮
// │                                               │
// │                    FONCTION                   │
// │                                               │
// ╰───────────────────────────────────────────────╯

// ██████  ███████ ███████ ███████  ██████  ██    ██ ██████   ██████ ███████ ███████
// ██   ██ ██      ██      ██      ██    ██ ██    ██ ██   ██ ██      ██      ██
// ██████  █████   ███████ ███████ ██    ██ ██    ██ ██████  ██      █████   ███████
// ██   ██ ██           ██      ██ ██    ██ ██    ██ ██   ██ ██      ██           ██
// ██   ██ ███████ ███████ ███████  ██████   ██████  ██   ██  ██████ ███████ ███████

pub fn generate_ressources(
    mut commands: Commands,
    mut meshes: ResMut<Assets<Mesh>>,
    mut materials: ResMut<Assets<StandardMaterial>>,
    asset_server: Res<AssetServer>,
    graphs: ResMut<Assets<AnimationGraph>>,
) {
    commands.insert_resource(MyResources {
        // wallpaper
        wallpaper: asset_server.load("wallpaper.png"),

        // map
        map: load_maps(&asset_server),

        // viseur
        viseur: (
            meshes.add(Cuboid::new(0.02, 0.02, 0.0)),
            materials.add(StandardMaterial {
                base_color_texture: Some(asset_server.load("textures/weapon/viseur.png")),
                unlit: true,
                alpha_mode: AlphaMode::Blend,
                ..default()
            }),
        ),

        // arme
        // weapon: load_modele_3d(&asset_server, "textures/weapon/blasterC.glb", 4), // blaster
        weapon: load_modele_3d(&asset_server, "textures/weapon/wand.glb", 1), // wand

        // projectile de l'arme
        projectile: (
            meshes.add(Cylinder::new(0.05, 0.1)),
            materials.add(StandardMaterial {
                base_color: Color::hsl(298.0, 0.44, 0.78),
                unlit: true, // désactive l'éclairage
                ..default()
            }),
        ),

        audio: Audio {
            ambiance: asset_server.load("audio/Heroic_Demise.ogg"),
            // shoot: asset_server.load("audio/laser.ogg"), // blaster
            shoot: asset_server.load("audio/spell3.ogg"), // wand
        },
        ennemy: load_ennemy(asset_server, graphs),
    });
}

pub fn load_modele_3d(
    asset_server: &Res<AssetServer>,
    path: &str,
    nb_primitive_nb_materiel: usize,
) -> Vec<(Handle<Mesh>, Handle<StandardMaterial>)> {
    let mut modele_handles: Vec<(Handle<Mesh>, Handle<StandardMaterial>)> = Vec::new();
    for i in 0..nb_primitive_nb_materiel {
        // Charger les mesh
        let mesh_path = format!("{}#Mesh0/Primitive{}", path, i);
        let mesh_handle: Handle<Mesh> = asset_server.load(&mesh_path);

        // Charger les matériau correspondant
        let material_path = format!("{}#Material{}", path, i);
        let material_handle: Handle<StandardMaterial> = asset_server.load(&material_path);

        modele_handles.push((mesh_handle, material_handle));
    }
    modele_handles
}

pub fn explain_text(mut commands: Commands) {
    commands
        .spawn(NodeBundle {
            style: Style {
                position_type: PositionType::Absolute,
                bottom: Val::Px(12.0),
                left: Val::Px(12.0),
                ..default()
            },
            ..default()
        })
        .insert(VisibilityBundle {
            visibility: Visibility::default(),
            inherited_visibility: InheritedVisibility::default(),
            view_visibility: ViewVisibility::default(),
        })
        .with_children(|parent| {
            parent.spawn(TextBundle::from_section(
                concat!(
                    "Press 'Escape' to quit the game.\n",
                    "Move the camera with your mouse.\n",
                    "Press 'arrow up' or 'Z' to advance.\n",
                    "Press 'arrow down' or 'S' to move back.\n",
                    "Press 'Left mouse' to shoot.\n",
                    "Press 'R' to reload",
                ),
                TextStyle {
                    font_size: 15.0,
                    ..default()
                },
            ));
        });
}

// ██      ██  ██████  ██   ██ ████████
// ██      ██ ██       ██   ██    ██
// ██      ██ ██   ███ ███████    ██
// ██      ██ ██    ██ ██   ██    ██
// ███████ ██  ██████  ██   ██    ██

pub fn spawn_lights(mut commands: Commands) {
    // Ajouter la ressource AmbientLight
    commands.insert_resource(AmbientLight {
        color: Color::WHITE, // Couleur de la lumière ambiante
        brightness: 10000.0, // Intensité de la lumière ambiante
    });

    // Ajoute un timer pour contrôler le clignotement de la lumière (facultatif)
    commands.insert_resource(LightTimer {
        duration: Timer::from_seconds(3.0, TimerMode::Once), // Durée du clignotement
        timer: Timer::from_seconds(0.1, TimerMode::Repeating), // Fréquence du clignotement
        active: false,                                       // Activation du clignotement
    });
}

pub fn light_clignotement(
    time: Res<Time>,
    mut timer: ResMut<LightTimer>,
    mut ambient_light: ResMut<AmbientLight>,
) {
    // Si le clignotement n'est pas actif
    if !timer.active {
        return;
    }

    // Compte à rebours pour la durée totale
    if timer.duration.tick(time.delta()).finished() {
        // Réinitialise la lumière ambiante à la couleur par défaut
        ambient_light.color = Color::WHITE;
        return; // Arrête le clignotement
    }

    // Vérifie si le timer de clignotement est terminé
    if timer.timer.tick(time.delta()).just_finished() {
        // Modifie la couleur de l'AmbientLight
        if ambient_light.color == Color::hsl(0.0, 1.0, 0.5) {
            ambient_light.color = Color::WHITE;
        } else {
            ambient_light.color = Color::hsl(0.0, 1.0, 0.5); // Rouge
        }
    }
}

// ███    ███ ██    ██ ███    ██ ██ ████████ ██  ██████  ███    ██
// ████  ████ ██    ██ ████   ██ ██    ██    ██ ██    ██ ████   ██
// ██ ████ ██ ██    ██ ██ ██  ██ ██    ██    ██ ██    ██ ██ ██  ██
// ██  ██  ██ ██    ██ ██  ██ ██ ██    ██    ██ ██    ██ ██  ██ ██
// ██      ██  ██████  ██   ████ ██    ██    ██  ██████  ██   ████

pub fn display_munition(mut commands: Commands, asset_server: Res<AssetServer>) {
    commands
        .spawn(NodeBundle {
            style: Style {
                position_type: PositionType::Absolute,
                bottom: Val::Px(12.0),
                right: Val::Px(12.0),
                ..default()
            },
            ..default()
        })
        .with_children(|parent| {
            parent.spawn((
                TextBundle::from_section(
                    format!("Munition\n{} / {}", MAX_MUNITION, MAX_MUNITION),
                    TextStyle {
                        font: asset_server.load(FONTS_MENU),
                        font_size: 25.0,
                        color: Color::WHITE,
                    },
                ),
                MunitionText,
            ));
        });
}

pub fn update_munition(
    mut query: Query<&mut Text, With<MunitionText>>,
    frame_counter: Local<u32>,
    players: ResMut<Players>,
) {
    // Mettre à jour le texte tous les XX frames
    if *frame_counter % 120 == 0 {
        for mut text in &mut query {
            // Mettre à jour la valeur de la section
            text.sections[0].value = format!(
                "Munition\n{} / {}",
                players.player_private.munition, MAX_MUNITION
            );
        }
    }
}

// ███████  ██████  ███    ██ ███████     ██ ███    ██ ███████  ██████  ██████  ███    ███  █████  ████████ ██  ██████  ███    ██
//    ███  ██    ██ ████   ██ ██          ██ ████   ██ ██      ██    ██ ██   ██ ████  ████ ██   ██    ██    ██ ██    ██ ████   ██
//   ███   ██    ██ ██ ██  ██ █████       ██ ██ ██  ██ █████   ██    ██ ██████  ██ ████ ██ ███████    ██    ██ ██    ██ ██ ██  ██
//  ███    ██    ██ ██  ██ ██ ██          ██ ██  ██ ██ ██      ██    ██ ██   ██ ██  ██  ██ ██   ██    ██    ██ ██    ██ ██  ██ ██
// ███████  ██████  ██   ████ ███████     ██ ██   ████ ██       ██████  ██   ██ ██      ██ ██   ██    ██    ██  ██████  ██   ████

#[derive(Component)]
pub struct LogTextTag;

#[derive(Debug, Resource)]
pub struct LogsResource {
    pub content: Vec<String>, // Stocke les lignes de log
    pub max_lines: usize,     // Nombre maximum de lignes visibles
}

impl Default for LogsResource {
    fn default() -> Self {
        Self {
            content: Vec::new(),
            max_lines: 10, // 10 lignes maximum de log
        }
    }
}

// Texte temporaire d'information
pub fn log_text(mut commands: Commands, asset_server: Res<AssetServer>) {
    commands
        .spawn(NodeBundle {
            style: Style {
                position_type: PositionType::Absolute,
                top: Val::Px(50.0),
                left: Val::Px(10.0),
                width: Val::Px(300.0),
                height: Val::Px(200.0),
                padding: UiRect {
                    right: Val::Px(2.0),
                    ..default()
                },
                align_items: AlignItems::FlexEnd,
                ..default()
            },
            ..default()
        })
        .with_children(|parent| {
            parent.spawn((
                TextBundle::from_section(
                    "texte initial",
                    TextStyle {
                        font: asset_server.load(FONTS_MENU),
                        font_size: 15.0,
                        color: Color::WHITE,
                    },
                ),
                LogTextTag,
            ));
        });
}

pub fn update_logs(mut query: Query<&mut Text, With<LogTextTag>>, logs: ResMut<LogsResource>) {
    if let Ok(mut text) = query.get_single_mut() {
        let reversed_logs: Vec<String> = logs.content.iter().rev().cloned().collect();
        text.sections[0].value = reversed_logs.join("\n");
    }
}

pub fn add_log(logs: &mut ResMut<LogsResource>, new_log: String) {
    println!("new_log: {}", new_log);
    logs.content.push(new_log);

    // Supprimez les anciens logs si le nombre dépasse `max_lines`
    if logs.content.len() > logs.max_lines {
        logs.content.remove(0);
    }
}

// ███████  ██████  ██████  ██████  ███████ ██████   ██████   █████  ██████  ██████
// ██      ██      ██    ██ ██   ██ ██      ██   ██ ██    ██ ██   ██ ██   ██ ██   ██
// ███████ ██      ██    ██ ██████  █████   ██████  ██    ██ ███████ ██████  ██   ██
//      ██ ██      ██    ██ ██   ██ ██      ██   ██ ██    ██ ██   ██ ██   ██ ██   ██
// ███████  ██████  ██████  ██   ██ ███████ ██████   ██████  ██   ██ ██   ██ ██████

pub fn display_scoreboard(
    mut commands: Commands,
    asset_server: Res<AssetServer>,
    players: Res<Players>,
) {
    let width_percent = 80.0;
    let left_percent = (100.0 - width_percent) / 2.0; // Centrer horizontalement
    let height_percent = 80.0;
    let top_percent = (100.0 - height_percent) / 2.0; // Centrer verticalement

    commands
        .spawn(NodeBundle {
            style: Style {
                width: Val::Percent(width_percent),
                left: Val::Percent(left_percent),
                height: Val::Percent(height_percent),
                top: Val::Percent(top_percent),
                position_type: PositionType::Absolute,
                justify_content: JustifyContent::Center,
                align_items: AlignItems::Center,
                border: UiRect::all(Val::Px(2.)),
                ..default()
            },
            visibility: Visibility::Hidden,
            background_color: Color::BLACK.into(),
            border_color: Color::WHITE.into(),
            border_radius: BorderRadius::new(
                Val::Px(30.),
                Val::Px(30.),
                Val::Px(30.),
                Val::Px(30.),
            ),
            ..default()
        })
        .insert(ScoreTable)
        .with_children(|parent| {
            parent.spawn((
                TextBundle::from_section(
                    get_scores_string(&players), // Met à jour pour afficher les scores
                    TextStyle {
                        font: asset_server.load(FONTS_MENU),
                        font_size: 40.0,
                        color: Color::WHITE,
                    },
                ),
                ScoreText,
            ));
        });
}

pub fn update_scoreboard(
    time: Res<Time>,
    mut send_rate_timer: ResMut<SendRateTimer>,
    keyboard_input: Res<ButtonInput<KeyCode>>,
    players: Res<Players>,
    mut score_table: Query<&mut Visibility, With<ScoreTable>>,
    mut score_text: Query<&mut Text, With<ScoreText>>,
) {
    let is_tab_pressed = keyboard_input.pressed(KeyCode::Tab);
    // Met à jour le texte des scores si le tableau est visible
    if is_tab_pressed
        && send_rate_timer
            .scoreboard_timer
            .tick(time.delta())
            .finished()
    {
        for mut text in score_text.iter_mut() {
            text.sections[0].value = get_scores_string(&players);
        }
    }

    // Met à jour la visibilité du tableau des scores
    for mut visibility in score_table.iter_mut() {
        *visibility = if is_tab_pressed {
            Visibility::Visible
        } else {
            Visibility::Hidden
        };
    }
}

pub fn get_scores_string(players: &Res<Players>) -> String {
    let mut score_display = "Score\n".to_string();
    // Convertir le HashMap en Vec et trier les joueurs par score décroissant
    let mut sorted_players: Vec<_> = players.all_players.iter().collect();
    sorted_players.sort_by(|a, b| {
        b.1.info_player_in_game
            .score
            .cmp(&a.1.info_player_in_game.score)
    });

    for (_, player) in sorted_players {
        println!(
            "sorted players : {}  name: {}\n--------------------------",
            player.info_player_in_game.score, player.info_player.username
        );
        let score = format!(
            "{}: {}\n",
            player.info_player.username, player.info_player_in_game.score
        );
        score_display.push_str(&score);
    }
    score_display
}

// ██      ██ ███████ ███████ ██████   █████  ██████
// ██      ██ ██      ██      ██   ██ ██   ██ ██   ██
// ██      ██ █████   █████   ██████  ███████ ██████
// ██      ██ ██      ██      ██   ██ ██   ██ ██   ██
// ███████ ██ ██      ███████ ██████  ██   ██ ██   ██
// pub fn display_lifebar(mut commands: Commands, asset_server: Res<AssetServer>) {
//     let width_percent = 35.0;
//     let left_percent = (90.0 - width_percent) / 2.0; // Centrer horizontalement
//     let height_percent = 3.0;

//     commands
//         .spawn(NodeBundle {
//             style: Style {
//                 width: Val::Percent(width_percent),
//                 left: Val::Percent(left_percent),
//                 height: Val::Percent(height_percent),
//                 bottom: Val::Px(24.0),
//                 position_type: PositionType::Absolute,
//                 justify_content: JustifyContent::Center,
//                 align_items: AlignItems::Center,
//                 border: UiRect::all(Val::Px(2.)),
//                 ..default()
//             },
//             background_color: Color::hsl(97.0, 1.0, 0.5).into(),
//             border_color: Color::BLACK.into(),
//             border_radius: BorderRadius::new(
//                 Val::Px(15.),
//                 Val::Px(15.),
//                 Val::Px(15.),
//                 Val::Px(15.),
//             ),
//             ..default()
//         })
//         .with_children(|parent| {
//             // Barre de vie restante : couleur verte (initialement pleine)
//             parent
//                 .spawn(NodeBundle {
//                     style: Style {
//                         width: Val::Percent(100.0),
//                         height: Val::Percent(98.0),
//                         ..default()
//                     },
//                     background_color: Color::hsl(97.0, 1.0, 0.5).into(),
//                     border_radius: BorderRadius::new(
//                         Val::Px(15.),
//                         Val::Px(15.),
//                         Val::Px(15.),
//                         Val::Px(15.),
//                     ),
//                     ..default()
//                 })
//                 .insert(LifeBar {
//                     bar_type: BarType::Green,
//                 });

//             // Barre de vie perdue, couleur rouge (initialement invisible)
//             parent
//                 .spawn(NodeBundle {
//                     style: Style {
//                         width: Val::Percent(0.0),
//                         height: Val::Percent(100.0),
//                         ..default()
//                     },
//                     background_color: Color::hsl(0.0, 1.0, 0.5).into(),
//                     border_radius: BorderRadius::new(
//                         Val::Px(0.),
//                         Val::Px(15.),
//                         Val::Px(15.),
//                         Val::Px(0.),
//                     ),
//                     ..default()
//                 })
//                 .insert(LifeBar {
//                     bar_type: BarType::Red,
//                 });

//             // Texte affichant la vie
//             parent.spawn((
//                 TextBundle {
//                     style: Style {
//                         width: Val::Percent(50.0),
//                         right: Val::Percent(22.0),
//                         justify_content: JustifyContent::Center,
//                         align_items: AlignItems::Center,
//                         ..default()
//                     },
//                     text: Text::from_section(
//                         format!("{:.0} / {:.0}", MAX_LIFE, MAX_LIFE),
//                         TextStyle {
//                             font: asset_server.load(FONTS_MENU),
//                             font_size: 18.0,
//                             color: Color::BLACK,
//                         },
//                     ),
//                     ..default()
//                 },
//                 LifeBar {
//                     bar_type: BarType::Text,
//                 },
//             ));
//         });
// }

// pub fn update_lifebar(
//     mut query: Query<(&mut Style, Option<&mut Text>, &LifeBar)>,
//     player_health: u32,
// ) {
//     // Calculer le pourcentage de vie du player
//     let health_percent = (player_health as f32 * 100.0) / MAX_LIFE;

//     for (mut style, text, lifebar) in query.iter_mut() {
//         match lifebar.bar_type {
//             BarType::Green => {
//                 style.width = Val::Percent(health_percent);
//             }

//             BarType::Red => {
//                 style.width = Val::Percent(100.0 - health_percent);
//             }

//             BarType::Text => {
//                 if let Some(mut text) = text {
//                     text.sections[0].value = format!("{:.0} / {}", player_health, MAX_LIFE);
//                 }
//             }
//         }
//     }
// }

// ████████ ██ ███    ███ ███████ ██████       ██████   █████  ███    ███ ███████
//    ██    ██ ████  ████ ██      ██   ██     ██       ██   ██ ████  ████ ██
//    ██    ██ ██ ████ ██ █████   ██████      ██   ███ ███████ ██ ████ ██ █████
//    ██    ██ ██  ██  ██ ██      ██   ██     ██    ██ ██   ██ ██  ██  ██ ██
//    ██    ██ ██      ██ ███████ ██   ██      ██████  ██   ██ ██      ██ ███████

// pub fn display_timer(mut commands: Commands, asset_server: Res<AssetServer>) {
//     // Conversion du timer en minutes et secondes
//     let minutes = (TIMER_GAME as i32) / 60;
//     let seconds = (TIMER_GAME as i32) % 60;

//     commands
//         .spawn(NodeBundle {
//             style: Style {
//                 position_type: PositionType::Absolute,
//                 top: Val::Px(12.0),
//                 left: Val::Percent(50.0),
//                 justify_content: JustifyContent::Center,
//                 align_items: AlignItems::Center,
//                 ..default()
//             },
//             ..default()
//         })
//         .with_children(|parent| {
//             parent.spawn((
//                 TextBundle::from_section(
//                     format!("{:02}:{:02}", minutes, seconds),
//                     TextStyle {
//                         font: asset_server.load(FONTS_MENU),
//                         font_size: 40.0,
//                         color: Color::WHITE,
//                     },
//                 ),
//                 TimerText,
//             ));
//         });
// }

// pub fn update_timer(
//     mut query: Query<&mut Text, With<TimerText>>,
//     time: Res<Time>,
//     mut players: ResMut<Players>,
// ) {
//     for mut text in &mut query {
//         // Mise à jour du timer
//         let timer = &mut players.player_public.common_information.timer_map;
//         *timer -= time.delta_seconds();

//         // Si le timer atteint zéro, on peut arrêter ou définir d'autres actions
//         if *timer < 0.0 {
//             *timer = 0.0; // Assure que le timer ne passe pas sous zéro
//         }

//         // Conversion du timer en minutes et secondes
//         let minutes = (*timer as i32) / 60;
//         let seconds = (*timer as i32) % 60;

//         // Met à jour le texte affiché
//         text.sections[0].value = format!("{:02}:{:02}", minutes, seconds);
//     }
// }
