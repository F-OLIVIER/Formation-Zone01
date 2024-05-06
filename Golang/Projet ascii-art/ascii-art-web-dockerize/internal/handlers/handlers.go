package handlers

import (
	ascii "ascii_web/cmd"
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

// fonction du handler "home"
func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" { // Definition de la méthode GET
		templ, err := template.ParseFiles("template/home.page.html") // Analyse du fichier template
		if err != nil {                                              // Gestion de l'erreur 500
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		templ.Execute(w, nil)
	} else if r.Method == "POST" { // Définition de la méthode POST
		txt := r.FormValue("txt")                    // Récupération de la valeur txt entrée
		police := r.FormValue("Polices")             // Récupération de la valeur Polices sélectionnée
		txt = strings.ReplaceAll(txt, "\r\n", "\\n") // Gestion des \n \r
		char := false
		for _, k := range txt { // Gestion de l'erreur 500 ( mauvaise input à transformer )
			if k < ' ' || k > '~' {
				char = true
				// w.WriteHeader(http.StatusInternalServerError)
				// // fmt.Fprintf(w, "500-Internal-Server-Error")
				// http.Redirect(w, r, "/500-Internal-Error", http.StatusInternalServerError)
				// break
			}
		}

		var page string
		if txt == "" {
			page = "template/home.page.html"
		} else {
			page = "template/result.page.html"
		}
		if char == true {
			page = "template/500.page.html"
			fmt.Println("500: Internal Server Error")
		}

		templ, err := template.ParseFiles(page)
		if err != nil { // Gestion de l'erreur 500
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !char { // Lancement de la commande importée (ascii-art-fs)
			tmp := ascii.Ascii_Return(txt, police)
			if tmp != "" && tmp != "\n" {
				templ.Execute(w, ascii.Ascii_Return(txt, police))
			} else {
				templ.Execute(w, nil)
			}
		} else {
			templ.Execute(w, nil)
		}
	} else {
		// Gestion de l'erreur 400
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("400: Bad Request")
		fmt.Fprintf(w, "400-Bad-Request")
	}
}
