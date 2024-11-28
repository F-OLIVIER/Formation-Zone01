use chrono::Local;
use indicatif::{ProgressBar, ProgressStyle};
use reqwest::Client;
use std::{
    error::Error,
    fs::OpenOptions,
    fs::{self, File},
    io,
    io::Write,
    path::PathBuf,
};
use tokio::time::{sleep, Duration};

pub async fn download_file(
    url: &str,
    output: Option<String>,
    path: Option<String>,
    rate_limit: Option<u64>,
    log_file: bool,
) -> Result<(), Box<dyn std::error::Error>> {
    // Affichage de l'heure de début
    let start_time = Local::now();

    match write_log(
        log_file,
        &format!(
            "───────────────────────────────────────────────\nstart at {}\nURL: {}",
            start_time.format("%Y-%m-%d %H:%M:%S"),
            url
        ),
    ) {
        Ok(_) => {}
        Err(e) => eprintln!("Failed to write start download : {}", e),
    }

    // Création d'un client HTTP
    let client = Client::new();

    // Téléchargement du fichier
    let mut response = client.get(url).send().await?;

    // Vérification du contenu de la réponse
    let content_length = response.content_length();
    let total_size = content_length.unwrap_or(0); // Taille totale du fichier
    let size_in_mb = total_size as f64 / (1024.0 * 1024.0);
    let size_in_kb = total_size as f64 / 1024.0;

    // Affichage du status
    match write_log(
        log_file,
        &format!(
            "sending request, awaiting response... status {}",
            response.status()
        ),
    ) {
        Ok(_) => {}
        Err(e) => eprintln!("Failed to write status : {}", e),
    }

    // Affichage de la taille du fichier
    match write_log(
        log_file,
        &format!("content size: {:.2} [{:.2}MB]", size_in_kb, size_in_mb),
    ) {
        Ok(_) => {}
        Err(e) => eprintln!("Failed to write size : {}", e),
    }

    // Extraction du nom de fichier à partir de l'URL si non spécifié
    let filename = output.unwrap_or_else(|| {
        url.split('/')
            .last()
            .unwrap_or("downloaded_file")
            .to_string()
    });

    // Expansion du chemin spécifié
    let save_dir = path
        .map(|p| expand_tilde(&p))
        .unwrap_or_else(|| PathBuf::from("."));
    fs::create_dir_all(&save_dir)?; // Création du répertoire si nécessaire

    let full_path = save_dir.join(filename.clone());
    match write_log(
        log_file,
        &format!("saving file to: {}", full_path.display()),
    ) {
        Ok(_) => {}
        Err(e) => eprintln!("Failed to write saving file : {}", e),
    }

    // Ouvrerture du fichier pour écrire les données téléchargées
    let mut file = File::create(full_path.clone())?;

    // Création de la barre de progression
    let progress_bar = init_progress_bar(total_size);

    // Gestion du taux de téléchargement
    let mut total_downloaded = 0;

    // Lecture et écrirure par morceaux
    while let Some(bytes_read) = response.chunk().await? {
        let bytes = bytes_read.len();

        if bytes == 0 {
            break; // Si aucun byte
        }

        // Si la vitesse est limité
        if let Some(rate_limit) = rate_limit {
            if total_downloaded >= rate_limit * 1024 {
                sleep(Duration::from_secs(1)).await;
                total_downloaded = 0;
            }
        }

        // Écriture dans le fichier local
        file.write_all(&bytes_read)?;
        total_downloaded += bytes as u64;

        // Mise à jour de la barre de progression
        if !log_file {
            progress_bar.inc(bytes as u64);
        }
    }

    progress_bar.finish_and_clear();
    if !log_file {
        progress_bar.finish_with_message("Download completed");
    } else {
        // ecrire la progress_bar dans le fichier de log
        let progress_message = "Download completed".to_string();
        match write_log(log_file, &progress_message) {
            Ok(_) => {}
            Err(e) => eprintln!("Failed to write progress message: {}", e),
        }
    }

    // Affichage du message du fichier télécharger
    match write_log(log_file, &format!("Downloaded [{}]", url)) {
        Ok(_) => {}
        Err(e) => eprintln!("Failed to write url : {}", e),
    }

    // Affichage de la fin du téléchargement
    let end_time = Local::now();
    match write_log(
        log_file,
        &format!(
            "finished at {}\n───────────────────────────────────────────────",
            end_time.format("%Y-%m-%d %H:%M:%S")
        ),
    ) {
        Ok(_) => {}
        Err(e) => eprintln!("Failed to write stop timestamps : {}", e),
    }

    Ok(())
}

// Barre de progression
pub fn init_progress_bar(total_size: u64) -> ProgressBar {
    let pb = ProgressBar::new(total_size);
    pb.set_style(
        ProgressStyle::default_bar()
            .template(
                "{spinner:.green} {bytes:>7} / {total_bytes:7} [{bar:40.cyan/blue}] {percent:>3}% {bytes_per_sec} {eta}",
            )
            .unwrap()
            .progress_chars("=>-"),
    );
    pb
}

// Chemin complet fichier (gestion du ~)
fn expand_tilde(path: &str) -> PathBuf {
    if path.starts_with("~/") {
        if let Some(home_dir) = dirs_next::home_dir() {
            return home_dir.join(&path[2..]);
        }
    }
    PathBuf::from(path)
}

// parse du rate limit
pub fn parse_rate_limit(rate_limit: &str) -> Result<u64, String> {
    let rate_limit = rate_limit.trim().to_lowercase();
    if rate_limit.ends_with('k') {
        rate_limit[..rate_limit.len() - 1]
            .parse::<u64>()
            .map(|v| v * 1024)
            .map_err(|_| format!("Invalid rate limit: {}", rate_limit))
    } else if rate_limit.ends_with('m') {
        rate_limit[..rate_limit.len() - 1]
            .parse::<u64>()
            .map(|v| v * 1024 * 1024)
            .map_err(|_| format!("Invalid rate limit: {}", rate_limit))
    } else {
        rate_limit
            .parse::<u64>()
            .map_err(|_| format!("Invalid rate limit: {}", rate_limit))
    }
}

pub async fn download_from_file(
    input_file: &str,
    output: Option<String>,
    path: Option<String>,
    rate_limit: Option<u64>,
    log_file: bool,
) -> Result<(), Box<dyn Error>> {
    // Lire le contenu du fichier
    let urls = std::fs::read_to_string(input_file)?;

    // Parcourir chaque URL du fichier
    for url in urls.lines() {
        if url.trim().is_empty() {
            continue; // Ignorer les lignes vides
        }

        // Appeler la fonction de téléchargement pour chaque URL
        let result = download_file(
            url,
            output.clone(),
            path.clone(),
            rate_limit.clone(),
            log_file,
        )
        .await;

        if let Err(e) = result {
            eprintln!("Erreur pour {} : {}", url, e);
        }
    }

    Ok(())
}

// gestion de l'ecriture des informations
fn write_log(log_file: bool, message: &str) -> io::Result<()> {
    if log_file {
        // Ouvrir ou créer le fichier en mode ajout (append)
        let mut file = OpenOptions::new()
            .create(true) // Créer le fichier s'il n'existe pas
            .append(true) // Ajouter au fichier s'il existe déjà
            .open("wget-log")?; // Ouvrir le fichier

        // Écrire le message dans le fichier
        writeln!(file, "{}", message)?;
    } else {
        println!("{}", message);
    }

    Ok(())
}
