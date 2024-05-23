package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAllGroups(w http.ResponseWriter, r *http.Request) {
	var groups []Group

	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
		return
	}
	defer db.Close()

	var MyId int
	err, MyId = WhoAmI(db, w, r)
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := db.Query("SELECT IDGroup,NameGroup, AboutUs, UserID_Creator FROM LISTGROUPS WHERE UserID_Creator = ?", MyId)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var group Group
		err := rows.Scan(&group.IdGroup, &group.NameGroup, &group.Description, &group.UserID_Creator)
		if err != nil {
			fmt.Println(err)
			return
		}

		groups = append(groups, group)
	}
	fmt.Println(groups)

	// Set the content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the status code to the response and the JSON data
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(groups)
	fmt.Println("GetAllGroups called")
}

func WhoAmI(db *sql.DB, w http.ResponseWriter, r *http.Request) (error, int) {
	var missing int

	cookie, err := r.Cookie("session")
	if err != nil {
		// Gérer l'erreur si le cookie n'est pas trouvé
		http.Error(w, "Session cookie not found", http.StatusUnauthorized)
		return err, missing
	}

	fmt.Println("Session cookie value:", cookie.Value)

	var userID int
	err = db.QueryRow("SELECT UserID FROM SESSIONS WHERE SessionToken = ?", cookie.Value).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found for this sessionToken")
		} else {
			fmt.Println("Error when searching for the user:", err)
		}
		return err, missing
	}

	return nil, userID

}
