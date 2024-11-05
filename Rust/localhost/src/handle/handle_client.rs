use crate::config;
use chrono::{Duration, Utc};
use std::io::{self, Read, Write};
use std::net::TcpStream;
use std::process::{Command, Output};
use std::{fs, usize};
use urlencoding::decode; // Ajouter cette ligne pour le décodage des données URL-encodées

// ---------------------Fonction pour gérer les clients---------------------

pub fn handle_client(stream: &mut TcpStream) {
    let mut buffer = [0; 1024]; // Buffer de 1 Ko pour recevoir la requête
                                // Lire la requête du client
    match stream.read(&mut buffer) {
        Ok(_) => {
            // Convertir la requête en String pour analyse
            let request = String::from_utf8_lossy(&buffer[..]);
            // Traiter la requête HTTP
            handle_http_request(request.as_ref(), stream);
        }
        Err(ref e) if e.kind() == std::io::ErrorKind::WouldBlock => {
            println!("WouldBlock error: {:?}", e);
        }
        Err(e) => eprintln!("Erreur lors de la lecture de la requête du client: {}", e),
    }
}

// ---------------------Fonction pour analyser les requêtes HTTP---------------------

fn parse_http_request(request: &str) -> Option<(String, String, String, Option<String>)> {
    let mut lines = request.lines();

    // Definition des variables
    let (mut method, mut path, mut version) = (String::new(), String::new(), String::new());
    let mut domain = None;

    // La première ligne est la ligne de requête (ex: "GET /index.html HTTP/1.1")
    if let Some(request_line) = lines.next() {
        let parts: Vec<&str> = request_line.split_whitespace().collect();
        if parts.len() == 3 {
            method = parts[0].to_string();
            path = parts[1].to_string();
            version = parts[2].to_string();
        }
    }

    // Parcourir les lignes suivantes pour trouver l'en-tête Host optionnel
    for line in lines {
        if line.starts_with("Host:") {
            // Extraire le domaine en supprimant "Host: "
            domain = Some(line[5..].trim().to_string());
            break;
        }
    }

    // Vérifie si les éléments principaux de la requête sont présents
    if !method.is_empty() && !path.is_empty() && !version.is_empty() {
        Some((method, path, version, domain))
    } else {
        None
    }
}

// ---------------------Fonction pour traiter les requêtes HTTP---------------------

