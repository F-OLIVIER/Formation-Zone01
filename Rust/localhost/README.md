# LOCALHOST

## 📝 Descriptif
Le projet consiste à créer un serveur web local en protocole `HTTP/1.1 RFC` qui gère différents types de requêtes ainsi que différents types d'erreurs.

**Méthode de requete géré :**
- Get
- Post
- Delete

**Type d'erreur géré :**
- 400: Bad Request
- 403: Forbidden
- 404: Not Found
- 405: Method Not Allowed
- 413: Payload Too Large
- 500: Internal error server 

___
## ⚙️ Installation & usage

**Paramétrage du serveur**</br>
⌨️ Avant de demarrer le serveur, il faut le configurer. Pour cela ouvrir le ficher `/src/config/config.txt` et saisisez les informations d'`IP` (obligatoire), de `port` (obligatoire) et de `nom de domaine` (facultatif) de la maniére suivantes : `IP:Port Domain_name`.
__Voici un exemple :__
```sh
127.0.0.2:7878 i.moi
127.0.0.2:7879 freddy.dieu 
127.0.0.3:4000
```

**Mise en route du serveur**</br>
⚠️ Si vous avez saisi un nom de domaine dans le fichier de configuration, le programme va automatiquement modifier le fichier système `/etc/hosts` de votre machine au démarrage, pour cela votre mot de passe vous sera demandé car c'est une opération administrateur.
```sh
cargo run
```

**Tester le serveur**
```sh
# Si necessaire "siege" doit être installer
#  Pour Linux : 
sudo apt install siege
#  Pour MacOS : 
brew install siege

#  Exécution du test (durée de 1 minute)
siege -b -c 10 -t 1M http://127.0.0.2:7878
```

Voici un exemple de résultats de test serveur
```sh
New configuration template added to /home/fabien/.siege
Run siege -C to view the current settings in that file

{	
"transactions":             2771689,   # Nombre total de transactions (requêtes HTTP) traitées
"availability":              100.00,   # Pourcentage de disponibilité du serveur (en %)
"elapsed_time":               59.15,   # Durée totale du test (en secondes)
"data_transferred":         3267.11,   # Quantité totale de données transférées (en Mo)
"response_time":               0.00,   # Temps de réponse moyen pour chaque requête (en secondes)
"transaction_rate":        46858.64,   # Nombre moyen de transactions par seconde
"throughput":                 55.23,   # Débit moyen du serveur en termes de données transférées par seconde
"concurrency":                 9.04,   # Nombre moyen d'utilisateurs simultanés
"successful_transactions":  2771690,   # Nombre total de transactions réussies
"failed_transactions":            0,   # Nombre total de transactions échouées
"longest_transaction":         0.01,   # Durée de la transaction la plus longue (en secondes)
"shortest_transaction":        0.00    # durée de la transaction la plus courte (en secondes)
}
```

___
## 🔗 Dépendences
Le programme utilise `Rust` avec la `version 1.79.0` de [cargo](https://www.rust-lang.org/fr) et les imports suivants :<br>
- [libc](https://docs.rs/libc/latest/libc/) version `0.2.161`
- [lazy_static](https://docs.rs/lazy_static/latest/lazy_static/) version `1.5.0`
- [urlencoding](https://docs.rs/urlencoding/latest/urlencoding/) version `2.1.3`
- [log](https://docs.rs/log/latest/log/) version `0.4.22`
- [env_logger](https://docs.rs/env_logger/latest/env_logger/) version `0.11.5`
- [chrono](https://docs.rs/chrono/latest/chrono/) version `0.4.38`

___
## 🧑‍💻 Authors
+ Fabien OLIVIER
+ Jean-Frédéric NANGY
+ Toufik HIMOUM
