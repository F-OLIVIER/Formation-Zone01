use crate::connexion_gestion::connexion_gestion::path_server;
use lazy_static::lazy_static;
use std::fs::File;
use std::io::{BufRead, BufReader};

lazy_static! {
    pub static ref PATH_SERVER: String = path_server();
    pub static ref RETURN_BUTTON: String = "<button onclick=\"window.history.back()\" style=\"background-color: #4CAF50; color: white; padding: 10px 20px; border: none; border-radius: 5px; cursor: pointer;\">Retour</button>".to_string();

}

pub const MAX_EVENTS: usize = 1024;

// -----------------------Fonction pour lire le fichier de configuration-----------------------

pub fn read_config(file_path: &str) -> Vec<(String, u16, String)> {
    let mut entries = Vec::new(); // Vecteur pour stocker les tuples (nom de domaine, port, adresse IP)

    // Tente d'ouvrir le fichier, retourne un vecteur vide en cas d'erreur
    let file = match File::open(file_path) {
        Ok(f) => f,
        Err(_) => {
            eprintln!("Erreur: Impossible d'ouvrir le fichier {}", file_path);
            return entries; // Retourne un vecteur vide
        }
    };

    let reader = BufReader::new(file);

    // Lecture ligne par ligne
    for line in reader.lines() {
        if let Ok(line) = line {
            let parts: Vec<&str> = line.split_whitespace().collect();

            // Vérifie que la ligne contient au moins 2 éléments
            if parts.len() == 1 {
                let ip_port = parts[0];
                
                // Extraction de l'adresse IP et du port
                let parts: Vec<&str> = ip_port.split(':').collect();
                if parts.len() == 2 {
                    let ip = parts[0].to_string();
                    let port_str = parts[1];

                    // Conversion du port en u16
                    if let Ok(port) = port_str.parse::<u16>() {
                        // Ajout au vecteur
                        entries.push(("".to_string(), port, ip));
                    } else {
                        eprintln!("Erreur lors de la conversion du port : {}", port_str);
                    }
                }
            } else if parts.len() == 2 {
                let ip_port = parts[0];
                let domain_name = parts[1];

                // Extraction de l'adresse IP et du port
                let parts: Vec<&str> = ip_port.split(':').collect();
                if parts.len() == 2 {
                    let ip = parts[0].to_string();
                    let port_str = parts[1];

                    // Conversion du port en u16
                    if let Ok(port) = port_str.parse::<u16>() {
                        // Ajout au vecteur
                        entries.push((domain_name.to_string(), port, ip));
                    } else {
                        eprintln!("Erreur lors de la conversion du port : {}", port_str);
                    }
                }
            } else {
                eprintln!("Ligne invalide dans le fichier de configuration : {}", line);
            }
        }
    }

    println!("List entries : {:?}", entries);
    entries // Retourne le vecteur des tuples
}
