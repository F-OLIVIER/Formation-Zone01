use crate::bevy_client::player_manager::*;
use crate::bevy_client::struct_manager::*;
use crate::interaction_server::config::{PlayerPublic, Players};
use crate::{
    bevy_client::{
        environment::*,
        setup_ennemy::*,
        struct_manager::{EnemyTag, LightTimer, MyResources, SendRateTimer},
    },
    interaction_server::config::*,
};
use bevy::utils::HashMap;
use bevy_rapier3d::prelude::Velocity;
use bevy_renet::renet::{ConnectionConfig, RenetClient};
use renet::transport::{ClientAuthentication, NetcodeClientTransport, NetcodeTransportError};
use std::{
    net::SocketAddr,
    net::UdpSocket,
    time::{Duration, SystemTime},
};

use super::config::EventMessage;

// Création du client Renet
pub fn setup_renet_client(mut commands: Commands) {
    // RenetClient::new(ConnectionConfig::default())
    let client = RenetClient::new(ConnectionConfig::default());
    commands.insert_resource(client);
}

// Création du transport Renet
pub fn setup_transport_layer(mut commands: Commands) {
    // Convertir l'adresse du serveur en `SocketAddr`
    // let server_addr: SocketAddr = "185.216.27.11:5001"
    //     .parse()
    //     .expect("Invalid server address");
    let server_addr: SocketAddr = "172.20.10.3:6000".parse().expect("Invalid server address");
    let socket = UdpSocket::bind("0.0.0.0:4001").unwrap();

    let current_time = SystemTime::now()
        .duration_since(SystemTime::UNIX_EPOCH)
        .unwrap();
    let client_id = current_time.as_millis() as u64;
    let authentication = ClientAuthentication::Unsecure {
        server_addr,
        client_id,
        user_data: None,
        protocol_id: 7,
    };

    let channel_transport = NetcodeClientTransport::new(current_time, authentication, socket)
        .expect("Failed to init NetcodeClientTransport");

    commands.insert_resource(channel_transport);
}

// Mise a jour du Transport Renet une fois le serveur sélectionner par l'utilisateur
pub fn update_transport_layer(commands: &mut Commands, players: &mut ResMut<Players>) {
    // Convertir l'adresse du serveur en `SocketAddr`
    let adress_socket = format!(
        "{}:6000",
        players.player_public.machine_information.connectedto
    );
    let server_addr: SocketAddr = adress_socket.parse().expect("Invalid server address");
    let socket = UdpSocket::bind("0.0.0.0:4000").unwrap();

    let current_time = SystemTime::now()
        .duration_since(SystemTime::UNIX_EPOCH)
        .unwrap();
    let client_id = current_time.as_millis() as u64;
    let authentication = ClientAuthentication::Unsecure {
        server_addr,
        client_id,
        user_data: None,
        protocol_id: 7,
    };

    let transport = NetcodeClientTransport::new(current_time, authentication, socket)
        .expect("Failed to init NetcodeClientTransport");

    commands.insert_resource(transport);
}

// Déconnection du client
pub fn disconnect_client(
    client: &mut ResMut<RenetClient>,
    transport: &mut ResMut<NetcodeClientTransport>,
    // mut commands: Commands
) {
    client.disconnect();
    transport.disconnect();
}

// Système de mise à jour du client pour gérer la connexion au serveur
pub fn client_update_system(
    players: Res<Players>,
    client: Option<ResMut<RenetClient>>,
    mut transport: ResMut<NetcodeClientTransport>,
) {
    // Condition pour vérifier si la connexion Client existe
    if let Some(mut client) = client {
        // Calcule la durée écoulée depuis la dernière mise à jour
        let duration = Duration::from_millis(8);

        client.update(duration);
        match transport.update(duration, &mut client) {
            Ok(_) => (),
            Err(e) => {
                warn!("Transport error: {:?}", e);
                send_player_info_to_server(&mut client, &mut transport, &players);
            }
        }

        let data_to_send = match bincode::serialize(&players.player_public) {
            Ok(msg) => msg,
            Err(e) => {
                eprintln!("Failed to serialize players.player_public: {}", e);
                return;
            }
        };

        client.send_message(1, data_to_send);
    } else {
        eprintln!("RenetClient resource is missing");
    }
}

// Fonction d'envoi des informations au serveur
pub fn handle_sending(
    mut client: ResMut<RenetClient>,
    mut transport: ResMut<NetcodeClientTransport>,
    players: ResMut<Players>,
    time: Res<Time>,                            // Pour gérer le temps écoulé
    mut send_rate_timer: ResMut<SendRateTimer>, // Le timer pour limiter l'envoi
) {
    // 1 envoie tout les x fps
    if send_rate_timer.send_timer.tick(time.delta()).finished() {
        send_player_info_to_server(&mut client, &mut transport, &players);
    }
}

