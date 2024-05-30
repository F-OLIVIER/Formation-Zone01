#!/bin/bash

echo "╭──────────────────────────────────────────────╮"
echo "│      Ajout d'un commentaire a hello.sh       │"
sed -i '3s/.*/# This is a bad comment. We want to revert it./' ./hello/hello.sh
echo "│           Contenu du fichier hello           │"
echo "╰──────────────────────────────────────────────╯"
cat ./hello/hello.sh 
echo "╭──────────────────────────────────────────────╮"
echo "│         Annulation des modifications         │"
git checkout -- ./hello/hello.sh 
echo "│           Contenu du fichier hello           │"
echo "╰──────────────────────────────────────────────╯"
cat ./hello/hello.sh 


echo "╭──────────────────────────────────────────────╮"
echo "│   Ajout commentaire + git add de hello.sh    │"
sed -i '3s/.*/# This is an unwanted but staged comment./' ./hello/hello.sh
git add ./hello/hello.sh
echo "│           Contenu du fichier hello           │"
echo "╰──────────────────────────────────────────────╯"
cat ./hello/hello.sh 
echo "╭──────────────────────────────────────────────╮"
echo "│         Annulation des modifications         │"
git restore --staged ./hello/hello.sh
git restore ./hello/hello.sh
echo "│           Contenu du fichier hello           │"
echo "╰──────────────────────────────────────────────╯"
cat ./hello/hello.sh 


echo "╭──────────────────────────────────────────────╮"
echo "│  Ajout commentaire + git commit de hello.sh  │"
sed -i '3s/.*/# This is an unwanted but committed change./' ./hello/hello.sh
git add ./hello/hello.sh
git commit -m "Commit a annuler"
echo "│           Contenu du fichier hello           │"
echo "╰──────────────────────────────────────────────╯"
cat ./hello/hello.sh 
echo "╭──────────────────────────────────────────────╮"
echo "│         Annulation des modifications         │"
git reset --hard HEAD^
echo "│           Contenu du fichier hello           │"
echo "╰──────────────────────────────────────────────╯"
cat ./hello/hello.sh 


echo "╭──────────────────────────────────────────────╮"
echo "│                Ajout d'un tag                │"
sed -i '3s/.*/# tag et commit oops./' ./hello/hello.sh
git tag oops
git add ./hello/hello.sh 
git commit -m "oops"
echo "│           Contenu du fichier hello           │"
echo "╰──────────────────────────────────────────────╯"
cat ./hello/hello.sh 
echo "╭──────────────────────────────────────────────╮"
echo "│         Annulation des modifications         │"
git log --oneline                                           # identifier le commit avec "oops
git show-ref --tags | grep v1                               # affichage du hash du tag v1
git reset --hard 40ab8d94cdecf667520af7f490db2e503f9cddff   # reset sur le tag v1
echo "│           Contenu du fichier hello           │"
echo "╰──────────────────────────────────────────────╯"
cat ./hello/hello.sh 

echo "╭──────────────────────────────────────────────╮"
echo "│     Affichage du commit avec tag \"oops\"      │"
echo "╰──────────────────────────────────────────────╯"
git reflog show --all | grep "oops"

echo "╭──────────────────────────────────────────────╮"
echo "│        Commit \"oops\" non référencé         │"
echo "╰──────────────────────────────────────────────╯"
git fsck --unreachable
echo "╭──────────────────────────────────────────────╮"
echo "│     Suppression des objets inatteignables     │"
echo "╰──────────────────────────────────────────────╯"
git gc --prune=now
echo "╭──────────────────────────────────────────────╮"
echo "│    Vérification des commit non référencé     │"
echo "╰──────────────────────────────────────────────╯"
git fsck --unreachable

