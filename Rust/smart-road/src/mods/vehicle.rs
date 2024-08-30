use chrono::{DateTime, Utc};
use rand::{self, Rng};
use crate::can_destroy;

#[derive(Debug, Clone, PartialEq)]
pub struct Vehicule {
    pub name: String,
    pub direction: Direction,
    pub route: Route,
    pub turning: bool,
    pub texture: String,
    pub position: (u32, u32),
    pub velocity: u32,

    // pour les stats
    pub time_pop: DateTime<Utc>,
    pub time_depop: DateTime<Utc>,
}

#[derive(Debug, Clone, PartialEq)]
pub enum Direction {
    North,
    South,
    East,
    West,
}

#[derive(Debug, Clone, PartialEq)]
pub enum Route {
    Left,
    Right,
    Straight,
}

pub struct Comptage {
    pub left_cars: usize,
    pub straight_cars: usize,
    pub right_cars: usize,
}

pub fn add_texture(route: &Route, direction: &Direction) -> String {
    match direction {
        Direction::North => match route {
            Route::Left => "./src/assets/Ambulance.png".to_string(),
            Route::Right => "./src/assets/Audi.png".to_string(),
            Route::Straight => "./src/assets/Black_viper.png".to_string(),
        },
        Direction::South => match route {
            Route::Left => "./src/assets/Truck.png".to_string(),
            Route::Right => "./src/assets/Audi.png".to_string(),
            Route::Straight => "./src/assets/Black_viper.png".to_string(),
        },
        Direction::East => match route {
            Route::Left => "./src/assets/Police.png".to_string(),
            Route::Right => "./src/assets/Audi.png".to_string(),
            Route::Straight => "./src/assets/Black_viper.png".to_string(),
        },
        Direction::West => match route {
            Route::Left => "./src/assets/Mini_van.png".to_string(),
            Route::Right => "./src/assets/Audi.png".to_string(),
            Route::Straight => "./src/assets/Black_viper.png".to_string(),
        },
    }

    // match route {
    //     Route::Left => "./src/assets/Ambulance.png".to_string(),
    //     Route::Right => "./src/assets/Audi.png".to_string(),
    //     Route::Straight => "./src/assets/Black_viper.png".to_string(),
    // }

    // Sprite aléatoire
    // match createrandom(0, 8) {
    //     0 => "./src/assets/Ambulance.png".to_string(),
    //     1 => "./src/assets/Audi.png".to_string(),
    //     2 => "./src/assets/Black_viper.png".to_string(),
    //     3 => "./src/assets/Car.png".to_string(),
    //     4 => "./src/assets/Mini_truck.png".to_string(),
    //     5 => "./src/assets/Mini_van.png".to_string(),
    //     6 => "./src/assets/Police.png".to_string(),
    //     7 => "./src/assets/Taxi.png".to_string(),
    //     _ => "./src/assets/Truck.png".to_string(),
    // }
}

pub fn select_route(cars_in_intersection: &Comptage) -> Route {
    let mut rng = rand::thread_rng();
    if cars_in_intersection.left_cars > 3 {
        match rng.gen_range(0..2) {
            0 => Route::Right,
            _ => Route::Straight,
        }
    } else {
        match rng.gen_range(0..3) {
            0 => Route::Left,
            1 => Route::Right,
            _ => Route::Straight,
        }
    }
}

impl Direction {
    pub fn random_direction() -> Direction {
        let mut rng = rand::thread_rng();
        match rng.gen_range(0..=3) {
            0 => Direction::East,
            1 => Direction::North,
            2 => Direction::South,
            _ => Direction::West,
        }
    }
}

impl Vehicule {
    pub fn spawn_cars(cars: &mut Vec<Vehicule>, direction: Direction, width: u32, height: u32, num_name: &mut usize, cars_in_intersection: &Comptage) {
        let route = select_route(cars_in_intersection);
        // let vehicle_color = add_color(route.clone());
        let texture = add_texture(&route, &direction);

        let position = match direction {
            Direction::South => match route {
                Route::Left => ((width / 2) - 25, 0),
                Route::Straight => ((width / 2) - 55, 0),
                Route::Right => ((width / 2) - 85, 0),
            },
            Direction::East => match route {
                Route::Left => (0, (height / 2) + 5),
                Route::Straight => (0, (height / 2) + 35),
                Route::Right => (0, (height / 2) + 65),
            },
            Direction::North => match route {
                Route::Left => ((width / 2) + 5, height - 20),
                Route::Straight => ((width / 2) + 35, height - 20),
                Route::Right => ((width / 2) + 65, height - 20),
            },
            Direction::West => match route {
                Route::Left => (width - 20, (height / 2) - 25),
                Route::Straight => (width - 20, (height / 2) - 55),
                Route::Right => (width - 20, (height / 2) - 85),
            },
        };
        /*Direction::South => (val, match => )*/

        let time = Utc::now();

        cars.push(Vehicule {
            name: format!("Car number {}", num_name),
            position,
            direction,
            turning: false,
            route,
            texture,
            velocity: 5,
            time_pop: time,
            time_depop: time,
        });
        *num_name += 1;
    }

