package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type Target struct {
	UserID int `json:"user_id"`
}

func TargetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data Target
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		http.Error(w, "Erreur lors de l'ouverture de la base de données", http.StatusInternalServerError)
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
		return
	}
	defer db.Close()

	users := WhoDisplayNotif(data.UserID, 0, "post", db)

	jsonResponse := map[string]interface{}{
		"success":    true,
		"listTarget": users,
	}

	if err := json.NewEncoder(w).Encode(jsonResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
