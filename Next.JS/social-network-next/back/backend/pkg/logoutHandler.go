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

	// for crossplatform
	// lecture du cookie pour supprimer l'utilisateur de la liste
	oldcookie, err := r.Cookie("session")
	if err != nil {
		return
	}
	curr, err := CurrentUser(oldcookie.Value)
	if err != nil {
		fmt.Println("no cookie.Value for ws")
		return
	}
	ListOnline = RemoveSliceInt(ListOnline, curr.Id)
	// fmt.Println("ListOnline logout : ", ListOnline)

	// delete de la session dans la db
	_, err = db.Exec(`DELETE FROM SESSIONS WHERE UserID = ?`, curr.Id)
	// delete du cookie
	DeleteCookie(w)

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

// Supprimer le cookie de session
func DeleteCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // Définit une date d'expiration passée pour supprimer le cookie
	})
}
