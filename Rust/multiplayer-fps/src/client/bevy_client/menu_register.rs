// ███    ███ ███████ ███    ██ ██    ██     ██████  ███████  ██████  ██ ███████ ████████ ███████ ██████
// ████  ████ ██      ████   ██ ██    ██     ██   ██ ██      ██       ██ ██         ██    ██      ██   ██
// ██ ████ ██ █████   ██ ██  ██ ██    ██     ██████  █████   ██   ███ ██ ███████    ██    █████   ██████
// ██  ██  ██ ██      ██  ██ ██ ██    ██     ██   ██ ██      ██    ██ ██      ██    ██    ██      ██   ██
// ██      ██ ███████ ██   ████  ██████      ██   ██ ███████  ██████  ██ ███████    ██    ███████ ██   ██

use crate::{bevy_client::struct_menu::*, interaction_server::config::*};
use bevy::{
    color::palettes::css::CRIMSON,
    input::{keyboard::KeyboardInput, ButtonState},
    prelude::*,
};

use super::struct_manager::MyResources;

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

// Cette fonction est généralement appelée au début pour initialiser l'état du menu.
pub fn register_menu(
    mut commands: Commands,
    asset_server: Res<AssetServer>,
    resources: Res<MyResources>,
) {
    println!("On Register Menu");
    let button_style = Style {
        width: Val::Px(200.0),
        height: Val::Px(65.0),
        margin: UiRect::all(Val::Px(20.0)),
        justify_content: JustifyContent::Center,
        align_items: AlignItems::Center,
        ..default()
    };
    let button_text_style = TextStyle {
        font_size: 40.0,
        color: TEXT_COLOR,
        ..default()
    };

    // Arrière-plan (pour couvrir le jeu)
    commands
        .spawn((
            NodeBundle {
                style: Style {
                    width: Val::Percent(100.0),
                    height: Val::Percent(100.0),
                    align_items: AlignItems::Center,
                    justify_content: JustifyContent::Center,
                    ..default()
                },
                ..default()
            },
            OnRegisterScreen,
        ))
        .with_children(|parent| {
            // Chargement image background
            parent.spawn(ImageBundle {
                image: UiImage {
                    texture: resources.wallpaper.clone(),
                    ..default()
                },
                style: Style {
                    width: Val::Percent(100.0),
                    height: Val::Percent(100.0),
                    position_type: PositionType::Absolute,
                    ..default()
                },
                ..default()
            });

            // Conteneur principal alignant les inputs et les boutons côte à côte
            parent
                .spawn(NodeBundle {
                    style: Style {
                        flex_direction: FlexDirection::Row, // Colonne pour aligner à gauche (inputs) et à droite (boutons)
                        align_items: AlignItems::Center,
                        justify_content: JustifyContent::SpaceBetween, // Espace entre les deux colonnes
                        width: Val::Percent(80.0), // Taille globale de la section (80% de l'écran)
                        height: Val::Auto,
                        ..default()
                    },
                    background_color: CRIMSON.into(),
                    ..default()
                })
                .with_children(|parent| {
                    // Colonne des inputs à gauche
                    parent
                        .spawn(NodeBundle {
                            style: Style {
                                flex_direction: FlexDirection::Column, // Empile les inputs verticalement
                                width: Val::Percent(30.0),
                                height: Val::Percent(100.0),
                                align_items: AlignItems::Center,
                                justify_content: JustifyContent::Center,
                                ..default()
                            },
                            ..default()
                        })
                        .with_children(|parent| {
                            // Titre "Enter your username"
                            parent.spawn(TextBundle::from_section(
                                "Enter your username",
                                TextStyle {
                                    font: asset_server.load(FONTS_MENU),
                                    font_size: 30.0,
                                    color: Color::WHITE,
                                },
                            ));
                            // Champ de texte pour l'input du nom d'utilisateur
                            parent
                                .spawn((
                                    NodeBundle {
                                        style: Style {
                                            width: Val::Px(250.0),
                                            height: Val::Px(65.0),
                                            margin: UiRect::all(Val::Px(10.0)),
                                            justify_content: JustifyContent::Center,
                                            align_items: AlignItems::Center,
                                            max_width: Val::Px(230.0),
                                            ..default()
                                        },
                                        background_color: BackgroundColor(Color::srgba(
                                            1.0, 1.0, 1.0, 1.0,
                                        )),
                                        ..default()
                                    },
                                    UsernameInputField,
                                ))
                                .with_children(|parent| {
                                    parent.spawn((
                                        TextBundle::from_section(
                                            "", // Contenu initial vide
                                            TextStyle {
                                                font: asset_server.load(FONTS_MENU),
                                                font_size: 30.0,
                                                color: Color::BLACK,
                                            },
                                        ),
                                        UsernameInputText,
                                    ));
                                });
                        });

                    // Gestion de la selection du serveur
                    parent
                        .spawn(NodeBundle {
                            style: Style {
                                flex_direction: FlexDirection::Column,
                                width: Val::Percent(30.0),
                                height: Val::Percent(100.0),
                                align_items: AlignItems::Center,
                                justify_content: JustifyContent::Center,
                                ..default()
                            },
                            ..default()
                        })
                        .with_children(|parent| {
                            // Sélecteur de serveur
                            parent
                                .spawn((
                                    ButtonBundle {
                                        style: button_style.clone(),
                                        background_color: NORMAL_BUTTON.into(),
                                        visibility: Visibility::Visible,
                                        ..default()
                                    },
                                    MenuButtonAction::ServersDisplay,
                                    ServerInputField,
                                ))
                                .with_children(|parent| {
                                    parent.spawn((
                                        TextBundle::from_section(
                                            "Select server",
                                            TextStyle {
                                                font: asset_server.load(FONTS_MENU),
                                                font_size: 30.0,
                                                color: Color::WHITE,
                                            },
                                        ),
                                        ServerInputText,
                                    ));

                                    // Menu déroulant des serveurs (caché par défaut)
                                    parent
                                        .spawn((
                                            NodeBundle {
                                                style: Style {
                                                    position_type: PositionType::Absolute,
                                                    display: Display::None,
                                                    flex_direction: FlexDirection::Column,
                                                    align_items: AlignItems::Center,
                                                    justify_content: JustifyContent::Center,
                                                    margin: UiRect::top(Val::Px(0.1)),
                                                    ..default()
                                                },
                                                background_color: Color::BLACK.into(),
                                                ..default()
                                            },
                                            ServerDropdownMenu,
                                        ))
                                        .with_children(|dropdown| {
                                            let servers = vec!["Local", "Yggdrasil"];
                                            for server in servers {
                                                dropdown
                                                    .spawn((
                                                        ButtonBundle {
                                                            style: button_style.clone(),
                                                            background_color: Color::WHITE.into(),
                                                            ..default()
                                                        },
                                                        ServerOption(server.to_string()),
                                                        MenuButtonAction::ServersDisplay,
                                                    ))
                                                    .with_children(|parent| {
                                                        parent.spawn(TextBundle::from_section(
                                                            server,
                                                            TextStyle {
                                                                font: asset_server.load(FONTS_MENU),
                                                                font_size: 25.0,
                                                                color: Color::BLACK,
                                                            },
                                                        ));
                                                    });
                                            }
                                        });
                                });
                        });

                    // Colonne des boutons à droite
                    parent
                        .spawn(NodeBundle {
                            style: Style {
                                flex_direction: FlexDirection::Column,
                                width: Val::Percent(30.0),
                                height: Val::Percent(100.0),
                                align_items: AlignItems::Center,
                                justify_content: JustifyContent::Center,
                                ..default()
                            },
                            ..default()
                        })
                        .with_children(|parent| {
                            // Conteneur parent pour les boutons
                            parent
                                .spawn(NodeBundle {
                                    style: Style {
                                        flex_direction: FlexDirection::Column, // Empile les boutons verticalement
                                        align_items: AlignItems::Center, // Centrer les boutons horizontalement
                                        ..default()
                                    },
                                    ..default()
                                })
                                .with_children(|parent| {
                                    // Boucle sur les boutons avec leurs actions respectives
                                    for (action, text) in [
                                        (MenuButtonAction::UsernameAndServer, "Start"),
                                        (MenuButtonAction::BackToMainMenu, "Back"),
                                    ] {
                                        parent
                                            .spawn((
                                                ButtonBundle {
                                                    style: button_style.clone(),
                                                    background_color: NORMAL_BUTTON.into(),
                                                    ..default()
                                                },
                                                action,
                                            ))
                                            .with_children(|parent| {
                                                parent.spawn(TextBundle::from_section(
                                                    text,
                                                    button_text_style.clone(),
                                                ));
                                            });
                                    }
                                });
                        });
                });
        });
}

