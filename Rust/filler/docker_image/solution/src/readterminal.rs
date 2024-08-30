// use std::io::{self, Write};
// use std::fs::{File, OpenOptions};

use std::io::{self};

pub fn read_line_stdin() {
    // crée un fichier de log (pour debug)
    // let mut file = OpenOptions::new()
    //     .write(true) // Autorise l'écriture dans le fichier
    //     .create(true) // Crée le fichier s'il n'existe pas
    //     .truncate(true) // Vide le fichier s'il existe déjà
    //     .open("output.txt") // Ouvre le fichier "output.txt"
    //     .expect("Error 0: Unable to open or create file"); // affiche l'erreur si probléme

    // Récupérer les données du player à jouer (p1 ou p2)
    let mut player_input = String::new();
    io::stdin().read_line(&mut player_input).expect("Error: player_input");
    // Ecrit le nom du joueur dans le fichier
    // writeln!(file, "{}",player_input).expect("Error 1: Unable to write to file");

    // let mut number_loop = 0;
    
    loop {
        // Lit les données de la map ("Anfield xx xx")
        let mut input_anfield = String::new();
        io::stdin().read_line(&mut input_anfield).expect("Error: input_anfield");
        // Vérification pour ignorer les lignes vide 
        if input_anfield.contains("Anfield") { 

            // Ecrire dans le fichier le n° de loop
            // writeln!(file, "__________________________________________________________________________________________\n-> Loop Number: {}", number_loop).expect("Error 2: Unable to write to file");


            let anfield: Vec<&str> = input_anfield.split_whitespace().collect();
            // Taille de la map
            let taille_map: (i32, i32) = (
                anfield[1].parse().unwrap(), 
                anfield[2].replace(":", "").parse().unwrap()
            );
            // Lit les lignes de la map
            let mut map: Vec<String> = vec![];
            for i in 0..taille_map.1+1 {
                let mut line_map = String::new();
                io::stdin().read_line(&mut line_map).expect("Error: line_map");
                // map.push(line_map);
                if i != 0 {
                    let to_push = &line_map[4..];
                    map.push(to_push.to_string());
                }
            }
            // Ecrit la map dans le fichier de log
            // writeln!(file, "Anfield : {}/{}\n{}",taille_map.0, taille_map.1, map.join("")).expect("Error 3: Unable to write to file");

            // Lit les données de la pièce "piece xx xx"
            let mut input_piece = String::new();
            io::stdin().read_line(&mut input_piece).expect("Error: input_piece");
            let piece_decode: Vec<&str> = input_piece.split_whitespace().collect();
            // Taille de la piece
            let taille_piece: (i32, i32) = (
                piece_decode[1].parse().unwrap(), 
                piece_decode[2].replace(":", "").parse().unwrap()
            );
            // Lit les lignes de la piece
            let mut piece: Vec<String> = vec![];
            for _ in 0..taille_piece.1 {
                let mut line_piece = String::new();
                io::stdin().read_line(&mut line_piece).expect("Error: line_piece");
                piece.push(line_piece);
            }
            // Écrit la piece dans le fichier de log
            // writeln!(file, "piece : {}/{}\n{}",taille_piece.0, taille_piece.1, piece.join("")).expect("Error 4: Unable to write to file");

            // Traitement ici de la piéce et print la position de la piéce donné
            let position_piece = traitement(&player_input, taille_map, map, taille_piece, piece);
            // let position_piece = traitement(&player_input, taille_map, map, taille_piece, piece, &mut file);
            // Écrit l'emplacement de la piece dans le fichier de log
            // writeln!(file, "\nPiece place à l'emplacement : {}", position_piece).expect("Error 5: Unable to write to file");

            // Print du résultat pour le robot
            println!("{}", position_piece);

            // number_loop += 1;
        }
    }
}

