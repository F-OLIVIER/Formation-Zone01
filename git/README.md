# GIT

## 📝 Descriptif

Le projet consiste à apprendre à utiliser [git](https://git-scm.com). Lien vers le [sujet de l'excercice](https://ytrack.learn.ynov.com/git/root/public/src/branch/master/subjects/git)

Voici une liste de commande git qui couvre les commandes les plus essentielles pour la plupart des workflows de développement :

__Configuration de Git__</br>
<table>
    <tr>
        <td><code>git config --global user.name "Nom Utilisateur"</code></td>
        <td>Configure le nom d'utilisateur pour tous les dépôts Git locaux.</td>
    </tr>
    <tr>
        <td><code>git config --global user.email "email@example.com"</code></td>
        <td>Configure l'email de l'utilisateur pour tous les dépôts Git locaux.</td>
    </tr>
    <tr>
        <td><code>git config --list</code></td>
        <td>Affiche toutes les configurations Git.</td>
    </tr>
</table>

__Commandes de base__</br>
<table>
    <tr>
        <td><code>git init</code></td>
        <td>Initialise un nouveau dépôt Git.</td>
    </tr>
    <tr>
        <td><code>git clone &lt;url&gt;</code></td>
        <td>Clone un dépôt distant.</td>
    </tr>
    <tr>
        <td><code>git status</code></td>
        <td>Affiche l'état des fichiers dans le répertoire de travail et la zone de staging.</td>
    </tr>
        <tr>
        <td><code>git add &lt;fichier&gt;</code></td>
        <td>Ajoute un fichier à la zone de staging.</td>
    </tr>
        <tr>
        <td><code>git commit -m "Message"</code></td>
        <td>Enregistre les modifications ajoutées avec un message de commit.</td>
    </tr>
        <tr>
        <td><code>git log</code></td>
        <td>Affiche l'historique des commits.</td>
    </tr>
        <tr>
        <td><code>git diff</code></td>
        <td>Affiche les différences entre les fichiers modifiés et la dernière version enregistrée.</td>
    </tr>
        <tr>
        <td><code>git diff --staged</code></td>
        <td>Affiche les différences entre les fichiers en staging et la dernière version enregistrée.</td>
    </tr>
</table>


__Branches et Merging__
<table>
    <tr>
        <td><code>git branch</code></td>
        <td>Liste toutes les branches dans le dépôt local.</td>
    </tr>
    <tr>
        <td><code>git branch &lt;nom_de_branche&gt;</code></td>
        <td>Crée une nouvelle branche.</td>
    </tr>
    <tr>
        <td><code>git checkout &lt;nom_de_branche&gt;</code></td>
        <td>Bascule vers une branche spécifique.</td>
    </tr>
    <tr>
        <td><code>git checkout -b &lt;nom_de_branche&gt;</code></td>
        <td>Crée et bascule vers une nouvelle branche.</td>
    </tr>
    <tr>
        <td><code>git merge &lt;nom_de_branche&gt;</code></td>
        <td>Fusionne une branche dans la branche courante.</td>
    </tr>
    <tr>
        <td><code>git branch -d &lt;nom_de_branche&gt;</code></td>
        <td>Supprime une branche localement.</td>
    </tr>
</table>

__Collaboration__
<table>
    <tr>
        <td><code>git remote add origin <url&gt;</code></td>
        <td>Ajoute un dépôt distant.</td>
    </tr>
    <tr>
        <td><code>git remote -v</code></td>
        <td>Affiche les dépôts distants.</td>
    </tr>
    <tr>
        <td><code>git fetch</code></td>
        <td>Récupère les modifications depuis le dépôt distant sans les fusionner.</td>
    </tr>
    <tr>
        <td><code>git pull</code></td>
        <td>Récupère et fusionne les modifications depuis le dépôt distant.</td>
    </tr>
    <tr>
        <td><code>git push</code></td>
        <td>Envoie les modifications locales vers le dépôt distant.</td>
    </tr>
    <tr>
        <td><code>git push -u origin &lt;nom_de_branche&gt;</code></td>
        <td>Envoie la branche locale vers le dépôt distant et la suit.</td>
    </tr>
</table>

__Annulation et réinitialisation__</br>
<table>
    <tr>
        <td><code>git reset &lt;fichier&gt;</code></td>
        <td>Retire un fichier de la zone de staging.</td>
    </tr>
    <tr>
        <td><code>git reset --hard</code></td>
        <td>Réinitialise le répertoire de travail et l'index à la dernière version commitée.</td>
    </tr>
    <tr>
        <td><code>git revert &lt;hash_commit&gt;</code></td>
        <td>Crée un nouveau commit qui annule les modifications d'un commit spécifique.</td>
    </tr>
    <tr>
        <td><code>git stash</code></td>
        <td>Enregistre temporairement les modifications en cours pour les rétablir plus tard.</td>
    </tr>
    <tr>
        <td><code>git stash pop</code></td>
        <td>Récupère les modifications stockées dans le stash.</td>
    </tr>
</table>

__Rebase et autres commandes avancées__</br>
<table>
    <tr>
        <td><code>git rebase &lt;branche&gt;</code></td>
        <td>Applique les commits de la branche courante sur une autre branche.</td>
    </tr>
    <tr>
        <td><code>git cherry-pick &lt;hash_commit&gt;</code></td>
        <td>Applique les modifications d'un commit spécifique dans la branche courante.</td>
    </tr>
    <tr>
        <td><code>git tag &lt;nom_tag&gt;</code></td>
        <td>Crée un tag pour marquer une version spécifique.</td>
    </tr>
    <tr>
        <td><code>git tag -d &lt;nom_tag&gt;</code></td>
        <td>Suprimer un tag.</td>
    </tr>
</table>

__Affichage et recherche__</br>
<table>
    <tr>
        <td><code>git show &lt;hash_commit&gt;</code></td>
        <td>Affiche les détails d'un commit spécifique.</td>
    </tr>
    <tr>
        <td><code>git grep &lt;terme_de_recherche&gt;</code></td>
        <td>Recherche un terme dans les fichiers du dépôt.</td>
    </tr>
</table>

___
## ⚙️ Usage

```sh
# History
sh ./History.sh

# Check it out
sh ./Check_it_out.sh

# git/Reverting_changes
sh ./Reverting_changes.sh

# Tag
sh ./Tag.sh
```

___
## 🧑‍💻 Authors

+ Fabien OLIVIER
