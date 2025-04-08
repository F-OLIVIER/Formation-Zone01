use super::player_manager::*;
use super::struct_manager::*;
use crate::interaction_server::config::*;
use bevy::prelude::*;
use bevy::render::camera::Exposure;
use bevy_rapier3d::prelude::*;
use std::f32::consts::TAU;

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

pub fn setup_player(mut commands: Commands, resources: Res<MyResources>) {
    // Note that we have two entities for the player
    // One is a "logical" player that handles the physics computation and collision
    // The other is a "render" player that is what is displayed to the user
    // This distinction is useful for later on if you want to add multiplayer,
    // where often time these two ideas are not exactly synced up
    let height = 3.0;
    let logical_entity = commands
        .spawn((
            Collider::cylinder(height, 0.5),
            // A capsule can be used but is NOT recommended
            // If you use it, you have to make sure each segment point is
            // equidistant from the translation of the player transform
            // Collider::capsule_y(height / 2.0, 0.5),
            Friction {
                coefficient: 0.0,
                combine_rule: CoefficientCombineRule::Min,
            },
            Restitution {
                coefficient: 0.0,
                combine_rule: CoefficientCombineRule::Min,
            },
            ActiveEvents::COLLISION_EVENTS,
            Velocity::zero(),
            RigidBody::Dynamic,
            Sleeping::disabled(),
            LockedAxes::ROTATION_LOCKED,
            AdditionalMassProperties::Mass(1.0),
            GravityScale(0.0),
            Ccd { enabled: true }, // Prevent clipping when going fast
            TransformBundle::from_transform(Transform::from_translation(respawn_point().into())),
            LogicalPlayer,
            FpsControllerInput {
                pitch: -TAU / 12.0,
                yaw: TAU * 5.0 / 8.0,
                ..default()
            },
            FpsController {
                air_acceleration: 80.0,
                ..default()
            },
        ))
        .insert(CameraConfig {
            height_offset: -0.5,
        })
        .id();

    // Camera player
    commands
        .spawn((
            Player,
            Camera3dBundle {
                projection: Projection::Perspective(PerspectiveProjection {
                    fov: TAU / 5.0,
                    ..default()
                }),
                exposure: Exposure::SUNLIGHT,
                ..default()
            },
            RenderPlayer { logical_entity },
        ))
        .insert(VisibilityBundle {
            visibility: Visibility::default(),
            inherited_visibility: InheritedVisibility::default(),
            view_visibility: ViewVisibility::default(),
        })
        .with_children(|camera_parent| {
            // Arme du joueur
            for (weapon_mesh, weapon_material) in resources.weapon.iter() {
                camera_parent.spawn((MaterialMeshBundle {
                    mesh: weapon_mesh.clone(),
                    material: weapon_material.clone(),
                    // wand
                    transform: Transform::from_xyz(0.4, -0.9, -1.0).with_scale(Vec3::splat(0.05)),
                    ..default()
                },));
            }

            // viseur du joueur
            camera_parent.spawn(MaterialMeshBundle {
                mesh: resources.viseur.0.clone(),
                material: resources.viseur.1.clone(),
                transform: Transform::from_xyz(0.03, -0.02, -0.30),
                ..default()
            });
        });
}
