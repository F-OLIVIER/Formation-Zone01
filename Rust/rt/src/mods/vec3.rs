use std::fmt::{Display, Formatter, Result};
use std::ops::{Add, AddAssign, Div, DivAssign, Mul, MulAssign, Neg, Sub};
 
pub use crate::common;
pub use rand::*;

#[derive(Copy, Clone, Default)]
pub struct Vec3 {
    e: [f64; 3],
}
 
impl Vec3 {
    pub fn new(x: f64, y: f64, z: f64) -> Vec3 {
        Vec3 { e: [x, y, z] }
    }

    pub fn random() -> Vec3 {
        Vec3::new(
            common::random_double(),
            common::random_double(),
            common::random_double(),
        )
    }
 
    pub fn random_range(min: f64, max: f64) -> Vec3 {
        Vec3::new(
            common::random_double_range(min, max),
            common::random_double_range(min, max),
            common::random_double_range(min, max),
        )
    }
 
    pub fn x(&self) -> f64 {
        self.e[0]
    }
 
    pub fn y(&self) -> f64 {
        self.e[1]
    }
 
    pub fn z(&self) -> f64 {
        self.e[2]
    }
 
    pub fn length(&self) -> f64 {
        f64::sqrt(self.length_squared())
    }
 
    pub fn length_squared(&self) -> f64 {
        self.e[0] * self.e[0] + self.e[1] * self.e[1] + self.e[2] * self.e[2]
    }

    pub fn near_zero(&self) -> bool {
        const EPS: f64 = 1.0e-8;
        // Return true if the vector is close to zero in all dimensions
        self.e[0].abs() < EPS && self.e[1].abs() < EPS && self.e[2].abs() < EPS
    }

    pub fn normalize(&self) -> Self {
        let len = self.length();
        if len == 0.0 {
            Vec3::new(0.0, 0.0, 0.0) // Éviter la division par zéro
        } else {
            Vec3::new(self.e[0] / len, self.e[1] / len, self.e[2] / len)
        }
    }
    pub fn cross(&self, other: Vec3) -> Vec3 {
        Vec3::new(
            self.e[1] * other.e[2] - self.e[2] * other.e[1],
            self.e[2] * other.e[0] - self.e[0] * other.e[2],
            self.e[0] * other.e[1] - self.e[1] * other.e[0]
        )
    }
}
 
// Type alias
pub type Point3 = Vec3;
 
// Output formatting
impl Display for Vec3 {
    fn fmt(&self, f: &mut Formatter) -> Result {
        write!(f, "{} {} {}", self.e[0], self.e[1], self.e[2])
    }
}
 
// -Vec3
impl Neg for Vec3 {
    type Output = Vec3;
 
    fn neg(self) -> Vec3 {
        Vec3::new(-self.x(), -self.y(), -self.z())
    }
}
 
// Vec3 += Vec3
impl AddAssign for Vec3 {
    fn add_assign(&mut self, v: Vec3) {
        *self = *self + v;
    }
}
 
// Vec3 *= f64
impl MulAssign<f64> for Vec3 {
    fn mul_assign(&mut self, t: f64) {
        *self = *self * t;
    }
}
 
// Vec3 /= f64
impl DivAssign<f64> for Vec3 {
    fn div_assign(&mut self, t: f64) {
        *self = *self / t;
    }
}
 
// Vec3 + Vec3
impl Add for Vec3 {
    type Output = Vec3;
 
    fn add(self, v: Vec3) -> Vec3 {
        Vec3::new(self.x() + v.x(), self.y() + v.y(), self.z() + v.z())
    }
}
 
// Vec3 - Vec3
impl Sub for Vec3 {
    type Output = Vec3;
 
    fn sub(self, v: Vec3) -> Vec3 {
        Vec3::new(self.x() - v.x(), self.y() - v.y(), self.z() - v.z())
    }
}
 
// Vec3 * Vec3
impl Mul for Vec3 {
    type Output = Vec3;
 
    fn mul(self, v: Vec3) -> Vec3 {
        Vec3::new(self.x() * v.x(), self.y() * v.y(), self.z() * v.z())
    }
}
 
// f64 * Vec3
impl Mul<Vec3> for f64 {
    type Output = Vec3;
 
    fn mul(self, v: Vec3) -> Vec3 {
        Vec3::new(self * v.x(), self * v.y(), self * v.z())
    }
}
 
