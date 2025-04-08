use bevy::{prelude::*, utils::HashMap};
use serde::{Deserialize, Serialize};

pub const ADRESS_SERVER: &str = "127.0.0.1:4000";

#[allow(dead_code)]
pub struct IdGenerator {
    current_id: u64,
}
#[allow(dead_code)]
impl IdGenerator {
    fn new() -> Self {
        Self { current_id: 0 }
    }

    fn next_id(&mut self) -> u64 {
        self.current_id += 1;
        self.current_id
    }
}

#[derive(Serialize, Deserialize, Debug)]
pub struct NetworkInfo {
    pub ipaddress_local: String,
    pub ipaddress_public: String,
    pub connected_to: String,
}

#[derive(Resource, Debug, Clone)]
pub struct PlayerLobby {
    pub clients: HashMap<u64, PlayerPublic>,
}

impl PlayerLobby {
    pub fn new() -> Self {
        PlayerLobby {
            clients: HashMap::new(),
        }
    }
}

#[derive(Resource, Debug, Serialize, Deserialize, Clone)]
pub struct PlayerPublic {
    // pub common_information: CommonInformation, // information commune à tous les joueurs
    pub info_player: InfoPlayer,               //Informations basiques du joueur
    pub info_player_in_game: InfoPlayerInGame, // Information sur le player qui dois être envoyer a tous les autres joueur
    pub machine_information: MachineInformation, // Information pour pour le traitement serveur
}
//Initialiser ressource pour pas que seerveur panique (but c'est de le refaire passer par update)
impl PlayerPublic {
    pub fn new() -> PlayerPublic {
        PlayerPublic {
            // common_information: CommonInformation::new(),
            info_player: InfoPlayer::new(),
            info_player_in_game: InfoPlayerInGame::default(),
            machine_information: MachineInformation::new(),
        }
    }
}
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct MachineInformation {
    // adresse ip du joueur
    pub ipaddress_public: String,
    // adresse du serveur auquel il est connecté
    pub connectedto: String,
}
impl MachineInformation {
    fn new() -> Self {
        MachineInformation {
            ipaddress_public: String::new(),
            connectedto: String::new(),
        }
    }
}

// #[derive(Debug, Clone, serde::Serialize, Deserialize)]
// pub struct CommonInformation {
//     pub map: String,
//     pub timer_map: f32,
// }
// impl CommonInformation {
//     fn new() -> Self {
//         CommonInformation {
//             map: "".to_string(),
//             timer_map: 0.0,
//         }
//     }
// }

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct InfoPlayer {
    pub id: u64,
    pub username: String,
}
impl InfoPlayer {
    fn new() -> Self {
        InfoPlayer {
            id: 0,
            username: "".to_string(),
        }
    }
}

#[derive(Serialize, Deserialize, Debug, Clone, Resource)]
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

#[derive(Resource, Debug, Serialize, Deserialize)]
pub struct AllPlayerPublic {
    pub all_players: Vec<PlayerPublic>,
}
impl AllPlayerPublic {
    pub fn new() -> AllPlayerPublic {
        AllPlayerPublic {
            all_players: Vec::new(),
        }
    }
}

#[derive(Resource, Debug)]

pub struct SendRateTimer(pub Timer);

#[derive(Debug, Serialize, Deserialize)]

pub struct EventMessage {
    pub types: String,
    pub id: u64,
    pub username: String,
    pub information: String,
}
