# Push swap

### Descriptif
_______
Le programme utilise un algorithme de tri sur deux piles avec un nombre d'opérations limité.

Il y a deux programmes :
- **push-swap**, qui calcule et affiche les étapes de tri.
- **checker**, qui calcule et affiche si le tri est possible ou non


### Liste des opérations autorisé dans l'exercice
- `pa` pousse le premier élément de la pile `b` vers la pile `a`
- `pb` pousse le premier élément de la pile `a` vers la pile `b`
- `sa` échange les 2 premiers éléments de la pile `a`
- `sb` échange les 2 premiers éléments de la pile `b`
- `ss` exécute `sa` et `sb`
- `ra` rotation de la pile `a` (décalage vers le haut de tous les éléments de la pile `a` de 1, le premier élément devient le dernier)
- `rb` rotation de la pile `b` (décalage vers le haut de tous les éléments de la pile `b` de 1, le premier élément devient le dernier)
- `rr` exécute `ra` et `rb`
- `rra` rotation inverse de la pile `a` (décale vers le bas tous les éléments de la pile `a` de 1, le dernier élément devient le premier)
- `rrb` rotation inverse de la pile `b`
- `rrr` exécute `rra` et `rrb`

### Usage
_______
Programme push-swap 
```go
go run ./push_swap/cmd/main.go "2 1 3 6 5 8"
```

Programme checkers 
```go
go run ./checkers/cmd/main.go "3 2 one 0"
```

### Authors
_______
+ Fabien OLIVIER
+ Rapahel LOVERGNE
+ Jean-Frédéric NANGY
