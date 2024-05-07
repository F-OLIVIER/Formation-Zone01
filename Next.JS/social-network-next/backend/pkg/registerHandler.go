package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var registerData User
	if err := json.NewDecoder(r.Body).Decode(&registerData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
		return
	}
	defer db.Close()

	var isError bool

	if registerData.Nickname != "" {

		var existingUsername string
		err = db.QueryRow("SELECT Nickname FROM USERS WHERE Nickname = ?", registerData.Nickname).Scan(&existingUsername)
		if err == nil {
			isError = true
			jsonResponse := map[string]interface{}{
				"success": false,
				"message": "Nickname déjà utilisé",
			}
			err := json.NewEncoder(w).Encode(jsonResponse)
			if err != nil {
				return
			}
		}
	}

	var existingEmail string
	err = db.QueryRow("SELECT Email FROM USERS WHERE Email = ?", registerData.Email).Scan(&existingEmail)
	if err == nil {
		isError = true
		jsonResponse := map[string]interface{}{
			"success": false,
			"message": "Email déjà utilisé",
		}
		err := json.NewEncoder(w).Encode(jsonResponse)
		if err != nil {
			return
		}
	}

	//rep si ok
	if !isError {
		var Privacy int
		Privacy = 1
		if registerData.PrivateProfile == "yes" {
			Privacy = 0
		}

		encryptPassword, _ := bcrypt.GenerateFromPassword([]byte(registerData.Password), bcrypt.DefaultCost)

		_, err = db.Exec("INSERT INTO USERS (Email, Password, FirstName, LastName, DateOfBirth, Avatar, Nickname, AboutMe, PrivateProfile) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", registerData.Email, encryptPassword, registerData.Firstname, registerData.Lastname, registerData.DateOfBirth, registerData.Avatar, registerData.Nickname, registerData.AboutMe, Privacy)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(registerData.Email)
		jsonResponse := map[string]interface{}{
			"success": true,
			"message": "",
			"email":   registerData.Email,
		}
		err = json.NewEncoder(w).Encode(jsonResponse)
		if err != nil {
			return
		}
	}
}
