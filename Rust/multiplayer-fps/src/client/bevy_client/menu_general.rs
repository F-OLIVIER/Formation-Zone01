// ███    ███ ███████ ███    ██ ██    ██      ██████  ███████ ███    ██ ███████ ██████   █████  ██
// ████  ████ ██      ████   ██ ██    ██     ██       ██      ████   ██ ██      ██   ██ ██   ██ ██
// ██ ████ ██ █████   ██ ██  ██ ██    ██     ██   ███ █████   ██ ██  ██ █████   ██████  ███████ ██
// ██  ██  ██ ██      ██  ██ ██ ██    ██     ██    ██ ██      ██  ██ ██ ██      ██   ██ ██   ██ ██
// ██      ██ ███████ ██   ████  ██████       ██████  ███████ ██   ████ ███████ ██   ██ ██   ██ ███████

use super::{menu_register::handle_server_dropdown, struct_menu::*};
use crate::bevy_client::moteur::exit_system;
use crate::bevy_client::struct_manager::MyResources;
use crate::interaction_server::udpsocket::*;
use crate::interaction_server::{config::*, get_ip::*};
#[allow(private_interfaces)]
pub use bevy::{app::AppExit, color::palettes::css::CRIMSON, prelude::*};
use renet::transport::NetcodeClientTransport;
use renet::RenetClient;
use std::net::{IpAddr, Ipv4Addr};
use std::time::SystemTime;

// ╭───────────────────────────────────────────────╮
// │                                               │
// │                   STRUCTURE                   │
// │                                               │
// ╰───────────────────────────────────────────────╯
#[derive(Component)]
pub struct SelectedOption;

// ╭───────────────────────────────────────────────╮
// │                                               │
// │                    FONCTION                   │
// │                                               │
// ╰───────────────────────────────────────────────╯
// This system handles changing all buttons color based on mouse interaction
pub fn button_system(
    mut interaction_query: Query<
        (&Interaction, &mut UiImage, Option<&SelectedOption>),
        (Changed<Interaction>, With<Button>),
    >,
) {
    for (interaction, mut image, selected) in &mut interaction_query {
        image.color = match (*interaction, selected) {
            (Interaction::Pressed, _) | (Interaction::None, Some(_)) => PRESSED_BUTTON,
            (Interaction::Hovered, Some(_)) => HOVERED_PRESSED_BUTTON,
            (Interaction::Hovered, None) => HOVERED_BUTTON,
            (Interaction::None, None) => NORMAL_BUTTON,
        }
    }
}

// This system updates the settings when a new value for a setting is selected, and marks
// the button as the one currently selected
pub fn setting_button<T: Resource + Component + PartialEq + Copy>(
    interaction_query: Query<(&Interaction, &T, Entity), (Changed<Interaction>, With<Button>)>,
    mut selected_query: Query<(Entity, &mut UiImage), With<SelectedOption>>,
    mut commands: Commands,
    mut setting: ResMut<T>,
) {
    for (interaction, button_setting, entity) in &interaction_query {
        if *interaction == Interaction::Pressed && *setting != *button_setting {
            let (previous_button, mut previous_image) = selected_query.single_mut();
            previous_image.color = NORMAL_BUTTON;
            commands.entity(previous_button).remove::<SelectedOption>();
            commands.entity(entity).insert(SelectedOption);
            *setting = *button_setting;
        }
    }
}

/// Configure l'état du menu en définissant l'état suivant comme étant `Main`.
///
/// # Paramètre
/// - `menu_state`: Une ressource mutable de type `NextState<MenuState>`.
///   Elle permet de définir l'état suivant du jeu.
///
/// Cette fonction est généralement appelée au début pour initialiser l'état du menu.
pub fn menu_setup(mut menu_state: ResMut<NextState<MenuState>>, mut windows: Query<&mut Window>) {
    // Définit l'état suivant du jeu comme étant `MenuState::Main`.
    // Cela signifie que lorsque la transition d'état sera effectuée, le jeu passera à l'état `Main`.
    menu_state.set(MenuState::Main);
    // affichage du curseur
    let mut window = windows.single_mut();
    window.cursor.visible = true;
}

// chargement de l'image de background si le joueur et dans le menu initial
fn backgroung_img(
    parent: &mut ChildBuilder<'_>,
    resources: Res<MyResources>,
    etat_paused_state: Res<State<PausedState>>,
) {
    match etat_paused_state.get() {
        PausedState::On => (),
        PausedState::Off => {
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
        }
    }
}

