package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func AcceptGroupNotification(w http.ResponseWriter, r *http.Request) {
	var group Group

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		return
	}

	err = json.Unmarshal(body, &group)
	if err != nil {
		fmt.Println("Error unmarshalling request body:", err)
		http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
		return
	}

	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec(`INSERT INTO MEMBERSGROUPS (IDGroup, UserID) VALUES (?, ?)`, group.ID, group.IdWhoIsInvited)
	if err != nil {
		fmt.Println("Error while inserting into the database:", err)
		return
	}

	jsonResponse := map[string]interface{}{
		"success": true,
		"message": "You have joined the group",
	}
	if err := json.NewEncoder(w).Encode(jsonResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
