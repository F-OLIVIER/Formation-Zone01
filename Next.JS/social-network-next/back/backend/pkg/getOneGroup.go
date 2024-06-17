package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetOneGroup(w http.ResponseWriter, r *http.Request) {
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

	err = db.QueryRow("SELECT NameGroup, AboutUs FROM LISTGROUPS WHERE IDGroup = ?", group.IdGroup).Scan(&group.Title, &group.AboutGroup)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows were returned!")
		} else {
			fmt.Println(err)
		}
		return
	}

	fmt.Println(group.Title, group.AboutGroup)
	fmt.Println("this id the id ", group.IdGroup)

	// Set the content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the status code to the response and the JSON data
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(group)
}