/// Configure l'interface utilisateur pour le menu principal du jeu.
/// Ajoute un conteneur principal avec un titre et trois boutons : "New Game", "Settings", et "Quit".
///
/// # Paramètres
/// - `commands`: Un objet `Commands` utilisé pour créer et configurer des entités dans la scène.
/// - `asset_server`: Un objet `AssetServer` utilisé pour charger des assets comme des images.
pub fn main_menu_setup(
    mut commands: Commands,
    asset_server: Res<AssetServer>,
    resources: Res<MyResources>,
    etat_paused_state: Res<State<PausedState>>,
) {
    // Définition du style commun pour tous les boutons du menu
    let button_style = Style {
        width: Val::Px(250.0),
        height: Val::Px(65.0),
        margin: UiRect::all(Val::Px(20.0)),
        justify_content: JustifyContent::Center,
        align_items: AlignItems::Center,
        ..default() // Valeurs par défaut pour les autres propriétés
    };
    // Définition du style spécifique pour les icônes des boutons
    let button_icon_style = Style {
        width: Val::Px(30.0),
        position_type: PositionType::Absolute, // Positionnement absolu pour placer l'icône précisément
        left: Val::Px(10.0),                   // L'icône sera proche du bord gauche du bouton
        ..default()                            // Valeurs par défaut pour les autres propriétés
    };
    // Définition du style pour le texte dans les boutons
    let button_text_style = TextStyle {
        font_size: 40.0,
        color: TEXT_COLOR, // Couleur du texte
        ..default()        // Valeurs par défaut pour les autres propriétés
    };
    // changement de texte en fonction de l'etat de la pause
    let text_play = match etat_paused_state.get() {
        PausedState::On => "Resume",
        PausedState::Off => "Play",
    };

    // Création du conteneur principal du menu
    commands
        .spawn((
            NodeBundle {
                style: Style {
                    width: Val::Percent(100.0),      // Remplit toute la largeur de l'écran
                    height: Val::Percent(100.0),     // Remplit toute la hauteur de l'écran
                    align_items: AlignItems::Center, // Aligne les éléments enfants au centre verticalement
                    justify_content: JustifyContent::Center, // Aligne les éléments enfants au centre horizontalement
                    ..default() // Valeurs par défaut pour les autres propriétés
                },

                ..default() // Valeurs par défaut pour les autres propriétés
            },
            OnMainMenuScreen, // Composant personnalisé pour identifier cet élément comme étant l'écran du menu principal
        ))
        .with_children(|parent| {
            // Chargement image background si menu initial
            backgroung_img(parent, resources, etat_paused_state);

            parent
                .spawn(NodeBundle {
                    style: Style {
                        flex_direction: FlexDirection::Column, // Dispose les enfants en colonne
                        align_items: AlignItems::Center, // Aligne les enfants au centre horizontalement
                        ..default() // Valeurs par défaut pour les autres propriétés
                    },
                    background_color: CRIMSON.into(), // Couleur de fond du conteneur
                    ..default() // Valeurs par défaut pour les autres propriétés
                })
                .with_children(|parent| {
                    // Affiche le nom du jeu
                    parent.spawn(
                        TextBundle::from_section(
                            "Game Menu",
                            TextStyle {
                                font_size: 80.0,
                                color: TEXT_COLOR, // Couleur du texte
                                ..default()        // Valeurs par défaut pour les autres propriétés
                            },
                        )
                        .with_style(Style {
                            margin: UiRect::all(Val::Px(50.0)), // Marge autour du texte
                            ..default() // Valeurs par défaut pour les autres propriétés
                        }),
                    );
                    // Création du bouton "New Game"
                    parent
                        .spawn((
                            ButtonBundle {
                                style: button_style.clone(), // Utilisation du style défini pour les boutons
                                background_color: NORMAL_BUTTON.into(), // Couleur de fond du bouton
                                ..default() // Valeurs par défaut pour les autres propriétés
                            },
                            MenuButtonAction::Play, // Action associée à ce bouton
                        ))
                        .with_children(|parent| {
                            let icon = asset_server.load("textures/game_icons/right.png"); // Chargement de l'icône du bouton
                            parent.spawn(ImageBundle {
                                style: button_icon_style.clone(), // Utilisation du style défini pour les icônes
                                image: UiImage::new(icon), // Définition de l'image de l'icône
                                ..default() // Valeurs par défaut pour les autres propriétés
                            });
                            parent.spawn(TextBundle::from_section(
                                text_play,                 // Texte du bouton
                                button_text_style.clone(), // Utilisation du style de texte défini pour les boutons
                            ));
                        });
                    // Création du bouton "Settings"
                    parent
                        .spawn((
                            ButtonBundle {
                                style: button_style.clone(), // Utilisation du style défini pour les boutons
                                background_color: NORMAL_BUTTON.into(), // Couleur de fond du bouton
                                ..default() // Valeurs par défaut pour les autres propriétés
                            },
                            MenuButtonAction::Settings, // Action associée à ce bouton
                        ))
                        .with_children(|parent| {
                            let icon = asset_server.load("textures/game_icons/wrench.png"); // Chargement de l'icône du bouton
                            parent.spawn(ImageBundle {
                                style: button_icon_style.clone(), // Utilisation du style défini pour les icônes
                                image: UiImage::new(icon), // Définition de l'image de l'icône
                                ..default() // Valeurs par défaut pour les autres propriétés
                            });
                            parent.spawn(TextBundle::from_section(
                                "Settings",                // Texte du bouton
                                button_text_style.clone(), // Utilisation du style de texte défini pour les boutons
                            ));
                        });
                    // Création du bouton "Quit"
                    parent
                        .spawn((
                            ButtonBundle {
                                style: button_style, // Utilisation du style défini pour les boutons
                                background_color: NORMAL_BUTTON.into(), // Couleur de fond du bouton
                                ..default()          // Valeurs par défaut pour les autres propriétés
                            },
                            MenuButtonAction::Quit, // Action associée à ce bouton
                        ))
                        .with_children(|parent| {
                            let icon = asset_server.load("textures/game_icons/exitRight.png"); // Chargement de l'icône du bouton
                            parent.spawn(ImageBundle {
                                style: button_icon_style,  // Utilisation du style défini pour les icônes
                                image: UiImage::new(icon), // Définition de l'image de l'icône
                                ..default() // Valeurs par défaut pour les autres propriétés
                            });
                            parent.spawn(TextBundle::from_section(
                                "Quit",            // Texte du bouton
                                button_text_style, // Utilisation du style de texte défini pour les boutons
                            ));
                        });
                });
        });
}