pub fn send_player_info_to_server(
    client: &mut RenetClient,
    transport: &mut NetcodeClientTransport,
    players: &Players,
) {
    // println!("Client is connected. Preparing to send player info...");

    let player_entity = players.player_public.clone();
    let converted_msg = match bincode::serialize(&player_entity) {
        Ok(msg) => msg,
        Err(e) => {
            eprintln!("Failed to serialize player entity: {}", e);
            return;
        }
    };
    // Envoi du message (ne retourne pas de `Result`, donc pas besoin de vérifier)
    client.send_message(0, converted_msg.clone());

    // Envoyez les paquets via le transport
    let _ = transport.send_packets(client);
}

// If any error is found we just panic
#[allow(clippy::never_loop)]
pub fn panic_on_error_system(mut renet_error: EventReader<NetcodeTransportError>) {
    for e in renet_error.read() {
        panic!("{}", e);
    }
}

// Fonction de réception des messages courant (mise à jour des players => channel 1)
pub fn handle_incoming_message(
    mut client: ResMut<RenetClient>,
    mut players: ResMut<Players>,
    commands: Commands,
    ressource: Res<MyResources>,
    query: ParamSet<(
        Query<(Entity, &EnemyTag, &mut Transform)>, // Query des ennemis
        Query<(Entity, &ProjectileInfo, &mut Transform)>, // Query des projectiles
    )>,
) {
    handle_incoming_message_traitement(&mut client, &mut players, commands, &ressource, query);
}

pub fn handle_incoming_message_traitement(
    client: &mut ResMut<RenetClient>,
    players: &mut ResMut<Players>,
    mut commands: Commands,
    ressource: &Res<MyResources>,
    mut query: ParamSet<(
        Query<(Entity, &EnemyTag, &mut Transform)>, // Query des ennemis
        Query<(Entity, &ProjectileInfo, &mut Transform)>, // Query des projectiles
    )>,
) {
    while let Some(message) = client.receive_message(1) {
        match bincode::deserialize::<HashMap<u64, PlayerPublic>>(&message) {
            Ok(player_list) => {
                // Mise à jour de chaque joueur dans la liste reçue
                for current_player in player_list.iter() {
                    if current_player.1.info_player.id == players.player_public.info_player.id {
                        // Le joueur local, ne rien faire
                        continue;
                    }

                    // 1. Mise à jour des ennemis
                    update_or_create_enemy(
                        current_player.1,
                        &mut query.p0(),
                        &mut commands,
                        ressource,
                    );

                    // 2. Mise à jour des projectiles
                    update_or_create_projectiles(
                        current_player.1,
                        &mut query.p1(),
                        &mut commands,
                        ressource,
                    );

                    // 3. Suppression des projectiles disparu de la HashMap
                    remove_orphan_projectiles(current_player.1, &mut query.p1(), &mut commands);
                }

                // Mise à jour des données locales des joueurs
                players.all_players = player_list;
            }
            Err(e) => {
                eprintln!(
                    "Erreur lors de la désérialisation (Function : handle_incoming_message_traitement) : {:?}",
                    e
                );
            }
        }
    }
}

// Fonction l'envoi d'événement sporadique (channel 2)
pub fn send_event(client: &mut RenetClient, event: EventMessage) {
    let data_to_send = match bincode::serialize(&event) {
        Ok(data) => data,
        Err(e) => {
            println!("Failed to serialize event: {}", e);
            return;
        }
    };
    // Envoi de l'event
    client.send_message(2, data_to_send);
}

// Réception des événements (channel 2)
pub fn handle_incoming_event(
    mut client: ResMut<RenetClient>,
    mut logs: ResMut<LogsResource>,
    mut commands: Commands,
    mut timer: ResMut<LightTimer>,
    mut query: Query<(Entity, &EnemyTag)>,
    mut player_query: Query<(Entity, &mut Transform, &mut Velocity)>,
    mut players: ResMut<Players>,
) {
    while let Some(message) = client.receive_message(2) {
        match bincode::deserialize::<EventMessage>(&message) {
            Ok(event) => {
                // println!("event reçu : {:?}", event);

                match event.types.as_str() {
                    "player_join" | "player_leave" => {
                        // log arrivé et depart joueur
                        add_log(&mut logs, event.information);
                        if event.types.as_str() == "player_leave" {
                            for (entity, _) in query.iter_mut() {
                                commands.entity(entity).despawn();
                            }
                        }
                    }
                    "player_dead" => {
                        // L'event reçu dois correspondre au joueur qui joue
                        if let Some((_, mut transform, mut velocity)) = player_query
                            .iter_mut()
                            .find(|(_, _, _)| players.player_public.info_player.id == event.id)
                        {
                            // Ecran clignotant rouge
                            timer.active = true;
                            timer.duration.reset();
                            // Immobilisation du joueur
                            velocity.linvel = Vec3::ZERO;
                            // Changement de sa position
                            transform.translation = respawn_point().into();
                            // Rechargement de son arme
                            players.player_private.munition = MAX_MUNITION;
                        }
                    }
                    _ => {
                        println!("event non géré : {:?}", event);
                    }
                };
            }
            Err(e) => {
                eprintln!(
                    "Erreur lors de la désérialisation (Function : handle_incoming_event) : {:?}",
                    e
                );
            }
        }
    }
}
