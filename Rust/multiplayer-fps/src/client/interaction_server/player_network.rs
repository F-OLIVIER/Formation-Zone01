// à implémenter dans udpsocket côté client :
pub fn client_connection_system(
    mut players: ResMut<Players>,
    client: ResMut<RenetClient>, // Le client de renet
) {
    if client.is_connecting() {
        println!("Le client tente de se connecter...");
    } else if client.is_connected() {
        println!("Client connecté, envoi des informations au serveur...");

        // Met à jour les informations du joueur local avec l'ID du client
        players.player_public.info_player.id = client.client_id().unwrap_or_default();

        // Prépare les données à envoyer au serveur
        let player_data = players.player_public.clone();
        // Envoie les informations publiques au serveur (à implémenter avec renet)
        send_player_info_to_server(&player_data);
    }
}

// Fonction d'envoi des données du joueur au serveur
fn send_player_info_to_server(player_data: PlayerPublic) {
    // Logique pour envoyer les informations publiques au serveur via Renet
    // client.send_message(channel_id, player_data); -> Exemple
}

//mise à jour peut-être à faire aussi côté serveur 
pub fn update_all_players_from_server(
    mut players: ResMut<Players>,
    server_messages: Vec<PlayerPublic>, // Les messages reçus du serveur contenant les infos des autres joueurs
) {
    for player_info in server_messages {
        // Vérifie si le joueur est déjà dans la liste
        if let Some(existing_player) = players.all_players.iter_mut().find(|p| p.info_player.id == player_info.info_player.id) {
            // Met à jour les infos du joueur existant
            *existing_player = player_info;
        } else {
            // Sinon, ajoute le nouveau joueur
            players.all_players.push(player_info);
        }
    }

    println!("All players updated: {:?}", players.all_players);
}





use bevy::prelude::*;

use serde::{Serialize, Deserialize};

use bincode;

use std::{collections::HashMap, net::{SocketAddr, UdpSocket}};
use bevy::ecs::system::Res;

#[derive(Serialize, Deserialize, Debug)]
pub struct PlayerState {
    pub id: u32,
    pub position: (f32, f32, f32),
    pub rotation: (f32, f32),
}

#[derive(Resource)]
pub struct NetworkResource {
    socket: UdpSocket,
    player_id: u32,
}


pub fn send_player_state(query: Query<&Transform>, network: Res<NetworkResource>) {
    for transform in query.iter() {
        let player_state = PlayerState {
            id: network.player_id,
            position: (transform.translation.x, transform.translation.y, transform.translation.z),
            rotation: (transform.rotation.x, transform.rotation.y),
        };

        println!("{:?}", player_state);

        let serialized = bincode::serialize(&player_state).unwrap();
        network.socket.send_to(&serialized, "127.0.0.1:3000").unwrap();
    }
}

pub fn setup_network() -> NetworkResource {
    let socket = UdpSocket::bind("127.0.0.1:34254").unwrap();  // Ex: adresse locale pour ton client
    NetworkResource {
        socket,
        player_id: 1,  // Exemple d'ID de joueur
    }
}

pub fn receive_game_state(network: Res<NetworkResource>) {
    let mut buffer = [0u8; 1024];
    if let Ok((amt, _src)) = network.socket.recv_from(&mut buffer) {
        if let Ok(player_states) = bincode::deserialize::<HashMap<SocketAddr, PlayerState>>(&buffer[..amt]) {
            for (addr, player_state) in player_states {
                println!("Player {} at {:?}", player_state.id, player_state.position);
            }
        }
    }
}