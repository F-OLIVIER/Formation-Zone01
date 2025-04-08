use crate::{EventMessage, SendRateTimer};

use super::config::{PlayerLobby, PlayerPublic};
use bevy::prelude::*;
use renet::transport::{
    NetcodeServerTransport, NetcodeTransportError, ServerAuthentication, ServerConfig,
};
use renet::{ConnectionConfig, RenetServer, ServerEvent};
use std::net::{IpAddr, Ipv4Addr, SocketAddr, UdpSocket};
use std::time::{Duration, SystemTime};

pub fn setup_renet_server() -> (RenetServer, NetcodeServerTransport) {
    let public_addr = SocketAddr::new(IpAddr::V4(Ipv4Addr::new(0, 0, 0, 0)), 6000);
    let socket = UdpSocket::bind(public_addr).expect("Failed to bind socket");

    // Création du serveur Renet avec le socket UDP
    let server = RenetServer::new(ConnectionConfig::default());
    let current_time: std::time::Duration = SystemTime::now()
        .duration_since(SystemTime::UNIX_EPOCH)
        .unwrap();

    let server_config = ServerConfig {
        max_clients: 16,
        protocol_id: 7,
        public_addresses: vec![public_addr],
        authentication: ServerAuthentication::Unsecure,
        current_time,
    };

    println!("{:?}", server_config.public_addresses);

    let transport = NetcodeServerTransport::new(server_config, socket).unwrap();
    (server, transport)
}

// Système de mise à jour du transport
pub fn server_update_system(
    mut server: ResMut<RenetServer>,
    mut _commands: Commands,
    time: Res<Time>,
    mut transport: ResMut<NetcodeServerTransport>,
) {
    let duration = Duration::from_secs_f32(time.delta_seconds());

    server.update(duration);
    match transport.update(duration, &mut server) {
        Ok(_) => (),
        Err(e) => panic!("{e}"),
    }
}

// Réception des messages client (channel 1)
pub fn handle_incoming_messages(mut server: ResMut<RenetServer>, mut lobby: ResMut<PlayerLobby>) {
    // println!("----------------------------------------------------------------");
    for client_id in server.clients_id() {
        // Mise à jours des informations clients
        while let Some(message) = server.receive_message(client_id, 1) {
            match bincode::deserialize::<PlayerPublic>(&message) {
                Ok(player_entity) => {
                    if player_entity.info_player.id != 0 {
                        //
                        if !lobby.clients.contains_key(&player_entity.info_player.id) {
                            // Envoi aux clients le message d'arrivé d'un player
                            let event = EventMessage {
                                types: "player_join".to_string(),
                                id: player_entity.info_player.id,
                                username: player_entity.info_player.username.clone(),
                                information: format!(
                                    "{} join the game",
                                    player_entity.info_player.username
                                ),
                            };
                            let data_to_send = bincode::serialize(&event).unwrap();
                            server.broadcast_message(2, data_to_send);

                            // Crée les informations du players dans la ressource
                            lobby
                                .clients
                                .insert(player_entity.info_player.id, player_entity);
                        } else {
                            // Met à jour les informations du players dans la ressource
                            lobby
                                .clients
                                .insert(player_entity.info_player.id, player_entity);
                        }
                    }
                }
                Err(e) => {
                    eprintln!("Failed to deserialize player_entity: {:?}", e);
                }
            }
        }
    }

    // affiche la liste des utilisateurs de la map
    // for (_, user) in lobby.clients.iter() {
    //     println!("user : {}", user.info_player.username);
    // }
}

// Envois des messages au clients connecté (channel 1)
pub fn handle_send_allplayer(
    mut server: ResMut<RenetServer>,
    lobby: Res<PlayerLobby>,
    time: Res<Time>,                            // Pour gérer le temps écoulé
    mut send_rate_timer: ResMut<SendRateTimer>, // Le timer pour limiter l'envoi
) {
    if send_rate_timer.0.tick(time.delta()).finished() {
        // envoi au clients des informations
        let data_to_send = bincode::serialize(&lobby.clients).unwrap();
        server.broadcast_message(1, data_to_send);
    }
}

// Réception des evenement client (channel 2)
pub fn handle_incoming_event(mut server: ResMut<RenetServer>, lobby: ResMut<PlayerLobby>) {
    for client_id in server.clients_id() {
        while let Some(message) = server.receive_message(client_id, 2) {
            match bincode::deserialize::<EventMessage>(&message) {
                Ok(event) => {
                    println!("event recu : {:?}", event);

                    let data_to_send: Vec<u8> = match event.types.as_str() {
                        "player_dead" => {
                            let mut result = Vec::new(); // Crée un vecteur pour stocker les données sérialisées

                            // Itère sur les clients et vérifie si l'événement correspond à un joueur
                            if let Some(current_player) = lobby
                                .clients
                                .get(&event.information.parse::<u64>().unwrap())
                            {
                                let transform_event = EventMessage {
                                    types: "player_dead".to_string(),
                                    id: current_player.info_player.id,
                                    username: current_player.info_player.username.to_string(),
                                    information: format!(
                                        "{} has killed {}",
                                        event.username, current_player.info_player.username
                                    ),
                                };

                                // println!("transform event: {:?}", transform_event);

                                // Sérialisation de l'événement unique
                                match bincode::serialize(&transform_event) {
                                    Ok(serialized) => {
                                        result = serialized; // Stocke les données sérialisées dans le vecteur
                                    }
                                    Err(e) => {
                                        println!("Erreur de sérialisation: {:?}", e);
                                    }
                                }
                            } else {
                                println!("Aucun joueur trouvé pour l'ID : {}", event.information);
                            }

                            result // Retourne le vecteur de données sérialisées
                        }
                        _ => {
                            println!("Événement non géré : {:?}", event);
                            Vec::new() // Retourne un vecteur vide pour les événements non gérés
                        }
                    };

                    // renvoi de l'event aux clients
                    if !data_to_send.is_empty() {
                        server.broadcast_message(2, data_to_send);
                    }
                }
                Err(e) => {
                    eprintln!("Failed to deserialize event: {:?}", e);
                }
            }
        }
    }
}

// If any error is found we just panic
#[allow(clippy::never_loop)]
pub fn panic_on_error_system(mut renet_error: EventReader<NetcodeTransportError>) {
    for e in renet_error.read() {
        panic!("{}", e);
    }
}

pub fn handle_server_events(
    mut server_events: EventReader<ServerEvent>,
    mut server: ResMut<RenetServer>,
    mut lobby: ResMut<PlayerLobby>,
) {
    for event in server_events.read() {
        match event {
            ServerEvent::ClientConnected { client_id } => {
                println!("Client {} connected", client_id);
            }
            ServerEvent::ClientDisconnected { client_id, reason } => {
                println!("Client {} disconnected (reason: {})", client_id, reason);

                let id_client = client_id.raw();
                if let Some(user_disconnected) = lobby.clients.get(&id_client) {
                    let event = EventMessage {
                        types: "player_leave".to_string(),
                        id: user_disconnected.info_player.id,
                        username: user_disconnected.info_player.username.clone(),
                        information: format!(
                            "{} quit the game",
                            user_disconnected.info_player.username
                        ),
                    };
                    let data_to_send = bincode::serialize(&event).unwrap();
                    server.broadcast_message(2, data_to_send);

                    // Crée les informations du players dans la ressource
                    lobby.clients.remove(&event.id);
                }
            }
        }
    }
}
