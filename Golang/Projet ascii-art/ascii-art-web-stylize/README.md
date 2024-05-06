# ascii-art-web-stylize

## Usage

Pour lancer le serveur, il suffit d'ouvrir un terminal à la racine et d'entrer : 
```sh
go run cmd/web/ascii-art-web.go
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