    pub fn move_cars(&mut self, width: u32, height: u32) -> bool {
        if !can_destroy(self, width, height) {
            match self.direction {
                Direction::North => self.position.1 -= self.velocity,
                Direction::South => self.position.1 += self.velocity,
                Direction::East => self.position.0 += self.velocity,
                Direction::West => self.position.0 -= self.velocity,
            }
            return true;
        } else {
            return false;
        }
    }

    pub fn update_route(&mut self, width: u32, height: u32) {
        match self.route {
            // tourne a droite
            Route::Right => {
                self.direction = match self.position {
                    // dir North
                    (x, y) if x == width / 2 + 65 && y == height / 2 + 63 ||
                                        x == width / 2 + 65 && y == height / 2 + 64 || 
                                        x == width / 2 + 65 && y == height / 2 + 65 || 
                                        x == width / 2 + 65 && y == height / 2 + 66 ||
                                        x == width / 2 + 65 && y == height / 2 + 67 => 
                                        Direction::East, 
                    // dir South
                    (x, y) if x == width / 2 - 85 && y == height / 2 - 83 ||
                                        x == width / 2 - 85 && y == height / 2 - 84 ||
                                        x == width / 2 - 85 && y == height / 2 - 85 || 
                                        x == width / 2 - 85 && y == height / 2 - 86 ||
                                        x == width / 2 - 85 && y == height / 2 - 87 => 
                                        Direction::West,
                    // dir East
                    (x, y) if x == width / 2 - 83 && y == height / 2 + 65 ||
                                        x == width / 2 - 84 && y == height / 2 + 65 || 
                                        x == width / 2 - 85 && y == height / 2 + 65 || 
                                        x == width / 2 - 86 && y == height / 2 + 65 ||
                                        x == width / 2 - 87 && y == height / 2 + 65 => 
                                        Direction::South,
                    // dir West
                    (x, y) if x == width / 2 + 63 && y == height / 2 - 85 || 
                                        x == width / 2 + 64 && y == height / 2 - 85 || 
                                        x == width / 2 + 65 && y == height / 2 - 85 || 
                                        x == width / 2 + 66 && y == height / 2 - 85 ||
                                        x == width / 2 + 67 && y == height / 2 - 85 => 
                                        Direction::North,
                    // Conserver la direction actuelle si aucune correspondance
                    _ => self.direction.clone(), 
                }
            }

            // tourne a gauche
            Route::Left => {
                self.direction = match self.position {
                    // dir North
                    (x, y) if x == width / 2 + 5 && y == height / 2 - 27 ||
                                        x == width / 2 + 5 && y == height / 2 - 26 || 
                                        x == width / 2 + 5 && y == height / 2 - 25 || 
                                        x == width / 2 + 5 && y == height / 2 - 24 ||
                                        x == width / 2 + 5 && y == height / 2 - 23 => 
                                        Direction::West, 
                    // dir South
                    (x, y) if x == width / 2 - 25 && y == height / 2 + 7 ||
                                        x == width / 2 - 25 && y == height / 2 + 6 ||
                                        x == width / 2 - 25 && y == height / 2 + 5 || 
                                        x == width / 2 - 25 && y == height / 2 + 4 ||
                                        x == width / 2 - 25 && y == height / 2 + 3 => 
                                        Direction::East,
                    // dir East
                    (x, y) if x == width / 2 + 3 && y == height / 2 + 5 ||
                                        x == width / 2 + 4 && y == height / 2 + 5 || 
                                        x == width / 2 + 5 && y == height / 2 + 5 || 
                                        x == width / 2 + 6 && y == height / 2 + 5 ||
                                        x == width / 2 + 7 && y == height / 2 + 5 => 
                                        Direction::North,
                    // dir West
                    (x, y) if x == width / 2 - 23 && y == height / 2 - 25 || 
                                        x == width / 2 - 24 && y == height / 2 - 25 || 
                                        x == width / 2 - 25 && y == height / 2 - 25 || 
                                        x == width / 2 - 26 && y == height / 2 - 25 ||
                                        x == width / 2 - 27 && y == height / 2 - 25 => 
                                        Direction::South,
                    // Conserver la direction actuelle si aucune correspondance
                    _ => self.direction.clone(), 
                }
            }

            // va en face
            Route::Straight => {
                self.direction = match self.position {
                    _ => self.direction.clone(), // Conserver la direction actuelle
                }
            }
        }
    }
}
