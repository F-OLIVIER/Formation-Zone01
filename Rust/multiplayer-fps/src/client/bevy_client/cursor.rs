//  ██████ ██    ██ ██████  ███████  ██████  ██████
// ██      ██    ██ ██   ██ ██      ██    ██ ██   ██
// ██      ██    ██ ██████  ███████ ██    ██ ██████
// ██      ██    ██ ██   ██      ██ ██    ██ ██   ██
//  ██████  ██████  ██   ██ ███████  ██████  ██   ██

use bevy::{prelude::*, window::CursorGrabMode};

// ╭───────────────────────────────────────────────╮
// │                                               │
// │                   STRUCTURE                   │
// │                                               │
// ╰───────────────────────────────────────────────╯

// ╭───────────────────────────────────────────────╮
// │                                               │
// │                    FONCTION                   │
// │                                               │
// ╰───────────────────────────────────────────────╯

// Modification du curseur
pub fn cursor(mut windows: Query<&mut Window>) {
    let mut window = windows.single_mut();
    // Changer l'apparence du curseur
    window.cursor.icon = CursorIcon::Crosshair; 
    
    // Capture le curseur dans la fenetre en fonction du systéme d'exploitation
    #[cfg(target_os = "linux")]
    {
        window.cursor.grab_mode = CursorGrabMode::None;
    };
    #[cfg(target_os = "windows")]
    {
        window.cursor.grab_mode = CursorGrabMode::Confined;
    };
    // Mac ne permet pas la prise en charge de la capture de la souris
    #[cfg(target_os = "macos")]
    {
        window.cursor.grab_mode = CursorGrabMode::None;
    };

}
