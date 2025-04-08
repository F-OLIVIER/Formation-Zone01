use bevy_client::moteur::init_bevy;

pub mod bevy_client;
pub mod interaction_server;

fn main() {
    // init_udp_client("127.0.0.1:8080",  "127.0.0.1:8081");
    init_bevy();
    // interaction_server::udpsocket::init_udp();
}
