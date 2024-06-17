package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
		return
	}
	defer db.Close()

	// Supprimer le cookie de session
	cookie := &http.Cookie{
		Name:     "session",
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		MaxAge:   -1, // Définit une date d'expiration passée pour supprimer le cookie
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}
	http.SetCookie(w, cookie)
	/*postDataLogin := WebsocketMessage{Type: "login", Data: UserData}
	broadcast <- postDataLogin*/
	jsonResponse := map[string]interface{}{
		"success": true,
		"message": "Logout successful",
	}
	err = json.NewEncoder(w).Encode(jsonResponse)
	if err != nil {
		return
	}
}