// Vec3 * f64
impl Mul<f64> for Vec3 {
    type Output = Vec3;
 
    fn mul(self, t: f64) -> Vec3 {
        Vec3::new(self.x() * t, self.y() * t, self.z() * t)
    }
}
 
// Vec3 / f64
impl Div<f64> for Vec3 {
    type Output = Vec3;
 
    fn div(self, t: f64) -> Vec3 {
        Vec3::new(self.x() / t, self.y() / t, self.z() / t)
    }
}
 
pub fn dot(u: Vec3, v: Vec3) -> f64 {
    u.e[0] * v.e[0] + u.e[1] * v.e[1] + u.e[2] * v.e[2]
}
 
pub fn cross(u: Vec3, v: Vec3) -> Vec3 {
    Vec3::new(
        u.e[1] * v.e[2] - u.e[2] * v.e[1],
        u.e[2] * v.e[0] - u.e[0] * v.e[2],
        u.e[0] * v.e[1] - u.e[1] * v.e[0],
    )
}
 
pub fn unit_vector(v: Vec3) -> Vec3 {
    v / v.length()
}

pub fn random_in_unit_sphere() -> Vec3 {
    loop {
        let p = Vec3::random_range(-1.0, 1.0);
        if p.length_squared() >= 1.0 {
            continue;
        }
        return p;
    }
}

pub fn random_in_unit_cube() -> Vec3 {
    Vec3::random_range(-1.0, 1.0)
}

#[derive(Debug, PartialEq, Eq, PartialOrd, Ord)]
pub enum Shape {
    Sphere,
    Cube,
    Cylinder, // Ajout du cylindre comme nouvelle forme
    Plane,
}

pub fn random_in_shape(shape: Shape) -> Vec3 {
    match shape {
        Shape::Plane => {
            // Génère un point aléatoire dans un carré limité [-size, size] dans le plan
            Vec3::new(
                common::random_double_range(-10.0, 10.0), // x dans [-size, size]
                0.0,                                     // y reste 0 pour un plan horizontal
                common::random_double_range(-10.0, 10.0), // z dans [-size, size]
            )
        }
        Shape::Sphere => {
            // Génère un point aléatoire dans une sphère de rayon 1
            loop {
                let p = Vec3::random_range(-1.0, 1.0);
                if p.length_squared() < 1.0 {
                    return p;
                }
            }
        }
        Shape::Cube => {
            // Génère un point aléatoire dans un cube de côté 2
            Vec3::random_range(-1.0, 1.0)
        }
        Shape::Cylinder => {
            // Génère un point aléatoire dans un cylindre de rayon 1 et de hauteur 2
            loop {
                let xz = Vec3::new(
                    common::random_double_range(-1.0, 1.0), // Coordonnée x aléatoire dans [-1, 1]
                    0.0,                                    // Ignorer l'axe y pour le moment
                    common::random_double_range(-1.0, 1.0), // Coordonnée z aléatoire dans [-1, 1]
                );
                if xz.length_squared() < 1.0 {
                    // Si le point est dans le disque de rayon 1 (dans le plan xz)
                    let y = common::random_double_range(-1.0, 1.0); // Génère une hauteur aléatoire dans [-1, 1]
                    return Vec3::new(xz.x(), y, xz.z()); // Retourne un point valide dans le cylindre
                }
            }
        }
    }
}

pub fn random_unit_vector(shape: Shape) -> Vec3 {
    unit_vector(random_in_shape(shape))
}

pub fn random_in_unit_disk() -> Vec3 {
    loop {
        let p = Vec3::new(
            common::random_double_range(-1.0, 1.0),
            common::random_double_range(-1.0, 1.0),
            0.0,
        );
        if p.length_squared() >= 1.0 {
            continue;
        }
        return p;
    }
}

pub fn reflect(v: Vec3, n: Vec3) -> Vec3 {
    v - 2.0 * dot(v, n) * n
}

pub fn refract(uv: Vec3, n: Vec3, etai_over_etat: f64) -> Vec3 {
    let cos_theta = f64::min(dot(-uv, n), 1.0);
    let r_out_perp = etai_over_etat * (uv + cos_theta * n);
    let r_out_parallel = -f64::sqrt(f64::abs(1.0 - r_out_perp.length_squared())) * n;
    r_out_perp + r_out_parallel
}