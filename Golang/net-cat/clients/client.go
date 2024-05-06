package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/jroimartin/gocui"
)

var wg sync.WaitGroup
var Conn net.Conn
var Pseudo string

func main() {
	adresse, inter := checkEnter(os.Args[1:])

	// Connexion au serveur
	var errD error
	Conn, errD = net.Dial("tcp", adresse)
	erreur(errD, "Probléme Conn")

	if inter { // client avec interface
		// debut de l'interface utilisateur (gocui)
		g, err := gocui.NewGui(gocui.OutputNormal)
		if err != nil {
			// wg.Done()
			log.Fatalln(err)
		}
		g.Cursor = true
		// g.Mouse = true

		// Mettez à jour les vues lorsque le terminal change de taille.
		g.SetManagerFunc(layout)

		if err := initKeybindings(g); err != nil {
			log.Fatalln(err)
		}
		_, err = g.SetViewOnTop("scroll")

		// lancement du client NC
		go clientNCai(g, Conn)

		if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
			// wg.Done()
			log.Fatalln(err)
		}

		defer g.Close()
		// fin de l'interface

		// chargement client (udate de l'interface)
	} else { // client sans interface
		clientNCsi(Conn)
	}
	// wg.Wait()
}

// -----------------------------------------------------------------------
// ------------------------ client sans interface ------------------------
// -----------------------------------------------------------------------
func clientNCsi(conn net.Conn) {
	for {
		go func() { // goroutine qui gére l'a réception depuis le serveur
			_, errE := io.Copy(os.Stdout, conn)
			erreur(errE, "Réception depuis le serveur, (os.Stdout)")
		}() // par defaut émission d'information d'information
		_, errR := io.Copy(conn, os.Stdin)
		erreur(errR, "émission vers le serveur, (os.Stdin)")
	}
}

// -----------------------------------------------------------------------
// ------------------------ client avec interface ------------------------
// -----------------------------------------------------------------------
func clientNCai(g *gocui.Gui, conn net.Conn) {
	// Attente du chargement de l'interface
	time.Sleep(200 * time.Millisecond)
	var information []string

	go func() { // goroutine qui gére la réception depuis le serveur
		// *************** update discussion ***************
		scanner := bufio.NewScanner(conn)
		scanner.Split(ScanLinesWithCR)
		for scanner.Scan() { // true
			line := scanner.Text()
			var member, attenteLine bool
			line = supColor(line)
			information, member, attenteLine = traitement_donnee(line, information, g)
			if member { // update encart information
				updateInfo(g, information)
			} else if attenteLine { // line d'attente, ne rien faire pour la supprimer

			} else { // Print des informations serveur tel quel
				line = supColor(line)
				updateDiscussion(g, line)
				// Time d'animation de la réception des informations serveur
				time.Sleep(40 * time.Millisecond)
			}
		}
		// *********************************************
	}()
}

// Fonction de récupération de l'input via le buffer lorsque "Enter" est saisie au clavier
func recuperationBuffer(g *gocui.Gui, v *gocui.View) error {
	var buf string
	// récupération du buffer (saisie de l'utilisateur)
	v, err := g.View("input")
	erreur(err, "input")
	buf = v.Buffer()

	// clean de l'encart d'input suite au "Enter"
	updateInput(g)

	// Ajout de l'horodatage dans la fenêtre de l'utilisateur avec pseudo (ou pas)
	horaire := time.Now().Format("[2006-01-02 15:04:05]")
	var horodatage string
	if Pseudo == "" {
		horodatage = "\x1b[36m" + horaire + ": \x1b[0m"
	} else {
		horodatage = "\x1b[36m" + horaire + "[" + Pseudo + "]:" + "\x1b[0m "
	}

	// Copie de l'input dans la fenêtre de l'utilisateur
	tmpbuf := strings.ReplaceAll(buf, "\n", "")
	updateDiscussion(g, horodatage+tmpbuf)
	time.Sleep(40 * time.Millisecond)

	// Envoie de l'input au serveur
	Conn.Write([]byte(string(buf)))

	return nil
}

