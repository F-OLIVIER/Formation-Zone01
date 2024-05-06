# Ascii Art Program

### Descriptif
Ce programme permet de généré un ascii art dans le terminal. Cet ascii art peux être affiché en couleurs

### Option de couleur :
L'ensembles des couleurs de type `HEX`, `RGB` et `HSL` de la palette de couleur HTML suivante sont prise en charge <br>
Plusieurs option de saisie sont possible :
+ Soit la couleur est saisie sans espace, avec le `#` ou les `,` et les `%`
+ Soit la couleur est saisie sans espace avec le type de couleur suivi entre parenthése des numéros de la couleur, c'est à dire : `hex(#5FC53D)` ou `rgb(95,197,61)` ou `hsl(105,69,51)` 

**Cliquer sur l'image pour ouvrir un site internet avec une palette de couleur intéractive**
<a href="https://htmlcolorcodes.com/fr/" title="Lien vers la palette de couleur HTML" target="blank">
	<img alt="Palette de couleur HTML" src="https://i43.servimg.com/u/f43/15/76/70/95/image_14.png"/>
</a>

### Usage :
___
```go
go run main.go [--color=<color>] [<filter>] <text>
```

**Exemple d'utilisation :** <br>
`go run main.go --color=red "te" "test"` </br>
"test" sera affiché au format ASCII dans le terminal dans la couleur par défaut, et chaque caractère de "test" correspondant à l'une des caractères de "te" sera affiché en rouge

`go run main.go "test"` </br>
"test" sera simplement affiché au format ASCII dans le terminal 

`go run main.go --color=red "test"` </br>
"test" sera affiché au format ASCII dans le terminal en rouge 


### Structure du script :
___

- Vérification de l'ordre et de la validité des arguments indiqués en paramètres 

- Récupération du code couleur RGB la couleur en la comparant avec l'Array `availableColors` pour les couleur saisie en toutes lettre et en convertissant pour les autres format

- Chaque item de l'array `asciiLines` contient une ligne du fichier `standard.txt` (en résultat de l'utilisation de la fonction `Split` de package `strings` sur le fichier `standard.txt` pour chaque saut de ligne `\n`) 

- Une boucle est appliquée pour chaque caractère de "text" indiqué en paramètre `<text>` de la commande pour exécuter le script
    . Pour chacun de ces caractères (en ignorant le combo de runes `\n`), on applique la fonction `getAsciiChar` qui s'occupe de créer une array contenant le caractère en ascii
    . Le résultat de la fonction `getAsciiChar` est ensuite ajouté les couleurs si demandé
    . Enfin une boucle permet d'afficher le caractére ascii dans la console

### Authors
___
+ Tanguy CHATIGNY
+ Fabien OLIVIER