// ███████ ██████  ███████ 
// ██      ██   ██ ██      
// █████   ██████  ███████ 
// ██      ██           ██ 
// ██      ██      ███████ 

use bevy::diagnostic::{DiagnosticsStore, FrameTimeDiagnosticsPlugin};
use crate::interaction_server::config::*;
use bevy::prelude::*;

// ╭───────────────────────────────────────────────╮
// │                                               │
// │                   STRUCTURE                   │
// │                                               │
// ╰───────────────────────────────────────────────╯
#[derive(Debug, Component)]
pub struct FpsText;

#[derive(Debug, Resource)]
pub struct FrameTimer {
    pub elapsed: f32,
}

// ╭───────────────────────────────────────────────╮
// │                                               │
// │                    FONCTION                   │
// │                                               │
// ╰───────────────────────────────────────────────╯
pub fn display_fps(mut commands: Commands, asset_server: Res<AssetServer>) {
    // Text with multiple sections
    commands.spawn((
        // Create a TextBundle that has a Text with a list of sections.
        TextBundle::from_sections([
            TextSection::new(
                "FPS: ",  
                TextStyle {
                    font: asset_server.load(FONTS_MENU),
                    font_size: 30.0,
                    ..default()
                },
            ),
            #[allow(unexpected_cfgs)]
            TextSection::from_style(if cfg!(feature = "default_font") {
                TextStyle {
                    font_size: 30.0,
                    color: Color::WHITE,
                    ..default()
                }
            } else {
                TextStyle {
                    font: asset_server.load(FONTS_MENU),
                    font_size: 30.0,
                    color: Color::WHITE,
                }
            }),
        ]),
        FpsText,
    ));
}

pub fn update_fps(
    diagnostics: Res<DiagnosticsStore>,
    mut query: Query<&mut Text, With<FpsText>>,
    mut frame_counter: Local<u32>,
) {
    *frame_counter += 1;

    // Mettre à jour le texte tous les XX frames
    if *frame_counter % 120 == 0 {
        for mut text in &mut query {
            if let Some(fps) = diagnostics.get(&FrameTimeDiagnosticsPlugin::FPS) {
                if let Some(value) = fps.smoothed() {
                    // Mettre à jour la valeur de la section
                    text.sections[1].value = format!("{:.2}", value);
                }
            }
        }
    }
}
