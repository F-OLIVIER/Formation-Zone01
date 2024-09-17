pub mod mods;

pub use mods::*;
pub use cube::Cube;
pub use std::rc::Rc;
pub use crate::mods::color::Color;
pub use camera::Camera;
pub use common;
pub use hittable::{HitRecord, Hittable};
pub use hittable_list::HittableList;
pub use materials::{Dielectric, Lambertian, Metal};
pub use ray::Ray;
pub use sphere::Sphere;
pub use vec3::{Point3, Vec3};

fn main() {
    generate_img();
}