fn handle_http_request(request: &str, client: &mut TcpStream) {
    if let Some((method, path, _version, domain)) = parse_http_request(request) {
        println!("Traitement request");
        if let Some(domain_name) = domain {
            println!("Domain détecté : {}", domain_name);
        }

        if method == "GET" {
            if path.starts_with("/ressources") {
                let resource_path = &path[11..]; // Retire "/ressources"
                handle_list_resources(client, resource_path);
                return; // Quitter ici car nous avons déjà envoyé la réponse
            }
            let file_path = if path == "/" {
                String::from("static/index.html")
            } else {
                format!("static/{}", &path[1..])
            };
            traitement_get(file_path.as_str(), client);
            println!("GET request for path: {}", path);
        } else if method == "POST" && path == "/delete-script" {
            println!("POST request for path: {}", path);

            if path == "/delete-script" {
                // Séparer les en-têtes et le corps de la requête
                let body = request.split("\r\n\r\n").nth(1).unwrap_or("");
                let params_with_caract = body.replace('\0', "").replace('+', " ");
                let decoded_body = decode(&params_with_caract).expect("Failed to decode body");

                // Extraire le nom du fichier à supprimer

                let params: Vec<&str> = decoded_body.split('&').collect();
                let mut delete_filename = "";

                for param in params {
                    if param.starts_with("delete-filename=") {
                        delete_filename = &param["delete-filename=".len()..];
                    }
                }

                // Vérifier si le nom du fichier est valide
                if delete_filename.is_empty() {
                    let error_response = format!("HTTP/1.1 400 BAD REQUEST\r\n\r\n<h1>400 Bad Request: Missing filename</h1> {}", *config::config::RETURN_BUTTON);

                    client.write_all(error_response.as_bytes()).unwrap();
                    return;
                }

                // Tenter de supprimer le fichier
                let result = fs::remove_file(format!("ressources/{}", delete_filename));
                match result {
                    Ok(_) => {
                        let response = format!(
                            "HTTP/1.1 200 OK\r\n\r\n<h1>File deleted successfully</h1> {}",
                            *config::config::RETURN_BUTTON
                        );
                        client.write_all(response.as_bytes()).unwrap();
                    }
                    Err(e) => {
                        eprintln!("Error deleting file: {}", e);
                        let error_response = format!("HTTP/1.1 500 INTERNAL SERVER ERROR\r\n\r\n<h1>500 Internal Server Error: Could not delete: file not exist file</h1> {}", *config::config::RETURN_BUTTON);
                        client.write_all(error_response.as_bytes()).unwrap();
                    }
                }
            } else {
                send_error_response(404, client);
            }
        } else if method == "POST" {
            if path == "/submit-script" {
                // Séparer les en-têtes et le corps de la requête
                let body = request.split("\r\n\r\n").nth(1).unwrap_or("");

                // Remplacer les "+" par des espaces car en http les espaces sont remplacer par des "+" (les vrai + ne sont pas impacté car encodé en "%2B")
                let body_with_spaces = body.replace('+', " ");

                // Décoder le corps de la requête
                let decoded_body = decode(&body_with_spaces).expect("Failed to decode body");

                // Extraire les différents paramètres du corps de la requête
                let params: Vec<&str> = decoded_body.split('&').collect();
                let mut filename = "";
                let mut extension = "";
                let mut script_content = "".to_string();

                // Extraire les valeurs des paramètres
                for param in params {
                    if param.starts_with("filename=") && !param.contains("/") {
                        filename = &param["filename=".len()..];
                    } else if param.starts_with("extension=") {
                        extension = &param["extension=".len()..];
                    } else if param.starts_with("script_content=") {
                        let script_brut = &param["script_content=".len()..];

                        // Nettoyer le contenu du script pour supprimer les bytes nuls
                        script_content = script_brut
                            .replace("\u{0}", "") // Supprime les bytes nuls
                            .replace("\r", ""); // Remplace les retours chariot (CR) par rien
                    }
                }

                // Vérifier si le nom du fichier et l'extension sont valides
                if filename.is_empty() || extension.is_empty() {
                    let error_response = format!("HTTP/1.1 400 BAD REQUEST\r\n\r\n<h1>400 Bad Request: Missing filename or extension</h1> {}", *config::config::RETURN_BUTTON);
                    client.write_all(error_response.as_bytes()).unwrap();
                    return;
                }

                // Construire le nom complet du fichier avec l'extension
                let script_filename = format!("ressources/{}{}", filename, extension);
                println!("Script filename: {}", script_filename);

                if script_filename.ends_with(".py") || script_filename.ends_with(".php") {
                    // Écrire le contenu du script dans le fichier spécifié
                    println!("Script file written: {}", script_filename);
                    fs::write(&script_filename, script_content)
                        .expect("Unable to write script file");
                    // Exécuter le script soumis
                    handle_cgi(&script_filename, client);
                } else {
                    let error_response = format!("HTTP/1.1 400 BAD REQUEST\r\n\r\n<h1>400 Bad Request: Not Autorized script extension</h1> {}", *config::config::RETURN_BUTTON);
                    client.write_all(error_response.as_bytes()).unwrap();
                }
            } else {
                send_error_response(405, client); // Méthode non autorisée
            }
        } else if method == "DELETE" {
            if path.starts_with("/delete-script") {
                // Extraire le nom du fichier à supprimer
                let params = path.split('?').nth(1).unwrap_or("").split('&');
                let mut delete_filename = "";

                for param in params {
                    if param.starts_with("delete-filename=") {
                        delete_filename = &param["delete-filename=".len()..];
                    }
                }

                // Vérifier si le nom du fichier est valide
                if delete_filename.is_empty() {
                    let error_response = format!("HTTP/1.1 400 BAD REQUEST\r\n\r\n<h1>400 Bad Request: Missing filename</h1> {}", *config::config::RETURN_BUTTON);
                    client.write_all(error_response.as_bytes()).unwrap();
                    return;
                }

                // Tenter de supprimer le fichier
                let result = fs::remove_file(format!("ressources/{}", delete_filename));
                match result {
                    Ok(_) => {
                        let response = format!(
                            "HTTP/1.1 200 OK\r\n\r\n<h1>File deleted successfully</h1> {}",
                            *config::config::RETURN_BUTTON
                        );
                        client.write_all(response.as_bytes()).unwrap();
                    }
                    Err(e) => {
                        eprintln!("Error deleting file: {}", e);
                        let error_response = format!("HTTP/1.1 500 INTERNAL SERVER ERROR\r\n\r\n<h1>500 Internal Server Error: Could not delete: file not exist file</h1> {}", *config::config::RETURN_BUTTON);
                        client.write_all(error_response.as_bytes()).unwrap();
                    }
                }
            } else {
                send_error_response(404, client);
            }
        } else {
            send_error_response(405, client);
        }
    } else {
        send_error_response(400, client);
    }
}