// Fonction qui gére le caractére de fin de ligne '\n' par defaut et '~' si '\n' absent
func ScanLinesWithCR(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		return i + 1, data[0:i], nil
	}

	if i := bytes.IndexByte(data, '~'); i >= 0 {
		return i + 2, data[0:i], nil // i+2 = suppression de l'espace aprés le '~'
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

// ---------------------------------------------------------------------------------
// --------------- Fonction de traitement des données reçu du serveur --------------
// ---------------------------------------------------------------------------------

var inDiscussion bool

// permet d'orienter dans quel encart de l'interface ce qui est reçu du serveur
func traitement_donnee(line string, information []string, g *gocui.Gui) ([]string, bool, bool) {
	if !inDiscussion {
		if strings.Contains(line, "[YOU HAVE JOINED THE FOLLOWING GROUP]:") { // gestion du groupe
			tmpline := strings.ReplaceAll(line, "[YOU HAVE JOINED THE FOLLOWING GROUP]:", "")
			tmp := "\x1b[35mDISCUSSION GROUP\x1b[0m\n" + tmpline + "\n\n\x1b[35mLIST OF MEMBERS\x1b[0m"
			information = append(information, tmp)
			return information, false, false

		} else if strings.Contains(line, "[ENTER YOUR NAME]") || strings.Contains(line, "[CHOICE OF GROUP]") { // Ajout couleur question posé à l'utilisateur
			lineColor := "\x1b[46m" + line + "\x1b[0m"
			updateDiscussion(g, lineColor)
			return information, false, true

		} else if strings.Contains(line, "│----------- !name your nickname -----------│") { // gestion du groupe
			inDiscussion = true
			return information, false, false
		} else if strings.Contains(line, "[LIST OF MEMBERS IN THE GROUP]:") { // liste des membres deja connecté a l'arrivé de l'utilisateur
			tmpline := strings.ReplaceAll(line, "[LIST OF MEMBERS IN THE GROUP]:", "")
			tmp := ""
			for _, value := range tmpline {
				if value == ';' {
					information = append(information, tmp)
					// récupération du pseudo utilisateur
					Pseudo = tmp
					// supression des espaces éventuelle en trop debut et fin du pseudo
					if len(Pseudo) > 0 && Pseudo[0] == ' ' {
						Pseudo = Pseudo[1:]
					}
					if len(Pseudo) > 0 && Pseudo[len(Pseudo)-1] == ' ' {
						Pseudo = Pseudo[:len(Pseudo)-1]
					}
					// re-initialisation de tmp si un autres utilisateur dans la liste
					tmp = ""
				} else {
					tmp += string(value)
				}
			}
			return information, true, false

		}
	} else {
		if strings.Contains(line, "["+Pseudo+"]:") { // Détection des ligne d'attente a supprimer
			return information, false, true

		} else if strings.Contains(line, "has changed nickname to") { // Ajout couleur Changement de pseudo
			lineColor := "\x1b[32m" + line + "\x1b[0m"
			updateDiscussion(g, lineColor)
			indexnameMember := strings.Index(line, "changed nickname to")
			oldNameMenber := line[:indexnameMember]
			newNameMember := line[indexnameMember+19:]
			// Parcourir information pour mettre à jour l'encart information
			for i, v := range information {
				if strings.ReplaceAll(oldNameMenber, " ", "") == strings.ReplaceAll(v, " ", "") {
					information[i] = newNameMember
				}
			}
			return information, true, false

		} else if strings.Contains(line, "has joined our chat...") { // Nouvelle utilisateur qui rejoint
			linecolor := "\x1b[32m" + line + "\x1b[0m "
			updateDiscussion(g, linecolor)
			nameMember := line
			nameMember = strings.ReplaceAll(nameMember, "User", "")
			nameMember = strings.ReplaceAll(nameMember, " has joined our chat...", "")
			information = append(information, nameMember)
			return information, true, false

		} else if strings.Contains(line, "has left our chat") { // utilisateur qui part du chat
			linecolor := "\x1b[31m" + line + "\x1b[0m "
			updateDiscussion(g, linecolor)
			nameMember := line
			nameMember = strings.ReplaceAll(nameMember, "User ", "")
			nameMember = strings.ReplaceAll(nameMember, "has left our chat...", "")
			nameMember = strings.ReplaceAll(nameMember, " ", "")
			nameMember = strings.ReplaceAll(nameMember, "\n", "")

			var tmp_information []string
			for i := 0; i < len(information); i++ {
				tmp := strings.ReplaceAll(information[i], " ", "")
				if tmp == nameMember {
					tmp_information = information[:i]
					for j := i + 1; j < len(information); j++ {
						tmp_information = append(tmp_information, information[j])
					}
					time.Sleep(10 * time.Millisecond)
					return tmp_information, true, false
				}
			}
		}
	}
	return information, false, false
}

// Fonction de suppression des couleurs (cause des erreurs)
func supColor(line string) string {
	var colorUser = "\x1b[38;2;0;240;255m"
	line = strings.ReplaceAll(line, colorUser, "")
	var colorMessageBienvenue = "\x1b[38;2;0;255;4m"
	line = strings.ReplaceAll(line, colorMessageBienvenue, "")
	var colorHistorique = "\x1b[38;2;255;151;0m"
	line = strings.ReplaceAll(line, colorHistorique, "")
	var colorArriveUser = "\x1b[38;2;0;255;4m"
	line = strings.ReplaceAll(line, colorArriveUser, "")
	var colorDepartUser = "\x1b[38;2;255;0;0m"
	line = strings.ReplaceAll(line, colorDepartUser, "")
	var yellow = "\x1b[38;2;255;255;0m"
	line = strings.ReplaceAll(line, yellow, "")
	var fincolor = "\x1b[0m"
	line = strings.ReplaceAll(line, fincolor, "")
	line = strings.ReplaceAll(line, "[38;", "")
	return line
}

// -----------------------------------------------------------------------

func checkEnter(args []string) (string, bool) {
	var adresse string
	port := ":8989"
	if len(args) > 3 || len(args) <= 1 { // check taille entrée
		fmt.Println("Erreur nombre d'argument en saisie")
		fmt.Println("Exemple : ./client <adresse> <port> <option \"-i\" si interface souhaité>")
		// wg.Done()
		os.Exit(1)
	}

	adresse = args[0]
	for _, v := range adresse { // check adresse
		if (v >= '0' && v <= '9') || v == '.' {
			// ne rien faire
		} else {
			fmt.Println("Erreur check adresse")
			fmt.Println("Exemple : ./client <adresse> <port> <option \"-i\" si interface souhaité>")
			// wg.Done()
			os.Exit(1)
		}
	}

	if len(args[1]) > 0 { // check port
		port = ":"
		for _, v := range args[1] {
			if v >= '0' && v <= '9' {
				// ne rien faire
				port += string(v)
			} else {
				fmt.Println("Erreur check port")
				fmt.Println("Exemple : ./client <adresse> <port> <option \"-i\" si interface souhaité>")
				// wg.Done()
				os.Exit(1)
			}
		}
	}
	adresse += port

	if len(args) == 3 && (args[2] == "-i" || args[2] == "-I") {
		return adresse, true
	}

	return adresse, false
}

func erreur(err error, str string) {
	if err != nil {
		fmt.Printf("Erreur, %v : \n%v\n", str, err)
		// wg.Done()
		os.Exit(1)
	}
}

// -----------------------------------------------------
// ----------------- Interface (gocui) -----------------
// -----------------------------------------------------
func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	// Espace du haut
	if v, err := g.SetView("Title", 0, 0, maxX-1, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		msg := " TCP Chat in golang "
		// calcul du centrage du titre
		debut := (maxX - len(msg)) / 2
		for debut > 0 {
			msg = " " + msg
			debut--
		}
		// Print d'un titre centré
		fmt.Fprintln(v, msg)
	}

	// Espace latéral Gauche (zone d'informations)
	if v, err := g.SetView("Info", 0, 3, 22, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = " Informations "
	}

	// Espace latéral milieu droite (zone de discussion)
	if v, err := g.SetView("discussion", 23, 3, maxX-1, maxY-4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = " discussion "
		v.Autoscroll = true
	}

	// Espace latéral bas droite (zone de saisie)
	if v, err := g.SetView("input", 23, maxY-3, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if _, err := g.SetCurrentView("input"); err != nil {
			return err
		}
		v.Title = " Zone de saisie "
		v.Editable = true
		v.Wrap = true
	}

	// Espace latéral haut droite (zone d'explication scroll)
	if v, err := g.SetView("scroll", maxX-20, 4, maxX-2, 10); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "\n\x1b[35m  Key to scroll\n")
		fmt.Fprintln(v, "  ↑      Up ")
		fmt.Fprintln(v, "  ↓     Down ")
		fmt.Fprintln(v, " ← →  Autoscroll\x1b[0m")
	}

	return nil
}

