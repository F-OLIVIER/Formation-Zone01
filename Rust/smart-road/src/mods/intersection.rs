use chrono::Utc;
use sdl2::rect::Rect;
use crate::Vehicule;

use super::{Comptage, Direction, Route};
//Système de cases pour départager l'intersection et voir quelles cases ont un potentiel risque de collision
pub fn case(num_case: i8, width: u32, height: u32) -> Rect {
    match num_case {
        // avant le carrefour
        -2 => Rect::new((width / 2 - 60).try_into().unwrap(), (height / 2 - 120).try_into().unwrap(), 30, 30),
        -3 => Rect::new((width / 2 - 30).try_into().unwrap(), (height / 2 - 120).try_into().unwrap(), 30, 30),
        -19 => Rect::new((width / 2 - 120).try_into().unwrap(), (height / 2).try_into().unwrap(), 30, 30),
        -25 => Rect::new((width / 2 - 120).try_into().unwrap(), (height / 2 + 30).try_into().unwrap(), 30, 30),
        -12 => Rect::new((width / 2 + 90).try_into().unwrap(), (height / 2 - 60).try_into().unwrap(), 30, 30),
        -18 => Rect::new((width / 2 - 90).try_into().unwrap(), (height / 2 - 30).try_into().unwrap(), 30, 30),
        -34 => Rect::new((width / 2).try_into().unwrap(), (height / 2 + 120).try_into().unwrap(), 30, 30),
        -35 => Rect::new((width / 2 + 30).try_into().unwrap(), (height / 2 + 120).try_into().unwrap(), 30, 30),

        // dans le carrefour
        2 => Rect::new((width / 2 - 60).try_into().unwrap(), (height / 2 - 90).try_into().unwrap(), 30, 30),
        3 => Rect::new((width / 2 - 30).try_into().unwrap(), (height / 2 - 90).try_into().unwrap(), 30, 30),
        8 => Rect::new((width / 2 - 60).try_into().unwrap(), (height / 2 - 60).try_into().unwrap(), 30, 30),
        9 => Rect::new((width / 2 - 30).try_into().unwrap(), (height / 2 - 60).try_into().unwrap(), 30, 30),
        10 => Rect::new((width / 2).try_into().unwrap(), (height / 2 - 60).try_into().unwrap(), 30, 30),
        11 => Rect::new((width / 2 + 30).try_into().unwrap(), (height / 2 - 60).try_into().unwrap(), 30, 30),
        12 => Rect::new((width / 2 + 60).try_into().unwrap(), (height / 2 - 60).try_into().unwrap(), 30, 30),
        14 => Rect::new((width / 2 - 60).try_into().unwrap(), (height / 2 - 30).try_into().unwrap(), 30, 30),
        15 => Rect::new((width / 2 - 30).try_into().unwrap(), (height / 2 - 30).try_into().unwrap(), 30, 30),
        16 => Rect::new((width / 2).try_into().unwrap(), (height / 2 - 30).try_into().unwrap(), 30, 30),
        17 => Rect::new((width / 2 + 30).try_into().unwrap(), (height / 2 - 30).try_into().unwrap(), 30, 30),
        18 => Rect::new((width / 2 + 60).try_into().unwrap(), (height / 2 - 30).try_into().unwrap(), 30, 30),
        19 => Rect::new((width / 2 - 90).try_into().unwrap(), (height / 2).try_into().unwrap(), 30, 30),
        20 => Rect::new((width / 2 - 60).try_into().unwrap(), (height / 2).try_into().unwrap(), 30, 30),
        21 => Rect::new((width / 2 - 30).try_into().unwrap(), (height / 2).try_into().unwrap(), 30, 30),
        22 => Rect::new((width / 2).try_into().unwrap(), (height / 2).try_into().unwrap(), 30, 30),
        23 => Rect::new((width / 2 + 30).try_into().unwrap(), (height / 2).try_into().unwrap(), 30, 30),
        25 => Rect::new((width / 2 - 90).try_into().unwrap(), (height / 2 + 30).try_into().unwrap(), 30, 30),
        26 => Rect::new((width / 2 - 60).try_into().unwrap(), (height / 2 + 30).try_into().unwrap(), 30, 30),
        27 => Rect::new((width / 2 - 30).try_into().unwrap(), (height / 2 + 30).try_into().unwrap(), 30, 30),
        28 => Rect::new((width / 2).try_into().unwrap(), (height / 2 + 30).try_into().unwrap(), 30, 30),
        29 => Rect::new((width / 2 + 30).try_into().unwrap(), (height / 2 + 30).try_into().unwrap(), 30, 30),
        34 => Rect::new((width / 2).try_into().unwrap(), (height / 2 + 60).try_into().unwrap(), 30, 30),
        35 => Rect::new((width / 2 + 30).try_into().unwrap(), (height / 2 + 60).try_into().unwrap(), 30, 30),
        _ => Rect::new(0, 0, 0, 0),
    }
}
//fn to check if there could be a collision by implementing security squares around the cars
pub fn prediction_to_velocity(cars: &mut Vec<Vehicule>, list_cars_passed: &mut Vec<String>, cars_in_intersection: &mut Comptage, width: u32, height: u32) {
    let copy_cars = cars.clone();
    let mut second_vehicule: Vec<(usize, usize)> = vec![];

    // comptage des tutures
    let mut nb_cars_left = 0;
    let mut nb_cars_straight = 0;
    let mut nb_cars_right = 0;
    
    for (index_current_car, current_car) in &mut cars.iter_mut().enumerate() {

        let rect_current_car = Rect::new((current_car.position.0).try_into().unwrap(), (current_car.position.1).try_into().unwrap(), 20, 20);
        // Dès que voitures sont dans l'intersection
        // Start time and end time for stats
        if enter_intersection(current_car.position, width, height) {
            current_car.time_pop = Utc::now();
        } else if exit_intersection(current_car.position, width, height) {
            current_car.time_depop = Utc::now();
        }

        // Comptage des tutures
        match current_car.route {
            Route::Left => nb_cars_left += 1,
            Route::Straight => nb_cars_straight += 1,
            Route::Right => nb_cars_right += 1,
        }

        let mut collision = false;
    
        // aprés le carrefour (vitesse rapide)
        let after_rect_top = Rect::new((width / 2).try_into().unwrap(), 0, 90, (height / 2 - 90).try_into().unwrap());
        let after_rect_bottom = Rect::new((width / 2 - 90).try_into().unwrap(), (height / 2 + 90).try_into().unwrap(), 90, (height / 2 - 90).try_into().unwrap());
        let after_rect_right = Rect::new((width / 2 + 90).try_into().unwrap(), (height / 2).try_into().unwrap(), (width / 2 - 90).try_into().unwrap(), 90);
        let after_rect_left = Rect::new(0, (height / 2 - 90).try_into().unwrap(), (width / 2 - 90).try_into().unwrap(), 90);

        //Cars have entered their exit zone, danger is gone so speeding they will be
        if is_rect_inside_case(after_rect_top, rect_current_car) ||
        is_rect_inside_case(after_rect_bottom, rect_current_car) ||
        is_rect_inside_case(after_rect_left, rect_current_car) ||
        is_rect_inside_case(after_rect_right, rect_current_car) {
            current_car.velocity = 10;
            //list used in mods.rs to delete the cars when they reach the borders
            if !list_cars_passed.contains(&current_car.name) {
                list_cars_passed.push(current_car.name.clone());
            }
            //Cars are entering big central square
        } else if (30 < rect_current_car.x as u32 && (width - 30) > rect_current_car.x as u32) && 
                  (30 < rect_current_car.y as u32 && (height - 30) > rect_current_car.y as u32) {
            
            // Rectangle de sécurité de la voiture (voiture entouré de 5px)
            let rect_security_current_car = match current_car.direction {
                Direction::North => Rect::new((current_car.position.0 - 5).try_into().unwrap(), (current_car.position.1 - 10).try_into().unwrap(), 30, 35),
                Direction::South => Rect::new((current_car.position.0 - 5).try_into().unwrap(), (current_car.position.1 - 5).try_into().unwrap(), 30, 35),
                Direction::East => Rect::new((current_car.position.0 - 5).try_into().unwrap(), (current_car.position.1 - 5).try_into().unwrap(), 35, 30),
                Direction::West => Rect::new((current_car.position.0 - 10).try_into().unwrap(), (current_car.position.1 - 5).try_into().unwrap(), 35, 30),
            };

            let rect_detection: Rect;

            if ((height / 2 - 30) < current_car.position.0 as u32 && (height / 2 + 6) > current_car.position.0 as u32) && 
                ((width / 2 - 30) < current_car.position.1 as u32 && (width / 2 + 6) > current_car.position.1 as u32) {
                    // si la tuture est dans le carré centrale (petit carré case 15 à 22)
                    //Depending on the route, the car in the XS square won't have to look at the same places
                rect_detection = match current_car.direction {
                    Direction::North => {
                        //car has turned so watching only in front
                        if current_car.turning {
                            Rect::new((current_car.position.0 - 5).try_into().unwrap(), (current_car.position.1 - 30).try_into().unwrap(), 30, 30)
                            //just before the turning square : looking at case 15 in case of another car
                        } else if case(22, width, height).has_intersection(rect_current_car) || is_rect_inside_case(case(22, width, height), rect_current_car) {
                            Rect::new((current_car.position.0 - 35).try_into().unwrap(), (current_car.position.1 - 30).try_into().unwrap(), 55, 30)
                        } else { // si carré 16
                            Rect::new((current_car.position.0 - 35).try_into().unwrap(), (current_car.position.1 - 10).try_into().unwrap(), 55, 20)
                        }
                    }
                    Direction::South => {
                        if current_car.turning {
                            Rect::new((current_car.position.0 - 5).try_into().unwrap(), (current_car.position.1 + 20).try_into().unwrap(), 30, 30)
                        } else if case(15, width, height).has_intersection(rect_current_car) || is_rect_inside_case(case(15, width, height), rect_current_car){
                            Rect::new((current_car.position.0 - 5).try_into().unwrap(), (current_car.position.1 + 20).try_into().unwrap(), 55, 30)
                        } else { // si carré 21
                            Rect::new((current_car.position.0 - 5).try_into().unwrap(), (current_car.position.1 + 10).try_into().unwrap(), 55, 20)
                        }
                    }
                    Direction::East =>{
                        if current_car.turning {
                            Rect::new((current_car.position.0 + 20).try_into().unwrap(), (current_car.position.1 - 5).try_into().unwrap(), 30, 30)
                        } else if case(21, width, height).has_intersection(rect_current_car) || is_rect_inside_case(case(21, width, height), rect_current_car){
                            Rect::new((current_car.position.0 + 20).try_into().unwrap(), (current_car.position.1 - 35).try_into().unwrap(), 30, 55)
                        } else { // si carré 22
                            Rect::new((current_car.position.0 + 10).try_into().unwrap(), (current_car.position.1 - 35).try_into().unwrap(), 20, 55)
                        }
                    }
                    Direction::West => {
                        if current_car.turning {
                            Rect::new((current_car.position.0 - 30).try_into().unwrap(), (current_car.position.1 - 5).try_into().unwrap(), 30, 30)
                        } else if case(16, width, height).has_intersection(rect_current_car) || is_rect_inside_case(case(16, width, height), rect_current_car){
                            Rect::new((current_car.position.0 - 30).try_into().unwrap(), (current_car.position.1 - 5).try_into().unwrap(), 30, 55)
                        } else { // si carré 15
                            Rect::new((current_car.position.0 - 10).try_into().unwrap(), (current_car.position.1 - 5).try_into().unwrap(), 20, 55)
                        }
                    }
                };
                
            } else if ((width / 2 - 60) < current_car.position.0 as u32 && (width / 2 + 36) > current_car.position.0 as u32) && 
                    ((height / 2 - 60) < current_car.position.1 as u32 && (height / 2 + 36) > current_car.position.1 as u32) {
                // si la tuture est dans le carré centrale (case 8 à 29)

                rect_detection = match current_car.direction {
                    Direction::North => {
                        // BF et BG
                        Rect::new((current_car.position.0 - 5).try_into().unwrap(), (current_car.position.1 - 30).try_into().unwrap(), 30, 30)
                    },
                    Direction::South => {
                        // HF et HG
                        Rect::new((current_car.position.0 - 5).try_into().unwrap(), (current_car.position.1 + 20).try_into().unwrap(), 30, 30)
                    },
                    Direction::East =>  {
                        // GF et GG
                        Rect::new((current_car.position.0 + 20).try_into().unwrap(), (current_car.position.1 - 5).try_into().unwrap(), 30, 30)
                    },
                    Direction::West =>  {
                        // DF et DG
                        Rect::new((current_car.position.0 - 30).try_into().unwrap(), (current_car.position.1 - 5).try_into().unwrap(), 30, 30)
                    },
                };

            } else {
                // si la tuture n'est pas dans le carré central : regarde devant
                rect_detection = match current_car.direction {
                    Direction::North => Rect::new((current_car.position.0 - 5).try_into().unwrap(), (current_car.position.1 - 30).try_into().unwrap(), 30, 30),
                    Direction::South => Rect::new((current_car.position.0 - 5).try_into().unwrap(), (current_car.position.1 + 20).try_into().unwrap(), 30, 30),
                    Direction::East => Rect::new((current_car.position.0 + 20).try_into().unwrap(), (current_car.position.1 - 5).try_into().unwrap(), 30, 30),
                    Direction::West => Rect::new((current_car.position.0- 30).try_into().unwrap(), (current_car.position.1 - 5).try_into().unwrap(), 30, 30),
                };
            }

                    // carré central (case 1 à 36)
                for (index_test_car, test_car) in copy_cars.iter().enumerate() {
                    if current_car != test_car &&
                        ((width / 2 - 90) < test_car.position.0 as u32 && (width / 2 + 66) > test_car.position.0 as u32) && 
                        ((height / 2 - 90) < test_car.position.1 as u32 && (height / 2 + 66) > test_car.position.1 as u32) {

                        let rect_test_car = Rect::new((test_car.position.0).try_into().unwrap(), (test_car.position.1).try_into().unwrap(), 20, 20);
                        let security_rect_test_car = Rect::new((test_car.position.0 - 5).try_into().unwrap(), (test_car.position.1 - 5).try_into().unwrap(), 30, 30);
                        
                        if rect_security_current_car.has_intersection(security_rect_test_car) {
                            //bool to check if current and testing car have been tested together already
                            let mut is_in_second_vehicule = false;
                            for vehicule in &second_vehicule {
                                if (vehicule.0 == index_current_car && vehicule.1 == index_test_car) || (vehicule.1 == index_current_car && vehicule.0 == index_test_car) {
                                    is_in_second_vehicule = true;
                                    break;
                                }
                                
                            }
                            if !is_in_second_vehicule {
                                second_vehicule.push((index_current_car, index_test_car));
                            }
                        } else if rect_detection.has_intersection(rect_test_car) {
                            collision = true;
                            break;
                        }
                    }
                }

                if collision {
                    // Présence véhicule (vitesse lente)
                    current_car.velocity = 1;
                } else {
                    // Absence véhicule (vitesse moyenne)
                    current_car.velocity = 5;
                }

        }
    }

    // println!("second_vehicule :\n{:?}", second_vehicule);

    // collision imminente, arret du véhicule
    for vehicule in &second_vehicule {
        // check si le 1er véhicule est à l'arrêt
        if cars[vehicule.0].velocity != 0 {
            // si non, on le stoppe et le deuxième véhicule peut continuer à rouler jusqu'à ce qu'il rentre dans la zone de collision
            cars[vehicule.0].velocity = 0;
            cars[vehicule.1].velocity = 5;
            // Si le 1er véhicule est à l'arrêt :
        } else if cars[vehicule.0].velocity == 0 {
            // On stoppe l'autre véhicule
            cars[vehicule.1].velocity = 0;
            // Si problème de collision se règle, on redémarre doucement le premier véhicule
        } else {
            cars[vehicule.0].velocity = 1;
            // Si le premier véhicule n'a pas atteint la vitesse moyenne : possibilité d'un autre arrêt
            if cars[vehicule.0].velocity < 6 {
                // donc on attribue la vitesse de l'autre véhicule au premier -1 pour qu'il reste derrière (ou ne rentre pas une nouvelle fois dedans)
                cars[vehicule.1].velocity = cars[vehicule.0].velocity -1;
                // Si premier véhicule à la vitesse moyenne : on redémarre doucement le deuxième pour déboucher la zone
            } else {
                cars[vehicule.1].velocity = 2;
            }
        }
    }

    cars_in_intersection.left_cars = nb_cars_left;
    cars_in_intersection.straight_cars = nb_cars_straight;
    cars_in_intersection.right_cars = nb_cars_right;
}

