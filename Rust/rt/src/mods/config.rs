use std::io;

use crate::ASPECT_RATIO;

#[derive(Debug, Clone)]
pub struct Config {
    pub filename: String,
    pub taille_image: TailleImg,
    pub brightness: f64,
    pub camera_position: (f64, f64, f64),
    pub camera_angle: f64, 
}

#[derive(Debug, Clone)]
pub struct TailleImg {
    pub width: i32,
    pub height: i32
}

pub fn generate_config() -> Config {
    let mut config = Config {
        filename: "output".to_string(),
        taille_image: TailleImg{width: 600, height: 400},
        brightness: 1.0,
        camera_position: (15.0, 4.0, 4.0),
        camera_angle: 90.0,
    };

    println!("\n╭─────────────────────────────────────────────────────────╮");
    println!("│                                                         │");
    println!("│        Information required to produce the image        │");
    println!("│                                                         │");
    println!("╰─────────────────────────────────────────────────────────╯");

    // Nom fichier de sortie
    println!("Output filename. If empty, default is 'output' : ");
    let mut input_filename = String::new();
    io::stdin().read_line(&mut input_filename).expect("Error: input_filename");
    if !input_filename.trim().is_empty() {
        config.filename = input_filename.trim().to_string();
        println!("Using provided filename: {}\n", config.filename);
    } else {
        println!("Using default: {}\n", config.filename);
    }

    // Taille de l'image
    println!("───────────────────────────────────────────────────────────");
    println!("Image size (width only, format 16/9). If empty, default is '600' : ");
    let mut input_size = String::new();
    io::stdin().read_line(&mut input_size).expect("Error: input_size");
    if !input_size.trim().is_empty() {
        if let Ok(width) = input_size.trim().parse::<usize>() {
            config.taille_image = TailleImg{
                width: width as i32,
                height: (width as f64 / ASPECT_RATIO) as i32
            };
            println!("Image size: {:?}\n", config.taille_image);
        } else {
            println!("Invalid input for image size, using default: {:?}\n", config.taille_image);
        }
    } else {
        println!("Using default: {:?}\n", config.taille_image);
    }

    // Luminosité (brightness)
    println!("───────────────────────────────────────────────────────────");
    println!("Brightness. If empty, default is '1.0' (min: 0.2, max: 2) : ");
    let mut input_brightness = String::new();
    io::stdin().read_line(&mut input_brightness).expect("Error: input_brightness");
    if !input_brightness.trim().is_empty() {
        if let Ok(brightness) = input_brightness.trim().parse::<f64>() {
            if brightness > 0.1 && brightness < 2.1 {
                config.brightness = brightness;
                println!("");
            } else {
                println!("Invalid input for brightness, using default: {:?}\n", config.brightness);
            }
        } else {
            println!("Invalid input for brightness, using default: {:?}\n", config.brightness);
        }
    } else {
        println!("Using default: {:?}\n", config.brightness);
    }
    
    // Position de la camera (x, y, z)
    println!("───────────────────────────────────────────────────────────");
    println!("Camera position (x, y, z). If empty, default is '15.0, 4.0, 4.0' : ");
    let mut input_position = String::new();
    io::stdin().read_line(&mut input_position).expect("Error: input_position");
    if !input_position.trim().is_empty() {
        let position: Vec<&str> = input_position.trim().split_whitespace().collect();
        if position.len() == 3 {
            if let (Ok(x), Ok(y), Ok(z)) = (position[0].parse::<f64>(), position[1].parse::<f64>(), position[2].parse::<f64>()) {
                config.camera_position = (x, y, z);
                println!("");
            } else {
                println!("Invalid input for camera position, using default: {:?}\n", config.camera_position);
            }
        } else {
            println!("Invalid input for camera position, using default: {:?}\n", config.camera_position);
        }
    } else {
        println!("Using default: {:?}\n", config.camera_position); 
    }

    // Angle de la caméra
    println!("───────────────────────────────────────────────────────────");
    println!("Camera angle. If empty, default is '90.0' (min: 60, max: 150)  : ");
    let mut input_angle = String::new();
    io::stdin().read_line(&mut input_angle).expect("Error: input_angle");
    if !input_angle.trim().is_empty() {
        if let Ok(angle) = input_angle.trim().parse::<f64>() {
            if angle > 59.9 && angle < 150.1 {
                config.camera_angle = angle;
                println!("");
            } else {
                println!("Invalid input for camera angle, using default: {:?}\n", config.camera_angle);
            }
        } else {
            println!("Invalid input for camera angle, using default: {:?}\n", config.camera_angle);
        }
    } else {
        println!("Using default: {:?}\n", config.camera_angle); 
    }

    println!("╭─────────────────────────────────────────────────────────╮");
    println!("│                                                         │");
    println!("│           Rendering in progress, please wait            │");
    println!("│                                                         │");
    println!("╰─────────────────────────────────────────────────────────╯");

    config
}