use crate::{common, HitRecord, Hittable, Ray};
use crate::vec3::Vec3;
use super::vec3::{self, Shape};

pub type Color = Vec3;

pub fn write_color(pixel_color: Color, samples_per_pixel: i32, brightness: f64) -> String {
    // Divide the color by the number of samples and gamma-correct for gamma=2.0
    let scale = 1.0 / samples_per_pixel as f64;
    let r = (f64::sqrt(scale * pixel_color.x()) * brightness).clamp(0.0, 0.999);
    let g = (f64::sqrt(scale * pixel_color.y()) * brightness).clamp(0.0, 0.999);
    let b = (f64::sqrt(scale * pixel_color.z()) * brightness).clamp(0.0, 0.999);
 
    // Write the translated [0, 255] value of each color component
    format!(
        "{} {} {}\n",
        (256.0 * common::clamp(r, 0.0, 0.999)) as i32,
        (256.0 * common::clamp(g, 0.0, 0.999)) as i32,
        (256.0 * common::clamp(b, 0.0, 0.999)) as i32,
    )
}

pub fn ray_color(r: &Ray, world: &dyn Hittable, depth: i32) -> Color {
    if depth <= 0 {
        return Color::new(0.0, 0.0, 0.0);
    }
    let mut rec = HitRecord::new();
    if world.hit(r, 0.001, common::INFINITY, &mut rec) {
        let mut attenuation = Color::default();
        let mut scattered = Ray::default();
        if rec
            .mat
            .as_ref()
            .unwrap()
            .scatter(r, &rec, &mut attenuation, &mut scattered, Shape::Cube)
        {
            return attenuation * ray_color(&scattered, world, depth - 1);
        }
        return Color::new(0.0, 0.0, 0.0);
    }
    let unit_direction = vec3::unit_vector(r.direction());
    let t = 0.5 * (unit_direction.y() + 1.0);
    (1.0 - t) * Color::new(1.0, 1.0, 1.0) + t * Color::new(0.5, 0.7, 1.0)
}