# WGET

## 📝 Descriptif
Ce projet consiste à créer en `Rust` un programme qui reproduit le comportement de la commande `wget`.


### Options (flags) prises en charge :
- `-u` ou `--url` : URL à télécharger.
- `-B` ou `--log` : La sortie est effectuée dans le fichier `wget-log`.
- `-O` ou `--output` : Renommer le fichier de sortie.
- `-P` ou `--path` : Spécifier le chemin de sauvegarde du fichier téléchargé.
- `-i` ou `--input` : Télécharger tous les liens d'un fichier.
- `--rate-limit` : Limiter la vitesse de téléchargement (en Ko/s).
- `--mirror` : Télécharger l'intégralité d'un site web (scraping et mise en miroir).

### Option (flag) spécifiquement pour le miroir (scraping) :
- `--convert-links` : Convertir les liens dans le HTML en liens absolus.
- `--reject` : Rejeter des fichiers avec des extensions spécifiées (ex. : `gif, jpg`).
- `-X` ou `--reject-dir` : Rejeter des dossiers spécifiés (ex. : `/img`, `/css`).

___
## ⚙️ Installation & usage

### Compilez le projet avec `cargo` :
```sh
cargo build --release
```

### Exécutez le programme avec les options et l'URL de votre choix :
```sh
cargo run --release -- <url>
```

__Exemple d'utilisation :__
```sh
cargo run --release -- --url https://pbs.twimg.com/media/EMtmPFLWkAA8CIS.jpg --output "image.jpg" --path "./downloads" --rate-limit 500 --log
```
___
## 🔗 Dépendances

Le programme utilise `Rust` version 1.79.0 avec [cargo](https://www.rust-lang.org/fr), ainsi que les bibliothèques suivantes :

- [`reqwest`](https://docs.rs/reqwest/latest/reqwest/) version `0.12` - Pour effectuer les requêtes HTTP avec TLS.
- [`indicatif`](https://docs.rs/indicatif/latest/indicatif/) version `0.17.9` - Pour les barres de progression.
- [`tokio`](https://docs.rs/tokio/latest/tokio/) version `1` (features = ["full"]) - Pour gérer la programmation asynchrone.
- [`chrono`](https://docs.rs/chrono/latest/chrono/) version `0.4` - Pour manipuler et formater les dates et heures.
- [`clap`](https://docs.rs/clap/latest/clap/) version `4.5.21` - Pour gérer les arguments de la ligne de commande.
- [`scraper`](https://docs.rs/scraper/latest/scraper/) version `0.21.0` - Pour analyser et scraper des pages HTML.
- [`url`](https://docs.rs/url/latest/url/) version `2.2` - Pour travailler avec des URL.
- [`regex`](https://docs.rs/regex/latest/regex/) version `1.9` - Pour la gestion des expressions régulières.
- [`dirs-next`](https://docs.rs/dirs-next/latest/dirs_next/) version `2.0` - Pour obtenir des chemins de répertoires standard.

___
## 🧑‍💻 Authors
+ Fabien OLIVIER
+ Raphaël LOVERGNE 
+ Axelle FOUQUEMBERG
+ Jean-Frederic NANGY
