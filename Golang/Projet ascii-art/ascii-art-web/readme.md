# Ascii Art Program

## Descriptif
_______
Programme qui génére un Ascii-Art sur une page internet.

## Détails de l'implémentation : algorithme

L'algorithme implémenté dans ce code est responsable de la gestion des requêtes HTTP et de la génération de pages HTML dynamiques en fonction des entrées reçues. Le code utilise le langage de programmation Go et le package standard `net/http` pour créer un serveur HTTP.

### Aperçu de l'algorithme

1. La fonction principale (`main`) configure le serveur pour écouter le port `9999` et enregistre la fonction `homeHandler` pour traiter les requêtes faites à l'endpoint `/home`.

2. La fonction `homeHandler` est responsable du traitement des requêtes GET et POST vers l'endpoint `/home`. Elle récupère les paramètres `Word` et `Option` de la requête et effectue différentes opérations en fonction de la méthode de la requête.

3. Si la requête est une requête GET, les paramètres `Word` et `Option` sont extraits de la chaîne de requête URL.

4. Si la requête est une requête POST, les données du formulaire sont analysées à l'aide de la fonction `ParseForm()`. Le paramètre `Word` est vérifié pour détecter la présence de caractères non pris en charge à l'aide d'une expression régulière, et s'il est valide, il est assigné à la variable `Input.Word`. Sinon, un message d'erreur est assigné.

5. La valeur de `Input.Word` est ensuite traitée en remplaçant les caractères de nouvelle ligne (`\r\n`) par la séquence d'échappement (`\\n`). Si le paramètre `Word` n'est pas vide, la fonction `PrintAll` du package `fs` est appelée pour générer la valeur de `AsciiArt`.

6. Le template `index.html` est analysé, et la page HTML résultante est générée en exécutant le template avec la structure `Input` en tant que données.

7. La fonction `codestatus` est appelée pour gérer les codes de statut en fonction du type de requête et des éventuelles erreurs survenues lors du traitement de la requête.

### Dépendances et bibliothèques

L'implémentation repose sur les packages suivants :
- `"fmt"` : Fournit des opérations d'entrée/sortie formatée pour l'affichage.
- `"fs"` : Fait référence à un package personnalisé, mis en œuvre pour la génération d'art ASCII.
- `"html/template"` : Permet l'analyse et l'exécution de templates HTML.
- `"net/http"` : Offre des fonctionnalités de client et de serveur HTTP.
- `"regexp"` : Permet l'utilisation d'expressions régulières pour la recherche de motifs.
- `"strings"` : Fournit des fonctions de manipulation de chaînes de caractères.

### Conclusion

Cet algorithme gère efficacement les requêtes GET et POST vers l'endpoint `/home`, récupère les paramètres d'entrée, génère de l'art ASCII et produit des pages HTML dynamiques. Il inclut la gestion des erreurs pour les caractères non pris en charge et les scénarios de méthode non autorisés. L'algorithme mis en œuvre permet aux utilisateurs d'interagir avec le site Web et de recevoir des représentations visuelles de leur entrée.


## Usage
_______
```go
go run ./main/
```
Puis ouvrir le lien HTML local du serveur [http://localhost:9999/home](http://localhost:9999/home)

Entrez le texte à transformer en Ascii-Art, selectionner l'option désirée, puis cliquez sur le bouton "Réaliser l'Ascii-Art"

## Authors
_______
+ Ludovic UNTEREINER
+ Chloé MASSE
+ Fabien OLIVIER
