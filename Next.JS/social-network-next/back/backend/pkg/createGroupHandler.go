package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	var registerData Group

	if err := json.NewDecoder(r.Body).Decode(&registerData); err != nil {
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		bodyString := string(bodyBytes)
		fmt.Printf("Request body: %s\n", bodyString) // Print the request body
		http.Error(w, fmt.Sprintf("Error decoding request body: %v. Body: %s", err, bodyString), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
		return
	}
	defer db.Close()

	err = VerifDataBase(registerData, db)
	if err != nil {
		fmt.Println("Error while verifying the database:", err)
		jsonResponse := map[string]interface{}{
			"success": false,
			"message": err.Error(),
		}
		_ = json.NewEncoder(w).Encode(jsonResponse)
		return
	}

	err, registerData.ID = GetName(db, w, r)

	/*_, err = db.Exec(`INSERT INTO MEMBERSGROUPS (IDGroup, UserID) VALUES (?, ?)`, data.ID, data.IdWhoIsInvited)
	if err != nil {
		fmt.Println("Error while inserting into the database:", err)
		return
	}*/

	err = InsertIntoDataBase(registerData, db, registerData.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonResponse := map[string]interface{}{
			"success": false,
			"message": err.Error(),
		}
		_ = json.NewEncoder(w).Encode(jsonResponse)
		return
	}

	err = AddingOwnerAsAMember(registerData, db, registerData.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonResponse := map[string]interface{}{
			"success": false,
			"message": err.Error(),
		}
		_ = json.NewEncoder(w).Encode(jsonResponse)
		return
	}

	fmt.Println("Group created successfully")

	w.WriteHeader(http.StatusOK)
	jsonResponse := map[string]interface{}{
		"success": true,
		"message": "Group created successfully",
	}
	_ = json.NewEncoder(w).Encode(jsonResponse)

}

func VerifDataBase(registerData Group, db *sql.DB) error {
	if registerData.Title != "" {
		err := db.QueryRow("SELECT NameGroup FROM LISTGROUPS WHERE NameGroup = ?", registerData.Title).Scan(&registerData.ID)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		return fmt.Errorf("Group name already exists")

	}
	return nil
}

func GetName(db *sql.DB, w http.ResponseWriter, r *http.Request) (error, int) {
	var missing int
	fmt.Println("Checking for session cookie...")
	cookie, err := r.Cookie("session")
	if err != nil {
		// Gérer l'erreur si le cookie n'est pas trouvé
		http.Error(w, "Session cookie not found", http.StatusUnauthorized)
		return err, missing
	}
	fmt.Println("Session cookie value:", cookie.Value)

	// SQL query to find the UserID corresponding to the sessionToken
	var userID int
	fmt.Println("cookie.Value", cookie.Value)
	fmt.Println("ownerOfTheGroup", cookie.Value)

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

func InsertIntoDataBase(registerData Group, db *sql.DB, IDownerOfTheGroup int) error {
	_, err := db.Exec(`INSERT INTO LISTGROUPS (NameGroup, AboutUs, UserID_Creator) VALUES (?, ?, ?)`, registerData.Title, registerData.AboutGroup, IDownerOfTheGroup)
	if err != nil {
		fmt.Println("Error while inserting into the database:", err)
		return err
	}
	return nil
}

func AddingOwnerAsAMember(registerData Group, db *sql.DB, IDownerOfTheGroup int) error {

	err := db.QueryRow("SELECT IDGroup FROM LISTGROUPS WHERE NameGroup = ?", registerData.Title).Scan(&registerData.IdGroup)
	if err != nil {
		fmt.Println("Error while fetching group ID from the database:", err)
		return err
	}

	_, err = db.Exec(`INSERT INTO MEMBERSGROUPS (IDGroup, UserID) VALUES (?, ?)`, registerData.IdGroup, IDownerOfTheGroup)
	if err != nil {
		fmt.Println("Error while inserting into the database:", err)
		return err
	}
	return nil
}
