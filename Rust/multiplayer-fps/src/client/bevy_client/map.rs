// ███    ███  █████  ██████
// ████  ████ ██   ██ ██   ██
// ██ ████ ██ ███████ ██████
// ██  ██  ██ ██   ██ ██
// ██      ██ ██   ██ ██

use super::struct_manager::*;
use bevy::prelude::*;
use std::collections::HashMap;

// ╭───────────────────────────────────────────────╮
// │                                               │
// │                   STRUCTURE                   │
// │                                               │
// ╰───────────────────────────────────────────────╯

#[derive(Component)]
pub struct MinimapMarker;

// ╭───────────────────────────────────────────────╮
// │                                               │
// │                    FONCTION                   │
// │                                               │
// ╰───────────────────────────────────────────────╯

pub fn display_map(mut commands: Commands, resources: Res<MyResources>) {
    if let Some((map, minimap)) = resources.map.get("map_z01") {
        // Spawn de la map
        commands.insert_resource(MainScene {
            handle: map.clone(),
            is_loaded: false,
        });

        // Minimap
        let minimap_entity = commands
            .spawn(NodeBundle {
                style: Style {
                    position_type: PositionType::Absolute,
                    top: Val::Px(10.0),
                    right: Val::Px(10.0),
                    width: Val::Percent(20.0),
                    height: Val::Percent(30.0),
                    ..default()
                },
                ..default()
            })
            .with_children(|parent| {
                parent.spawn(ImageBundle {
                    image: UiImage {
                        texture: minimap.clone(),
                        ..default()
                    },
                    style: Style {
                        width: Val::Percent(100.0),
                        height: Val::Percent(100.0),
                        ..default()
                    },
                    ..default()
                });
            })
            .id();

        // Ajouter un marqueur pour la minimap
        commands.entity(minimap_entity).with_children(|parent| {
            parent.spawn((
                NodeBundle {
                    style: Style {
                        width: Val::Px(10.0),
                        height: Val::Px(10.0),
                        position_type: PositionType::Absolute,
                        ..default()
                    },
                    border_radius: BorderRadius {
                        top_left: Val::Percent(50.0),
                        top_right: Val::Percent(50.0),
                        bottom_left: Val::Percent(50.0),
                        bottom_right: Val::Percent(50.0),
                    },
                    #[allow(deprecated)]
                    background_color: BackgroundColor(Color::rgb(0.0, 255.0, 0.0)),
                    ..default()
                },
                MinimapMarker,
            ));
        });
    } else {
        info!("Problème avec la HashMap dans resources (carte non trouvée)");
    }
}

pub fn load_maps(
    asset_server: &Res<AssetServer>,
) -> HashMap<String, (Handle<Gltf>, Handle<Image>)> {
    let name_map = ["map_z01"];

    let mut list_map: HashMap<String, (Handle<Gltf>, Handle<Image>)> = HashMap::new();

    // Ajouter les éléments des maps à la HashMap
    for map in name_map {
        list_map.insert(
            map.to_string(),
            (
                asset_server.load(format!("map/{}.glb", map)), // Scene 3D de la map
                asset_server.load(format!("map/{}.png", map)), // Minimap 2D de la map
            ),
        );
    }

    list_map
}

use crate::bevy_client::player_manager::Player;

pub fn update_minimap_marker(
    player_query: Query<(&Transform, &Player)>,
    mut marker_query: Query<(&mut Style, &MinimapMarker)>,
    mut windows: Query<&mut Window>,
    time: Res<Time>,                            // Pour gérer le temps écoulé
    mut send_rate_timer: ResMut<SendRateTimer>, // Le timer pour limiter l'envoi
) {
    if send_rate_timer.minimap_timer.tick(time.delta()).finished() {
        // println!("enter update_minimap_marker");
        let window = windows.get_single_mut().unwrap();
        let window_width = window.width();
        let window_height = window.height();

        if let Ok((player_transform, _)) = player_query.get_single() {
            let player_position = player_transform.translation;

            // Convertir la position du joueur en coordonnées de la minimap
            let minimap_position =
                convert_to_minimap_position(player_position, window_width, window_height);
            // println!("minimap_position : {}", minimap_position);

            if let Ok((mut marker_style, _)) = marker_query.get_single_mut() {
                marker_style.left = Val::Px(minimap_position.x);
                marker_style.top = Val::Px(minimap_position.y);
            }
        }
    }
}

fn convert_to_minimap_position(
    player_position: Vec3,
    window_width: f32,
    window_height: f32,
) -> Vec2 {
    // Taille de la carte principale
    let map_size = Vec2::new(250.0, 250.0);
    // Taille de la minimap (20% de la largeur et 30% de la hauteur de la fenêtre)
    let minimap_size = Vec2::new(window_width * 0.2, window_height * 0.3);
    // Conversion des positions
    let x = ((player_position.x / map_size.x) * minimap_size.x) + (minimap_size.x / 2.0);
    let y = ((player_position.z / map_size.y) * minimap_size.y) + (minimap_size.y / 2.0);

    Vec2::new(x, y)
}