pub fn text_input_system(
    mut key_evr: EventReader<KeyboardInput>,
    mut username: ResMut<UsernameInput>,
    mut query: Query<&mut Text, With<UsernameInputText>>,
) {
    for ev in key_evr.read() {
        let key_code = ev.key_code;
        match key_code {
            KeyCode::Backspace => {
                // Si l'utilisateur appuie sur "Backspace", on supprime le dernier caractère
                username.0.pop();
            }
            KeyCode::Enter => {
                // Si l'utilisateur appuie sur "Entrée", vous pouvez gérer cela ici
                // println!("L'utilisateur a appuyé sur Entrée : {}", username.0);
            }
            _ => {
                // Capturer les touches alphabétiques et numériques
                if ev.state == ButtonState::Pressed {
                    if username.0.len() < MAX_CHARACTERS {
                        if let Some(char) = key_code_to_char(key_code) {
                            username.0.push(char);
                        }
                    }
                }
            }
        }
    }
    // Mettre à jour l'affichage du texte avec le contenu du nom d'utilisateur
    for mut text in &mut query {
        text.sections[0].value = username.0.clone();
    }
}

// Fonction auxiliaire pour convertir KeyCode en char
pub fn key_code_to_char(key_code: KeyCode) -> Option<char> {
    match key_code {
        KeyCode::KeyA => Some('a'),
        KeyCode::KeyB => Some('b'),
        KeyCode::KeyC => Some('c'),
        KeyCode::KeyD => Some('d'),
        KeyCode::KeyE => Some('e'),
        KeyCode::KeyF => Some('f'),
        KeyCode::KeyG => Some('g'),
        KeyCode::KeyH => Some('h'),
        KeyCode::KeyI => Some('i'),
        KeyCode::KeyJ => Some('j'),
        KeyCode::KeyK => Some('k'),
        KeyCode::KeyL => Some('l'),
        KeyCode::KeyM => Some('m'),
        KeyCode::KeyN => Some('n'),
        KeyCode::KeyO => Some('o'),
        KeyCode::KeyP => Some('p'),
        KeyCode::KeyQ => Some('q'),
        KeyCode::KeyR => Some('r'),
        KeyCode::KeyS => Some('s'),
        KeyCode::KeyT => Some('t'),
        KeyCode::KeyU => Some('u'),
        KeyCode::KeyV => Some('v'),
        KeyCode::KeyW => Some('w'),
        KeyCode::KeyX => Some('x'),
        KeyCode::KeyY => Some('y'),
        KeyCode::KeyZ => Some('z'),
        KeyCode::Digit0 => Some('0'),
        KeyCode::Digit1 => Some('1'),
        KeyCode::Digit2 => Some('2'),
        KeyCode::Digit3 => Some('3'),
        KeyCode::Digit4 => Some('4'),
        KeyCode::Digit5 => Some('5'),
        KeyCode::Digit6 => Some('6'),
        KeyCode::Digit7 => Some('7'),
        KeyCode::Digit8 => Some('8'),
        KeyCode::Digit9 => Some('9'),
        _ => None, // Vous pouvez étendre cette liste si nécessaire
    }
}
pub fn handle_server_dropdown(
    interaction_query: &mut Query<
        (&Interaction, &ServerOption),
        (Changed<Interaction>, With<Button>),
    >,
    selected_server_text: &mut Query<&mut Text, With<ServerInputText>>,
    dropdown_visibility: &mut Query<&mut Style, With<ServerDropdownMenu>>,
) {
    for (&interaction, server_option) in interaction_query.iter_mut() {
        if interaction == Interaction::Pressed {
            // Met à jour le texte du serveur sélectionné
            for mut text in selected_server_text.into_iter() {
                text.sections[0].value = server_option.0.clone(); // Met à jour avec le nom du serveur
            }

            // Cache la liste déroulante après sélection
            for mut style in dropdown_visibility.into_iter() {
                style.display = Display::None;
            }
        }
    }
}
