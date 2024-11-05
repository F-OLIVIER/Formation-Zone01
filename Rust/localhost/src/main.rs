pub mod config;
pub mod connexion_gestion;
pub mod handle;

use crate::connexion_gestion::connexion_gestion::*;
use config::config::*;
use log::{error, info};
use std::net::TcpListener;
use std::process::{exit, Command};

fn main() {
    info!("Votre mot de passe machine est nécessaire afin de mettre à jour le fichier système '/etc/hosts'");

    // Exécuter le script .sh
    exec_script_hosts_domain();

    env_logger::init();
    let config_path = format!("{}/src/config/config.txt", *PATH_SERVER);
    let configs = read_config(&config_path);
    // println!("configs : {:?}", configs);

    let mut listeners = Vec::new();

    // Créer un listener pour chaque configuration
    for (domain_name, port, ip) in configs {
        let address = format!("{}:{}", ip, port);
        // println!("address : {:?}", address);

        match TcpListener::bind(&address) {
            Ok(listener) => {
                if domain_name.is_empty() {
                    println!("Serveur démarré sur http://{} sans domaine", address);
                } else {
                    println!(
                        "Serveur démarré sur http://{} pour le domaine http://{}:{}",
                        address, domain_name, port
                    );
                }
                listeners.push(listener); // Conserver les listeners
            }
            Err(e) => {
                eprintln!("Erreur lors de la liaison à l'adresse {}: {}", address, e);
            }
        }
    }

    // Assurer que nous avons des listeners avant de continuer
    if listeners.is_empty() {
        eprintln!("Aucun listener n'a été créé. Fermeture du serveur.");
        exit(1);
    }

    #[cfg(target_os = "macos")]
    run_kqueue_macos(listeners); // Passer tous les listeners à kqueue

    #[cfg(target_os = "linux")]
    run_epoll_linux(listeners); // Passer tous les listeners à epoll
}


// ---------------------Fonction pour exécuter le script setup_hosts_file.sh---------------------

fn exec_script_hosts_domain() {
    let script_output = Command::new("sudo")
        .arg("./setup_hosts_file.sh")
        .output()
        .expect("Échec de l'exécution du script");

    if script_output.status.success() {
        info!("Script exécuté avec succès.");
    } else {
        error!(
            "Échec du script. Erreur : {}",
            String::from_utf8_lossy(&script_output.stderr)
        );
        exit(1); // Arrêter le programme
    }
}
