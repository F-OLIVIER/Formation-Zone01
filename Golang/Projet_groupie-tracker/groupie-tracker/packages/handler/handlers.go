package HD

import (
	"fmt"
	GT "groupie-tracker/packages/api"
	"html/template"
	"net/http"
	"strconv"
)

// Fonction qui génére la home page
func HomeHandler(write http.ResponseWriter, read *http.Request) {
	var templates = template.Must(template.ParseFiles("./templates/index.html"))
	err2 := templates.ExecuteTemplate(write, "index.html", GT.ArtistAPI)
	if err2 != nil {
		fmt.Println("Erreur 500 : ", err2)
		fmt.Fprintf(write, "Erreur 500, Internal erreur serveur : %v", err2)
		http.Error(write, err2.Error(), http.StatusInternalServerError)
	}
}

// Fonction qui génére la page d'erreur 404
func Error404Handler(write http.ResponseWriter, read *http.Request) {
	write.WriteHeader(404)
	var templates = template.Must(template.ParseFiles("./templates/404.html"))
	err2 := templates.ExecuteTemplate(write, "404.html", GT.ArtistAPI)
	if err2 != nil {
		fmt.Println("Erreur 500 : ", err2)
		fmt.Fprintf(write, "Erreur 500, Internal erreur serveur : %v", err2)
		http.Error(write, err2.Error(), http.StatusInternalServerError)
	}
}

// Fonction qui génére la page d'erreur 404
func Error500Handler(write http.ResponseWriter, read *http.Request) {
	write.WriteHeader(500)
	var templates = template.Must(template.ParseFiles("./templates/500.html"))
	err2 := templates.ExecuteTemplate(write, "500.html", GT.ArtistAPI)
	if err2 != nil {
		fmt.Println("Erreur 500 : ", err2)
		fmt.Fprintf(write, "Erreur 500, Internal erreur serveur : %v", err2)
		http.Error(write, err2.Error(), http.StatusInternalServerError)
	}
}

// Fonction pour la page groupe
func GroupeHandler(write http.ResponseWriter, read *http.Request) {
	// Lecture des informations des formulaires
	switch read.Method {
	case "POST":
		// Récupération des argument Word et Option si méthode 'Post' utilisé
		// Appelez ParseForm() pour analyser la requête brute et mettre à jour mes données
		if err := read.ParseForm(); err != nil {
			fmt.Fprintf(write, "ParseForm() err: %v", err)
		}

		recherche := read.FormValue("recherche")
		var id int
		if recherche == "" { //si cliqué sur un groupe
			id, _ = strconv.Atoi(read.FormValue("groupe"))
		} else { // si une recherche est faite
			nom_groupe := extraction_nom_groupe(recherche)
			// recherche de l'id
			id = GT.RechercheID(nom_groupe)
		}
		GT.RequestGroupe(id)
	default:
		// Erreur 405, methode non valide
		fmt.Fprintf(write, "Error 405 Method Not Allowed")
	}

	var templates = template.Must(template.ParseFiles("./templates/groupe.html"))
	err2 := templates.ExecuteTemplate(write, "groupe.html", GT.Globalgroupe)
	if err2 != nil {
		fmt.Println(err2)
		http.Error(write, err2.Error(), http.StatusInternalServerError)
	}
}

// Fonction qui récupére le nom du groupe quand une recherche est faite
func extraction_nom_groupe(recherche string) (nom_groupe string) {
	var group []string
	var mot string
	// Mise dans une array de chaque mot de la recherche
	for l, a := range recherche {
		if a == ' ' {
			group = append(group, mot)
			mot = ""
		} else if l == len(recherche)-1 {
			mot += string(a)
			group = append(group, mot)
		} else {
			mot += string(a)
		}
	}

	motok := false
	count := 0
	// Extraction du nom du groupe
	for i := 0; i < len(group); i++ {
		if motok == true {
			if count == 0 {
				nom_groupe = group[i]
			} else {
				nom_groupe = nom_groupe + " " + group[i]
			}
			count++
		}
		if group[i] == "Groupe" || group[i] == "membre" {
			motok = true
		}
	}

	return nom_groupe
}
