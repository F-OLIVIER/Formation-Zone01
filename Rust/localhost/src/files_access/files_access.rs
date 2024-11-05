use std::fs;
use std::path::Path;


#[derive(Serialize)]
struct Fichier {
    nom: String,
}

// --------------------Fonction pour lister les fichiers du répertoire----------------------

async fn lister_repertoire() -> impl Responder {
    let chemin = Path::new("./ressources"); // Le répertoire à lister
    let mut fichiers = Vec::new();

    if chemin.is_dir() {
        for entry in fs::read_dir(chemin).unwrap() {
            let entry = entry.unwrap();
            let path = entry.path();
            if let Some(nom_fichier) = path.file_name() {
                fichiers.push(Fichier {
                    nom: nom_fichier.to_string_lossy().into_owned(),
                });
            }
        }
    }

    web::Json(fichiers)  // Renvoie la liste des fichiers au format JSON
}
