package GC

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	Annexe "net-cat/Packages/annexe"
	ST "net-cat/Packages/struct"
)

// création d'un mutex pour les problémes de message en même temps
var M sync.Mutex

// Lancement d'une goroutine pour chaque client
func ProcessClient(conn net.Conn) {
	buf := bufio.NewReader(conn)

	// Print du message de bienvenue
	conn.Write([]byte(Annexe.PenguinWelcome()))

	// Gestion du groupe utilisateur, création du groupe [0]global et gestion du pseudo
	pseudo, idgroupe := ST.AddClient(conn, buf)

	nameGroupeSansEspace := strings.TrimSpace(ST.Group[idgroupe])
	// ouverture du fichier historique du groupe
	historique, errhisto := os.OpenFile("./log/historique"+nameGroupeSansEspace+".txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if errhisto != nil {
		fmt.Println("Error group history file "+ST.Group[idgroupe]+" :", errhisto)
		os.Exit(1)
	}

	// Chargement du fichier historique du groupe
	data, errreadfile := os.ReadFile("./log/historique" + nameGroupeSansEspace + ".txt")
	if errreadfile != nil {
		fmt.Println("Error open group history file "+ST.Group[idgroupe]+" :", errreadfile)
	}
	if string(data) != "" {
		Annexe.SendHistorique(conn, data) // print de l'historique du groupe rejoint
	}

	// Message pour dire comment changer de pseudo
	conn.Write([]byte(Annexe.MessageChangeName()))

	// print de l'arrivé de quelqu'un aux autres clients
	M.Lock() // Blocage du mutexe
	newClient(idgroupe, conn, pseudo, historique)
	M.Unlock() // Déblocage du mutexe

	// Boucle d'attente d'émission de message par un client
	for {
		msg, err := buf.ReadString('\n')
		for msg == "\n" { // gestion d'erreur du client "gocui" qui envoie des ligne vide "\n"
			msg, err = buf.ReadString('\n')
		}

		horaire := time.Now().Format("[2006-01-02 15:04:05]")
		M.Lock() // Blocage du mutexe

		if err != nil {
			// Si err, le client a quitté la console, depart du clients
			clientParti(idgroupe, conn, horaire, pseudo, historique)
			// Debloque le Mutex avant de casser la goroutine
			M.Unlock() // Deblocage du mutexe
			break
		}

		// Check si changement de pseudo "!name" inclus
		changement, newname := Annexe.CheckName(msg)
		if changement == true {
			// print du changement de pseudo aux clients
			if strings.TrimSpace(newname) != "" {
				// si !name non vide
				changementPseudo(idgroupe, conn, newname, horaire, pseudo, historique)
				pseudo = newname
			} else {
				// si !name vide
				conn.Write([]byte("New pseudo incorrect"))
			}
		} else {
			// print d'un message envoyer au autres clients
			printMessageClient(idgroupe, conn, horaire, pseudo, historique, msg)
		}
		M.Unlock() // Déblocage du mutexe
	}
}

func newClient(idgroupe int, conn net.Conn, pseudo string, historique *os.File) {
	for i := 0; i < len(ST.Clients); i++ {
		horaire := time.Now().Format("[2006-01-02 15:04:05]")
		// Si le client est dans le même groupe
		if ST.Clients[i].Group == idgroupe {
			if ST.Clients[i].Connexion != conn {
				// si client différent de l'utilisateur
				ST.Clients[i].Connexion.Write([]byte(Annexe.ColorArriveUser + "\nUser " + pseudo + " has joined our chat...\n" + Annexe.Fincolor)) // Print a chaque client qu'un utilisateur à rejoint le groupe du chat
				ST.Clients[i].Connexion.Write([]byte(Annexe.ColorUser + horaire + "[" + ST.Clients[i].Pseudo + "]: " + Annexe.Fincolor))           // Print d'une ligne d'attente de message à chaque client
			} else {
				// si client est l'utilisateur
				ST.Clients[i].Connexion.Write([]byte(Annexe.ColorUser + horaire + "[" + pseudo + "]: " + Annexe.Fincolor)) // Print d'une ligne d'attente de message
				log.Printf("[Groupe " + ST.Group[idgroupe] + "]" + pseudo + " has joined our chat...\n")                   // Print dans le fichier log
				historique.WriteString("│ [Group " + ST.Group[idgroupe] + "]" + pseudo + "  has joined our chat...\n")     // Print dans l'historique de discussion du groupe
				// fmt.Printf("│" + Annexe.ColorArriveUser + " [Group " + ST.Group[idgroupe] + "] " + pseudo + " has joined our chat...\n" + Annexe.Fincolor) // Print dans la console serveur
			}
		}
	}
}

