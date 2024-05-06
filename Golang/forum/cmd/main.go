package main

import (
	"fmt"
	"forum/handlers"
	data "forum/internal"
	"net/http"
)

const port = ":8080"

func main() {
	// Initialisation de la database
	data.Createdb()

	// En fonction de l'URL, chargement du Handler correspondant
	http.HandleFunc("/", handlers.HomeHandler)

	// page du forum
	// http.HandleFunc("/ViewPostsList", handlers.ViewPostsListHandler)
	http.HandleFunc("/ViewPost", handlers.ViewPostHandler)
	http.HandleFunc("/PostEditor", handlers.PostEditorHandler)

	// Page de gestion utilisateurs
	http.HandleFunc("/Login", handlers.LoginHandler)
	http.HandleFunc("/Register", handlers.RegisterHandler)
	http.HandleFunc("/Compte", handlers.CompteHandler)
	http.HandleFunc("/ForgetPassword", handlers.ForgetPasswordHandler)

	// Appel des fichiers annexes
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Mise en écoute du serveur
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(port, nil)
}
