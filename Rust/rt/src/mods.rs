pub mod color;
pub mod vec3;
pub mod ray;
pub mod hittable;
pub mod sphere;
pub mod hittable_list;
pub mod common;
pub mod camera;
pub mod materials;
pub mod cube;
pub mod cylinder;
pub mod config;
pub mod plane;

use std::{fs::{File, OpenOptions}, sync::Arc};
use std::io::Write;
use crate::{Camera, Color, Dielectric, HittableList, Lambertian, Metal, Sphere};
use indicatif::{ProgressBar, ProgressStyle};
use plane::Plane;
use rayon::iter::{IntoParallelIterator, ParallelIterator};
use color::*;
use config::*;
use cube::*;
use cylinder::*;
use vec3::*;

pub const ASPECT_RATIO: f64 = 16.0 / 9.0;

pub fn generate_img() {
    // récupération des configs souhaité par l'utilisateur
    let config: Config = generate_config();
    // println!("Config : {:?}", config);
    let mut file = OpenOptions::new()
        .write(true) // Autorise l'écriture dans le fichier
        .create(true) // Crée le fichier s'il n'existe pas
        .truncate(true) // Vide le fichier s'il existe déjà
        .open(config.filename.to_owned() + ".ppm") // Ouvre le fichier "filename.ppm"
        .expect("Error 0: Unable to open or create file"); // affiche l'erreur si probléme

    // Traitement de l'image
    traitement(&mut file, &config);
}


fn traitement(file: &mut File, config: &Config) {
    // Constante de config
    const SAMPLES_PER_PIXEL: i32 = 500;
    const MAX_DEPTH: i32 = 100;

    // Initialiser la barre de progression
    let total_pixels = config.taille_image.width * config.taille_image.height;
    let progress_bar = ProgressBar::new(total_pixels as u64);
    progress_bar.set_style(
        ProgressStyle::with_template("{spinner:.green} [{bar:80.green/blue}] {elapsed_precise} ({percent}%)").unwrap().progress_chars("※⏵…")
    );
    
    // Position de la caméra
    let lookfrom = Point3::new(config.camera_position.0, config.camera_position.1, config.camera_position.2);
    let lookat = Point3::new(0.0, 0.0, 0.0);
    let direction = (lookat - lookfrom).normalize();
    // Calculer `vup` en fonction de la direction pour eviter les caméras aberrante
    let default_vup = Point3::new(0.0, 1.0, 0.0);
    let right = direction.cross(default_vup).normalize();
    let vup = right.cross(direction).normalize();
 
    // Camera
    let cam = Camera::new(
        lookfrom, // Position de la caméra
        lookat, // Point regardé par la caméra
        vup, // Direction du haut de l'image
        config.camera_angle, // angle de vision
        ASPECT_RATIO,
        0.1 + (90.0 - config.camera_angle).max(0.0) / 90.0 * 0.1, // Contrôle la profondeur de champ et le flou
        (lookfrom - lookat).length(), // Distance de mise au point,
    );

    // Création des matériaux
    let mut world = HittableList::new();
    
    let material_ground = Arc::new(Lambertian::new(Color::new(0.0, 0.8, 0.5)));
    world.add(Box::new(Plane::new( Point3::new(0.0, -1.0, 0.0), Vec3::new(0.0, 3.0, 0.0), 6.0, material_ground,)));

    // let material_center = Arc::new(Metal::new(Color::new(0.8, 0.6, 0.2), 0.1));
    // world.add(Box::new(Sphere::new( Point3::new(0.3, 0.3, 1.0), 0.5, material_center,)));
    let material_center = Arc::new(Dielectric::new(0.8));
    world.add(Box::new(Sphere::new( Point3::new(0.3, 0.3, 1.0), 0.2, material_center,)));

    let material_left = Arc::new(Metal::new(Color::new(0.8, 0.8, 0.8), 0.3));
    world.add(Box::new(Cube::new( Point3::new(0.0, 0.3, 0.0), 1.0, material_left,)));

    let material_right = Arc::new(Metal::new(Color::new(0.8, 0.6, 0.2), 1.0));
    world.add(Box::new(Cylinder::new( Point3::new(2.0, 0.3, -0.5), 0.5, 0.2, material_right,)));

    // Render (Header)
    writeln!(file, "P3\n{} {}\n255\n", config.taille_image.width, config.taille_image.height).expect("Error: Unable to write to file (img)");

    // Render (Content)
    for j in (0..config.taille_image.height).rev() {
        let mut res_img = String::new();

        let pixel_colors: Vec<_> = (0..config.taille_image.width)
            .into_par_iter()
            .map(|i| {
                let mut pixel_color = Color::new(0.0, 0.0, 0.0);
                for _ in 0..SAMPLES_PER_PIXEL {
                    let u = ((i as f64) + common::random_double()) / (config.taille_image.width - 1) as f64;
                        let v = ((j as f64) + common::random_double()) / (config.taille_image.height - 1) as f64;
                    let r = cam.get_ray(u, v);
                    pixel_color += ray_color(&r, &world, MAX_DEPTH);
                }
                pixel_color
            })
            .collect();

        for pixel_color in pixel_colors {
            res_img += &color::write_color( pixel_color, SAMPLES_PER_PIXEL, config.brightness);
        }

        writeln!(file, "{}", res_img).expect("Error: Unable to write to file (img)");

        // Mettre à jour la barre de progression
        progress_bar.inc(config.taille_image.width as u64);
    }

    // Mettre fin à la barre de progression
    progress_bar.finish_with_message("done");
}

