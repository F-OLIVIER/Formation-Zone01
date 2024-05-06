# Lem-in

### Descriptif
_______
Le programme prend en entrée un fichier contenant les éléments suivants :
- un nombre de fourmis
- une liste de salle (nom de la salle et coordonnée x et y dans l'espace)
- une liste de lien entre les salles

Le programme `lem-in` permet via un backtrakking trouver la liste des chemins afin de faire circuler les fourmis dans les salles depuis la salle `start` jusqu'à la salle `end` sans jamais avoir 2 fourmis dans la même salle et en ayant un minimum de déplacement des fourmis exécuté.

Le programme `vizualizer-lemin` prend en entrée la sortie du premier programme et permet de voir visuellement via un visualisateur 3D la fourmilière.

### Usage
_______
```go
go run . example/example00.txt 
```

### Authors
_______
+ Quentin DELOOZE
+ Agustin GUILLEN
+ Fabien OLIVIER
+ Fabien FANISSE

