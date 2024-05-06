package utils

import (
	"database/sql"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

// Fonction qui gère les formulaires de connexion pour enregistrer les données de l'utilisateur, le booléen est uniquement destiné au gestionnaire
func RegisterUser(w http.ResponseWriter, r *http.Request, database *sql.DB) bool {
	r.ParseForm()
	mail := r.FormValue("mail")
	pass := r.FormValue("pass")
	username := r.FormValue("username")
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(pass), 8)
	confirmPass := r.FormValue("confirmPass")

	// il vérifie si le nom d'utilisateur et l'e-mail sont déjà utilisés
	usedEmail := database.QueryRow("SELECT Email FROM Users WHERE Email = ?", mail).Scan()
	usedUsername := database.QueryRow("SELECT Username FROM Users WHERE Username = ?", username).Scan()
	if pass == confirmPass && usedEmail == sql.ErrNoRows && usedUsername == sql.ErrNoRows && CheckFormFull(r) {
		stmt, err := database.Prepare("INSERT INTO Users(Email,Password,Username,Photo) Values(?,?,?,?)")
		CheckErr(err, "regiisteruser")
		_, err = stmt.Exec(mail, hashedPass, username, "default/absencePhoto.jpg")
		CheckErr(err, "registeruser")
		return true
	} else if usedEmail != sql.ErrNoRows {
		ErrorMessage(w, r, "email")
	} else if pass != confirmPass {
		ErrorMessage(w, r, "nomatch")
	} else if usedUsername != sql.ErrNoRows {
		ErrorMessage(w, r, "username")
	} else if !CheckFormFull(r) {
		ErrorMessage(w, r, "form")
	}
	return false
}