// fn traitement(player_input: &String, taille_map: (i32, i32), map: Vec<String>, taille_piece: (i32, i32), piece: Vec<String>, file: &mut File) -> String {
fn traitement(player_input: &String, taille_map: (i32, i32), map: Vec<String>, taille_piece: (i32, i32), piece: Vec<String>) -> String {
    let new_piece: &str;
    let piece_placer: &str;
    if player_input.contains("p1") {
        new_piece = "a";
        piece_placer = "@";
    } else {
        new_piece = "s";
        piece_placer = "$";
    }
    let mut emplacement_depart_p1: (usize, usize) = (0, 0);
    let mut emplacement_depart_p2: (usize, usize) = (0, 0);

    // création de la matrice de traitement
    let mut matrix: Vec<Vec<String>> = vec![];
    let mut list_emplacement_map: Vec<(usize, usize)> = vec![];
    for (index_line, current_line_map) in map.iter().enumerate() {
        let mut line: Vec<String> = vec![];
        for (index_char, char) in current_line_map.chars().enumerate() {
            let current_string_map = char.to_string();
            line.push(current_string_map.clone());
            if current_string_map == piece_placer || current_string_map == new_piece {
                list_emplacement_map.push((index_line, index_char)); // (y, x)
            }

            // récupération de la position de depart de P1
            if current_string_map == "@" {
                emplacement_depart_p1 = (index_char, index_line); // (x, y)
            }

            // récupération de la position de depart de P2
            if current_string_map == "$" {
                emplacement_depart_p2 = (index_char, index_line); // (x, y)
            }
        }
        matrix.push(line);
    }

    // Récupération de la liste des emplacements "O" de la piéce
    let mut list_emplacement_piece: Vec<(usize, usize)> = vec![];
    for (index_line, current_line_piece) in piece.iter().enumerate() {
        for (index_char, char) in current_line_piece.chars().enumerate() {
            if char == 'O' {
                list_emplacement_piece.push((index_line, index_char));
            }
        }
    }

    // writeln!(file, "list_emplacement_piece : {:?}", list_emplacement_piece).expect("Error 5: Unable to write to file");

    // Test de placement de la piece
    let mut list_emplacement_possible_piece: Vec<(usize, usize)> = vec![];
    for emplacement_map in &list_emplacement_map {
        // let mut trouve = false;

        for emplacement_piece in &list_emplacement_piece {
            
            // Calcul des positions d'origine de la piece sur la map
            let offset_x = emplacement_map.1 as i32 - emplacement_piece.1 as i32;
            let offset_y = emplacement_map.0 as i32 - emplacement_piece.0 as i32;
            // writeln!(file, "offset : {}, {}", offset_x, offset_y).expect("Error 5: Unable to write to file");
 
            // Vérification si la piece rentre dans l'emplacement sans sortir de la map
            if offset_x >= 0 && offset_y >= 0 &&
               offset_x + taille_piece.0 <= taille_map.0 && offset_y + taille_piece.1 <= taille_map.1 {

                let mut possible = true;
                
                // Test des positions
                for current_emplacement_piece in &list_emplacement_piece {
                    if emplacement_piece != current_emplacement_piece {
                        let current_x_map = (offset_x + current_emplacement_piece.1 as i32) as usize;
                        let current_y_map = (offset_y + current_emplacement_piece.0 as i32) as usize;

                        if matrix[current_y_map][current_x_map] != "." {
                            possible = false;
                            break;
                        }
                    }
                }
                
                // enregistrement des emplacements possible
                if possible {
                    list_emplacement_possible_piece.push((offset_x as usize, offset_y as usize));
                    // writeln!(file, "list_emplacement_possible_piece : {:?}", list_emplacement_possible_piece).expect("Error 5: Unable to write to file");
                }
            }
        }
    }
    
    // envoi du dernier emplacement possible
    if list_emplacement_possible_piece.is_empty() {
        // writeln!(file, "list_emplacement_possible_piece est VIDE !!!").expect("Error 5: Unable to write to file");
        
        return format!("{} {}", 0, 0);
    } else {
        // writeln!(file, "list_emplacement_possible_piece : {:?}", list_emplacement_possible_piece).expect("Error 5: Unable to write to file");

        let taille_list = list_emplacement_possible_piece.len() -1;

        // fonction de la position des joueurs
        if new_piece == "a" { // je suis le player 1
            if emplacement_depart_p1.1 < emplacement_depart_p2.1 { // p1 en haut
                let futur_position = list_emplacement_possible_piece[taille_list];
                return format!("{} {}", futur_position.0, futur_position.1);
            } else { // p1 en bas  
                let futur_position = list_emplacement_possible_piece[0];
                return format!("{} {}", futur_position.0, futur_position.1);          
            }
        } else { // je suis le player 2
            if emplacement_depart_p1.1 < emplacement_depart_p2.1 { // p1 en haut
                let futur_position = list_emplacement_possible_piece[0];
                return format!("{} {}", futur_position.0, futur_position.1);           
            } else { // p1 en bas     
                let futur_position = list_emplacement_possible_piece[taille_list];
                return format!("{} {}", futur_position.0, futur_position.1);      
            }
        }
    }
}