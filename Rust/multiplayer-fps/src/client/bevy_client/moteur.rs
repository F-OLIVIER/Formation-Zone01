// ██████  ███████ ██    ██ ██    ██     ███    ███  ██████  ████████ ███████ ██    ██ ██████
// ██   ██ ██      ██    ██  ██  ██      ████  ████ ██    ██    ██    ██      ██    ██ ██   ██
// ██████  █████   ██    ██   ████       ██ ████ ██ ██    ██    ██    █████   ██    ██ ██████
// ██   ██ ██       ██  ██     ██        ██  ██  ██ ██    ██    ██    ██      ██    ██ ██   ██
// ██████  ███████   ████      ██        ██      ██  ██████     ██    ███████  ██████  ██   ██

use super::{
    audio::*, cursor::*, environment::*, fps::*, graphique_quality::*, map::*, menu_general::*,
    menu_register::*, player_manager::*, scene::*, setup_player::*, struct_manager::*,
    struct_menu::*,
};
use crate::bevy_client::setup_ennemy::*;
use crate::interaction_server::config::*;
use crate::interaction_server::udpsocket::*;
#[allow(unused_imports)]
use bevy::diagnostic::{FrameTimeDiagnosticsPlugin, LogDiagnosticsPlugin};
use bevy::prelude::App;
use bevy::window::{PresentMode, WindowTheme};
use bevy_rapier3d::prelude::*;
use bevy_renet::{client_connected, RenetClientPlugin, RenetServerPlugin};
use renet::transport::NetcodeClientTransport;
use renet::RenetClient;

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

pub fn init_bevy() {
    App::new()
        // Utiliser le pool de threads personnalisé
        // Ressources
        .insert_resource(ClearColor(Color::linear_rgb(0.1, 0.3, 0.96)))
        .insert_resource(DisplayQuality::Fullscreen)
        .insert_resource(AudioVolume(3))
        .insert_resource(Players::new())
        .insert_resource(UsernameInput::default())
        .insert_resource(SendRateTimer {
            send_timer: Timer::from_seconds(1.0 / 30.0, TimerMode::Repeating), // Timer pour limiter les envois au serveur (30 envois par seconde)
            scoreboard_timer: Timer::from_seconds(1.0, TimerMode::Repeating), // Timer pour l'actualisation du scoreboard (5 envois par seconde)
            minimap_timer: Timer::from_seconds(1.0 / 5.0, TimerMode::Repeating), // Timer pour l'actualisation de la minimap (5 envois par seconde)
        })
        .insert_resource(LogsResource::default())
        .insert_resource(ProjectileNumber { id: 0 })
        // Plugins
        .add_plugins((
            DefaultPlugins.set(WindowPlugin {
                primary_window: Some(Window {
                    title: "Multiplayer FPS de la mort qui tue tous !!!".into(),
                    present_mode: PresentMode::AutoVsync,
                    fit_canvas_to_parent: true,
                    prevent_default_event_handling: false,
                    window_theme: Some(WindowTheme::Dark),
                    enabled_buttons: bevy::window::EnabledButtons {
                        maximize: false,
                        ..Default::default()
                    },
                    ..default()
                }),
                ..default()
            }),
            RapierPhysicsPlugin::<NoUserData>::default(), // moteur de physique Rapier
            RenetClientPlugin,
            RenetServerPlugin,
            // LogDiagnosticsPlugin::default(), // affichage des diagnostics collectés dans la console
            FrameTimeDiagnosticsPlugin, // diagnostics spécifiques au temps de frame (fps et durée des frames)
            FpsControllerPlugin,        // Controleur pour les deplacements
            menu_plugin,                // Eléments systéme
        ))
        // Run app
        .run();
}

