package CL

import (
	"fmt"
	"log"
	"net"
	"os"

	Annexe "net-cat/Packages/annexe"
	GC "net-cat/Packages/goroutineClient"
	ST "net-cat/Packages/struct"
)

// Fonction de demarrage du serveur
func StartServer(port string) {
	// imprime l'IPv4 du serveur
	fmt.Printf("Server launched, use one of this option to connect client\n")
	fmt.Printf("Net-cat linux client              : \x1b[38;2;255;255;0m"+"nc %v %v \x1b[0m\n", Annexe.GetOutboundIP(), port)
	fmt.Printf("Custom client (without interface) : \x1b[38;2;255;255;0m"+"go run . %v %v \x1b[0m\n", Annexe.GetOutboundIP(), port)
	fmt.Printf("Custom client (with interface)    : \x1b[38;2;255;255;0m"+"go run . %v %v -i\x1b[0m\n", Annexe.GetOutboundIP(), port)

	// Initialisation du fichier de log
	logfile, errlog := os.OpenFile("./log/logs.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if errlog != nil {
		fmt.Println("Error fichier de log général :", errlog)
		os.Exit(1)
	}
	log.SetOutput(logfile)
	// Print du demmarrage du serveur dans le fichier de log
	log.Printf("--------------------------------")
	log.Printf("----- Demarrage du serveur -----")
	log.Printf("--------------------------------")

	// Création d'un groupe initial [0]Global
	ST.Group[1] = "Global"
	ST.NbMmebresGroupe["Global"] = 2
	// Création du fichier historique du groupe Global
	os.Create("./log/historiqueGlobal.txt")

	// Lancement du serveur TCP et écoute sur le port définie (port par defaut si non définie)
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error ouverture du listener : ", err)
		os.Exit(1)
	}

	// Boucle pour chaque nouvelle utilisateur
	for {
		// Acceptation de la connexion client (max 10 connexion)
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error acceptation listener : ", err)
			os.Exit(1)
		}
		// Gestion du nombre de connexion max (max 10 clients)
		if len(ST.Clients) > 9 {
			conn.Write([]byte("Nombre de client maximum atteint"))
			log.Printf("Client refused (maximum number of clients reached)")
			// si plus de 10 client, fermeture de la connexion
			conn.Close()
			continue
		}

		// création d'une goroutine par client
		go GC.ProcessClient(conn)
	}
}
