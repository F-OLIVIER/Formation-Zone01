# Local and remote repositories

```sh
# Aller dans le dossier work
cd work/

# Initialise le depot
git init cloned_hello

# naviguer dans le dossier
cd cloned_hello

# Ajouter un pointage distant vers le référentiel d'origine
git remote add origin <repository_url>

# Récupérer le contenu du repo d'origine
git fetch origin

# Réinitialiser la branche au contenu récupéré
git reset --hard origin/master
```
```sh
# Résultats du `git log`
commit 77bc561534878df1ec246d90da299cce54804dea (HEAD -> master, origin/master)
Author: folivier <fabien.olivier2@gmail.com>
Date:   Thu May 30 10:45:34 2024 +0200

    merge.md

commit 091cefda4aaa1a50826947d0a0d93f223448e5af
Author: folivier <fabien.olivier2@gmail.com>
Date:   Thu May 30 10:30:58 2024 +0200

    change lib/hello.sh

commit 04a0a4a3c4c7ef4b2f22cb6063f0b7f9992606d7
Author: folivier <fabien.olivier2@gmail.com>
Date:   Thu May 30 10:27:42 2024 +0200

    branching

commit 874442b27bc7b834e009abd967af5487b29b0b92
Author: folivier <fabien.olivier2@gmail.com>
Date:   Thu May 30 10:21:15 2024 +0200

    branching

commit 67f203d69ffb952b1d60baa4258a41a90a577df4
Author: folivier <fabien.olivier2@gmail.com>
Date:   Thu May 30 10:05:41 2024 +0200

    blobs, trees and commits

commit 4de4b9cae6dfc7ebddb2d5d0cda615e2dab3cdb7
Author: folivier <fabien.olivier2@gmail.com>
Date:   Thu May 30 10:05:12 2024 +0200

    blobs, trees and commits

commit e0e8f2c728c29e070432e5ce527712a33e1d76cb
Author: folivier <fabien.olivier2@gmail.com>
Date:   Thu May 30 09:59:03 2024 +0200

    blobs, trees and commits

commit 29789fb349f7efec107c1c6a7f4cec65c5dbc7b7
Author: folivier <fabien.olivier2@gmail.com>
Date:   Thu May 30 09:28:12 2024 +0200

    Move hello.sh to lib/ directory

commit 3ef48a8ffb61471e62a1b2b6c287c147fc349458
Author: folivier <fabien.olivier2@gmail.com>
Date:   Thu May 30 09:23:25 2024 +0200
```

__Affiche le nom du référentiel distant et fournit plus d'informations à son sujet.__
```sh
# Afficher le nom du repo distant
git remote

# Fournir plus d'informations sur le repo distant
git remote show origin
```
```sh
# Résultats de `git remote show origin`
* distant origin
  URL de rapatriement : https://zone01normandie.org/git/folivier/git.git
  URL push : https://zone01normandie.org/git/folivier/git.git
  Branche HEAD : master
  Branches distantes :
    greet  suivi
    master suivi
  Référence locale configurée pour 'git push' :
    master pousse vers master (à jour)
```

__Liste toutes les branches locales et distantes du dépôt cloned_hello.__
```sh
# List all local branches
git branch

# List all remote branches
git branch -r
```

Résultats du `git branch -r`
```sh
# Résultats du `git branch`
* master

# Résultats du `git branch -r`
  origin/greet
  origin/master
```

__Pull les modifications du repo d'origine ou du repo cloné__
```sh
git pull
```

__Ajouter une branche locale nommée greet qui suit la branche distante origin/greet.__
```sh
# Se placer dans le répertoire cloné (cloned_hello)
cd work/cloned_hello

# Créer une nouvelle branche locale nommée greet qui suit la branche distante origin/greet
git checkout -b greet --track origin/greet
```

__Ajoutez une branche distante à votre dépôt Git et poussez les branches main et greet vers la branche distante.__
```sh
# Se placer dans le répertoire cloné (cloned_hello)
cd work/cloned_hello

# Ajouter un dépôt distant nommé "origin" avec l'URL du dépôt distant
git remote add origin <url_du_depot>

# Pousser la branche main vers le dépôt distant
git push -u origin main

# Pousser la branche greet vers le dépôt distant
git push -u origin greet
```