// Fonction d'update de "discussion"
func updateDiscussion(g *gocui.Gui, line string) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("discussion")
		if err != nil {
			fmt.Fprintln(v, "Erreur détecté : autodestuction de l'interface dans 2 minutes")
		}
		// fmt.Fprintln(v, "%q", line)
		fmt.Fprintln(v, line)
		return nil
	})
}

// Fonction d'update de "Informations"
func updateInfo(g *gocui.Gui, information []string) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("Info")
		if err != nil {
			fmt.Fprintln(v, "Erreur détecté : autodestuction de l'interface dans 2 minutes")
		}
		v.Clear()
		fmt.Fprintln(v, "\n")
		for _, value := range information {
			fmt.Fprintln(v, value)
		}
		// fmt.Fprintln(v, "\x1b[0m")
		// fmt.Fprintln(v, "%q", information)
		return nil
	})
}

// Fonction d'update qui clean l'encart de saisie "input"
func updateInput(g *gocui.Gui) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("input")
		if err != nil {
			fmt.Fprintln(v, "Erreur détecté : autodestuction de l'interface dans 2 minutes")
		}
		v.Clear()       // clean la ligne sans déplacer le curseur
		v.EditNewLine() // raméne le curseur en début de ligne
		return nil
	})
}

func initKeybindings(g *gocui.Gui) error {
	// Quitter l'interface si "ctrl + c"
	if err := g.SetKeybinding("input", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	// récupération du buffer si "enter"
	if err := g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, recuperationBuffer); err != nil {
		log.Fatalln(err)
	}
	// activation autoscroll si fléche droite
	if err := g.SetKeybinding("input", gocui.KeyArrowRight, gocui.ModNone, autoscroll); err != nil {
		return err
	}
	// activation autoscroll si fléche gauche
	if err := g.SetKeybinding("input", gocui.KeyArrowLeft, gocui.ModNone, autoscroll); err != nil {
		return err
	}
	// scroll vers le haut si fléche up
	if err := g.SetKeybinding("input", gocui.KeyArrowUp, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			w, err := g.View("discussion")
			if err != nil {
				fmt.Fprintln(w, "Erreur détecté : probleme scroll KeyArrowDown")
			}
			scrollView(w, -1)
			return nil
		}); err != nil {
		return err
	}
	// scroll vers le bas si fléche down
	if err := g.SetKeybinding("input", gocui.KeyArrowDown, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			w, err := g.View("discussion")
			if err != nil {
				fmt.Fprintln(w, "Erreur détecté : probleme scroll KeyArrowDown")
			}
			scrollView(w, 1)
			return nil
		}); err != nil {
		return err
	}
	return nil
}

// Fonction qui permet de quitter l'interface via le "crtl + c"
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// Fonction de scroll (Attention : desactive autoscroll)
func scrollView(v *gocui.View, dy int) error {
	if v != nil {
		v.Autoscroll = false
		ox, oy := v.Origin()
		if err := v.SetOrigin(ox, oy+dy); err != nil {
			return err
		}
	}
	return nil
}

// Fonction pour réactiver l'autoscroll apres déplacement avec les fléches up ou down
func autoscroll(g *gocui.Gui, v *gocui.View) error {
	w, err := g.View("discussion")
	if err != nil {
		fmt.Fprintln(w, "Erreur détecté : probleme autoscroll")
	}
	w.Autoscroll = true
	return nil
}
