package main

import (
	"fmt"
	"fs"
	"html/template"
	"net/http"
	"regexp"
	"strings"
)

type Page struct {
	Word     string
	Option   string
	AsciiArt string
}

func main() {
	fmt.Println("http://127.0.0.1:9999/home")
	// Gestionnaire de serveur qui appel la fonction homeHandler en /home
	http.HandleFunc("/home", homeHandler)
	// On indique le port à utilisé par le serveur local
	http.ListenAndServe(":9999", nil)
}

func homeHandler(write http.ResponseWriter, read *http.Request) {
	var Input Page
	verif400 := true
	verif405 := true

	switch read.Method {
	case "GET":
		// Récupération des argument Word et Option si méthode 'Get' utilisé
		Input = Page{Word: read.URL.Query().Get("word"), Option: read.URL.Query().Get("option")}
	case "POST":
		// Récupération des argument Word et Option si méthode 'Post' utilisé
		// Appelez ParseForm() pour analyser la requête brute et mettre à jour mes données
		if err := read.ParseForm(); err != nil {
			fmt.Fprintf(write, "ParseForm() err: %v", err)
		}

		// gestion des caractéres spéciaux en entrée
		pattern := regexp.MustCompile(`^[ -~\r\n]+$`)
		// Récupération de l'argument Word en méthode 'Post' avec vérification de la saisie dans la case
		if pattern.MatchString(read.FormValue("word")) {
			// Si pas d'erreur, lecture des valeurs du formulaire
			Input.Word = read.FormValue("word")
		} else {
			// Si erreur, print de l'erreur en Ascii Art
			Input.Word = "Error 400\\nUnsupported character"
			verif400 = false
		}
		// Récupération de l'argument Option en méthode 'Post'
		Input.Option = read.FormValue("option")
	default:
		verif405 = false
		fmt.Fprintf(write, "Error 405 Method Not Allowed")
	}

	// gestion des retours a la ligne avec "enter" dans la zone de texte
	Input.Word = strings.ReplaceAll(Input.Word, "\r\n", "\\n")
	// Si Word n'est pas vide, on récupére le contenu de word et option
	if Input.Word != "" {
		Input.AsciiArt = fs.PrintAll(string(Input.Word), string(Input.Option))
	}

	// On génére un template golang du fichier index.html
	// Must == Gestion des erreurs
	// ParseFiles == crée un modéle html pour go
	var templates = template.Must(template.ParseFiles("./templates/index.html"))
	// On génére la page Html avec les Input
	err2 := templates.ExecuteTemplate(write, "index.html", Input)
	// Gestion des erreurs casse bonbon
	if err2 != nil {
		fmt.Println(err2)
		// Print de l'erreur sur la page HTML
		http.Error(write, err2.Error(), http.StatusInternalServerError)
	}

	// Gestion des code status
	// Vérification si post existe
	read.ParseForm()
	postok := read.PostForm.Has("word")
	if postok == true {
		// récupération du code status et retour des erreurs avec print en console
		codestatus(verif400, verif405, write)
	}
}

// Gestion du status code
func codestatus(verif400 bool, verif405 bool, write http.ResponseWriter) {
	// Récupération des erreurs généré automatiquement
	resp, err := http.Get("http://127.0.0.1:9999/home")
	if err != nil {
		fmt.Println(err)
	}
	// Enregistrement du code erreur
	erreur := resp.StatusCode

	// Check des erreurs manuel puis auto et print en console
	if verif400 == false {
		fmt.Println("The status code we got is: 400 Unsupported character")
		fmt.Fprintf(write, "<center><font size= \"30\"><mark style=\"background-color: red;\">Sorry: error 400 Unsupported character</mark></font></center>")
	} else if verif405 == false {
		fmt.Println("The status code we got is: 405 Bad Request")
		fmt.Fprintf(write, "<center><font size= \"30\"><mark style=\"background-color: red;\">Sorry, error 405 Method not allowed</mark></font></center>")
	} else {
		fmt.Println("The status code we got is:", erreur, http.StatusText(erreur))
		if erreur != 200 {
			// print des erreurs sur la page internet
			fmt.Fprintf(write, "<center><font size= \"30\"><mark style=\"background-color: red;\">"+string(erreur)+http.StatusText(erreur)+"</mark></font></center>")

		}
	}
}
