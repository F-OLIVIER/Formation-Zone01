#!/bin/bash

# Configurer les chemins des scripts CGI
PYTHON_SCRIPT="ressources/script.py"
PHP_SCRIPT="ressources/script.php"
RUST_SCRIPT="ressources/script.rs"

# Lancer le serveur CGI en arrière-plan (adaptez le chemin si nécessaire)
echo "Lancement du serveur CGI..."
cargo run &> server_log.txt &
SERVER_PID=$!
sleep 2 # Donne du temps au serveur pour démarrer

# Assurez-vous que le serveur est bien lancé
if ! ps -p $SERVER_PID > /dev/null; then
    echo "Erreur: Impossible de lancer le serveur CGI. Vérifiez server_log.txt pour les détails."
    exit 1
fi

# Créer un script Python CGI
echo "Création du script Python CGI..."
cat > $PYTHON_SCRIPT <<EOL
#!/usr/bin/env python3
print("Content-Type: text/html\n")
print("<html><body><h1>Hello, Python CGI!</h1></body></html>")
EOL
chmod +x $PYTHON_SCRIPT

# Créer un script PHP CGI
echo "Création du script PHP CGI..."
cat > $PHP_SCRIPT <<EOL
<?php
echo "Content-Type: text/html\n\n";
echo "<html><body><h1>Hello, PHP CGI!</h1></body></html>";
?>
EOL
chmod +x $PHP_SCRIPT

# Créer un script Rust CGI
echo "Création du script Rust CGI..."
cat > $RUST_SCRIPT <<EOL
fn main() {
    // En-tête CGI avec deux println pour bien séparer l'en-tête du corps
    println!("Content-Type: text/html"); 
    println!(); // Ceci ajoute une ligne vide après l'en-tête

    // Corps du message HTML
    println!("<html><body><h1>Hello, Rust CGI!</h1></body></html>");
}
EOL
chmod +x $RUST_SCRIPT

# Tester les scripts avec curl
echo "Test du script Python CGI..."
response=$(curl -s -w "%{http_code}" -o response.txt http://127.0.0.1:7878/script.py)
cat response.txt # Affiche le contenu de la réponse
if [ "$response" -eq 200 ]; then
    echo "Test du script Python CGI réussi!"
else
    echo "Test du script Python CGI échoué avec le code HTTP: $response"
fi

echo "Test du script PHP CGI..."
response=$(curl -s -w "%{http_code}" -o response.txt http://127.0.0.1:7878/script.php)
cat response.txt # Affiche le contenu de la réponse
if [ "$response" -eq 200 ]; then
    echo "Test du script PHP CGI réussi!"
else
    echo "Test du script PHP CGI échoué avec le code HTTP: $response"
fi

echo "Test du script Rust CGI..."
response=$(curl -s -w "%{http_code}" -o response.txt http://127.0.0.1:7878/script.rs)
cat response.txt # Affiche le contenu de la réponse
if [ "$response" -eq 200 ]; then
    echo "Test du script Rust CGI réussi!"
else
    echo "Test du script Rust CGI échoué avec le code HTTP: $response"
fi

# Nettoyer les processus et fichiers
echo "Arrêt du serveur CGI..."
kill $SERVER_PID

echo "Suppression des scripts CGI..."
rm -f $PYTHON_SCRIPT $PHP_SCRIPT $RUST_SCRIPT

echo "Test terminé."
