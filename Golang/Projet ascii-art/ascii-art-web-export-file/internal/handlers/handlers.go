package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"

	ascii "ascii_web/cmd"
)

var txt_temp string

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
		txt := r.FormValue("txt")        // Récupération de la valeur txt entrée
		police := r.FormValue("Polices") // Récupération de la valeur Polices sélectionnée
		policeVide := false
		if r.FormValue("telechargement") == "TelechargementValide" {
			w.Header().Set("Content-type", "text/html; charset=utf-8")
			w.Header().Set("Content-Disposition", "attachment; filename=result.txt")

			// récupération taille du fichier
			fi, err := os.Stat("template/result.txt")
			if err != nil { // Gestion de l'erreur 500
				templ, _ := template.ParseFiles("template/500.page.html")
				templ.Execute(w, nil)
				fmt.Println("500: Internal Server Error")
				return
			}
			size := fi.Size()
			w.Header().Set("Content-Length:", strconv.Itoa(int(size)))
			templ, _ := template.ParseFiles("template/result.download.html")
			templ.Execute(w, txt_temp)
			return
		}
		txt = strings.ReplaceAll(txt, "\r\n", "\\n") // Gestion des \n \r
		// telechargement := r.FormValue("telechargement")
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
		if char {
			page = "template/500.page.html"
			fmt.Println("500: Internal Server Error")
		}

		templ, err := template.ParseFiles(page)
		if err != nil { // Gestion de l'erreur 500
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !char && policeVide == false { // Lancement de la commande importée (ascii-art-fs)
			tmp := ascii.Ascii_Return(txt, police)
			if tmp != "" && tmp != "\n" {
				txt_temp = tmp
				down, _ := os.Create("template/result.txt")
				down.WriteString(tmp)
				err := os.Chmod("template/result.txt", 0777)
				// 0400 -r--------
				// 0600 -rw-------
				// 0644 -rw-r--r--
				// 0755 -rwxr-xr-x
				// 0777 -rwxrwxrwx

				if err != nil && !os.IsExist(err) {
					log.Fatal(err)
				}

				templ.Execute(w, tmp)
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
