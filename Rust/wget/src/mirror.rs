use regex::Regex;
use reqwest::blocking::get;
use scraper::{Html, Selector};
use std::{
    fs::{create_dir_all, File},
    io::{self, Read, Write},
    path::Path,
};
use url::Url;

pub fn download_and_mirror(
    url: &str,
    reject_ext: &str,
    reject_dir: &str,
    convert_links: bool,
) -> Result<(), Box<dyn std::error::Error>> {
    // Télécharger la page HTML
    let response = get(url)?;
    if !response.status().is_success() {
        return Err("Failed to fetch the page".into());
    }

    let body = response.text()?;

    // Convertir les liens dans le HTML (si convert-links present)
    let body = if convert_links {
        convert_links_in_html(&body, url)
    } else {
        body
    };

    let document = Html::parse_document(&body);

    // Extraire les liens classiques et les images
    let link_selector = Selector::parse("a[href], link[href], img[src], script[src]").unwrap();
    let style_selector = Selector::parse("style").unwrap();

    // Recherche et Collecte des liens comme dans les balises (a, img, etc.)
    let mut links_to_download = Vec::new();
    for element in document.select(&link_selector) {
        if let Some(link) = element
            .value()
            .attr("href")
            .or_else(|| element.value().attr("src"))
        {
            links_to_download.push(link.to_string());
        }
    }

    // Extraire les images depuis les balises <style> (image du css)
    for element in document.select(&style_selector) {
        if let Some(style_content) = element.text().next() {
            let re = Regex::new("url\\(['\\\"]?([^'\\\")]+)['\\\"]?\\)").unwrap();
            for cap in re.captures_iter(style_content) {
                let url = cap[1].to_string();
                links_to_download.push(url);
            }
        }
    }

    // Créer le répertoire du domaine
    let parsed_url = Url::parse(url)?;
    let domain = parsed_url.host_str().unwrap_or("unknown");
    let dir_path = Path::new(domain);

    if dir_path.exists() && dir_path.is_file() {
        return Err("Destination path is a file, not a directory".into());
    }

    create_dir_all(&dir_path)?;

    // Sauvegarde le fichier HTML principal (index.html)
    let html_file = dir_path.join("index.html");
    let mut file = File::create(&html_file)?;
    file.write_all(body.as_bytes())?;

    // Téléchargement et tri des ressources
    for link in links_to_download {
        let absolute_url = convert_link(&link, url);

        // Vérification si dossiers rejetés
        if !reject_dir.is_empty() {
            let reject_dirs: Vec<&str> = reject_dir.split(',').collect();
            if reject_dirs.iter().any(|dir| absolute_url.contains(dir)) {
                println!("Skipping resource in rejected directory: {}", absolute_url);
                continue;
            }
        }

        let file_name = get_file_name_from_url(&absolute_url);
        // Vérifier si le fichier est `index.html` et ignorer pour eviter les doublons
        if file_name == "index.html" {
            continue;
        }

        // Vérifier si l'extension de fichier est rejetée
        let current_extension = file_name.split('.').last();
        if reject_ext != "" && current_extension == Some(reject_ext) {
            continue;
        }

        // Déterminer le sous-dossier cible (basé sur l'extension)
        let extension = file_name.split('.').last().unwrap_or("").to_lowercase();
        let sub_dir = match extension.as_str() {
            "css" => "css",
            "js" => "js",
            "png" | "jpg" | "jpeg" | "gif" | "bmp" | "svg" => "img",
            _ => "template", // Fichiers sans extension ou autres extensions
        };

        // Déterminer le dossier cible
        let sub_dir_path = dir_path.join(sub_dir);
        // Créer le dossier si nécessaire
        create_dir_all(&sub_dir_path)?;
        // Chemin final
        let file_path = sub_dir_path.join(&file_name);

        // Vérifier si le fichier existe déjà
        if file_path.exists() {
            eprintln!("File already exists, skipping download: {:?}", file_path);
            continue;
        }

        println!("Téléchargement de : {}", absolute_url);

        // Téléchargement du fichier
        let file_response = get(&absolute_url)?;
        if file_response.status().is_success() {
            let mut file = File::create(&file_path)?;
            io::copy(&mut file_response.take(10_000_000), &mut file)?;
        } else {
            eprintln!("Failed to download: {}", absolute_url);
        }
    }

    Ok(())
}

// Fonction pour obtenir le nom de fichier à partir de l'URL
fn get_file_name_from_url(url: &str) -> String {
    let parsed_url = Url::parse(url).expect("Invalid URL");
    let path = parsed_url.path();

    if path.ends_with("/") {
        "index.html".to_string()
    } else {
        path.split('/').last().unwrap_or("default.html").to_string()
    }
}

// Fonction pour convertir les liens relatifs dans le HTML
fn convert_links_in_html(body: &str, base_url: &str) -> String {
    let document = Html::parse_document(body);
    let link_selector = Selector::parse("a[href], link[href], img[src], script[src]").unwrap();
    let style_selector = Selector::parse("style").unwrap();

    let mut updated_body = body.to_string();

    for element in document.select(&link_selector) {
        if let Some(link) = element
            .value()
            .attr("href")
            .or_else(|| element.value().attr("src"))
        {
            let absolute_link = convert_link(link, base_url);
            updated_body = updated_body.replace(link, &absolute_link);
        }
    }

    for element in document.select(&style_selector) {
        if let Some(style_content) = element.text().next() {
            let re = Regex::new("url\\(['\\\"]?([^'\\\")]+)['\\\"]?\\)").unwrap();
            let modified_style_content = re.replace_all(style_content, |caps: &regex::Captures| {
                let url = &caps[1];
                if url.starts_with("/") {
                    return format!("url('./{}')", &url[1..]);
                }
                format!("url('{}')", url)
            });

            updated_body = updated_body.replace(style_content, &modified_style_content);
        }
    }

    updated_body
}

// Fonction pour convertir un lien en lien absolu
fn convert_link(link: &str, base_url: &str) -> String {
    let base = Url::parse(base_url).expect("Invalid base URL");
    let resolved_url = base.join(link).expect("Failed to resolve URL");
    resolved_url.to_string()
}
