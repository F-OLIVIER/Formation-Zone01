package ST

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	Annexe "net-cat/Packages/annexe"
)

// Structure client
type Client struct {
	Pseudo    string
	Connexion net.Conn
	Group     int
}

var Clients []Client

var Group = make(map[int]string)
var NbMmebresGroupe = make(map[string]int)

// Fonction qui supprime un client de la base de la struct client
func RemoveClient(pseudo string) {
	var tempClient []Client
	for a := 0; a < len(Clients); a++ {
		if Clients[a].Pseudo != pseudo {
			tempClient = append(tempClient, Clients[a])
		}
	}
	Clients = tempClient
}

// Fonction qui demande au client un pseudo non vide, puis un nom de groupe de discussion pour ajoute le client à la struct
func AddClient(conn net.Conn, buf *bufio.Reader) (pseudo string, groupe int) {
	pseudoValide := false
	intGroupe := 1
	for pseudoValide == false {
		conn.Write([]byte(Annexe.ColorUser))
		conn.Write([]byte("[ENTER YOUR NAME]~ ")) // Print ligne d'attente du name
		conn.Write([]byte(Annexe.Fincolor))
		pseudoinitial, err := buf.ReadString('\n')

		if err != nil {
			fmt.Println("│ Client party when entering name")
			break
		}
		arraypseudo := []rune(pseudoinitial)
		pseudo = string(arraypseudo[0 : len(arraypseudo)-1])
		pseudo = strings.TrimSpace(pseudo)
		if len(pseudo) > 0 {
			// gestion du groupe de discussion
			// Print a l'utilisateur la liste des groupes
			// Affiche en console
			// ┌------------------------------------------------┐
			// │--► LISTE DES GROUPES DE DISCUSSION EXISTANT ◄--│
			// │------------------------------------------------│
			// GROUPE N° XX, nom du groupe : XXXXX
			// │------------------------------------------------│
			// |------ POUR CHOISIR UN GROUPE, AU CHOIX : ------|
			// |-- SAISSISEZ LE N° DU GROUPE POUR LE REJOINDRE -|
			// |---- CREER UN GROUPE EN SAISISSANT SON NOM -----|
			// └------------------------------------------------┘
			// conn.Write([]byte(Annexe.ColorHistorique))
			conn.Write([]byte(Annexe.ColorHistorique + "┌------------------------------------------------┐\n│------ LIST OF EXISTING DISCUSSION GROUPS ------│\n│------------------------------------------------│\n"))
			for a, b := range Group {
				conn.Write([]byte("│--►  GROUP N° " + Annexe.Fincolor + Annexe.Convstring(a) + Annexe.ColorHistorique + " --> GROUP NAME : " + Annexe.Fincolor + b + Annexe.ColorHistorique + "\n"))
			}
			conn.Write([]byte("│------------------------------------------------│\n"))
			conn.Write([]byte("│-------------- TO CHOOSE A GROUP: --------------│\n"))
			conn.Write([]byte("│-------- ENTER THE GROUP NUMBER TO JOIN --------│\n"))
			conn.Write([]byte("│----- CREATE A GROUP BY ENTERING ITS NAME ------│\n"))
			conn.Write([]byte("└------------------------------------------------┘\n"))
			conn.Write([]byte(Annexe.Fincolor))
			conn.Write([]byte(Annexe.ColorUser))
			conn.Write([]byte("[CHOICE OF GROUP]~ ")) // Print ligne d'attente du groupe
			conn.Write([]byte(Annexe.Fincolor))

			nameGroupeInitial, err := buf.ReadString('\n')
			if nameGroupeInitial == "\n" { // gestion d'erreur du client gocui qui envoie au serveur une ligne vide "\n"
				nameGroupeInitial, err = buf.ReadString('\n')
			}
			if err != nil {
				fmt.Println("│ Client party when entering group name")
				break
			}

			// récupération du nom du groupe
			arraynameGroupe := []rune(nameGroupeInitial)
			// retrait du \n est des espaces
			nameGroupe := string(arraynameGroupe[0 : len(arraynameGroupe)-1])
			nameGroupeSansEspace := strings.TrimSpace(nameGroupe)

			// Conversion en int si numéro du groupe saisie
			errConvInt := ""
			intGroupe, errConvInt = Annexe.Convint(nameGroupeSansEspace)
			// Gestion de la création ou non d'un groupe de discussion (si intGroupe non valide)
			if errConvInt != "" { // Création d'un groupe
				Group[len(Group)+1] = nameGroupe
				intGroupe = len(Group)
				// Création du fichier de log du groupe
				os.Create("./log/historique" + nameGroupeSansEspace + ".txt")
				// Création du compteur de nombre de membre dans le groupe
				NbMmebresGroupe[nameGroupe] = 1
			} else if nameGroupeSansEspace == "" || intGroupe > len(Group) { // Si pas de groupe saisie ou si intGroupe non valide, rejoindre le groupe Global
				intGroupe = 0
				NbMmebresGroupe["Global"]++
				conn.Write([]byte("[INCORRECT ENTRY, DEFAULT GROUP IS]: Global\n"))
			} else {
				// sinon intGroupe est dejà valide, pas de modification de intGroupe
				NbMmebresGroupe[nameGroupe]++
			}

			conn.Write([]byte(Annexe.ColorArriveUser + "[YOU HAVE JOINED THE FOLLOWING GROUP]: " + Annexe.Fincolor))
			conn.Write([]byte(Group[intGroupe] + "\n"))
			// fmt.Println("group : ", Group)
			// fmt.Println("Group[intGroupe] : ", Group[intGroupe])

			// Ajout du nouveau client dans la struct
			var newClient Client
			newClient.Pseudo = pseudo
			newClient.Connexion = conn
			newClient.Group = intGroupe
			Clients = append(Clients, newClient)

			// liste des utilisateurs présent dans le groupe
			listeMember := "[LIST OF MEMBERS IN THE GROUP]: "
			for _, v := range Clients {
				if v.Group == intGroupe {
					listeMember += v.Pseudo + "; "
				}
			}
			conn.Write([]byte(listeMember + "\n"))

			// Fermeture de la boucle
			pseudoValide = true
		} else {
			conn.Write([]byte("Incorrect username\n"))
		}
	}
	return pseudo, intGroupe
}

func ChangementPseudoClient(i int, newname string) {
	// Changement du pseudo client dans la struct
	var changementClient Client
	changementClient.Pseudo = newname
	changementClient.Connexion = Clients[i].Connexion
	changementClient.Group = Clients[i].Group
	Clients[i] = changementClient
}
