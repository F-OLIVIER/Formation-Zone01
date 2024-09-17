use crate::hittable::{HitRecord, Hittable, Material};
use crate::ray::Ray;
use crate::vec3::{self, Point3, Vec3};

// use std::rc::Rc;
pub use std::sync::Arc;

pub struct Cube {
    center: Point3,   // Point central du cube
    side_length: f64, // Longueur d'un côté du cube
    mat: Arc<dyn Material>, // Matériau du cube
}

impl Cube {
    // Constructeur du cube
    pub fn new(center: Point3, side_length: f64, mat: Arc<dyn Material>) -> Cube {
        Cube {
            center,
            side_length,
            mat,
        }
    }

    // Méthode pour obtenir les points "min" et "max" en fonction du centre et de la longueur
    fn get_min_max(&self) -> (Point3, Point3) {
        let half_side = self.side_length / 2.0;
        let min = self.center - Vec3::new(half_side, half_side, half_side);
        let max = self.center + Vec3::new(half_side, half_side, half_side);
        (min, max)
    }
}

impl Hittable for Cube {
    fn hit(&self, r: &Ray, t_min: f64, t_max: f64, rec: &mut HitRecord) -> bool {
        // Récupère les points "min" et "max" à partir du centre et de la longueur du cube
        let (min, max) = self.get_min_max();

        // Inverse de la direction du rayon
        let inv_d = Vec3::new(1.0 / r.direction().x(), 1.0 / r.direction().y(), 1.0 / r.direction().z());

        let oc = r.origin() - self.center;
        // Calcule la longueur au carré de la direction du rayon
        let a = r.direction().length_squared();
        let half_b = vec3::dot(oc, r.direction()); // Calcule le produit scalaire entre le vecteur oc et la direction du rayon
        let c = oc.length_squared() - self.side_length * self.side_length; // Calcule c dans l'équation quadratique : (oc² - r²)
        let discriminant = half_b * half_b - a * c; // Calcule le discriminant de l'équation quadratique

        if discriminant < 0.0 { // Si le discriminant est négatif, il n'y a pas d'intersection
            return false; // Retourne faux, pas de collision
        }

        // Calcul des intersections pour chaque axe
        let t0 = (min - r.origin()) * inv_d;
        let t1 = (max - r.origin()) * inv_d;

        // Calcul des valeurs min et max de l'intersection
        let tmin = Vec3::new(t0.x().min(t1.x()), t0.y().min(t1.y()), t0.z().min(t1.z()));
        let tmax = Vec3::new(t0.x().max(t1.x()), t0.y().max(t1.y()), t0.z().max(t1.z()));

        // Trouver les valeurs globales de t_min et t_max
        let t_min = t_min.max(tmin.x()).max(tmin.y()).max(tmin.z());
        let t_max = t_max.min(tmax.x()).min(tmax.y()).min(tmax.z());

        if t_min > t_max {
            return false; // Pas de collision
        }

        // Enregistrement de l'impact
        rec.t = t_min;
        rec.p = r.at(rec.t);
        // rec.mat = Some(self.mat.clone());

        // Calcul de la normale sortante en fonction de la face touchée
        let outward_normal = if (rec.p.x() - min.x()).abs() < 1e-5 {
            Vec3::new(-1.0, 0.0, 0.0) // Face gauche
        } else if (rec.p.x() - max.x()).abs() < 1e-5 {
            Vec3::new(1.0, 0.0, 0.0) // Face droite
        } else if (rec.p.y() - min.y()).abs() < 1e-5 {
            Vec3::new(0.0, -1.0, 0.0) // Face inférieure
        } else if (rec.p.y() - max.y()).abs() < 1e-5 {
            Vec3::new(0.0, 1.0, 0.0) // Face supérieure
        } else if (rec.p.z() - min.z()).abs() < 1e-5 {
            Vec3::new(0.0, 0.0, -1.0) // Face arrière
        } else {
            Vec3::new(0.0, 0.0, 1.0) // Face avant
        };

        // let outward_normal = (rec.p - self.center) / self.side_length;

        rec.set_face_normal(r, outward_normal);
        rec.mat = Some(self.mat.clone()); // Associe le matériau de la sphère à l'enregistrement de l'impact
        true // Intersection trouvée
    }
}

