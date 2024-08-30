pub mod libs;
pub mod vehicle;
pub mod intersection;
use sdl2::image::LoadTexture;
pub use vehicle::*;
pub use libs::*;
pub use intersection::*;

extern crate sdl2;
use sdl2::pixels::Color;
use sdl2::event::Event;
use sdl2::video::Window;
use sdl2::keyboard::Keycode;
use sdl2::rect::{Point, Rect};
use sdl2::render::{Canvas, WindowCanvas};
use std::time::Duration;

pub async fn launch_sdl2() {
    let sdl_context = sdl2::init().unwrap();
    let video_subsystem = sdl_context.video().unwrap();

    let window = video_subsystem.window("Smart road", 800, 600).position_centered().build().unwrap();

    let mut canvas = window.into_canvas().build().unwrap();
    canvas.set_draw_color(Color::RGB(0, 255, 255));
    canvas.clear();
    canvas.present();

    let mut event_pump = sdl_context.event_pump().unwrap();
    // initialisationd des variables
    let (width, height) = canvas.output_size().unwrap();
    let mut throttle = Throttle::new(350);
    let mut cars: Vec<Vehicule> = Vec::new();
    let mut stat = Stat::new();
    let mut inrunning = true;
    let button_rect = Rect::new(350, 500, 100, 50);
    let mut list_collision: Vec<(String, String)> = vec![];
    let mut list_cars_passed: Vec<String> = vec![];
    let mut num_name = 0;
    let mut cars_in_intersection = Comptage { left_cars: 0, straight_cars: 0, right_cars: 0 };

    'running: loop {
        // Couleur fond de map
        canvas.set_draw_color(Color::RGB(45, 194, 46));
        canvas.clear();

        for event in event_pump.poll_iter() {
            let (width, height) = canvas.output_size().unwrap();
            match event {
                | Event::Quit { .. }
                // Esc Escape: Affiche les statistiques
                | Event::KeyDown { keycode: Some(Keycode::Escape), .. } => {
                    inrunning = false;
                    map(&mut canvas);
                    stat.display_stat(&mut canvas, button_rect, list_cars_passed.len(), width);
                }
                // Click button: ends the simulation.
                Event::MouseButtonDown { x, y, .. } => {
                    if !inrunning && button_rect.contains_point((x, y)) {
                        break 'running;
                    }
                }
                // S: Affiche les statistiques
                Event::KeyDown { keycode: Some(Keycode::S), .. } => {
                   if inrunning {
                    inrunning = false;
                    map(&mut canvas);
                    stat.display_stat(&mut canvas, button_rect, list_cars_passed.len(), width);
                    //to continue playing the animation
                   } else {
                        inrunning = true;
                    }
                }
                // ↑ Up: moves towards the intersection from the south.
                Event::KeyDown { keycode: Some(Keycode::Up), .. } => {
                    if throttle.call_throttle() && inrunning && cars.len() < 10 {
                        Vehicule::spawn_cars(&mut cars, Direction::North, width, height, &mut num_name, &cars_in_intersection);
                    }
                }
                // ↓ Down: moves towards the intersection from the north.
                Event::KeyDown { keycode: Some(Keycode::Down), .. } => {
                    if throttle.call_throttle() && inrunning && cars.len() < 10 {
                        Vehicule::spawn_cars(&mut cars, Direction::South, width, height, &mut num_name, &cars_in_intersection);
                    }
                }
                // → Right: moves towards the intersection from the west.
                Event::KeyDown { keycode: Some(Keycode::Right), .. } => {
                    if throttle.call_throttle() && inrunning && cars.len() < 10 {
                        Vehicule::spawn_cars(&mut cars, Direction::East, width, height, &mut num_name, &cars_in_intersection);
                    }
                }
                // ← Left: moves towards the intersection from the east.
                Event::KeyDown { keycode: Some(Keycode::Left), .. } => {
                    if throttle.call_throttle() && inrunning && cars.len() < 10 {
                        Vehicule::spawn_cars(&mut cars, Direction::West, width, height, &mut num_name, &cars_in_intersection);
                        
                    }
                }
                // r: moves towards the intersection from a random direction.
                Event::KeyDown { keycode: Some(Keycode::R), .. } => {
                    if throttle.call_throttle() && inrunning && cars.len() < 10 {
                        Vehicule::spawn_cars(&mut cars, Direction::random_direction(), width, height, &mut num_name, &cars_in_intersection);
                    }
                }

                _ => {}
            }
        }

        if inrunning {
            let mut list_to_delete: Vec<usize> = vec![];

            for (i, car) in cars.iter_mut().enumerate() {
                
                if !car.move_cars(width, height) {
                    list_to_delete.push(i);
                }
                //check if direction is changed by its route
                if !car.turning {
                    let tmp_dir = car.direction.clone();
                    car.update_route(width, height);
                    if tmp_dir != car.direction {
                        car.turning = true;
                    }
                }
            }

            // Mise à jour des statistiques
            stat.update_stat(cars.clone(), list_to_delete.clone(), &mut list_collision);
            //Deleting cars that have 
            for num_car in list_to_delete.iter().rev() {
                cars.remove(*num_car);
            }
            
            prediction_to_velocity(&mut cars, &mut list_cars_passed, &mut cars_in_intersection, width, height);

            // Dessine la map de base
            map(&mut canvas);
            render(&mut canvas, &mut cars);
    
            canvas.present();
        }
        ::std::thread::sleep(Duration::new(0, 1_000_000_000u32 / 60));
    }
}

