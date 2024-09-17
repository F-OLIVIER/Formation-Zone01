use crate::hittable::{HitRecord, Hittable, Material};
use crate::ray::Ray;
use crate::vec3::{self, Point3};

// use std::rc::Rc;
use std::sync::Arc;
 
pub struct Sphere { // Déclare une structure nommée "Sphere"
    center: Point3, // Champ "center" de type Point3 qui représente le centre de la sphère
    radius: f64,    // Champ "radius" de type f64 qui représente le rayon de la sphère
    mat: Arc<dyn Material>, // Champ "mat" qui est un pointeur intelligent vers un objet implémentant le trait "Material"
}

impl Sphere { // Implémente des méthodes pour la structure Sphere
    pub fn new(cen: Point3, r: f64, m: Arc<dyn Material>) -> Sphere { // Fonction de constructeur qui prend un centre, un rayon et un matériau
        Sphere { // Crée et retourne une nouvelle instance de la structure Sphere
            center: cen, // Initialise le centre avec "cen"
            radius: r,   // Initialise le rayon avec "r"
            mat: m,      // Initialise le matériau avec "m"
        }
    }
}

impl Hittable for Sphere { // Implémente le trait "Hittable" pour la structure Sphere
    fn hit(&self, r: &Ray, t_min: f64, t_max: f64, rec: &mut HitRecord) -> bool { 
        // Fonction "hit" pour tester si un rayon touche la sphère, retourne un booléen
        let oc = r.origin() - self.center; // Calcule le vecteur du rayon au centre de la sphère
        let a = r.direction().length_squared(); // Calcule la longueur au carré de la direction du rayon
        let half_b = vec3::dot(oc, r.direction()); // Calcule le produit scalaire entre le vecteur oc et la direction du rayon
        let c = oc.length_squared() - self.radius * self.radius; // Calcule c dans l'équation quadratique : (oc² - r²)
        let discriminant = half_b * half_b - a * c; // Calcule le discriminant de l'équation quadratique

        if discriminant < 0.0 { // Si le discriminant est négatif, il n'y a pas d'intersection
            return false; // Retourne faux, pas de collision
        }

        let sqrt_d = f64::sqrt(discriminant); // Calcule la racine carrée du discriminant

        // Cherche la racine la plus proche dans la plage acceptable
        let mut root = (-half_b - sqrt_d) / a; // Calcule la première solution de l'équation quadratique
        if root <= t_min || t_max <= root { // Si la racine n'est pas dans l'intervalle [t_min, t_max]
            root = (-half_b + sqrt_d) / a; // Calcule la deuxième solution
            if root <= t_min || t_max <= root { // Si la deuxième solution n'est pas non plus valide
                return false; // Retourne faux, pas de collision
            }
        }

        rec.t = root; // Stocke la valeur de t (le point d'intersection) dans "rec"
        rec.p = r.at(rec.t); // Calcule le point d'intersection sur le rayon et le stocke dans "rec.p"
        let outward_normal = (rec.p - self.center) / self.radius; // Calcule la normale à la surface de la sphère à l'intersection
        rec.set_face_normal(r, outward_normal); // Définit la normale pour l'enregistrement de l'impact en fonction de l'orientation du rayon
        rec.mat = Some(self.mat.clone()); // Associe le matériau de la sphère à l'enregistrement de l'impact
        true // Retourne vrai, il y a eu une collision
    }
}
