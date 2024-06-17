package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func AskForJoinGroup(w http.ResponseWriter, r *http.Request) {
	var grp Group

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		return
	}

	err = json.Unmarshal(body, &grp)
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

	fmt.Println(grp.ID)
	fmt.Println(grp.IdGroup)

	// Insert the notification in the askgroup

	var UserIdCreator string
	now := time.Now()
	date := now.Format("2006-01-02 15:04:05")

	db.QueryRow("SELECT UserID_Creator FROM LISTGROUPS WHERE IDGroup = ?", grp.IdGroup).Scan(&UserIdCreator)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("USER ID CREATOR", UserIdCreator)

	//convert the string to int
	UserIdCreatorInt, err := strconv.Atoi(UserIdCreator)
	if err != nil {
		fmt.Println(err)
		return
	}

	alreadyAsked, err := AlreadyAsked(db, UserIdCreatorInt, grp.IdGroup, grp.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	alreadyInvited, err := AlreadyAreInvited(db, UserIdCreatorInt, grp.IdGroup)
	if err != nil {
		fmt.Println(err)
		return
	}

	if alreadyInvited {
		jsonResponse := map[string]interface{}{
			"success": false,
			"message": "Already invited",
		}
		err := json.NewEncoder(w).Encode(jsonResponse)
		if err != nil {
			fmt.Println("Error encoding JSON response:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		return
	}

	if alreadyAsked {
		jsonResponse := map[string]interface{}{
			"success": false,
			"message": "Already asked",
		}
		err := json.NewEncoder(w).Encode(jsonResponse)
		if err != nil {
			fmt.Println("Error encoding JSON response:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		return
	}

	_, err = db.Exec("INSERT INTO NOTIFICATIONS ( IDAsking, UserID_Sender,UserID_Receiver ,Date) VALUES (?, ?, ?, ?)", grp.IdGroup, grp.ID, UserIdCreator, date)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonResponse := map[string]interface{}{
		"success":  true,
		"message":  "Request sent successfully",
		"receiver": UserIdCreatorInt,
	}
	err = json.NewEncoder(w).Encode(jsonResponse)
	if err != nil {
		fmt.Println("Error encoding JSON response:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func AlreadyAsked(db *sql.DB, receiver int, groupId int, sender int) (bool, error) {
	query := `SELECT COUNT(*) FROM NOTIFICATIONS WHERE UserID_Receiver = ? AND IDAsking = ? AND UserID_Sender = ?`
	var count int
	err := db.QueryRow(query, receiver, groupId, sender).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func AlreadyAreInvited(db *sql.DB, receiver int, groupId int) (bool, error) {
	query := `SELECT COUNT(*) FROM NOTIFICATIONS WHERE IDGroup = ? AND UserID_Receiver = ?`
	var count int
	err := db.QueryRow(query, groupId, receiver).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
