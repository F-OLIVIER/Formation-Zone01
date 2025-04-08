//   ██████  ██████  ███    ██ ███████ ██  ██████       ██████  ██       ██████  ██████   █████  ██
//  ██      ██    ██ ████   ██ ██      ██ ██           ██       ██      ██    ██ ██   ██ ██   ██ ██
//  ██      ██    ██ ██ ██  ██ █████   ██ ██   ███     ██   ███ ██      ██    ██ ██████  ███████ ██
//  ██      ██    ██ ██  ██ ██ ██      ██ ██    ██     ██    ██ ██      ██    ██ ██   ██ ██   ██ ██
//   ██████  ██████  ██   ████ ██      ██  ██████       ██████  ███████  ██████  ██████  ██   ██ ███████

pub use bevy::prelude::*;
use bevy::utils::HashMap;
use rand::prelude::*;
use serde::{Deserialize, Serialize};

// ╭───────────────────────────────────────────────╮
// │                                               │
// │                   SERVER                      │
// │                                               │
// ╰───────────────────────────────────────────────╯
pub const ADRESS_SERVER: &str = "0.0.0.0:6000";

// ╭───────────────────────────────────────────────╮
// │                                               │
// │                  GAME MENU                    │
// │                                               │
// ╰───────────────────────────────────────────────╯
/////////////////////// CONSTS ///////////////////////
pub const NORMAL_BUTTON: Color = Color::srgb(0.15, 0.15, 0.15);
pub const HOVERED_BUTTON: Color = Color::srgb(0.25, 0.25, 0.25);
pub const HOVERED_PRESSED_BUTTON: Color = Color::srgb(0.25, 0.65, 0.25);
pub const PRESSED_BUTTON: Color = Color::srgb(0.35, 0.75, 0.35);
pub const TEXT_COLOR: Color = Color::srgb(0.9, 0.9, 0.9);
pub const CURSOR_HANDLE: Handle<Font> = Handle::weak_from_u128(10482756907980398621);
pub const DEFAULT_RENDER_LAYER: usize = 0;
pub const VIEW_MODEL_RENDER_LAYER: usize = 1;
pub const FONTS_MENU: &str = "../assets/fonts/FiraSans-Regular.ttf";

pub const TIMER_GAME: f32 = 600.0; // temps en secondes
pub const MAX_LIFE: f32 = 200.0;
pub const MAX_CHARACTERS: usize = 15;

// ╭───────────────────────────────────────────────╮
// │                                               │
// │                    PLAYER                     │
// │                                               │
// ╰───────────────────────────────────────────────╯
/////////////////////// CONSTS ///////////////////////
pub const ANGLE_EPSILON: f32 = 0.001953125;
pub const GROUNDED_DISTANCE: f32 = 0.125;
pub const SLIGHT_SCALE_DOWN: f32 = 0.9375;
// pub const CT_SPAWN : Vec3 = Vec3::new(5.14, -1.61, -61.62);
// pub const SPAWN_POINT_1: (f32, f32, f32) = (0.0, 5.0, 0.0);
// pub const SPAWN_POINT_2: (f32, f32, f32) = (63.474808, 3.230173, -90.172165);
// pub const SPAWN_POINT_3: (f32, f32, f32) = (20.49373, 3.2090104, 119.404625);
// pub const SPAWN_POINT_4: (f32, f32, f32) = (0.49373, 3.2090104, 119.404625);
pub const LIST_SPAWN: [(f32, f32, f32); 9] = [
    (75.74387, 3.255145, -94.71171),
    (-45.05372, 3.1925905, -16.16917),
    (-49.391758, 3.2922134, 1.2592349),
    (-60.067814, 3.2921305, 16.628012),
    (-113.20835, 3.2294383, 6.285963),
    (23.708323, 29.056433, 61.955837),
    (32.642612, 3.2168102, 57.111877),
    (88.25176, 3.2161734, 41.512383),
    (74.41631, 3.2240748, -7.814401),
];

pub fn respawn_point() -> (f32, f32, f32) {
    *LIST_SPAWN.choose(&mut rand::thread_rng()).unwrap()
}

pub const SHOOTING_COOLDOWN: f32 = 0.2;
pub const MAX_MUNITION: usize = 20;
////////////////////// RESSOURCES /////////////////////

/////////////////////// STRUCT ///////////////////////
///

#[derive(Resource, Debug, Clone, Serialize, Deserialize, Component)]
pub struct Players {
    pub player_private: PlayerPrivate, // information du joueur courant pour traitement local uniquement
    pub player_public: PlayerPublic,   // information du joueur courant pour envoie serveur
    pub all_players: HashMap<u64, PlayerPublic>, // Information sur tous les joueurs (reçu du serveur)
}
impl Players {
    pub fn new() -> Players {
        Players {
            player_private: PlayerPrivate {
                velocity: 0.0,
                munition: MAX_MUNITION,
            },

            player_public: PlayerPublic {
                // common_information: CommonInformation {
                //     map: "map_z01".to_string(),
                //     timer_map: TIMER_GAME,
                // },
                info_player: InfoPlayer {
                    id: 0,
                    username: "".to_string(),
                },
                info_player_in_game: InfoPlayerInGame::default(),
                machine_information: MachineInformation {
                    ipaddress_public: "".to_string(),
                    connectedto: "127.0.0.1".to_string(),
                },
            },

            all_players: HashMap::new(),
        }
    }
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlayerPrivate {
    pub velocity: f32, // Vitesse du joueur (pour la mise à jour de la vélocité des projectiles)
    pub munition: usize,
}

#[derive(Debug, Clone, Serialize, Deserialize, Resource)]
pub struct PlayerPublic {
    // pub common_information: CommonInformation, // Information commune à tous les joueurs
    pub info_player: InfoPlayer, // Informations basiques du joueur
    pub info_player_in_game: InfoPlayerInGame, // Information sur le player qui dois être envoyer a tous les autres joueur
    pub machine_information: MachineInformation, // Information pour pour le traitement serveur
}

// #[derive(Debug, Clone, Serialize, Deserialize)]
// pub struct CommonInformation {
//     pub map: String,
//     pub timer_map: f32,
// }

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct InfoPlayer {
    pub id: u64,
    pub username: String,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct InfoPlayerInGame {
    pub position: (f32, f32, f32),
    pub rotation: (f32, f32),
    // pub health: u32,
    pub ammo: u32,
    pub score: usize,
    pub projectile: HashMap<u128, ((f32, f32, f32), Quat)>, // Vec<((f32, f32, f32), (f32, f32))>,
}

impl Default for InfoPlayerInGame {
    fn default() -> Self {
        InfoPlayerInGame {
            position: (0.0, 0.0, 0.0),
            rotation: (0.0, 0.0),
            // health: 100,
            ammo: 20,
            score: 0,
            projectile: HashMap::new(), // Vec::new(),
        }
    }
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct MachineInformation {
    // adresse ip du joueur
    pub ipaddress_public: String,
    //adresse du serveur auquel il est connecté
    pub connectedto: String,
}

#[derive(SystemSet, Debug, Clone, Copy, PartialEq, Eq, Hash)]
pub struct Connected;

#[derive(Debug, Serialize, Deserialize)]

pub struct EventMessage {
    pub types: String,
    pub id: u64,
    pub username: String,
    pub information: String,
}