fn menu_plugin(app: &mut App) {
    println!("App running !");
    // Application sur le menu
    app
        // ---------------------------------------------------------
        // ------------------- Etat du menu/jeu --------------------
        // ---------------------------------------------------------
        .init_state::<MenuState>() // Etat menu
        .init_state::<GameState>() // Etat jeu
        .init_state::<PausedState>() // Etat menu : format menu pause (on), format menu initial (off)
        .configure_sets(Update, Connected.run_if(client_connected))
        // ---------------------------------------------------------
        // ---------------------- Partie menu ----------------------
        // ---------------------------------------------------------
        .add_systems(OnEnter(GameState::Menu), menu_setup)
        .add_systems(
            OnEnter(MenuState::Main),
            main_menu_setup.after(generate_ressources),
        )
        .add_systems(OnExit(MenuState::Main), despawn_screen::<OnMainMenuScreen>)
        // --------------------------------------------------------
        //-----------------------Menu settings---------------------
        //---------------------------------------------------------
        .add_systems(OnEnter(MenuState::Settings), settings_menu_setup)
        .add_systems(
            OnExit(MenuState::Settings),
            despawn_screen::<OnSettingsMenuScreen>,
        )
        // --------------------------------------------------------
        //------------------Menu Qualité graphique-----------------
        //---------------------------------------------------------
        .add_systems(
            OnEnter(MenuState::SettingsDisplay),
            display_settings_menu_setup,
        )
        .add_systems(
            OnExit(MenuState::SettingsDisplay),
            despawn_screen::<OnDisplaySettingsMenuScreen>,
        )
        // --------------------------------------------------------
        //------------------------Menu Audio-----------------------
        //---------------------------------------------------------
        .add_systems(OnEnter(MenuState::SettingsSound), sound_settings_menu_setup)
        .add_systems(
            OnExit(MenuState::SettingsSound),
            despawn_screen::<OnSoundSettingsMenuScreen>,
        )
        // --------------------------------------------------------
        //---------------------Menu Registration-------------------
        //---------------------------------------------------------
        .add_systems(OnEnter(MenuState::LaunchGame), register_menu)
        .add_systems(
            OnExit(MenuState::LaunchGame),
            despawn_screen::<OnRegisterScreen>,
        )
        // --------------------------------------------------------
        // -----------------------Update menu----------------------
        // --------------------------------------------------------
        .add_systems(
            Update,
            (
                (setting_button::<DisplayQuality>.run_if(in_state(MenuState::SettingsDisplay)),), // Menu qualité graphique
                setting_button::<AudioVolume>.run_if(in_state(MenuState::SettingsSound)), // Menu volume audio
                (menu_action, button_system).run_if(in_state(GameState::Menu)), // gestion des action des bouttons
                (text_input_system).run_if(in_state(MenuState::LaunchGame)),    // menu inputtext
                // --------------------------------------------------------
                // -------------- Update commune (menu & game) ------------
                // --------------------------------------------------------
                volume_ambiant_sound,
                quality_graphique,
            ),
        )
        // ---------------------------------------------------------
        // ------------------- Partie jeu (Game) -------------------
        // ---------------------------------------------------------
        .add_systems(OnEnter(GameState::Game), despawn_screen::<OnMainMenuScreen>)
        .add_systems(OnExit(GameState::Game), despawn_screen::<OnGameScreen>)
        .add_systems(
            Startup,
            (
                setup_renet_client,    // insére la ressource client pour la communication server
                setup_transport_layer, // insére le transport pour la communication server
                generate_ressources,   // Génère les ressources à utiliser
                (
                    cursor,             // Gestion du curseur
                    play_loop_audio,    // Audio en boucle
                    setup_player,       // Caméra du joueur
                    display_map,        // Génération et affichage de la carte
                    spawn_lights,       // Lumière de la scène
                    display_fps,        // Affichage des FPS
                    display_scoreboard, // Affichage du tableau des scores
                    display_munition,   // Affichage des munitions
                    // display_timer,      // Affichage du chronomètre
                    explain_text, // Texte explicatif des touches clavier
                    // display_lifebar,    // Affichage de la barre de vie
                    log_text, // Zone de texte log
                )
                    .after(generate_ressources), // Tous ces systèmes s'exécutent après `generate_ressources`
            ),
        )
        // ---------------------------------------------------------
        // ------------------- Update jeu (Game) -------------------
        // ---------------------------------------------------------
        .add_systems(
            Update,
            (
                client_update_system, // Update du transport pour envoi server
                update_fps,           // Mise à jour des FPS en temps réel
                update_munition,      // Mise à jour des munition en temps réel
                // update_timer,               // Mise à jour des munition en temps réel
                update_scoreboard,          // Mise à jour du scoreboard si "tab"
                scene_colliders,            // Gestion des colissions
                shoot,                      // Gestion des tir d'arme
                move_projectiles,           // Déplacement des projectiles de l'arme du player
                remove_expired_projectiles, // Retire les projectiles expiré
                light_clignotement,         // Clignotement rouge de la lumiére (mort du player)
                respawn,                    // ???
                play_animation,
                keyboard_animation_control,
                update_logs, // Mise à jour de la zone de texte log
                update_minimap_marker,
            )
                .run_if(in_state(GameState::Game)),
        )
        // ---------------------------------------------------------
        // ----------------- Communication SERVER ------------------
        // ---------------------------------------------------------
        .add_systems(
            Update,
            (
                handle_sending,          // Envoi des informations au serveur
                handle_incoming_message, // Réception des informations du serveur
                handle_incoming_event,   // Réception des events du serveur
            )
                .run_if(in_state(GameState::Game)),
        );
}

pub fn exit_system(
    client: &mut ResMut<RenetClient>,
    transport: &mut ResMut<NetcodeClientTransport>,
    players: &mut ResMut<Players>,
    exit: &mut EventWriter<AppExit>,
) {
    println!("You quit the game");

    send_event(
        client,
        EventMessage {
            types: "player_leave".to_string(),
            id: players.player_public.info_player.id,
            username: players.player_public.info_player.username.clone(),
            information: format!(
                "{} leave the game before disconnect_client",
                players.player_public.info_player.username
            ),
        },
    );

    disconnect_client(client, transport);
    exit.send(AppExit::Success);
}