// ---------------------Fonction pour gérer les ressources---------------------

fn handle_list_resources(client: &mut TcpStream, path: &str) {
    let resource_dir = format!("ressources/{}", path);
    let entries = fs::read_dir(&resource_dir);

    let mut file_list = String::new();

    // En-tête de la page
    file_list.push_str("<h1>Liste des fichiers</h1><ul>");

    // Lien pour remonter d'un niveau
    if path != "" {
        let parent_path = path.rsplit('/').nth(1).unwrap_or("");
        file_list.push_str(&format!(
            "<li><a href=\"/ressources/{}\">..</a></li>",
            parent_path
        ));
    }

    // Lister les fichiers et répertoires
    match entries {
        Ok(entries) => {
            for entry in entries {
                if let Ok(entry) = entry {
                    let file_name = entry.file_name().into_string().unwrap();
                    let entry_path = format!("{}/{}", path, file_name);

                    if entry.file_type().unwrap().is_dir() {
                        // Lien vers le répertoire
                        file_list.push_str(&format!(
                            "<li><a href=\"/ressources{}\">{}/</a></li>",
                            entry_path, file_name
                        ));
                    } else {
                        // Lien vers le fichier
                        file_list.push_str(&format!("<li>{}</li>", file_name));
                    }
                }
            }
        }
        Err(e) => {
            eprintln!("Erreur en lisant le dossier: {}", e);
            let error_response =
                "HTTP/1.1 500 INTERNAL SERVER ERROR\r\n\r\n<h1>500 Internal Server Error</h1>";
            client.write_all(error_response.as_bytes()).unwrap();
            return;
        }
    }

    file_list.push_str("</ul>");
    file_list.push_str(&config::config::RETURN_BUTTON);
    let response = format!(
        "HTTP/1.1 200 OK\r\nContent-Length: {}\r\nContent-Type: text/html\r\n\r\n{}",
        file_list.len(),
        file_list,
    );
    client.write_all(response.as_bytes()).unwrap();
}

// ---------------------Fonction pour envoyer une réponse d'erreur---------------------

fn send_error_response(code: usize, client: &mut TcpStream) {
    let (status_error, message) = match code {
        400 => (
            "400 BAD REQUEST",
            format!(
                "<h1>400 Bad Request</h1> {}",
                *config::config::RETURN_BUTTON
            ),
        ),
        403 => (
            "403 FORBIDDEN",
            format!("<h1>403 Forbidden</h1> {}", *config::config::RETURN_BUTTON),
        ),
        404 => (
            "404 NOT FOUND",
            format!("<h1>404 Not Found</h1> {}", *config::config::RETURN_BUTTON),
        ),
        405 => (
            "405 METHOD NOT ALLOWED",
            format!(
                "<h1>405 Method Not Allowed</h1> {}",
                *config::config::RETURN_BUTTON
            ),
        ),
        413 => (
            "413 PAYLOAD TOO LARGE",
            format!(
                "<h1>413 Payload Too Large</h1> {}",
                *config::config::RETURN_BUTTON
            ),
        ),
        500 => (
            "500 INTERNAL SERVER ERROR",
            format!(
                "<h1>500 Internal error server</h1> {}",
                *config::config::RETURN_BUTTON
            ),
        ),
        _ => (
            "000 UNKNOWN ERROR",
            format!(
                "<h1>000 Unknown error</h1> {}",
                *config::config::RETURN_BUTTON
            ),
        ),
    };

    let response = format!(
        "HTTP/1.1 {}\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n{}",
        status_error, message
    );

    if let Err(e) = client.write(response.as_bytes()) {
        eprintln!("500 Internal error server: {}", e);
    }
}

// ---------------------Fonction pour gérer les scripts CGI---------------------

