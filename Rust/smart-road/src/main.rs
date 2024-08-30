mod mods;
use mods::*;

#[tokio::main]
pub async fn main() {
    launch_sdl2().await;
}