pub fn settings_menu_setup(
    mut commands: Commands,
    resources: Res<MyResources>,
    etat_paused_state: Res<State<PausedState>>,
) {
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
            OnSettingsMenuScreen,
        ))
        .with_children(|parent| {
            // Chargement image background si menu initial
            backgroung_img(parent, resources, etat_paused_state);

            parent
                .spawn(NodeBundle {
                    style: Style {
                        flex_direction: FlexDirection::Column,
                        align_items: AlignItems::Center,
                        ..default()
                    },
                    background_color: CRIMSON.into(),
                    ..default()
                })
                .with_children(|parent| {
                    for (action, text) in [
                        (MenuButtonAction::SettingsDisplay, "Display"),
                        (MenuButtonAction::SettingsSound, "Sound"),
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
}

pub fn display_settings_menu_setup(
    mut commands: Commands,
    display_quality: Res<DisplayQuality>,
    resources: Res<MyResources>,
    etat_paused_state: Res<State<PausedState>>,
) {
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
            OnDisplaySettingsMenuScreen,
        ))
        .with_children(|parent| {
            // Chargement image background si menu initial
            backgroung_img(parent, resources, etat_paused_state);

            parent
                .spawn(NodeBundle {
                    style: Style {
                        flex_direction: FlexDirection::Column,
                        align_items: AlignItems::Center,
                        ..default()
                    },
                    background_color: CRIMSON.into(),
                    ..default()
                })
                .with_children(|parent| {
                    // Create a new `NodeBundle`, this time not setting its `flex_direction`. It will
                    // use the default value, `FlexDirection::Row`, from left to right.
                    parent
                        .spawn(NodeBundle {
                            style: Style {
                                align_items: AlignItems::Center,
                                ..default()
                            },
                            background_color: CRIMSON.into(),
                            ..default()
                        })
                        .with_children(|parent| {
                            // Display a label for the current setting
                            parent.spawn(TextBundle::from_section(
                                "Graphic size",
                                button_text_style.clone(),
                            ));
                            // Display a button for each possible value
                            for quality_setting in [
                                DisplayQuality::Windowed,
                                DisplayQuality::BorderlessFullscreen,
                                DisplayQuality::Fullscreen,
                            ] {
                                let color: Color;
                                if *display_quality == quality_setting {
                                    color = Color::srgb(0.04, 2.55, 0.0);
                                } else {
                                    color = NORMAL_BUTTON;
                                }

                                let contenttext = match quality_setting {
                                    DisplayQuality::Windowed => "Windowed",
                                    DisplayQuality::BorderlessFullscreen => {
                                        "Borderless\nFullscreen"
                                    }
                                    DisplayQuality::Fullscreen => "Fullscreen",
                                };
                                println!("contenttext : {}", contenttext);

                                let mut entity = parent.spawn((
                                    ButtonBundle {
                                        style: Style {
                                            width: Val::Px(150.0),
                                            height: Val::Px(65.0),
                                            ..button_style.clone()
                                        },
                                        background_color: color.into(),
                                        ..default()
                                    },
                                    quality_setting,
                                ));
                                entity.with_children(|parent| {
                                    parent.spawn(TextBundle::from_section(
                                        contenttext,
                                        TextStyle {
                                            font_size: 25.0,
                                            color: TEXT_COLOR,
                                            ..default()
                                        },
                                    ));
                                });
                                if *display_quality == quality_setting {
                                    entity.insert(SelectedOption);
                                }
                            }
                        });
                    // Display the back button to return to the settings screen
                    parent
                        .spawn((
                            ButtonBundle {
                                style: button_style,
                                background_color: NORMAL_BUTTON.into(),
                                ..default()
                            },
                            MenuButtonAction::BackToSettings,
                        ))
                        .with_children(|parent| {
                            parent.spawn(TextBundle::from_section("Back", button_text_style));
                        });
                });
        });
}

pub fn sound_settings_menu_setup(
    mut commands: Commands,
    volume: Res<AudioVolume>,
    resources: Res<MyResources>,
    etat_paused_state: Res<State<PausedState>>,
) {
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
            OnSoundSettingsMenuScreen,
        ))
        .with_children(|parent| {
            // Chargement image background si menu initial
            backgroung_img(parent, resources, etat_paused_state);

            parent
                .spawn(NodeBundle {
                    style: Style {
                        flex_direction: FlexDirection::Column,
                        align_items: AlignItems::Center,
                        ..default()
                    },
                    background_color: CRIMSON.into(),
                    ..default()
                })
                .with_children(|parent| {
                    parent
                        .spawn(NodeBundle {
                            style: Style {
                                align_items: AlignItems::Center,
                                ..default()
                            },
                            background_color: CRIMSON.into(),
                            ..default()
                        })
                        .with_children(|parent| {
                            parent.spawn(TextBundle::from_section(
                                "Volume",
                                button_text_style.clone(),
                            ));
                            for volume_setting in [0, 1, 2, 3, 4, 5, 6, 7, 8, 9] {
                                let color: Color;
                                if *volume >= AudioVolume(volume_setting) {
                                    color = Color::srgb(0.04, 2.55, 0.0);
                                } else {
                                    color = NORMAL_BUTTON;
                                }

                                let mut entity = parent.spawn((
                                    ButtonBundle {
                                        style: Style {
                                            width: Val::Px(30.0),
                                            height: Val::Px(65.0),
                                            ..button_style.clone()
                                        },
                                        background_color: color.into(),
                                        ..default()
                                    },
                                    AudioVolume(volume_setting),
                                ));
                                if *volume == AudioVolume(volume_setting) {
                                    entity.insert(SelectedOption);
                                }
                            }
                        });
                    parent
                        .spawn((
                            ButtonBundle {
                                style: button_style,
                                background_color: NORMAL_BUTTON.into(),
                                ..default()
                            },
                            MenuButtonAction::BackToSettings,
                        ))
                        .with_children(|parent| {
                            parent.spawn(TextBundle::from_section("Back", button_text_style));
                        });
                });
        });
}