fn handle_cgi(path: &str, client: &mut TcpStream) {
    // Déterminer quel interpréteur utiliser selon l'extension
    if path.ends_with(".py") {
        execute_script("python3", path, client);
    } else if path.ends_with(".php") {
        execute_script("php", path, client);
    } else {
        let error_500 =
            "HTTP/1.1 500 INTERNAL SERVER ERROR\r\n\r\n<h1>500 Internal Server Error dans tes dennnnnnts</h1>";
        client.write_all(error_500.as_bytes()).unwrap();
    }
}

// Fonction pour envoyer une réponse HTTP
fn send_response(client: &mut TcpStream, output: Output) {
    let response = format!(
        "HTTP/1.1 200 OK\r\nContent-Length: {}\r\n\r\n{}",
        output.stdout.len(),
        String::from_utf8_lossy(&output.stdout),
    );

    client.write_all(response.as_bytes()).unwrap();
}

// --------------------Fonction pour gérer l'exécution des scripts------------------

fn execute_script(command: &str, path: &str, client: &mut TcpStream) {
    println!("Executing script: {}", path);
    match Command::new(command).arg(path).output() {
        Ok(output) => {
            let out_output = String::from_utf8_lossy(&output.stdout);
            let err_output = String::from_utf8_lossy(&output.stderr);
            // Affiche la sortie du script dans le terminal
            println!("Script output: {}", out_output);
            println!("Script error output: {}", err_output);
            if !err_output.is_empty() {
                let error_400 =
                    "HTTP/1.1 400 BAD REQUEST\r\n\r\n<h1>400 Bad Request: incorrect script</h1>";
                client.write_all(error_400.as_bytes()).unwrap();
                return;
            }
            // Création d'un objet Output
            let output_result = Output {
                status: output.status,
                stdout: output.stdout,
                stderr: output.stderr,
            };
            // Appel de send_response après avoir créé l'Output
            send_response(client, output_result);
        }
        Err(e) => {
            eprintln!("Failed to execute command: {}", e);
            let error_500 =
                "HTTP/1.1 500 INTERNAL SERVER ERROR\r\n\r\n<h1>500 Internal Server Error</h1>";
            client.write_all(error_500.as_bytes()).unwrap();
        }
    }
}

// ------------------Fonction pour traiter les requêtes GET------------------

pub fn traitement_get(file_path: &str, client: &mut TcpStream) {
    // Vérifie si le fichier existe avant d'essayer de le lire ou de le canoniser
    if fs::metadata(file_path).is_ok() {
        // Convertit le chemin en chemin canonique (propre et sécurisé)
        if let Ok(canonical_path) = fs::canonicalize(file_path) {
            // Vérifie que le chemin canonique commence par "static"
            if canonical_path.starts_with(format!("{}/static/", *config::config::PATH_SERVER)) {
                // Tente de lire le fichier sous forme de chaîne de caractères.
                match fs::read_to_string(canonical_path) {
                    // Si la lecture du fichier réussit, construction de la réponse HTTP avec le statut 200 (OK).
                    Ok(contents) => {
                        // Création du cookies
                        let cookie_name = "mycookies";
                        let cookie_value = "123456";
                        let now = Utc::now();
                        let expires = now + Duration::hours(1);

                        let response = format!(
                            "HTTP/1.1 200 OK\r\n\
                            Set-Cookie: {}={}; Expires={}; Path=/; HttpOnly\r\n\
                            Content-Length: {}\r\n\r\n\
                            {}",
                            cookie_name,
                            cookie_value,
                            expires,
                            contents.len(),
                            contents
                        );
                        // Essaie d'envoyer la réponse au client.
                        if let Err(e) = client.write(response.as_bytes()) {
                            eprintln!("500 Internal server error: {}", e);
                        }
                    }

                    Err(e) => {
                        if e.kind() == io::ErrorKind::PermissionDenied {
                            // Si l'accès au fichier est refusé (problème de permissions), erreur 403 (Forbidden).
                            send_error_response(403, client);
                        } else {
                            // Pour toute autre erreur (problème d'entrée/sortie non prévu), erreur 500 (Internal Server Error).
                            send_error_response(500, client);
                        }
                    }
                }
            } else {
                // Si le chemin ne commence pas par "static", cela signifie que l'utilisateur essaie d'accéder à un fichier en dehors
                // du répertoire autorisé. Dans ce cas, on renvoie une erreur 403 (Forbidden) pour interdire cet accès.
                send_error_response(403, client);
            }
        } else {
            // Problème avec le chemin
            send_error_response(500, client);
        }
    } else {
        // Fichier non trouvé
        send_error_response(404, client);
    }
}
