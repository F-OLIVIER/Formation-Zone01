package utils

import (
	"database/sql"
	"fmt"
	"html"
	"strconv"
	"strings"

	data "forum/internal"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

// Fonction qui récupérer les informations utilisateur à partir de son uuid
func UserInfo(uuid string, database *sql.DB) (user_id, username, email, linkPhoto string) {
	if uuid != "" {
		var photo string
		usedcookie, errdb := database.Prepare("SELECT id, Username, Email, Photo FROM Users WHERE uuid = ?")
		CheckErr(errdb, "Requete DB")
		usedcookie.QueryRow(uuid).Scan(&user_id, &username, &email, &photo)

		linkPhoto = "/assets/photoCompte/" + photo
	}
	return user_id, username, email, linkPhoto
}

// Fonction qui récupérer les informations utilisateur de account
func UserAccount(uuid string, database *sql.DB) (user_id, username, email, linkPhoto, Age, Gender, FirstName, LastName string) {
	if uuid != "" {
		var photo string
		usedcookie, errdb := database.Prepare("SELECT id, Username, Email, Photo, Age, Gender, FirstName, LastName FROM Users WHERE uuid = ?")
		CheckErr(errdb, "Requete DB")
		usedcookie.QueryRow(uuid).Scan(&user_id, &username, &email, &photo, &Age, &Gender, &FirstName, &LastName)

		linkPhoto = "/assets/photoCompte/" + photo
	}
	return user_id, username, email, linkPhoto, Age, Gender, FirstName, LastName
}

// ----------------------------------------------------------
// ------- Fonction pour la gestion de l'utilisateur --------
// ----------------------------------------------------------

// Fonction qui met à jour le username (utilisé dans /Compte)
func UpdateUsername(user_id, newUsername string, database *sql.DB) bool {
	// check du nouveau username
	usedUsername := database.QueryRow("SELECT Username FROM Users WHERE Username = ?", newUsername).Scan()

	if usedUsername == sql.ErrNoRows {
		db, err := database.Prepare("UPDATE Users SET Username = ? WHERE ID = ?")
		CheckErr(err, "Prepare db")
		_, err = db.Exec(newUsername, user_id)
		CheckErr(err, "db Exec")
		return true
	}

	return false
}

// Fonction qui met à jour le mot de passe (utilisé dans /Compte)
func UpdatePassword(user_id, oldPass, newPass, confirmNewPass string, database *sql.DB) bool {
	// Obtenir le mot de passe haché du compte existant
	var pwdChecker string
	stmt, errPrep := database.Prepare("SELECT password FROM Users WHERE ID = ?")
	CheckErr(errPrep, "Prepare db")
	err := stmt.QueryRow(user_id).Scan(&pwdChecker)
	CheckErr(err, "queryrow fonction update : ")

	if newPass == confirmNewPass && bcrypt.CompareHashAndPassword([]byte(pwdChecker), []byte(oldPass)) == nil {
		hashedPass, _ := bcrypt.GenerateFromPassword([]byte(newPass), 8)
		db, err := database.Prepare("UPDATE Users SET Password = ? WHERE ID = ?")
		CheckErr(err, "Prepare db")
		_, err = db.Exec(hashedPass, user_id)
		CheckErr(err, "db Exec")
		return true
	}
	return false
}

// Fonction qui permet de supprimer un compte utilisateur (utilisé dans /Compte)
func DeleteAccount(user_id string, database *sql.DB) bool {
	db, err := database.Prepare("DELETE FROM Users WHERE ID = ?")
	CheckErr(err, "Prepare db")
	_, err = db.Exec(user_id)
	if err != nil {
		fmt.Printf("ERROR delete User account : %v\n", err)
	} else {
		return true
	}

	return false
}

// Fonction qui met à jour la photo de profils utilisateur (utilisé dans /Compte)
func UpdateProfilPicture(user_id, newName string, database *sql.DB) {
	db, err := database.Prepare("UPDATE Users SET Photo = ? WHERE ID = ?")
	CheckErr(err, "Prepare db")
	_, err = db.Exec(newName, user_id)
	CheckErr(err, "db Exec")
}

// ----------------------------------------------------
// ------------- Fonction pour les posts --------------
// ----------------------------------------------------

// Fonction retourne l'ensemble des catégories presente dans la table catégorie (utilisé dans /ViewPost)
func AllColumOfCategory(database *sql.DB) (res []data.StructCategorie) {
	var id, name string
	stmt, err := database.Prepare("SELECT ID, Name FROM Category")
	CheckErr(err, "db prepare")
	defer stmt.Close()
	rows, err := stmt.Query()
	CheckErr(err, "db query")
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &name)
		CheckErr(err, "db rows.Next scan")
		data := &data.StructCategorie{
			Id:   id,
			Name: name,
		}
		res = append(res, *data)
	}
	return res
}

func SearchComment(post_ID int, database *sql.DB) []*data.Post {
	var comment_ID, comment_Date, comment_Content, comment_user, comment_photo_user string
	rows, err := database.Query(`
	SELECT Comments.ID, Comments.Date, Comments.Content, Users.Username, Users.Photo
	FROM Comments 
	INNER JOIN Users ON Users.ID=Comments.User_ID
	WHERE Post_ID=?`, post_ID)
	CheckErr(err, "INNER JOIN SearchComment")

	var postTab []*data.Post
	if err != nil {
		CheckErr(err, "SearchComment")
		return nil
	} else {
		for rows.Next() {
			rows.Scan(&comment_ID, &comment_Date, &comment_Content, &comment_user, &comment_photo_user)

			if !strings.Contains(comment_photo_user, "http") {
				comment_photo_user = "/assets/photoCompte/" + comment_photo_user
			}
			comment_ID_int, err := strconv.Atoi(comment_ID)
			CheckErr(err, "Atoi SearchComment")
			PostsData := &data.Post{
				ID:          comment_ID_int,
				Date:        comment_Date,
				Author:      html.EscapeString(comment_user),
				PhotoAuthor: comment_photo_user,
				Content:     comment_Content,
			}
			postTab = append(postTab, PostsData)
		}
	}
	// fmt.Println("postTab : ", postTab)
	return postTab
}
