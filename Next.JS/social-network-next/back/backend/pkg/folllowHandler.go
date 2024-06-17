package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func FollowHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/follow" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	// FIXME: check du cookies ICI
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("Error cookie FollowHandler")
	}
	user, err := CurrentUser(cookie.Value)
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}

	// décodage du json
	var follow Followers
	errjson := json.NewDecoder(r.Body).Decode(&follow)
	if errjson != nil {
		http.Error(w, "500 internal server error: Failed to connect to database. "+errjson.Error(), http.StatusInternalServerError)
	}

	// Ouverture db
	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
		return
	}
	defer db.Close()

	// fmt.Println("follow : ", follow)
	switch follow.Action {
	case "follow":
		if testNewFollow(follow.UserId_Following, follow.UserId_Follower, db) { // test si le follow existe
			newFollow(follow.UserId_Following, follow.UserId_Follower, db)
		} else {
			deleteFollow(follow.UserId_Following, follow.UserId_Follower, db)
		}
	case "validatefollow":
		if follow.ValidateFollow {
			validateFollow(follow.UserId_Following, follow.UserId_Follower, db)
		} else {
			deleteFollow(follow.UserId_Following, follow.UserId_Follower, db)
		}

	default:
		fmt.Println("\nPourquoi fetch ? Il n'y a rien a changer")
	}

	// Retrieve the updated user information
	err = db.QueryRow("SELECT ID, Email, FirstName, LastName, DateOfBirth, Avatar, Nickname, AboutMe, PrivateProfile FROM USERS WHERE ID = ?", user.Id).Scan(&user.Id, &user.Email, &user.Firstname, &user.Lastname, &user.DateOfBirth, &user.Avatar, &user.Nickname, &user.AboutMe, &user.PrivateProfile)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des données utilisateur:", err)
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func testNewFollow(UserID_Following, UserID_Followers int, db *sql.DB) bool {
	stmt, errdb := db.Prepare("SELECT ID FROM FOLLOWERS WHERE UserID_Following = ? AND UserID_Follower = ?")
	CheckErr(errdb, "testNewFollow Requete DB")
	id := 0
	stmt.QueryRow(UserID_Following, UserID_Followers).Scan(&id)
	if id == 0 {
		return true
	}
	return false
}

func newFollow(UserID_Following, UserID_Followers int, db *sql.DB) {
	var currentTime = time.Now()
	englishDateTime := currentTime.Format("01/02/2006 15:04:05")
	// Vérification de la privacité du profil follow
	respPrivateProfile, errdb := db.Prepare("SELECT PrivateProfile FROM USERS WHERE ID = ?")
	CheckErr(errdb, "3-newFollow Requete DB")
	var privateProfile bool
	respPrivateProfile.QueryRow(UserID_Following).Scan(&privateProfile)

	// insertion dans la table FOLLOWERS
	stmt, err := db.Prepare("INSERT INTO FOLLOWERS(UserID_Following,UserID_Follower,DateFollow,ValidateFollow) Values(?,?,?,?)")
	CheckErr(err, "1-newFollow Prepare db")
	result, err := stmt.Exec(UserID_Following, UserID_Followers, englishDateTime, !privateProfile) // false insert '0',true insert '1'
	CheckErr(err, "1- newFollow db Exec")

	lastInsertID, err := result.LastInsertId()
	dt := time.Now().Format("01-02-2006 15:04:05")
	CheckErr(err, "2-newFollow Getting last insert ID")

	InsertNotif(int(lastInsertID), UserID_Following, dt, "follow", db)

}

func validateFollow(UserID_Following, UserID_Followers int, db *sql.DB) {
	stmt, err := db.Prepare("UPDATE FOLLOWERS SET ValidateFollow = 0 WHERE UserID_Following = ? AND UserID_Follower = ?")
	CheckErr(err, "validateFollow Prepare db")
	_, err = stmt.Exec(UserID_Following, UserID_Followers)
	CheckErr(err, "validateFollow db Exec")
}

func deleteFollow(UserID_Following, UserID_Followers int, db *sql.DB) {
	respPrivateProfile, errdb := db.Prepare("SELECT ID FROM FOLLOWERS WHERE UserID_Following = ? AND UserID_Follower = ?")
	CheckErr(errdb, "1-deleteFollow Requete DB")
	idFollow := 0
	respPrivateProfile.QueryRow(UserID_Following, UserID_Followers).Scan(&idFollow)

	if idFollow != 0 {
		DeleteNotif(idFollow, "follow", db)
	}

	stmt, err := db.Prepare("DELETE FROM FOLLOWERS WHERE UserID_Following = ? AND UserID_Follower = ?")
	CheckErr(err, "deleteFollow Prepare db")
	_, err = stmt.Exec(UserID_Following, UserID_Followers)
	CheckErr(err, "deleteFollow db Exec")
}

// Retourne la liste des followers ou des followings
// option = 'followers' ou 'followings'
// validate : 0 = list validé / 1 = utilisateur en attente de validation
func ListFollow(UserID int, option string, validate int) (listuser []User) {
	var stmt *sql.Stmt
	var err error
	// Ouverture db
	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
		return
	}
	defer db.Close()

	switch option {
	case "followings":
		stmt, err = db.Prepare(`SELECT USERS.ID, USERS.FirstName, USERS.LastName, USERS.Avatar, USERS.Nickname
										FROM USERS
										INNER JOIN FOLLOWERS ON FOLLOWERS.UserID_Following = USERS.ID
										WHERE FOLLOWERS.ValidateFollow = ? AND FOLLOWERS.UserID_Follower = ?;`)
	case "followers":
		stmt, err = db.Prepare(`SELECT USERS.ID, USERS.FirstName, USERS.LastName, USERS.Avatar, USERS.Nickname
										FROM USERS
										INNER JOIN FOLLOWERS ON FOLLOWERS.UserID_Follower = USERS.ID
										WHERE FOLLOWERS.ValidateFollow = ? AND FOLLOWERS.UserID_Following = ?;`)
	default:
		fmt.Println("\nY A DE LA MERDE ICI : vérifie t'on option à la fonction ListFollow\n Option correct : 'followers' ou 'followings'")
		return
	}

	CheckErr(err, "ListFollowers, db prepare")
	rows, err := stmt.Query(validate, UserID)
	CheckErr(err, "ListFollowers, db query")
	for rows.Next() {
		var currentUser User
		err = rows.Scan(&currentUser.Id, &currentUser.Firstname, &currentUser.Lastname, &currentUser.Avatar, &currentUser.Nickname)
		CheckErr(err, "ListFollowers, db rows.Next scan")
		listuser = append(listuser, currentUser)
	}
	return listuser
}

func CheckErr(err error, str string) {
	if err != nil {
		fmt.Printf("______________________________________\nERROR : %v\n%v\n", str, err)
	}
}
