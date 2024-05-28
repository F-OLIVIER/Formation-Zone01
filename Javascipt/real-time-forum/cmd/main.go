package main

import (
	"flag"
	"fmt"
	handlers "forum/handlers"
	data "forum/internal"
	utils "forum/middleware"
	"log"
	"net/http"
	"time"
)

const port = "8080"

var addr = flag.String("addr", ":"+port, "http service address")

func main() {
	// Initialisation de la database
	data.Createdb()

	flag.Parse()
	http.HandleFunc("/", handlers.ServeHome)

	endpoints := []string{
		"/api/home", "/api/ViewPost", "/api/PostEditor", "/api/NewComment", // page du Forum
		"/api/Login", "/api/Register", "/api/Compte", // page Utilisateurs
		"/api/historyMsg", // gestion du Chat
	}
	for _, endpoint := range endpoints {
		http.HandleFunc(endpoint, handlers.Handler)
	}

	// Appel des fichiers annexes
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./assets/static/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js/"))))
	http.Handle("/assets/img/", http.StripPrefix("/assets/img/", http.FileServer(http.Dir("./assets/img/"))))
	http.Handle("/assets/photoCompte/", http.StripPrefix("/assets/photoCompte/", http.FileServer(http.Dir("./assets/photoCompte/"))))

	server := &http.Server{
		Addr:              *addr,
		ReadHeaderTimeout: 3 * time.Second,
	}
	fmt.Println("Server started on port " + port + " : http://localhost:" + port)

	// mise en écoute du WebSocket
	http.HandleFunc("/ws", utils.HandleConnections)
	go utils.HandleMessages()

	http.ListenAndServe(":"+port, nil)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
