package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var loginData Login
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var trueEmail string
	var truePassword []byte
	var username string
	var id int
	var err error
	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
		return
	}
	defer db.Close()
	//fmt.Println("loginData:", loginData)
	err = db.QueryRow("SELECT ID, Nickname, Email, Password FROM USERS WHERE Nickname = ? OR Email = ?", loginData.Data, loginData.Data).Scan(&id, &username, &trueEmail, &truePassword)
	loginEmail := loginData.Data
	if err != nil {
		err = db.QueryRow("SELECT ID, Nickname, Email, Password FROM USERS WHERE Email = ?", loginEmail).Scan(&id, &username, &trueEmail, &truePassword)
		fmt.Println("Erreur lors de la récupération des données de l'utilisateur par nom d'utilisateur:", err) //Existe pas
		if err != nil {
			fmt.Println("Erreur lors de la récupération des données de l'utilisateur par email:", err) //Existe pas
		}
	}
	if err != nil {
		jsonResponse := map[string]interface{}{
			"success": false,
			"message": "Email ou username incorrect",
			"error":   "email",
		}
		err := json.NewEncoder(w).Encode(jsonResponse)
		if err != nil {
			return
		}
	} else {
		if err := bcrypt.CompareHashAndPassword(truePassword, []byte(loginData.Password)); err != nil {

			jsonResponse := map[string]interface{}{
				"success": false,
				"message": "Mot de passe incorrect",
				"error":   "password",
			}
			err := json.NewEncoder(w).Encode(jsonResponse)
			if err != nil {
				return
			}
		} else {
			// Delete si cookie existant.
			_, err = db.Exec(`DELETE FROM SESSIONS WHERE UserID = ?`, id)
			if err != nil {
				fmt.Println("error cookie supp")
				http.Error(w, "500 internal server error.", http.StatusInternalServerError)
				return
			}
			// Création d'une session et cookie avec protocole UUID.
			cookie, err := r.Cookie("session")
			if cookie != nil {
				// Remplace les valeurs du cookie existant
				sessionId, err := uuid.NewV4()
				if err != nil {
					fmt.Println("error cookie crea")
					http.Error(w, "500 internal server error.", http.StatusInternalServerError)
					return
				}
				cookie.Value = sessionId.String()
				cookie.MaxAge = 60 * 60 * 24
				// Mise à jour du cookie dans le navigateur
				http.SetCookie(w, cookie)
			} else {
				// Crée un nouveau cookie s'il n'existe pas
				sessionId, err := uuid.NewV4()
				if err != nil {
					fmt.Println("error cookie crea")
					http.Error(w, "500 internal server error.", http.StatusInternalServerError)
					return
				}
				cookie = &http.Cookie{
					Name:     "session",
					Value:    sessionId.String(),
					HttpOnly: true,
					Path:     "/",
					MaxAge:   60 * 60 * 24,
					SameSite: http.SameSiteNoneMode,
					Secure:   true,
				}
				// Ajoute le cookie à la réponse
				http.SetCookie(w, cookie)
			}
			_, err = db.Exec(`INSERT INTO SESSIONS (SessionToken, UserID) values(?, ?)`, cookie.Value, id)
			if err != nil {
				fmt.Println("error DB cookie")
				http.Error(w, "500 internal server error.", http.StatusInternalServerError)
				return
			}
			/*postDataLogin := WebsocketMessage{Type: "login", Data: UserData}
			broadcast <- postDataLogin*/

			var val string
			val = trueEmail
			Sessions[val] = SessionUser{
				email:  trueEmail,
				Cookie: cookie,
			}

			//fmt.Println("Create Sessions:", Sessions)
			jsonResponse := map[string]interface{}{
				"success": true,
				"message": "Connexion spécie",
				"id":      id,
			}
			err = json.NewEncoder(w).Encode(jsonResponse)
			if err != nil {
				return
			}
		}
	}
}