pub fn map(canvas: &mut WindowCanvas) {
    let (width, height) = canvas.output_size().unwrap();

    let (v_lane1, v_lane2) = ((width as i32) / 2 - 90, (width as i32) / 2 + 90);

    let (h_lane1, h_lane2) = ((height as i32) / 2 - 90, (height as i32) / 2 + 90);
    // valeur de depart pur (800 x 400) : h_lane1 : 280, h_lane2 : 320, v_lane1 : 380, v_lane2 : 420

    // Dessine les rectangles qui representent la route
    canvas.set_draw_color(Color::GRAY);
    // rectangle horizontal
    let _ = canvas.fill_rect(Rect::new(0, h_lane1, width, 180));
    // rectangle vertical
    let _ = canvas.fill_rect(Rect::new(v_lane1, 0, 180, height));

    // Dessine les lignes de bordure de la route
    canvas.set_draw_color(Color::WHITE);
    // --------------------- Ligne horizontale ---------------------
    // Partie haute, du centre (milieu de l'écran) vers droite.
    canvas.draw_line(Point::new(width as i32, h_lane1), Point::new(v_lane2, h_lane1)).unwrap();
    draw_dotted_line(canvas, Point::new(width as i32, h_lane1 + 30), Point::new(v_lane2, h_lane1 + 30), Color::WHITE, 10);
    draw_dotted_line(canvas, Point::new(width as i32, h_lane1 + 60), Point::new(v_lane2, h_lane1 + 60), Color::WHITE, 10);

    // Partie haute, de gauche vers le centre (milieu de l'écran).
    canvas.draw_line(Point::new(0, h_lane1), Point::new(v_lane1, h_lane1)).unwrap();
    draw_dotted_line(canvas, Point::new(0, h_lane1 + 30), Point::new(v_lane1, h_lane1 + 30), Color::WHITE, 10);
    draw_dotted_line(canvas, Point::new(0, h_lane1 + 60), Point::new(v_lane1, h_lane1 + 60), Color::WHITE, 10);

    // Partie basse, du centre (milieu de l'écran) vers droite.
    canvas.draw_line(Point::new(width as i32, h_lane2), Point::new(v_lane2, h_lane2)).unwrap();
    draw_dotted_line(canvas, Point::new(width as i32, h_lane2 - 30), Point::new(v_lane2, h_lane2 - 30), Color::WHITE, 10);
    draw_dotted_line(canvas, Point::new(width as i32, h_lane2 - 60), Point::new(v_lane2, h_lane2 - 60), Color::WHITE, 10);

    // Partie basse, de gauche vers le centre (milieu de l'écran).
    canvas.draw_line(Point::new(0, h_lane2), Point::new(v_lane1, h_lane2)).unwrap();
    draw_dotted_line(canvas, Point::new(0, h_lane2 - 30), Point::new(v_lane1, h_lane2 - 30), Color::WHITE, 10);
    draw_dotted_line(canvas, Point::new(0, h_lane2 - 60), Point::new(v_lane1, h_lane2 - 60), Color::WHITE, 10);
    
    // --------------------- Ligne verticale ---------------------
    // Partie gauche (milieu de l'écran vers haut).
    canvas.draw_line(Point::new(v_lane1, height as i32), Point::new(v_lane1, h_lane2)).unwrap();
    draw_dotted_line(canvas, Point::new(v_lane1 + 30, height as i32), Point::new(v_lane1 + 30, h_lane2), Color::WHITE, 10);
    draw_dotted_line(canvas, Point::new(v_lane1 + 60, height as i32), Point::new(v_lane1 + 60, h_lane2), Color::WHITE, 10);

    // Partie gauche (haut de l'écran vers milieu).
    canvas.draw_line(Point::new(v_lane1, 0), Point::new(v_lane1, h_lane1)).unwrap();
    draw_dotted_line(canvas, Point::new(v_lane1 + 30, 0), Point::new(v_lane1 + 30, h_lane1), Color::WHITE, 10);
    draw_dotted_line(canvas, Point::new(v_lane1 + 60, 0), Point::new(v_lane1 + 60, h_lane1), Color::WHITE, 10);

    // Partie droite (haut de l'écran vers milieu).
    canvas.draw_line(Point::new(v_lane2, h_lane1), Point::new(v_lane2, 0)).unwrap();
    draw_dotted_line(canvas, Point::new(v_lane2 - 30, h_lane1), Point::new(v_lane2 - 30, 0), Color::WHITE, 10);
    draw_dotted_line(canvas, Point::new(v_lane2 - 60, h_lane1), Point::new(v_lane2 - 60, 0), Color::WHITE, 10);

    // Partie droite (milieu de l'écran vers haut).
    canvas.draw_line(Point::new(v_lane2, height as i32), Point::new(v_lane2, h_lane2)).unwrap();
    draw_dotted_line(canvas, Point::new(v_lane2 - 30, height as i32), Point::new(v_lane2 - 30, h_lane2), Color::WHITE, 10);
    draw_dotted_line(canvas, Point::new(v_lane2 - 60, height as i32), Point::new(v_lane2 - 60, h_lane2), Color::WHITE, 10);

    // --------------------- Ligne horizontale ---------------------
    // Partie haute (pleine)
    canvas.draw_line(Point::new(0, (height as i32) / 2), Point::new((width as i32) / 2 - 90, (height as i32) / 2)).unwrap();

    // Partie basse (pleine)
    canvas.draw_line(Point::new((width as i32) / 2 + 90, (height as i32) / 2),Point::new(width as i32, (height as i32) / 2)).unwrap();

    // --------------------- Ligne verticale ---------------------
    // Partie gauche
    canvas.draw_line(Point::new((width as i32) / 2, 0),Point::new((width as i32) / 2, (height as i32) / 2 - 90)).unwrap();

    // partie droite
    canvas.draw_line(Point::new((width as i32) / 2, (height as i32) / 2 + 90),Point::new((width as i32) / 2, height as i32)).unwrap();

}

