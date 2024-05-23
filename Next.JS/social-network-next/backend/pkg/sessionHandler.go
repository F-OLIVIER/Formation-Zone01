package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// Appelé lors d'une ouverture de navigateur pour check si cookie valide et auto-log l'user.
func SessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/session" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
		return
	}
	defer db.Close()
	// Retrieve user ID from session cookie
	cookie, err := r.Cookie("session")
	if err != nil {
		jsonResponse := map[string]interface{}{
			"success": false,
			"message": "Error",
		}
		err = json.NewEncoder(w).Encode(jsonResponse)
		if err != nil {
			return
		}
		return
	}
	//fmt.Println("cookie:", cookie)
	isGood := false
	//fmt.Println("Sessions:", Sessions)
	for _, v := range Sessions {
		if v.Cookie.Value == cookie.Value {
			isGood = true
		}
	}
	//fmt.Println("isGood:", isGood)
	if !isGood {
		jsonResponse := map[string]interface{}{
			"success": false,
			"message": "Error you change your cookie",
		}
		err = json.NewEncoder(w).Encode(jsonResponse)
		if err != nil {
			return
		}
		return
	}
	foundVal := cookie.Value
	// Requête SQL pour rechercher l'UserID correspondant au sessionToken
	var userID int
	err = db.QueryRow("SELECT UserID FROM SESSIONS WHERE SessionToken = ?", foundVal).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Aucun utilisateur trouvé pour ce sessionToken")
		} else {
			fmt.Println("Erreur lors de la recherche de l'utilisateur:", err)
		}
		return
	}
	var email, username string
	err = db.QueryRow("SELECT u.Email, u.Nickname FROM USERS u INNER JOIN SESSIONS s ON u.ID = s.UserID WHERE u.ID = ?", userID).Scan(&email, &username)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Aucun utilisateur trouvé pour cet ID")
		} else {
			fmt.Println("Erreur lors de la récupération de l'utilisateur:", err)
		}
		return
	}
	jsonResponse := map[string]interface{}{
		"success":  true,
		"message":  "Re-Log successful",
		"email":    email,
		"username": username,
		"id":       userID,
	}
	err = json.NewEncoder(w).Encode(jsonResponse)
	if err != nil {
		return
	}
}
