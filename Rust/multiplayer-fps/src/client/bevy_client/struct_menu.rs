// ███████ ███    ██ ██    ██ ███    ██         ██     ███████ ████████ ██████  ██    ██  ██████ ████████
// ██      ████   ██ ██    ██ ████   ██        ██      ██         ██    ██   ██ ██    ██ ██         ██
// █████   ██ ██  ██ ██    ██ ██ ██  ██       ██       ███████    ██    ██████  ██    ██ ██         ██
// ██      ██  ██ ██ ██    ██ ██  ██ ██      ██             ██    ██    ██   ██ ██    ██ ██         ██
// ███████ ██   ████  ██████  ██   ████     ██         ███████    ██    ██   ██  ██████   ██████    ██
//
// ███╗   ███╗███████╗███╗   ██╗██╗   ██╗
// ████╗ ████║██╔════╝████╗  ██║██║   ██║
// ██╔████╔██║█████╗  ██╔██╗ ██║██║   ██║
// ██║╚██╔╝██║██╔══╝  ██║╚██╗██║██║   ██║
// ██║ ╚═╝ ██║███████╗██║ ╚████║╚██████╔╝
// ╚═╝     ╚═╝╚══════╝╚═╝  ╚═══╝ ╚═════╝
//
use bevy::prelude::*;
//
// ╭───────────────────────────────────────────────╮
// │                                               │
// │                     ENUM                      │
// │                                               │
// ╰───────────────────────────────────────────────╯
#[derive(Clone, Copy, Default, Eq, PartialEq, Debug, Hash, States)]
pub enum GameState {
    #[default]
    Menu,
    Game,
}

#[derive(Clone, Copy, Default, Eq, PartialEq, Debug, Hash, States)]
pub enum PausedState {
    On,
    #[default]
    Off,
}

#[derive(Resource, Debug, Component, PartialEq, Eq, Clone, Copy)]
pub enum DisplayQuality {
    Windowed,
    BorderlessFullscreen,
    Fullscreen,
}

#[derive(Default, Debug, Clone, Copy, PartialEq, Eq, Hash, States)]
pub enum MenuState {
    Main,
    Settings,
    SettingsDisplay,
    SettingsSound,
    LaunchGame,
    #[default]
    Disabled,
}

// All actions that can be triggered from a button click
#[derive(Component)]
pub enum MenuButtonAction {
    Play,
    Settings,
    SettingsDisplay,
    SettingsSound,
    BackToMainMenu,
    BackToSettings,
    Quit,
    UsernameAndServer, // Menu Registration
    ServersDisplay,    //Toggle liste déroulante du choix du serveur
}

// ╭───────────────────────────────────────────────╮
// │                                               │
// │                   STRUCTURE                   │
// │                                               │
// ╰───────────────────────────────────────────────╯
#[derive(Resource, Debug, Component, PartialEq, PartialOrd, Eq, Clone, Copy)]
pub struct AudioVolume(pub u32);

#[derive(Component)]
pub struct OnMainMenuScreen;

#[derive(Component)]
pub struct OnGameScreen;

// Tag component used to tag entities added on the settings menu screen
#[derive(Component)]
pub struct OnSettingsMenuScreen;

// Tag component used to tag entities added on the display settings menu screen
#[derive(Component)]
pub struct OnDisplaySettingsMenuScreen;

// Tag component used to tag entities added on the sound settings menu screen
#[derive(Component)]
pub struct OnSoundSettingsMenuScreen;

#[derive(Component)]
pub struct OnRegisterScreen;

// Récupère le champ d'entrée du username
#[derive(Component)]
pub struct UsernameInputField;

//Composant pour afficher le texte
#[derive(Component)]
pub struct UsernameInputText;

// Composant pour stocker l'input utilisateur dans le champ de texte
#[derive(Resource)]
pub struct UsernameInput(pub String);
impl Default for UsernameInput {
    fn default() -> Self {
        UsernameInput(String::new()) // Retourne une instance avec une chaîne vide
    }
}
//Composant pour afficher le texte
#[derive(Component)]
pub struct ServerInputText;

//Récupère le champ d'entrée du serveur choisi
#[derive(Component)]
pub struct ServerInputField;

// Composant pour identifier le menu déroulant
#[derive(Component)]
pub struct ServerDropdownMenu;

//Composant désignant une option de la liste (dans une boucle)
#[derive(Component)]
pub struct ServerOption(pub String);

// #[derive(Resource, Debug)]
// pub struct NetworkClient {
//     pub client: Option<RenetClient>,
// }
