use super::{struct_manager::*, struct_menu::AudioVolume};
use crate::interaction_server::config::*;
use bevy::{
    audio::{Volume, *},
    prelude::*,
};

// ╭───────────────────────────────────────────────╮
// │                                               │
// │                   STRUCTURE                   │
// │                                               │
// ╰───────────────────────────────────────────────╯
#[derive(Debug, Resource)]
pub struct Audio {
    pub ambiance: Handle<AudioSource>,
    pub shoot: Handle<AudioSource>,
}

#[derive(Debug, Component)]
pub struct SoundAmbiant;

// ╭───────────────────────────────────────────────╮
// │                                               │
// │                    FONCTION                   │
// │                                               │
// ╰───────────────────────────────────────────────╯

pub fn play_loop_audio(mut commands: Commands, resources: Res<MyResources>) {
    commands
        .spawn(AudioBundle {
            source: resources.audio.ambiance.clone(),
            settings: PlaybackSettings {
                mode: PlaybackMode::Loop,
                ..Default::default()
            },
            ..Default::default()
        })
        .insert(SoundAmbiant);
}

pub fn volume_ambiant_sound(
    music_controller: Query<&AudioSink, With<SoundAmbiant>>,
    volume: Res<AudioVolume>,
    mut button_query: Query<(&mut BackgroundColor, &AudioVolume)>,
) {
    // Ajuster le volume du AudioSink
    if let Ok(sink) = music_controller.get_single() {
        sink.set_volume(volume.0 as f32 / 10.0);
    }

    // Mettre à jour la couleur des boutons en fonction du volume
    for (mut background_color, button_volume) in button_query.iter_mut() {
        let color: Color;
        if volume.0 >= button_volume.0 {
            color = Color::srgb(0.04, 2.55, 0.0); // Couleur verte
        } else {
            color = NORMAL_BUTTON; // Couleur par défaut
        }

        // Mettre à jour la couleur
        *background_color = color.into();
    }
}

pub fn play_once_audio(
    mut commands: Commands,
    audio: Handle<AudioSource>,
    volume: Res<AudioVolume>,
) {
    commands.spawn(AudioBundle {
        source: audio,
        // settings: PlaybackSettings::ONCE,
        settings: PlaybackSettings {
            mode: PlaybackMode::Once,
            volume: Volume::new(volume.0 as f32 / 10.0),
            ..Default::default()
        },

        ..default()
    });
}