fn draw_dotted_line(canvas: &mut WindowCanvas, start: Point, end: Point, color: Color, segment_length: i32) {
    let dx = end.x - start.x;
    let dy = end.y - start.y;

    // Convertir en f64 pour effectuer le calcul de la racine carrée
    let distance = ((dx * dx + dy * dy) as f64).sqrt();
    let steps = (distance / segment_length as f64).ceil() as i32;
    
    for i in 0..steps {
        let t = i as f64 / steps as f64;
        let x = start.x as f64 + t * dx as f64;
        let y = start.y as f64 + t * dy as f64;
        
        let next_t = (i + 1) as f64 / steps as f64;
        let next_x = start.x as f64 + next_t * dx as f64;
        let next_y = start.y as f64 + next_t * dy as f64;

        // On ne dessine que les segments pairs (ou impairs) pour créer l'effet pointillé
        if i % 2 == 0 {
            canvas.set_draw_color(color);
            canvas.draw_line(
                Point::new(x as i32, y as i32), 
                Point::new(next_x as i32, next_y as i32)
            ).unwrap();
        }
    }
}

pub fn render(canvas: &mut Canvas<Window>, vehicles: &mut Vec<Vehicule>) {
    for v in vehicles {
        // vehicule carré
        // canvas.set_draw_color(v.color);
        // canvas.fill_rect(Rect::new(v.position.0 as i32, v.position.1 as i32, 20, 20)).unwrap();

        // bieau véhicule
        let texture_creator = canvas.texture_creator();
        let texture = texture_creator.load_texture(v.texture.clone()).unwrap();
        let (rotation, dest_rect) = match v.direction {
            Direction::South => (180.0, Rect::new((v.position.0 - 5).try_into().unwrap(), (v.position.1).try_into().unwrap(), 30, 30)),
            Direction::North => (0.0, Rect::new((v.position.0 - 5).try_into().unwrap(), (v.position.1).try_into().unwrap(), 30, 30)),
            Direction::West => (270.0, Rect::new((v.position.0).try_into().unwrap(), (v.position.1 - 5).try_into().unwrap(), 30, 30)),
            Direction::East => (90.0, Rect::new((v.position.0).try_into().unwrap(), (v.position.1 - 5).try_into().unwrap(), 30, 30)),
        };
        canvas
            .copy_ex(&texture, None, Some(dest_rect), rotation, None, false, false)
            .unwrap();
    }
}
//check if car is beyond the screen width OR height
pub fn can_destroy(car: &mut Vehicule, width:u32, height:u32) -> bool {
    let vertical = match car.direction {
        Direction::North => true,
        Direction::South => true,
        Direction::East => false,
        Direction::West => false,
    };

    let futurpositiony = match car.direction {
        Direction::North => car.position.1 as i32 - car.velocity as i32,
        Direction::South => car.position.1 as i32 + car.velocity as i32,
        _ => car.position.1 as i32,
    };
    if vertical && (futurpositiony > height  as i32 || futurpositiony < 0) {
        return true
    } 

    let futurpositionx = match car.direction {
        Direction::East => car.position.0 as i32 + car.velocity as i32,
        Direction::West => car.position.0 as i32 - car.velocity as i32,
        _ => car.position.0 as i32,
    };  
    if !vertical && (futurpositionx > width as i32 || futurpositionx < 0) {
        return true
    } 

    return false
}
