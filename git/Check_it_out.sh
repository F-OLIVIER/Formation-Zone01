#!/bin/bash

echo "╭──────────────────────────────────────────────╮"
echo "│           Contenu du fichier hello           │"
echo "╰──────────────────────────────────────────────╯"
cat ./hello/hello.sh 

echo "\n╭──────────────────────────────────────────────╮"
echo "│            restore le 1er commit             │"
echo "╰──────────────────────────────────────────────╯"
first_commit_hash=$(git rev-list --max-parents=0 HEAD)
git checkout $first_commit_hash -- ./hello/hello.sh
echo "╭──────────────────────────────────────────────╮"
echo "│           Contenu du fichier hello           │"
echo "╰──────────────────────────────────────────────╯"
cat ./hello/hello.sh 

echo "╭──────────────────────────────────────────────╮"
echo "│            restore le 2eme commit            │"
echo "╰──────────────────────────────────────────────╯"
second_recent_commit_hash=$(git rev-list --skip=2 -n 1 HEAD)
git checkout $second_recent_commit_hash -- ./hello/hello.sh
echo "╭──────────────────────────────────────────────╮"
echo "│           Contenu du fichier hello           │"
echo "╰──────────────────────────────────────────────╯"
cat ./hello/hello.sh 

echo "\n╭──────────────────────────────────────────────╮"
echo "│        restore commit le plus récent         │"
echo "╰──────────────────────────────────────────────╯"
git checkout master -- ./hello/hello.sh
echo "╭──────────────────────────────────────────────╮"
echo "│           Contenu du fichier hello           │"
echo "╰──────────────────────────────────────────────╯"
cat ./hello/hello.sh
