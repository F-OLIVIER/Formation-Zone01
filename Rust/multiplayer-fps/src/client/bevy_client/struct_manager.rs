// ███████ ███    ██ ██    ██ ███    ██         ██     ███████ ████████ ██████  ██    ██  ██████ ████████
// ██      ████   ██ ██    ██ ████   ██        ██      ██         ██    ██   ██ ██    ██ ██         ██
// █████   ██ ██  ██ ██    ██ ██ ██  ██       ██       ███████    ██    ██████  ██    ██ ██         ██
// ██      ██  ██ ██ ██    ██ ██  ██ ██      ██             ██    ██    ██   ██ ██    ██ ██         ██
// ███████ ██   ████  ██████  ██   ████     ██         ███████    ██    ██   ██  ██████   ██████    ██
//
//  ██████   █████  ███    ███ ███████
// ██       ██   ██ ████  ████ ██
// ██   ███ ███████ ██ ████ ██ █████
// ██    ██ ██   ██ ██  ██  ██ ██
//  ██████  ██   ██ ██      ██ ███████
//
use super::{audio::Audio, struct_menu::GameState};
use crate::bevy_client::player_movement::*;
use bevy::input::{gamepad, keyboard, mouse, touch};
use bevy::prelude::*;
use std::collections::HashMap;
use multiplayer_fps::{Deserialize, Serialize};

//
// ╭───────────────────────────────────────────────╮
// │                                               │
// │                     ENUM                      │
// │                                               │
// ╰───────────────────────────────────────────────╯
// ╭───────────────────────────────────────────────╮
// │                                               │
// │                   STRUCTURE                   │
// │                                               │
// ╰───────────────────────────────────────────────╯

#[derive(Resource)]
pub struct SendRateTimer {
    pub send_timer: Timer,
    pub scoreboard_timer: Timer,
    pub minimap_timer: Timer,
}

#[derive(Debug, Resource)]
pub struct MyResources {
    pub wallpaper: Handle<Image>,
    pub map: HashMap<String, (Handle<Gltf>, Handle<Image>)>,
    pub viseur: (Handle<Mesh>, Handle<StandardMaterial>),
    pub weapon: Vec<(Handle<Mesh>, Handle<StandardMaterial>)>,
    pub projectile: (Handle<Mesh>, Handle<StandardMaterial>),
    pub audio: Audio,
    pub ennemy: (Handle<Scene>, Animations),
}

#[derive(Resource, Clone, Debug)]
pub struct Animations {
    pub animations: Vec<AnimationNodeIndex>,
    #[allow(dead_code)]
    pub graph: Handle<AnimationGraph>,
}

#[derive(Resource)]
pub struct MainScene {
    pub handle: Handle<Gltf>,
    pub is_loaded: bool,
}
// Clignotement de la lumiére
#[derive(Debug, Resource)]
pub struct LightTimer {
    pub timer: Timer,
    pub duration: Timer,
    pub active: bool,
}

// Ennemis
#[derive(Component)]
pub struct EnemyTag {
    pub id: u64,
}

// Projectile du joueur
#[derive(Component)]
pub struct ProjectileTag {
    pub id: u64,
}

// Projectile des autres joueurs
#[derive(Serialize, Deserialize, Debug, Component)]
pub struct ProjectileInfo {
    pub id_projectile: u128,
    pub position: (f32, f32, f32),
    pub rotation: (f32, f32),
}

#[derive(Serialize, Deserialize, Debug, Component)]
pub struct PlayerProjectiles {
    pub player_id: u64,
    pub projectiles: Vec<ProjectileInfo>,
}

#[derive(Debug, Resource)]
pub struct Resourcesprojectile {
    pub id_projectile: usize,
    pub list_id: Vec<usize>,
}

// ██████  ██       █████  ██    ██ ███████ ██████
// ██   ██ ██      ██   ██  ██  ██  ██      ██   ██
// ██████  ██      ███████   ████   █████   ██████
// ██      ██      ██   ██    ██    ██      ██   ██
// ██      ███████ ██   ██    ██    ███████ ██   ██
//
// ╭───────────────────────────────────────────────╮
// │                                               │
// │                     ENUM                      │
// │                                               │
// ╰───────────────────────────────────────────────╯
#[derive(PartialEq)]
pub enum MoveMode {
    Noclip,
    Ground,
}
// ╭───────────────────────────────────────────────╮
// │                                               │
// │                   STRUCTURE                   │
// │                                               │
// ╰───────────────────────────────────────────────╯

