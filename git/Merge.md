# Merge

__Fusion de Main dans la branche Greet (Conflit) :__
```sh
# Se placer sur la branche greet
git checkout greet

# Fusionner la branche main dans la branche greet
git merge master
```
Résoudre le conflit de merge
```sh
# Ajouter les fichiers résolus
git add --all

# Valider les modifications
git commit -m "Merge master"
git push
```

__Rebase de la branche Greet :__
Rebasage : Réapplique les commits de la branche source sur le dessus de la branche cible, réécrivant l'historique de la branche source. Cela crée une histoire linéaire et nette, mais peut provoquer des conflits si des modifications similaires ont été apportées dans les deux branches.
```sh
# Hash du dernier commit
git log -1
```
```sh
# Revenir au commit avant la fusion initiale
git reset --hard <commit_avant_fusion>

# Rebase de la branche greet sur la branche main
git rebase main
```

__Fusion de Greet dans Main :__
Fusion : Crée un nouveau commit de fusion qui combine les modifications de deux branches. Les branches impliquées dans la fusion restent inchangées.
```sh
# Se placer sur la branche main
git checkout main

# Fusionner la branche greet dans la branche main
git merge greet
```