pub fn menu_action(
    mut commands: Commands,
    interaction_query: Query<
        (&Interaction, &MenuButtonAction),
        (Changed<Interaction>, With<Button>),
    >,
    mut client: ResMut<RenetClient>,
    mut transport: ResMut<NetcodeClientTransport>,
    mut menu_state: ResMut<NextState<MenuState>>,
    mut game_state: ResMut<NextState<GameState>>,
    mut paused_state: ResMut<NextState<PausedState>>, // Pour changer l'état
    etat_paused_state: Res<State<PausedState>>,       // Pour lire l'état actuel
    username: Res<UsernameInput>,
    mut players: ResMut<Players>,
    mut windows: Query<&mut Window>,
    mut dropdown_query: Query<&mut Style, With<ServerDropdownMenu>>, // Query pour cibler le menu déroulant
    mut server_option_query: Query<
        (&Interaction, &ServerOption),
        (Changed<Interaction>, With<Button>),
    >, // Query pour les interactions avec les options du serveur
    mut selected_server_text: Query<&mut Text, With<ServerInputText>>, // Query pour mettre à jour le texte du serveur sélectionné
    mut exit: EventWriter<AppExit>,
) {
    for (interaction, menu_button_action) in &interaction_query {
        if *interaction == Interaction::Pressed {
            match menu_button_action {
                MenuButtonAction::Quit => {
                    exit_system(&mut client, &mut transport, &mut players, &mut exit);
                }
                MenuButtonAction::Play => {
                    match etat_paused_state.get() {
                        PausedState::On => {
                            // update des états
                            menu_state.set(MenuState::Disabled);
                            game_state.set(GameState::Game);

                            // desactive le curseur
                            let mut window = windows.single_mut();
                            window.cursor.visible = false;
                        }
                        PausedState::Off => {
                            menu_state.set(MenuState::LaunchGame);
                        }
                    }
                }
                MenuButtonAction::Settings => menu_state.set(MenuState::Settings),
                MenuButtonAction::SettingsDisplay => {
                    menu_state.set(MenuState::SettingsDisplay);
                }
                MenuButtonAction::SettingsSound => {
                    menu_state.set(MenuState::SettingsSound);
                }
                MenuButtonAction::BackToMainMenu => menu_state.set(MenuState::Main),
                MenuButtonAction::BackToSettings => {
                    menu_state.set(MenuState::Settings);
                }
                MenuButtonAction::UsernameAndServer => {
                    if username.0.len() == 0 {
                        return;
                    } else {
                        println!("Game starting...");
                        let ipaddress_public = match get_client_ip_public() {
                            Ok(ip) => ip, // ip du client
                            Err(e) => {
                                eprintln!("Erreur lors de la récupération de l'IP : {}", e);
                                IpAddr::V4(Ipv4Addr::LOCALHOST) // Valeur par défaut en cas d'erreur
                            }
                        };

                        // TODO: Update des info initial du player ici
                        let current_time = SystemTime::now()
                            .duration_since(SystemTime::UNIX_EPOCH)
                            .unwrap();
                        let client_id = current_time.as_millis() as u64;
                        players.player_public.info_player.id = client_id;
                        players.player_public.info_player.username = username.0.clone();
                        players.player_public.machine_information.ipaddress_public =
                            ipaddress_public.to_string();
                        if selected_server_text.single().sections[0].value.clone() == "Local" {
                            players.player_public.machine_information.connectedto =
                                "127.0.0.1".to_string();
                        } else {
                            //Adresse vps
                            players.player_public.machine_information.connectedto =
                                "90.25.70.155".to_string();
                            //Adresse ip test sur tel
                            // players.player_public.machine_information.connectedto =
                            //     "172.20.10.2".to_string();
                        }

                    // mise à jour du transport avec l'ip du server selectionné
                    update_transport_layer(&mut commands, &mut players);

                        // update des états
                        menu_state.set(MenuState::Disabled);
                        game_state.set(GameState::Game);
                        paused_state.set(PausedState::On);

                        // desactive le curseur
                        let mut window = windows.single_mut();
                        window.cursor.visible = false;
                        // println!("players : {:?}", players);
                    }
                }
                MenuButtonAction::ServersDisplay => {
                    // Toggle de la liste déroulante
                    for mut style in &mut dropdown_query {
                        if style.display == Display::None {
                            style.display = Display::Flex; // Affiche la liste déroulante
                        } else {
                            style.display = Display::None; // Cache la liste déroulante
                        }
                    }

                    // Appel de la fonction handle_server_dropdown pour gérer la sélection de serveur
                    handle_server_dropdown(
                        &mut server_option_query,
                        &mut selected_server_text,
                        &mut dropdown_query,
                    );
                }
            }
        }
    }
}

/// Cette fonction supprime toutes les entités ayant un composant spécifique,
/// ainsi que toutes leurs entités enfants, de manière récursive.
///
/// # Paramètres
/// - `to_despawn`: Une requête (`Query`) qui récupère toutes les entités ayant le composant `T`.
/// - `commands`: Un objet `Commands` utilisé pour émettre des commandes pour ajouter ou retirer des entités.
///
/// # Type Paramétré
/// - `T`: Un type qui implémente le trait `Component`, ce qui permet de spécifier le composant dont les entités seront supprimées.
pub fn despawn_screen<T: Component>(to_despawn: Query<Entity, With<T>>, mut commands: Commands) {
    // Parcourt toutes les entités qui possèdent le composant `T`
    for entity in &to_despawn {
        // Supprime l'entité ainsi que toutes ses entités enfants de manière récursive
        commands.entity(entity).despawn_recursive();
    }
}
