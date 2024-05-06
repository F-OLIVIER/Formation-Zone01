#!/bin/bash

# ----------------------------------------------------------
# ------ Script de désinstallation de l'image Docker -------
# ----------------------------------------------------------

echo
echo "--------------> Suppression du container <----------------"
docker container rm -f web
echo "------------------------> Terminé <-----------------------"
echo
echo "----------------> Suppression de l'image <----------------"
docker image rm -f exercice_dockerize
echo "------------------------> Terminé <-----------------------"
echo
echo "-------------> Liste des conteneurs docker <--------------"
# option -a pour afficher tout les containers (par defaut que ceux qui run)
docker container ls -a
echo
echo "----------------------------------------------------------"
echo
echo "---------------> Liste des images docker <----------------"
# option -a pour afficher toutes les images (par defaut que celle qui run)
docker image ls -a
echo
echo "----------------------------------------------------------"
echo

# Pour une purge compléte de tout le contenu docker, décommenter les lignes 30 à 33
echo "---------> Purge compléte du cache de docker <------------"
echo "-----> Souhaitez-vous effectuer la purge ? (y/n) <--------"
docker system prune -a
echo "------------------------> Terminé <-----------------------"
echo

