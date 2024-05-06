#!/bin/bash

# ----------------------------------------------------------
# -------- Script d'installation de l'image Docker ---------
# ----------------------------------------------------------

# Build de l'image
echo "-----------------> Création de l'image <------------------"
# option t (tag) pour donner un nom à l'image
docker build -t ascii-art-web-dockerize .
echo "------------------------> Terminé <-----------------------"
echo
echo "----------------> Lancement de l'image <------------------"
# Execution du conteneur web contenant l'image ascii-art-web-dockerize
# option d (detache) pour tache de fond, p (publish) pour le port d'écoute
docker run -d -p 8080:8080 --name web ascii-art-web-dockerize
echo "------------------------> Terminé <-----------------------"
echo
echo "---------------> Liste des images docker <----------------"
# option a pour afficher tout (en run et stop)
docker image ls -a
echo "----------------------------------------------------------"
echo
echo "-------------> Liste des conteneurs docker <--------------"
# option a pour afficher tout (en run et stop)
docker container ls -a
echo "----------------------------------------------------------"
echo
echo "-------------------> Adresse du site <--------------------"
echo "                 http://127.0.0.1:8080/"
echo "----------------------------------------------------------"
# Ouverture automatique de la page web dans le navigateur (attention : bloque la console)
# open 'http://127.0.0.1:8080/'
echo
