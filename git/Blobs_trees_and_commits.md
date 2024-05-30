# blobs, trees and commits

## Exploration du répertoire `.git/`

Ce document décrit les étapes nécessaires pour explorer le répertoire `.git/` d'un projet, trouver le dernier hash d'objet, et afficher le contenu de certains répertoires et fichiers référencés par le dernier commit.

### 1. Naviguer dans le répertoire `.git/`

Dans votre projet, naviguez jusqu'au répertoire `.git/` et examinez son contenu. Vous devrez expliquer l'utilité de chaque sous-répertoire et fichier important.

- **objects/** : Contient tous les objets Git (blobs, trees, commits, et tags) qui composent l'historique du projet.
- **config** : Contient la configuration spécifique du dépôt, incluant les paramètres utilisateur, les alias de commandes, et les paramètres spécifiques au dépôt.
- **refs/** : Contient les références vers les objets commit, comme les branches (`refs/heads/`), les tags (`refs/tags/`), et les remotes (`refs/remotes/`).
- **HEAD** : Indique la référence actuelle de la branche sur laquelle vous travaillez. Il pointe souvent vers un fichier dans `refs/heads/`.

### 2. Dernier hash d'objet (type et contenu)

Pour trouver le dernier hash d'objet dans le répertoire `.git/objects/`, utilisez les commandes Git suivantes :
```sh
# Hash du dernier commit
git log -1

# Affiche le type d'objet du dernier commit
git cat-file -t <hash_objet>

# Affiche le contenu du dernier commit
git cat-file -p <hash_objet>
```

### 3. Déverser l'arborescence du répertoire

Déverser le contenu d'un répertoire fait référence à l'action de lister tous les fichiers et sous-répertoires contenus dans un répertoire spécifique d'un commit donné dans un dépôt Git.

1. **Afficher l'arborescence** :
```sh
# déverser le contenu du répertoire lib/ au commit <hash_de_commit>
git ls-tree <hash_de_commit> lib/
#  déverser le contenu du fichier hello.sh au commi <hash_de_commit>
git show <hash_de_commit>:hello.sh
```