#[derive(Component)]
pub struct LogicalPlayer;
#[derive(Component)]
pub struct RenderPlayer {
    pub logical_entity: Entity,
}
#[derive(Component)]
pub struct CameraConfig {
    pub height_offset: f32,
}
#[derive(Component, Default)]
pub struct FpsControllerInput {
    pub fly: bool,
    pub sprint: bool,
    pub jump: bool,
    pub crouch: bool,
    pub pitch: f32,
    pub yaw: f32,
    pub movement: Vec3,
}
#[derive(Component)]
pub struct FpsController {
    pub move_mode: MoveMode,
    pub radius: f32,
    pub gravity: f32,
    pub walk_speed: f32,
    pub run_speed: f32,
    pub forward_speed: f32,
    pub side_speed: f32,
    pub air_speed_cap: f32,
    pub air_acceleration: f32,
    pub max_air_speed: f32,
    pub acceleration: f32,
    pub friction: f32,
    pub traction_normal_cutoff: f32,
    pub friction_speed_cutoff: f32,
    pub jump_speed: f32,
    pub fly_speed: f32,
    pub crouched_speed: f32,
    pub crouch_speed: f32,
    pub uncrouch_speed: f32,
    pub height: f32,
    pub upright_height: f32,
    pub crouch_height: f32,
    pub fast_fly_speed: f32,
    pub fly_friction: f32,
    pub pitch: f32,
    pub yaw: f32,
    pub ground_tick: u8,
    pub stop_speed: f32,
    pub sensitivity: f32,
    pub enable_input: bool,
    pub step_offset: f32,
    pub key_forward: KeyCode,
    pub key_back: KeyCode,
    pub key_left: KeyCode,
    pub key_right: KeyCode,
    pub key_up: KeyCode,
    pub key_down: KeyCode,
    pub key_sprint: KeyCode,
    pub key_jump: KeyCode,
    pub key_fly: KeyCode,
    pub key_crouch: KeyCode,
}
impl Default for FpsController {
    fn default() -> Self {
        Self {
            move_mode: MoveMode::Ground,
            radius: 0.5,
            fly_speed: 10.0,
            fast_fly_speed: 30.0,
            gravity: 23.0,
            walk_speed: 9.0,
            run_speed: 14.0,
            forward_speed: 30.0,
            side_speed: 30.0,
            air_speed_cap: 2.0,
            air_acceleration: 20.0,
            max_air_speed: 15.0,
            crouched_speed: 5.0,
            crouch_speed: 6.0,
            uncrouch_speed: 8.0,
            height: 3.0,
            upright_height: 3.0,
            crouch_height: 1.5,
            acceleration: 10.0,
            friction: 10.0,
            traction_normal_cutoff: 0.7,
            friction_speed_cutoff: 0.1,
            fly_friction: 0.5,
            pitch: 0.0,
            yaw: 0.0,
            ground_tick: 0,
            stop_speed: 1.0,
            jump_speed: 8.5,
            step_offset: 0.25,
            enable_input: true,
            key_forward: KeyCode::KeyW,
            key_back: KeyCode::KeyS,
            key_left: KeyCode::KeyA,
            key_right: KeyCode::KeyD,
            key_up: KeyCode::KeyQ,
            key_down: KeyCode::KeyE,
            key_sprint: KeyCode::ShiftLeft,
            key_jump: KeyCode::Space,
            key_fly: KeyCode::KeyF,
            key_crouch: KeyCode::ControlLeft,
            sensitivity: 0.001,
        }
    }
}
pub struct FpsControllerPlugin;
impl Plugin for FpsControllerPlugin {
    fn build(&self, app: &mut App) {
        app.add_systems(
            PreUpdate,
            (
                fps_controller_input,
                fps_controller_look,
                fps_controller_move,
                fps_controller_render,
            )
                .chain()
                .after(mouse::mouse_button_input_system)
                .after(keyboard::keyboard_input_system)
                .after(gamepad::gamepad_axis_event_system)
                .after(gamepad::gamepad_button_event_system)
                .after(gamepad::gamepad_connection_system)
                .after(gamepad::gamepad_event_system)
                .after(touch::touch_screen_input_system)
                .run_if(in_state(GameState::Game)),
        );
    }
}
