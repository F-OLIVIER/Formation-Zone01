mod download;
mod mirror;

use clap::{Arg, ArgAction, Command};
use download::*;
use mirror::*;
use std::error::Error;

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    // Définir les arguments de la ligne de commande
    let matches = Command::new("Téléchargement avec options")
        .version("1.0")
        .author("Votre nom <votre-email@example.com>")
        .about("Un programme de téléchargement")
        .arg(
            Arg::new("url") // nouvel argument
                .short('u') // Définit le flag court '-u'
                .long("url") // Définit le flag long '--url'
                .value_name("URL") // Spécifie le nom de la valeur attendue pour cet argument, ici "URL", qui sera utilisé dans l'aide pour décrire la valeur que l'utilisateur doit fournir.
                .help("URL à télécharger") // Ajoute une description de l'argument qui sera affichée dans l'aide du programme.
                .value_parser(clap::value_parser!(String)), // Utilise la macro `value_parser!` pour définir le type de valeur attendu pour cet argument.
        )
        .arg(
            Arg::new("output")
                .short('O')
                .long("output")
                .value_name("FICHIER")
                .help("Nom du fichier de sortie")
                .value_parser(clap::value_parser!(String)),
        )
        .arg(
            Arg::new("input-file")
                .short('i')
                .long("input")
                .value_name("FICHIER")
                .help("Fichier contenant les URLs à télécharger")
                .value_parser(clap::value_parser!(String)),
        )
        .arg(
            Arg::new("path")
                .short('P')
                .long("path")
                .value_name("CHEMIN")
                .help("Chemin de sauvegarde")
                .value_parser(clap::value_parser!(String)),
        )
        .arg(
            Arg::new("rate-limit")
                .long("rate-limit")
                .value_name("VITESSE")
                .help("Limite la vitesse de téléchargement (en Ko/s)")
                .value_parser(clap::builder::ValueParser::new(download::parse_rate_limit)),
        )
        .arg(
            Arg::new("log")
                .short('B')
                .long("log")
                .help("Sortie dans le fichier wget-log")
                .action(ArgAction::SetTrue),
        )
        .arg(
            Arg::new("mirror")
                .long("mirror")
                .help("Download and mirror the website")
                .action(ArgAction::SetTrue),
        )
        .arg(
            Arg::new("convert-links")
                .long("convert-links")
                .help("Convert links in the HTML to absolute")
                .action(ArgAction::SetTrue),
        )
        .arg(
            Arg::new("reject-ext")
                .long("reject")
                .value_name("REJECT")
                .help("Reject files with specified extensions (comma separated, e.g., gif,jpg)")
                .value_parser(clap::value_parser!(String)),
        )
        .arg(
            Arg::new("reject-dir")
                .long("reject-dir")
                .short('X')
                .value_name("DOSSIERS")
                .help("Rejects specified folders (e.g. /img or /css)")
                .value_parser(clap::value_parser!(String)),
        )
        .get_matches();

    // Récupérer le/les URL
    let url = matches.get_one::<String>("url");
    let input_file = matches.get_one::<String>("input-file");

    // Récupérer les options de la ligne de commande
    let output = matches.get_one::<String>("output");
    let path = matches.get_one::<String>("path");
    let rate_limit = matches.get_one::<u64>("rate-limit");
    let log_file =  matches.get_one::<bool>("log").unwrap_or(&false);
    let convert_links = matches.contains_id("convert-links");
    let mirror = matches.get_one::<bool>("mirror").unwrap_or(&false);
    println!("mirror : {}", mirror);
    if *mirror {
        let url_mirror = matches
            .get_one::<String>("url")
            .expect("URL must be provided")
            .as_str();
        // Récupération d'extension de fichier rejetés
        let reject_extention = matches
            .get_one::<String>("reject")
            .map_or("", |s| s.as_str());
        // Récupération de dossiers rejetés
        let reject_dir = matches
            .get_one::<String>("reject-dir")
            .map_or("", |s| s.as_str());
        if let Err(e) = download_and_mirror(url_mirror, reject_extention, reject_dir, convert_links)
        {
            eprintln!("Error: {}", e);
        }
    }

    if let Some(input_file) = input_file {
        // Télécharger à partir d'un fichier
        download_from_file(
            &input_file,
            output.cloned(),
            path.cloned(),
            rate_limit.cloned(),
            *log_file,
        )
        .await?;
    } else if let Some(url) = url {
        // Télécharger une seule URL
        let result = download::download_file(
            url,
            output.cloned(),
            path.cloned(),
            rate_limit.cloned(),
            *log_file,
        )
        .await;
        // Vérifier le résultat
        if let Err(e) = result {
            if e.to_string().contains("Is a directory (os error 21)") {
                // Ignorer cette erreur
            } else {
                eprintln!("Erreur : {}", e);
            }
            return Ok(());
        }
    } else {
        eprintln!("Aucune URL ou fichier d'entrée spécifié.");
    }

    if *log_file {
        println!("Output will be written to ‘wget-log’.")
    }

    Ok(())
}
