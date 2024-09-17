use std::sync::Arc;
use crate::vec3::{Vec3, Point3};
use crate::ray::Ray;
use crate::hittable::{HitRecord, Hittable, Material};

pub struct Plane {
    origin: Point3,    // Un point d'origine sur le plan
    normal: Vec3,      // La normale au plan
    size: f64,         // Taille du plan (rayon pour un carré centré)
    mat: Arc<dyn Material>, // Matériau du plan
}

impl Plane {
    pub fn new(origin: Point3, normal: Vec3, size: f64, mat: Arc<dyn Material>) -> Self {
        Plane {
            origin,
            normal: normal.normalize(), // Normaliser la normale
            size,
            mat,
        }
    }
}

impl Hittable for Plane {
    fn hit(&self, r: &Ray, t_min: f64, t_max: f64, rec: &mut HitRecord) -> bool {
        let denom = self.normal.x() * r.direction().x() + self.normal.y() * r.direction().y() + self.normal.z() * r.direction().z();
        if denom.abs() > 1e-6 { // Vérifie si le rayon n'est pas parallèle au plan
            let t = (self.origin.x() - r.origin().x()) * self.normal.x() + (self.origin.y() - r.origin().y()) * self.normal.y() + (self.origin.z() - r.origin().z()) * self.normal.z();
            let t = t / denom;
            
            if t < t_min || t > t_max {
                return false; // Pas d'intersection dans l'intervalle donné
            }

            // Calculer le point d'intersection
            let hit_point = r.at(t);

            // Vérifier si le point d'intersection est dans les limites du plan
            let offset = hit_point - self.origin; // Vecteur de l'origine au point d'impact
            let x_dist = offset.x(); // Composante x
            let z_dist = offset.z(); // Composante z

            if x_dist.abs() > self.size || z_dist.abs() > self.size {
                return false; // Le point est en dehors des limites du plan
            }

            // Stocker les informations dans HitRecord
            rec.t = t;
            rec.p = hit_point;
            rec.set_face_normal(r, self.normal);
            rec.mat = Some(self.mat.clone());
            true
        } else {
            false // Le rayon est parallèle au plan
        }
    }
}
