package main

import (
	"fmt"
	"os"

	CL "net-cat/Packages/client"
)

// Command : go run ./TCPChat.go
// Executable : go build TCPChat.go --> puis --> ./TCPChat

func main() {
	// Récupération du port en entrée :
	// Si un port est saisie, vérification de la validité de l'entrée et récupération
	// Si pas de port saisie, le port par defaut est defini sur 8989
	port := "8989"
	if len(os.Args[1:]) == 1 {
		portint := 0
		for _, j := range os.Args[1] {
			// Check si un caractére n'est pas numérique
			if j < '0' || j > '9' {
				fmt.Println("Mauvais format d'entrée : le port contiens un caractére non numérique.")
				fmt.Println("Exemple d'usage : [USAGE]: ./TCPChat $port")
				os.Exit(1)
			}
			// récupération de l'int
			portint = portint*10 + int(j-'0')
		}
		// Check si le port est trop élevé (supérieur à 65535)
		if portint > 65535 {
			fmt.Println("Mauvais format d'entrée : le port choisie est trop élevé.")
			fmt.Println("Exemple d'usage : [USAGE]: ./TCPChat $port")
			os.Exit(1)
		}
		port = os.Args[1]
	} else if len(os.Args[1:]) > 1 {
		// Si trop d'argument en entrée
		fmt.Println("Mauvais format d'entrée : trop d'argument.")
		fmt.Println("Exemple d'usage : [USAGE]: ./TCPChat $port")
		os.Exit(1)
	}

	// demarrage du serveur
	CL.StartServer(port)
}