fn is_rect_inside_case(big_rect: Rect, low_rect: Rect) -> bool {
    low_rect.x() >= big_rect.x() &&
    low_rect.y() >= big_rect.y() &&
    low_rect.right() <= big_rect.right() &&
    low_rect.bottom() <= big_rect.bottom()
}

fn enter_intersection(car: (u32, u32), width: u32, height: u32) -> bool {
    let line_h = Rect::new(width as i32 / 2 - 90, height as i32 / 2 - 90, 90, 1);
    let line_b = Rect::new(width as i32 / 2, height as i32 / 2 + 90, 90, 1);
    let line_g = Rect::new(width as i32 / 2 - 90, height as i32 / 2, 1, 90);
    let line_d = Rect::new(width as i32 / 2 + 90, height as i32 / 2 - 90, 1, 90);

    let car = Rect::new(car.0 as i32, car.1 as i32, 5, 5);

    if car.has_intersection(line_h) || car.has_intersection(line_b) || car.has_intersection(line_g)|| car.has_intersection(line_d) {
        return true;
    }

    return false;
}

fn exit_intersection(car: (u32, u32), width: u32, height: u32) -> bool {
    let line_h = Rect::new(width as i32 / 2, height as i32 / 2 - 90, 90, 1);
    let line_b = Rect::new(width as i32 / 2 - 90, height as i32 / 2 + 90, 90, 1);
    let line_g = Rect::new(width as i32 / 2 - 90, height as i32 / 2 - 90, 1, 90);
    let line_d = Rect::new(width as i32 / 2 + 90, height as i32 / 2, 1, 90);

    let car = Rect::new(car.0 as i32, car.1 as i32, 5, 5);

    if car.has_intersection(line_h) || car.has_intersection(line_b) || car.has_intersection(line_g)|| car.has_intersection(line_d) {
        return true;
    }

    return false;
}