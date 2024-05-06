# ASCII-ART-WEB-DOCKERIZE

## Description

Le programme permet de générer une image docker qui sera mise en service dans un conteneur afin de pouvoir accéder à un site internet d'Ascii-art.

## Usage

Afin de pouvoir faire fonctionner le programme, le logiciel Docker dois être installé.
+ <a href="https://docs.docker.com" target="blank">documentation docker</a><br>
+ <a href="https://docs.docker.com/get-docker/" target="blank">procédure de téléchargement et d'installation de docker</a><br>

Une fois le logiciel docker installé sur votre machine, voici les commandes d'installation du programme à rentrer dans un terminal à la racine du dossier :

La commande suivante sert à créer une image, en ouvrir un conteneur et le lancer : 
```sh
sh ./Launcher.sh
```

Une fois lancé, la commande suivante permet de stopper le conteneur : 
```sh
sh ./StopContainer.sh
```

Si vous souhaitez relancer votre conteneur, utilisez la commande suivante : 
```sh
sh ./StartContainer.sh
```

Pour effacer le conteneur tout en gardant l'image, lancez la commande :
```sh
sh ./DeleteContainer.sh
```

Pour en créer un nouveau depuis l'image encore présente, faites :
```sh
sh ./CreateContainer.sh
```

Enfin voici la commande de désinstallation du programme :
```sh
sh ./Uninstaller.sh
```

## Fonctionnement

Une fois le serveur lancé, rendez-vous à l'adresse suivante : <a href="http://localhost:8080/" target="blank">http://localhost:8080/</a><br>
Vous avez devant vous la page d'accueil du serveur où vous trouverez une zone de texte à remplir.
Renseignez la zone de texte avec le texte que vous souhaitez transcrire en ascii-art.

Selectionez votre police à l'aide du menu déroulant, puis apuyez sur le bouton `Go`.

<table align= "center">
    <thead>
        <th align= "center" colspan="4">Exemples des polices d'Ascii-art possible</th>
    </thead>
    <tbody>
        <tr>
            <th align= "center">Standard</th>
            <th align= "center">Shadow</th>
            <th align= "center">Thinkertoy</th>
            <th align= "center">Varsity</th>
        </tr>
        <tr>
            <td><img src="https://i43.servimg.com/u/f43/15/76/70/95/image_11.png"></td>
            <td><img src="https://i43.servimg.com/u/f43/15/76/70/95/captur22.png"></td>
            <td><img src="https://i43.servimg.com/u/f43/15/76/70/95/image_12.png"></td>
            <td><img src="https://cdn.discordapp.com/attachments/1118901027920424980/1118901070979137547/Capture_decran_du_2023-06-15_15-51-24.png"></td>
        </tr>    
    </tbody>
</table>

## Contraintes

Seul les caractére Ascii sont pris en charge par les polices ascii-art.

Vous pouvez renseigner les caractères suivants : 
```
abcdefghijklmnopqrstuvwxyz
ABCDEFGHIJKLMNOPQRSTUVWXYZ
0123456789
{}!:;,?./%$+-<>'"~&[]@`\*=()_|#
[espace]
```
Le programme prend en compte la commande de retour à la ligne : 
`\n`

## Authors
+ Fabien OLIVIER
+ Mickael MARCHAIS
