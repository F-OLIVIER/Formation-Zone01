use super::struct_menu::DisplayQuality;
use crate::interaction_server::config::*;
use bevy::window::WindowMode;

pub fn quality_graphique(
    mut windows: Query<&mut Window>,
    display_quality: Res<DisplayQuality>,
    mut button_query: Query<(&mut BackgroundColor, &DisplayQuality)>,
) {
    // Ajuster la taille de la fenêtre
    if let Ok(mut window) = windows.get_single_mut() {
        match *display_quality {
            DisplayQuality::Windowed => {
                window.mode = WindowMode::Windowed;
            }
            DisplayQuality::BorderlessFullscreen => {
                window.mode = WindowMode::BorderlessFullscreen;
            }
            DisplayQuality::Fullscreen => {
                window.mode = WindowMode::Fullscreen;
            }
        }
    }

    // Mettre à jour la couleur des boutons en fonction de la qualité graphique
    for (mut background_color, current_quality) in button_query.iter_mut() {
        let color: Color;
        if *current_quality == *display_quality {
            color = Color::srgb(0.04, 2.55, 0.0);
        } else {
            color = NORMAL_BUTTON;
        }

        // Mettre à jour la couleur
        *background_color = color.into();
    }
}
