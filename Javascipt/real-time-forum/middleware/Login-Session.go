package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	data "forum/internal"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var Sessions = map[string]Session{}

type Session struct {
	email  string
	Cookie *http.Cookie
}

// Fonction qui gère le formulaire de connexion pour démarrer une session utilisateur, le booléen est uniquement destiné au gestionnaire
func Login(w http.ResponseWriter, r *http.Request, database *sql.DB) (bool, string) {
	// Récupération des informations du post
	var login data.Register
	err := json.NewDecoder(r.Body).Decode(&login)
	CheckErr(err, "Erreur de décodage JSON viewpost")
	// fmt.Println("login", login)

	email := login.Mail
	password := login.Pass
	if email != "" && password != "" {
		valid, maildb, username := CredentialsCheckerEmailOrUsername(email, password, database)
		if valid {
			userUUID := uuid.Must(uuid.NewV4())
			uuid := userUUID.String()
			cookie := &http.Cookie{
				Name:     "user_token",
				Path:     "/",
				Value:    uuid,
				Expires:  time.Now().Add(3600 * time.Second),
				HttpOnly: false, // Si actif, empéche la lecture du cookie en js (vulnérabilités potentielles)
				// Secure:   true, // pour l'HTTPS
			}
			SessionLogger(w, r, maildb, Sessions, cookie, database)
			return true, username
		}
	}
	return false, ""
}

// Fonction qui vérifie si les informations d'identification pour se connecter à la session sont correctes ou non (utilisez dans la fonction Login)
func CredentialsCheckerEmailOrUsername(login, password string, database *sql.DB) (bool, string, string) {
	usedEmail := database.QueryRow("SELECT Email, Username FROM Users WHERE Email = ? OR Username = ?", login, login).Scan()
	if usedEmail == sql.ErrNoRows {
		return false, "", ""
	} else {
		var pwdChecker, maildb, username string
		// get hashed password from existing account
		stmt, err := database.Prepare("SELECT password, Email, Username FROM Users WHERE Email = ? OR Username = ?")
		CheckErr(err, "credentialschecker")
		err = stmt.QueryRow(login, login).Scan(&pwdChecker, &maildb, &username) // pwdChecker now contains the hashed password
		CheckErr(err, "credentialschecker")
		// check for unhashed and hashed passwords match
		if bcrypt.CompareHashAndPassword([]byte(pwdChecker), []byte(password)) == nil {
			return true, maildb, username
		} else {
			return false, "", ""
		}
	}
}

// Fonction de déconnexion (supression du cookie) et suppression de l'uuid dans la db
func Logout(w http.ResponseWriter, r *http.Request, database *sql.DB) {
	c, err := r.Cookie("user_token")
	if err == http.ErrNoCookie {
		return
	}
	CheckErr(err, "logout")
	stmt, err := database.Prepare("UPDATE Users SET uuid = NULL WHERE uuid = ?")
	CheckErr(err, "logout")
	stmt.Exec(c.Value)
	delete(Sessions, c.Value)
	c.MaxAge = -1
	http.SetCookie(w, c)
	fmt.Println("Logged out successfully")
}

// Fonction qui crée le cookie et sa correspondance dans la db ainsi que dans la map
func SessionLogger(w http.ResponseWriter, r *http.Request, email string, s map[string]Session, cookie *http.Cookie, database *sql.DB) {
	var val string
	stmt, err := database.Prepare("SELECT ID FROM Users WHERE email = ?")
	CheckErr(err, "db Prepare sessionlogger")
	err1 := stmt.QueryRow(email).Scan(&val)
	CheckErr(err1, "QueryRow sessionlogger")

	Sessions[val] = Session{
		email:  email,
		Cookie: cookie,
	}
	stmt, err = database.Prepare("UPDATE Users SET uuid = ? WHERE Email = ?")
	CheckErr(err, "sessionlogger")
	stmt.Exec(cookie.Value, email)
	http.SetCookie(w, cookie)
	// fmt.Println(Sessions)
}

// Fonction qui compare le cookie utilisateur avec la map de gestion des cookies
func CheckToken(s map[string]Session, c *http.Cookie, email string, database *sql.DB) bool {
	var ID int
	stmt, err := database.Prepare("SELECT ID FROM Users WHERE email = ?")
	CheckErr(err, "sessionlogger")
	stmt.QueryRow(email).Scan(&ID)
	for _, v := range s {
		if v.Cookie.Value == c.Value {
			return true
		}
	}
	return false
}
