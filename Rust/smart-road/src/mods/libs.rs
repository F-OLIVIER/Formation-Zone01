// ------------------------------------------------------
// ---------------------- Throttle ----------------------
// ------------------------------------------------------
use std::time::{Duration, Instant};

pub struct Throttle {
    pub last_call: Instant,
    pub delay_ms: Duration,
}

impl Throttle {
    pub fn new(delay_ms: u64) -> Self {
        Throttle {
            last_call: Instant::now(),
            delay_ms: Duration::from_millis(delay_ms),
        }
    }

    // Vérifie si le temps spécifié n'est pas encore écoulé
    pub fn call_throttle(&mut self) -> bool {
        let now = Instant::now();
        if now.duration_since(self.last_call) >= self.delay_ms {
            self.last_call = now;
            true
        } else {
            false
        }
    }
}

// ------------------------------------------------------
// --------------------- Statistique --------------------
// ------------------------------------------------------
use chrono::TimeDelta;
use sdl2::{pixels::Color, rect::Rect, render::{TextureQuery, WindowCanvas}};

use super::Vehicule;

#[derive(Debug, Clone, PartialEq)]
pub struct Stat {
    pub velocity_max: u32,
    pub velocity_min: u32,
    pub max_time: TimeDelta,
    pub min_time: TimeDelta,
    pub close_calls: u32,
}

impl Stat {
    pub fn new() -> Stat {
        Stat {
            velocity_max: 0,
            velocity_min: 100,
            max_time: TimeDelta::new(0, 0).unwrap(),
            min_time: TimeDelta::new(60, 0).unwrap(),
            close_calls: 0, 
        }
    }

    pub fn update_stat(&mut self, cars: Vec<Vehicule>, list_to_delete: Vec<usize>, list_collision: &mut Vec<(String, String)>) {
        //Looping through all the cars to take corresponding stat
        for (i, current_car) in cars.iter().enumerate() {

            // velocity
            if self.velocity_min > current_car.velocity {
                self.velocity_min = current_car.velocity;
            }
            if self.velocity_max < current_car.velocity {
                self.velocity_max = current_car.velocity;
            } 

            // time
            if list_to_delete.contains(&i) {
                let duration = current_car.time_depop.signed_duration_since(current_car.time_pop);
                if self.min_time > duration {
                    self.min_time = duration;
                }
                if self.max_time < duration {
                    self.max_time = duration;
                } 
            }
            
            // nb_close_calls, considered as a collision

            let rect_current_car  = Rect::new((current_car.position.0).try_into().unwrap(), (current_car.position.1).try_into().unwrap(), 20, 20);
            for test_car in &cars {
                if current_car != test_car {
                    let rect_test_car  = Rect::new((test_car.position.0).try_into().unwrap(), (test_car.position.1).try_into().unwrap(), 20, 20);
    
                    if rect_current_car.has_intersection(rect_test_car) {
                        //check if car has already been tested with the other one
                        let mut is_in_list = false;
                        for colission in list_collision.iter() {
                            if current_car.name == colission.0 && test_car.name == colission.1 ||  
                                current_car.name == colission.1 && test_car.name == colission.0 {
                                is_in_list = true;
                                break;
                            }
                        }
                        if !is_in_list {
                            list_collision.push((current_car.name.clone(), test_car.name.clone()));
                            self.close_calls += 1;
                        }
                    }
                }
            }
        }
    }

    pub fn display_stat(&self, canvas: &mut WindowCanvas, button_rect: Rect, nb_car: usize, width_windows: u32) {
        // Initialisation de SDL2_ttf et chargement de la police
        let ttf_context = sdl2::ttf::init().unwrap();
        let texture_creator = canvas.texture_creator();
        let font = ttf_context.load_font("src/assets/Matemasie-Regular.ttf", 32).unwrap();
    
        let max_time_std = self.max_time.to_std().unwrap();
        let min_time_std = self.min_time.to_std().unwrap();
        
        let max_time_seconds = max_time_std.as_secs_f64();
        let min_time_seconds = min_time_std.as_secs_f64();
        // Créer le texte des statistiques
        let stats_text = format!(
            "STATISTIQUES\n \nMax number of vehicules: {}\nMax velocity: {}\nMin velocity: {}\nMax time: {:.2}s\nMin time: {:.2}s\nClose calls: {}",
            nb_car, self.velocity_max, self.velocity_min, max_time_seconds, min_time_seconds, self.close_calls
        );
        // println!("stats_text : \n{}", stats_text);
    
        // Créer une surface pour le texte
        let surface = font.render(&stats_text).solid(Color::RGB(255, 255, 255)).unwrap();
        let texture = texture_creator.create_texture_from_surface(&surface).unwrap();
        // Obtenir les dimensions de la texture
        let TextureQuery { width, height, .. } = texture.query();
    
        // Définir la position et la taille de l'encart
        let text_rect = Rect::new((width_windows as i32 - 600) / 2, 30, width, height);
        let background_rect = Rect::new(
            text_rect.x(),
            text_rect.y(),
            600,
            550,
        );
    
        // Dessiner l'encart
        canvas.set_draw_color(Color::RGB(0, 0, 0)); // Couleur de l'encart
        canvas.fill_rect(background_rect).unwrap();
        
        // texte
        let lines: Vec<&str> = stats_text.split('\n').collect();
        let mut y_offset = 50; // Position Y initiale
        for line in lines {
            let surface = font.render(line).blended(Color::RGB(255, 255, 255)).unwrap();
            let texture = texture_creator.create_texture_from_surface(&surface).unwrap();
            let TextureQuery { width, height, .. } = texture.query();
            
            let text_rect = Rect::new((width_windows as i32 - 600) / 2 + 50, y_offset, width, height);
            y_offset += height as i32 + 5; // espace entre les lignes
            canvas.copy(&texture, None, Some(text_rect)).unwrap();
        }

        // bouton
        canvas.set_draw_color(Color::RGB(255, 255, 255));
        canvas.fill_rect(button_rect).unwrap();

        // Créer une surface pour le texte
        let button_text = "EXIT";
        let surface_button = font.render(&button_text).solid(Color::RGB(0, 0, 0)).unwrap();
        let texture_button = texture_creator.create_texture_from_surface(&surface_button).unwrap();
        
        let text_rect = Rect::new(
            button_rect.x + (button_rect.width() as i32 - 50) / 2,
            button_rect.y + (button_rect.height() as i32 - 25) / 2,
            50,
            25,
        );
        canvas.copy(&texture_button, None, text_rect).unwrap();

        canvas.present();
    }
}
