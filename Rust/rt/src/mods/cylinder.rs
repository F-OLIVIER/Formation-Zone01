use crate::hittable::{HitRecord, Hittable, Material};
use crate::ray::Ray;
use crate::vec3::{Point3, Vec3};
use std::sync::Arc;

pub struct Cylinder {
    base_center: Point3, // Centre de la base du cylindre
    height: f64,         // Hauteur du cylindre
    radius: f64,         // Rayon du cylindre
    mat: Arc<dyn Material>, // Matériau du cylindre
}

impl Cylinder {
    pub fn new(base_center: Point3, height: f64, radius: f64, mat: Arc<dyn Material>) -> Cylinder {
        Cylinder {
            base_center,
            height,
            radius,
            mat,
        }
    }
}

impl Hittable for Cylinder {
    fn hit(&self, r: &Ray, t_min: f64, t_max: f64, rec: &mut HitRecord) -> bool {
        let mut hit_anything = false;
        let mut closest_so_far = t_max;

        // Test pour les côtés du cylindre
        let oc = r.origin() - self.base_center;
        let a = r.direction().x() * r.direction().x() + r.direction().z() * r.direction().z(); // Distance projetée dans le plan xy
        let half_b = oc.x() * r.direction().x() + oc.z() * r.direction().z();
        let c = oc.x() * oc.x() + oc.z() * oc.z() - self.radius * self.radius;
        let discriminant = half_b * half_b - a * c;

        if discriminant >= 0.0 {
            let sqrt_d = f64::sqrt(discriminant);

            // Chercher la racine la plus proche dans la plage acceptable
            let mut root = (-half_b - sqrt_d) / a;
            if root < t_min || root > t_max {
                root = (-half_b + sqrt_d) / a;
            }

            if root >= t_min && root <= t_max {
                let hit_y = r.origin().y() + root * r.direction().y(); // Calculer la hauteur d'impact sur le cylindre

                // Vérifier si l'intersection est dans la plage de hauteur du cylindre
                if hit_y >= self.base_center.y() && hit_y <= self.base_center.y() + self.height {
                    closest_so_far = root;
                    hit_anything = true;

                    rec.t = root;
                    rec.p = r.at(rec.t);
                    let outward_normal = Vec3::new(
                        (rec.p.x() - self.base_center.x()) / self.radius,
                        0.0,
                        (rec.p.z() - self.base_center.z()) / self.radius,
                    );
                    rec.set_face_normal(r, outward_normal);
                    rec.mat = Some(self.mat.clone());
                }
            }
        }

        // Test pour la base inférieure du cylindre (disque à y = base_center.y)
        let t_base_min = (self.base_center.y() - r.origin().y()) / r.direction().y();
        if t_base_min >= t_min && t_base_min <= closest_so_far {
            let p = r.at(t_base_min);
            let dist_from_center = (p.x() - self.base_center.x()) * (p.x() - self.base_center.x())
                + (p.z() - self.base_center.z()) * (p.z() - self.base_center.z());
            if dist_from_center <= self.radius * self.radius {
                closest_so_far = t_base_min;
                hit_anything = true;

                rec.t = t_base_min;
                rec.p = p;
                rec.set_face_normal(r, Vec3::new(0.0, -1.0, 0.0));
                rec.mat = Some(self.mat.clone());
            }
        }

        // Test pour la base supérieure du cylindre (disque à y = base_center.y + height)
        let t_base_max = (self.base_center.y() + self.height - r.origin().y()) / r.direction().y();
        if t_base_max >= t_min && t_base_max <= closest_so_far {
            let p = r.at(t_base_max);
            let dist_from_center = (p.x() - self.base_center.x()) * (p.x() - self.base_center.x())
                + (p.z() - self.base_center.z()) * (p.z() - self.base_center.z());
            if dist_from_center <= self.radius * self.radius {
                closest_so_far = t_base_max;
                hit_anything = true;

                rec.t = t_base_max;
                rec.p = p;
                rec.set_face_normal(r, Vec3::new(0.0, 1.0, 0.0));
                rec.mat = Some(self.mat.clone());
            }
        }

        hit_anything
    }
}
