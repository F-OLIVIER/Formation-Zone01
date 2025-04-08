pub mod mods;
#[allow(unused_imports)]
use bevy::{
    diagnostic::{FrameTimeDiagnosticsPlugin, LogDiagnosticsPlugin},
    prelude::*,
};
use multiplayer_fps::RenetServerPlugin;
use bevy_renet::transport::NetcodeServerPlugin;
use config::*;
pub use mods::*;
use udpsocket::*;

fn main() {
    let (server, transport) = setup_renet_server();

    App::new()
        .add_plugins(MinimalPlugins)
        .add_plugins(RenetServerPlugin)
        .add_plugins(NetcodeServerPlugin)
        .insert_resource(PlayerLobby::new())
        .insert_resource(PlayerPublic::new())
        .insert_resource(SendRateTimer(Timer::from_seconds(
            1.0 / 50.0, // frequence d'envoi des informations au clients (50 envois par seconde)
            TimerMode::Repeating,
        )))
        .insert_resource(server)
        .insert_resource(transport)
        .add_systems(
            Update,
            (
                server_update_system,
                handle_incoming_messages,
                handle_send_allplayer,
                handle_incoming_event,
                handle_server_events,
            ),
        )
        .run();
}
