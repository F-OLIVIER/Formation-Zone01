package HD

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	GT "groupie-tracker-filters/packages/api"
	BR "groupie-tracker-filters/packages/barreRecherche"
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
		}
		BR.RequestGroupe(id)
	default:
		// Erreur 405, methode non valide
		write.WriteHeader(405)
		fmt.Fprintf(write, "Error 405 Method Not Allowed")
	}

	var templates = template.Must(template.ParseFiles("./templates/groupe.html"))
	err2 := templates.ExecuteTemplate(write, "groupe.html", BR.Globalgroupe)
	if err2 != nil {
		fmt.Println(err2)
		http.Error(write, err2.Error(), http.StatusInternalServerError)
	}
}

// Fonction pour la page groupe
func SearchHandler(write http.ResponseWriter, read *http.Request) {
	var id int
	// Lecture des informations des formulaires
	switch read.Method {
	case "POST":
		// Récupération des argument Word et Option si méthode 'Post' utilisé
		// Appelez ParseForm() pour analyser la requête brute et mettre à jour mes données
		if err := read.ParseForm(); err != nil {
			fmt.Fprintf(write, "ParseForm() err: %v", err)
		}

		// Récupération des informations des inputs range
		rech_creation := read.FormValue("rech_creation")
		rech_album := read.FormValue("rech_album")
		rech_date_min, errAtoi1 := strconv.Atoi(read.FormValue("rech_date_min")) // int en entrée
		rech_date_max, errAtoi2 := strconv.Atoi(read.FormValue("rech_date_max")) // int en entrée
		if errAtoi1 != nil || errAtoi2 != nil {
			fmt.Println("Erreur conversion Atoi : ", errAtoi1, errAtoi2)
		}
		BR.RequestAnnees(rech_date_min, rech_date_max, rech_creation, rech_album)

		// Récupérations des informations des checkbox
		rech_nbmember1 := read.FormValue("rech_nbmember1") // si "on"
		rech_nbmember2 := read.FormValue("rech_nbmember2") // si "on"
		rech_nbmember3 := read.FormValue("rech_nbmember3") // si "on"
		rech_nbmember4 := read.FormValue("rech_nbmember4") // si "on"
		rech_nbmember5 := read.FormValue("rech_nbmember5") // si "on"
		rech_nbmember6 := read.FormValue("rech_nbmember6") // si "on"
		rech_nbmember7 := read.FormValue("rech_nbmember7") // si "on"
		rech_nbmember8 := read.FormValue("rech_nbmember8") // si "on"
		BR.Requestnbmembers(rech_nbmember1, rech_nbmember2, rech_nbmember3, rech_nbmember4, rech_nbmember5, rech_nbmember6, rech_nbmember7, rech_nbmember8)

		// Récupération de la case de recherche texte
		recherche := read.FormValue("recherche")
		if recherche != "" {
			BR.BarreRecherche(recherche)
		}

		// Récupération de l'id du groupe si l'utilisateur a cliqué sur une groupe
		idGroupe := read.FormValue("groupe")

		if recherche == "" && idGroupe != "" { //si cliqué sur un groupe
			id, _ = strconv.Atoi(idGroupe)
		}
		BR.RequestGroupe(id)
	default:
		// Erreur 405, methode non valide
		write.WriteHeader(405)
		fmt.Fprintf(write, "Error 405 Method Not Allowed")
	}

	var err2 error
	if len(BR.SearchArtists) == 1 { // si 1 seul résultats de recherche on retourne directement la page du groupe trouvé (page groupe)
		id = BR.SearchArtists[0].Art_id
		BR.RequestGroupe(id)
		var templates = template.Must(template.ParseFiles("./templates/groupe.html"))
		err2 = templates.ExecuteTemplate(write, "groupe.html", BR.Globalgroupe)
	} else { // si plusieurs résultats de recherche on retourne une page de résultats de recherche (page search)
		var templates = template.Must(template.ParseFiles("./templates/search.html"))
		err2 = templates.ExecuteTemplate(write, "search.html", BR.SearchArtists)
	}
	if err2 != nil {
		fmt.Println(err2)
		http.Error(write, err2.Error(), http.StatusInternalServerError)
	}
}
