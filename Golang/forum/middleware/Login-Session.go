package utils

import (
	"database/sql"
	"fmt"
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
func Login(w http.ResponseWriter, r *http.Request, database *sql.DB) bool {
	r.ParseForm()
	email := r.FormValue("mail")
	if CheckFormFull(r) {
		password := r.FormValue("pass")
		if CredentialsChecker(email, password, database) {
			userUUID := uuid.Must(uuid.NewV4())
			// userUUID, _ = uuid.NewV4()
			// var uuid string
			uuid := userUUID.String()
			cookie := &http.Cookie{
				Name:     "user_token",
				Value:    uuid,
				Expires:  time.Now().Add(3600 * time.Second),
				HttpOnly: true,
			}
			SessionLogger(w, r, email, Sessions, cookie, database)
			return true
		}
	}
	return false
}

// Fonction qui vérifie si les informations d'identification pour se connecter à la session sont correctes ou non (utilisez dans la fonction Login)
func CredentialsChecker(email, password string, database *sql.DB) bool {
	usedEmail := database.QueryRow("SELECT Email FROM Users WHERE Email = ?", email).Scan()
	if usedEmail == sql.ErrNoRows {
		return false
	} else {
		var pwdChecker string
		// get hashed password from existing account
		stmt, err := database.Prepare("SELECT password FROM Users WHERE email = ?")
		CheckErr(err, "credentialschecker")
		err = stmt.QueryRow(email).Scan(&pwdChecker) // pwdChecker now contains the hashed password
		CheckErr(err, "credentialschecker")
		// check for unhashed and hashed passwords match
		if bcrypt.CompareHashAndPassword([]byte(pwdChecker), []byte(password)) == nil {
			return true
		} else {
			return false
		}
	}
}

// Fonction de déconnexion (supression du cookie) et suppression de l'uuid dans la db
func Logout(w http.ResponseWriter, r *http.Request, database *sql.DB) {
	c, err := r.Cookie("user_token")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	CheckErr(err, "logout")
	stmt, err := database.Prepare("UPDATE Users SET uuid = NULL WHERE uuid = ?")
	CheckErr(err, "logout")
	stmt.Exec(c.Value)
	delete(Sessions, c.Value)
	c.MaxAge = -1
	http.SetCookie(w, c)
	fmt.Println("Logged out successfully")
	http.Redirect(w, r, "/?session=disconnected", http.StatusSeeOther)
}

// Fonction qui crée le cookie et sa correspondance dans la db ainsi que dans la map
func SessionLogger(w http.ResponseWriter, r *http.Request, email string, s map[string]Session, cookie *http.Cookie, database *sql.DB) {
	var val string
	stmt, err := database.Prepare("SELECT ID FROM Users WHERE email = ?")
	CheckErr(err, "sessionlogger")
	err1 := stmt.QueryRow(email).Scan(&val)
	fmt.Println("test", err1)
	// if err1 == sql.ErrNoRows {
	// fmt.Println("No rows")
	Sessions[val] = Session{
		email:  email,
		Cookie: cookie,
	}
	stmt, err = database.Prepare("UPDATE Users SET uuid = ? WHERE Email = ?")
	CheckErr(err, "sessionlogger")
	stmt.Exec(cookie.Value, email)
	http.SetCookie(w, cookie)
	fmt.Println(Sessions)
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

// ---------------------------------------------------
// --------------- Mot de passe oublié ---------------
// ---------------------------------------------------

// var CodeInitPassword = make(map[string]string)

var CodeResetPasswords = map[string]CodeResetPassword{}

type CodeResetPassword struct {
	email   string
	horaire time.Time // durée de validité du code
}

// Fonction qui permet de re-initialisé un mot de passe utilisateur '/ForgetPassword'
func ForgetPassword(r *http.Request, database *sql.DB) (string, bool, bool) {
	r.ParseForm()
	var Msg string
	forgetmail := r.FormValue("forgetmail")
	codeChangePassword := r.FormValue("codeChangePassword")
	newPass := r.FormValue("newPass")
	confirmNewPass := r.FormValue("confirmNewPass")
	if codeChangePassword != "" && newPass != "" && confirmNewPass != "" && newPass == confirmNewPass { // changement du mot de passe
		horaireReset := time.Since(CodeResetPasswords[codeChangePassword].horaire)
		if horaireReset > 300*time.Second {
			delete(CodeResetPasswords, codeChangePassword) // code périmé, supression de la map
			return "Délais de ré-initialisation depassé", false, false
		}
		mail := CodeResetPasswords[codeChangePassword].email
		if mail != "" { // Si code de reinitialisation valide, presence d'un e-mail
			_, username := SearchUser(mail, database)
			hashedPass, _ := bcrypt.GenerateFromPassword([]byte(newPass), 8)
			ReinitialisationPassword(mail, hashedPass, database)
			Msg = username + ", votre mot de passe a bien été ré-initialisé"
			delete(CodeResetPasswords, codeChangePassword) // code utilisé, supression de la map
			return Msg, false, true
		} else {
			return "Code de ré-initialisation incorrect", true, false
		}
	} else if forgetmail != "" { // send d'un mail avec le code de re-initialisation
		var username string
		var exist bool
		exist, username = SearchUser(forgetmail, database)
		if !exist { // si l'e-mail n'existe pas dans la db
			Msg = "Cet e-mail n'a pas de compte"
			return Msg, false, false
		} else {
			codeReinit := RandomFileName()
			// fmt.Println("codeReinit : ", codeReinit)
			CodeResetPasswords[codeReinit] = CodeResetPassword{
				email:   forgetmail,
				horaire: time.Now(),
			}
			// CodeInitPassword[codeReinit] = forgetmail
			// Contenu du message du mail
			Subject := "Re-initialisation de votre mot de passe"
			message := "Bonjour " + username +
				"\n\nVoici le code de re-initialisation de votre mot de passe : " + codeReinit +
				"\n\nSi ce n'est pas vous qui avez demandé une ré-initialisation de votre mot de passe, ignorer ce mail" +
				"\n\nL'équipe du Forum de la Couilliére Burgond"
			sended := SendMail(forgetmail, Subject, message)
			if sended {
				Msg = "Un e-mail avec votre code de re-initialisation vous a été envoyé"
			} else {
				Msg = "Erreur dans l'envoi de l'e-mail"
			}
		}
		return Msg, true, false
	}
	return Msg, false, false
}
