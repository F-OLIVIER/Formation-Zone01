# Bare

Un dépôt Git "bare" est un dépôt qui n'a pas de répertoire de travail. Il contient uniquement les informations de contrôle de version et n'a pas de copie des fichiers de votre projet. Les dépôts bare sont généralement utilisés comme dépôts centralisés vers lesquels plusieurs développeurs peuvent pousser et tirer des modifications. Ils sont utiles pour la collaboration car ils permettent aux développeurs de partager des modifications sans affecter leurs propres copies de travail du projet.

__Créez le dépôt bare__
```sh
git clone --bare hello hello.git
```

__Ajoutez le dépôt bare comme remote au dépôt original__
```sh
cd hello
git remote add shared ../hello.git
```

__Modifiez le fichier README.md dans le dépôt original avec le contenu fourni__
```sh
echo "Ceci est l'exemple Hello World du projet git. (Modifié dans l'original et poussé vers shared)" > README.md
```

__Validez les modifications et poussez-les vers le dépôt partagé__
```sh
git add README.md
git commit -m "Mise à jour README.md"
git push shared main
```

__Pull le depot partager sur le repo cloné__
```sh
cd ./work/cloned_hello
git pull origin main
```
