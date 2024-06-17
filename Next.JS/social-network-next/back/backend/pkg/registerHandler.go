package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var registerData User
	err := r.ParseMultipartForm(10 << 20) // 10MB max size
	if err != nil {
		http.Error(w, "400 bad request: "+err.Error(), http.StatusBadRequest)
		fmt.Println("here 2")
		return
	}

	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
		return
	}
	defer db.Close()

	var isError bool

	registerData.Nickname = r.FormValue("nickname")
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

	registerData.Email = r.FormValue("email")
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

		registerData.Password = r.FormValue("password")
		registerData.Firstname = r.FormValue("firstname")
		registerData.Lastname = r.FormValue("lastname")
		registerData.DateOfBirth = r.FormValue("dateofbirth")
		registerData.AboutMe = r.FormValue("aboutme")

		encryptPassword, _ := bcrypt.GenerateFromPassword([]byte(registerData.Password), bcrypt.DefaultCost)

		_, imageHeader, err := r.FormFile("avatar")
		if err != nil && err != http.ErrMissingFile {
			http.Error(w, "Error processing image", http.StatusInternalServerError)
			return
		}

		if imageHeader != nil {
			file, _, err := r.FormFile("avatar")
			if err != nil {
				fmt.Println("here")
				http.Error(w, "400 bad request: "+err.Error(), http.StatusBadRequest)
				return
			}
			defer file.Close()

			// Read the file data
			fileBytes, err := io.ReadAll(file)
			if err != nil {
				http.Error(w, "500 internal server error: "+err.Error(), http.StatusInternalServerError)
				return
			}

			registerData.Avatar = fileBytes
		}

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