func clientParti(idgroupe int, conn net.Conn, horaire string, pseudo string, historique *os.File) {
	for i := 0; i < len(ST.Clients); i++ {
		// Si le client est dans le même groupe
		if ST.Clients[i].Group == idgroupe {
			if ST.Clients[i].Connexion != conn {
				// si client différent de l'utilisateur
				ST.Clients[i].Connexion.Write([]byte(Annexe.ColorDepartUser + "\nUser " + pseudo + " has left our chat..." + Annexe.Fincolor + "\n")) // Print a chaque client qu'un utilisateur à quitté le groupe du chat
				ST.Clients[i].Connexion.Write([]byte(Annexe.ColorUser + horaire + "[" + ST.Clients[i].Pseudo + "]: " + Annexe.Fincolor))              // Print d'une ligne d'attente de message à chaque client
			} else {
				// si client est l'utilisateur
				log.Printf("[Group " + ST.Group[idgroupe] + "]" + pseudo + " has left our chat...\n")               // Print dans le fichier log
				historique.WriteString("│ [Group " + ST.Group[idgroupe] + "]" + pseudo + " has left our chat...\n") // Print dans l'historique de discussion du groupe
				// fmt.Printf("│" + Annexe.ColorDepartUser + " [Group " + ST.Group[idgroupe] + "] " + pseudo + " has left our chat...\n" + Annexe.Fincolor) // Print dans la console serveur
			}
		}
	}

	// Suppression du client de la liste du groupe
	ST.NbMmebresGroupe[ST.Group[idgroupe]]--
	// Si plus aucun client dans le groupe, suppression du groupe est du fichier historique du groupe
	if ST.NbMmebresGroupe[ST.Group[idgroupe]] == 0 {
		// Suppression du fichier historique du groupe
		nameGroupeSansEspace := strings.TrimSpace(ST.Group[idgroupe])
		os.Remove("./log/historique" + nameGroupeSansEspace + ".txt")

		// suppression du groupe
		delete(ST.NbMmebresGroupe, ST.Group[idgroupe])
		delete(ST.Group, idgroupe)
	}
	// Suppression du client de la structure "Clients"
	ST.RemoveClient(pseudo)
}

func changementPseudo(idgroupe int, conn net.Conn, newname string, horaire string, pseudo string, historique *os.File) {
	ancienPseudo := ""
	// Boucle confirmation de changement de pseudo
	for i := 0; i < len(ST.Clients); i++ {
		// Si le client est dans le même groupe
		if ST.Clients[i].Group == idgroupe {
			if ST.Clients[i].Connexion == conn {
				// si client est l'utilisateur
				ancienPseudo = ST.Clients[i].Pseudo
				// modification du pseudo dans la structure
				ST.ChangementPseudoClient(i, newname)

				ST.Clients[i].Connexion.Write([]byte(Annexe.ColorArriveUser + ancienPseudo + " has changed nickname to " + newname + Annexe.Fincolor + "\n")) // Print de la confirmation de changement de pseudo
				ST.Clients[i].Connexion.Write([]byte(Annexe.ColorUser + horaire + "[" + ST.Clients[i].Pseudo + "]: " + Annexe.Fincolor))                      // Print d'une ligne d'attente de message
				log.Printf("[Group " + ST.Group[idgroupe] + "] " + ancienPseudo + " has changed his nickname to " + newname)                                  // Print dans le fichier log
				historique.WriteString("│" + ancienPseudo + " has changed nickname to " + newname + "\n")                                                     // Print dans l'historique de discussion du groupe
				// fmt.Println("│" + Annexe.ColorUser + " [Group " + ST.Group[idgroupe] + "] " + ancienPseudo + " has changed nickname to " + newname + Annexe.Fincolor) // Print dans la console serveur
			}
		}
	}

	// Boucle de print du changement de pseudo aux clients
	for i := 0; i < len(ST.Clients); i++ {
		// M.Lock()
		// Si le client est dans le même groupe
		if ST.Clients[i].Group == idgroupe {
			if ST.Clients[i].Connexion != conn {
				// si client différent de l'utilisateur
				ST.Clients[i].Connexion.Write([]byte("\n" + Annexe.ColorArriveUser + ancienPseudo + " has changed nickname to " + newname + "\n" + Annexe.Fincolor)) // Print du changement de pseudo aux autres clients
				ST.Clients[i].Connexion.Write([]byte(Annexe.ColorUser + horaire + "[" + ST.Clients[i].Pseudo + "]: " + Annexe.Fincolor))                             // Print d'une ligne d'attente de message à chaque client
			}
		}
	}
}

func printMessageClient(idgroupe int, conn net.Conn, horaire string, pseudo string, historique *os.File, msg string) {
	for i := 0; i < len(ST.Clients); i++ {
		// Si le client est dans le même groupe
		if ST.Clients[i].Group == idgroupe {
			if ST.Clients[i].Connexion != conn {
				// si client différent de l'utilisateur
				ST.Clients[i].Connexion.Write([]byte("\n" + horaire + "[" + pseudo + "]: " + msg))                                       // Print horodatage, pseudo et message à chaque client
				ST.Clients[i].Connexion.Write([]byte(Annexe.ColorUser + horaire + "[" + ST.Clients[i].Pseudo + "]: " + Annexe.Fincolor)) // Print d'une ligne d'attente de message à chaque client
			} else {
				// si client est l'utilisateur
				ST.Clients[i].Connexion.Write([]byte(Annexe.ColorUser + horaire + "[" + pseudo + "]: " + Annexe.Fincolor)) // Print d'une ligne d'attente de message
				log.Printf("[Group " + ST.Group[idgroupe] + "]" + pseudo + " : " + msg)                                    // Print dans le fichier log
				historique.WriteString("│ " + horaire + "[" + pseudo + "]: " + msg)                                        // Print dans l'historique de discussion du groupe
				// fmt.Printf("│" + Annexe.ColorUser + " [Group " + ST.Group[idgroupe] + "] " + pseudo + " : " + msg + Annexe.Fincolor) // Print dans la console serveur
			}
		}
	}
}
