package utils

import (
	"database/sql"
	"encoding/json"
	data "forum/internal"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

// Fonction qui gère les formulaires de connexion pour enregistrer les données de l'utilisateur, le booléen est uniquement destiné au gestionnaire
func RegisterUser(w http.ResponseWriter, r *http.Request, database *sql.DB) bool {
	// Récupération des informations du post
	var register data.Register
	err := json.NewDecoder(r.Body).Decode(&register)
	CheckErr(err, "Erreur de décodage JSON viewpost")

	Age := register.Ageregister
	Gender := register.Genderregister
	FirstName := register.FirstNameregister
	LastName := register.LastNameregister
	username := register.Username
	mail := register.Mail
	pass := register.Pass
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(pass), 8)
	confirmPass := register.ConfirmPass

	// il vérifie si le nom d'utilisateur et l'e-mail sont déjà utilisés
	usedEmail := database.QueryRow("SELECT Email FROM Users WHERE Email = ?", mail).Scan()
	usedUsername := database.QueryRow("SELECT Username FROM Users WHERE Username = ?", username).Scan()
	if Age == "" || Gender == "" || FirstName == "" || LastName == "" || username == "" || mail == "" || pass == "" || confirmPass == "" {
		ErrorMessage(w, r, "form")
	} else if strings.Contains("@", username) {
		ErrorMessage(w, r, "badusername")
	} else if pass == confirmPass && usedEmail == sql.ErrNoRows && usedUsername == sql.ErrNoRows && mail != "" && pass != "" && username != "" && confirmPass != "" {
		stmt, err := database.Prepare("INSERT INTO Users(Email,Password,Username,Photo,Age,Gender,FirstName,LastName) Values(?,?,?,?,?,?,?,?)")
		CheckErr(err, "regiisteruser")
		_, err = stmt.Exec(mail, hashedPass, username, "default/absencePhoto.jpg", Age, Gender, FirstName, LastName)
		CheckErr(err, "registeruser")
		return true
	} else if usedEmail != sql.ErrNoRows {
		ErrorMessage(w, r, "email")
	} else if pass != confirmPass {
		ErrorMessage(w, r, "nomatch")
	} else if usedUsername != sql.ErrNoRows {
		ErrorMessage(w, r, "username")
	}
	return false
}
