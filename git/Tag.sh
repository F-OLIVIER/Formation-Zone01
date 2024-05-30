#!/bin/bash

echo "╭──────────────────────────────────────────────╮"
echo "│            Tag la version actuel             │"
echo "╰──────────────────────────────────────────────╯"
git tag v1

echo "╭──────────────────────────────────────────────╮"
echo "│          Tag la version précédente           │"
echo "╰──────────────────────────────────────────────╯"
git tag v1-beta HEAD~1

echo "╭──────────────────────────────────────────────╮"
echo "│           Aller sur la version v1            │"
echo "╰──────────────────────────────────────────────╯"
git checkout v1
echo "╭──────────────────────────────────────────────╮"
echo "│           Contenu du fichier hello           │"
echo "╰──────────────────────────────────────────────╯"
cat ./hello/hello.sh 

echo "╭──────────────────────────────────────────────╮"
echo "│        Aller sur la version v1-beta          │"
echo "╰──────────────────────────────────────────────╯"
git checkout v1-beta
echo "╭──────────────────────────────────────────────╮"
echo "│           Contenu du fichier hello           │"
echo "╰──────────────────────────────────────────────╯"
cat ./hello/hello.sh 

echo "╭──────────────────────────────────────────────╮"
echo "│      Retourner à la branche principale       │"
echo "╰──────────────────────────────────────────────╯"
git checkout master
echo "╭──────────────────────────────────────────────╮"
echo "│           Contenu du fichier hello           │"
echo "╰──────────────────────────────────────────────╯"
cat ./hello/hello.sh 

echo "╭──────────────────────────────────────────────╮"
echo "│      Afficher la liste de tous les tags      │"
echo "╰──────────────────────────────────────────────╯"
git tag

