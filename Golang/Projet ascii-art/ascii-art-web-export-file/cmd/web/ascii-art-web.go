package main

import (
	"fmt"
	"net/http"
	"text/template"

	"ascii_web/internal/handlers" // Import des handlers
)

// Définition du port à ouvrir
const port = ":8080"

func main() {
	fs := http.FileServer(http.Dir("template/"))
	http.Handle("/template/", http.StripPrefix("/template/", fs))

	// r.URL.Path

	// HandleFunc met en lien l'adresse URL et la func du template.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path { // Gestion de l'erreur 404.
		case "/":
			handlers.Home(w, r) // Affichage du handler "home".
		default:
			// w.WriteHeader(http.StatusNotFound)
			if r.URL.Path != "/favicon.ico" {
				templ, err := template.ParseFiles("template/404.page.html")
				fmt.Println("404 : Page Not Found")

				if err != nil { // Gestion de l'erreur 404
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				templ.Execute(w, nil)
			}
		}
	})
	fmt.Println("(http://localhost:8080/) - Server started on port", port)
	http.ListenAndServe(port, nil) // Ouverture du port renseigné.
}
