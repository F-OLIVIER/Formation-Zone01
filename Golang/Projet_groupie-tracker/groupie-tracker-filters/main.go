package main

import (
	"fmt"
	GT "groupie-tracker-filters/packages/api"
	HD "groupie-tracker-filters/packages/handler"
	"net/http"
)

func main() {

	// Appel du style.css
	fs := http.FileServer(http.Dir("templates/"))
	http.Handle("/templates/", http.StripPrefix("/templates/", fs))

	fmt.Println("http://localhost:8080/home")
	// Gestionnaire de serveur qui appel les Handlers
	http.HandleFunc("/", func(write http.ResponseWriter, read *http.Request) {
		// Appel des informations contenu dans les db
		erreur := GT.Appeldb()
		// gestion erreur 500 si db non charger
		if erreur == true {
			HD.Error500Handler(write, read)
			http.ListenAndServe(":8080", nil)
			return
		}
		// lecture des URL.Path
		switch read.URL.Path {
		case "/":
			HD.HomeHandler(write, read)
		case "/home":
			HD.HomeHandler(write, read)
		case "/groupe":
			HD.GroupeHandler(write, read)
		case "/search":
			HD.SearchHandler(write, read)
		default:
			// Gestion erreur 404
			HD.Error404Handler(write, read)
		}

	})
	http.ListenAndServe(":8080", nil)

}
