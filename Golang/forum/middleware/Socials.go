package utils

import (
	"database/sql"
	"fmt"
	"net/http"
)

func hasAlreadyLikedPost(postID int, username string, db *sql.DB) bool {
	var userID int
	err := db.QueryRow("SELECT ID FROM Users WHERE Username = ?", username).Scan(&userID)
	CheckErr(err, "select userID in hasalready liked post")
	var res int
	err = db.QueryRow("SELECT Socials FROM Socials WHERE User_ID = ? AND Post_ID = ?", userID, postID).Scan(&res)
	CheckErr(err, "select socials in hasalreadylikedpost")
	if res == 1 {
		return true
	} else {
		return false
	}
}

func hasAlreadyDislikedPost(postID int, username string, db *sql.DB) bool {
	var userID int
	err := db.QueryRow("SELECT ID FROM Users WHERE Username = ?", username).Scan(&userID)
	CheckErr(err, "select userID in hasalready DISliked post")
	var res int
	err = db.QueryRow("SELECT Socials FROM Socials WHERE User_ID = ? AND Post_ID = ?", userID, postID).Scan(&res)
	CheckErr(err, "select socials in has already DISlikedpost")
	if res == -1 {
		return true
	} else {
		return false
	}
}

func UpdateSocials(r *http.Request, username string, database *sql.DB, postID int) {
	if r.PostForm["socials"][0] == "liked" {
		//===================================
		// 		LIKES HANDLING			   //
		//===================================
		if !hasAlreadyLikedPost(postID, username, database) {
			//****************************
			// 			LIKING			//
			//****************************
			var userID string
			// var commentID string
			// var postID string
			err := database.QueryRow("SELECT ID FROM Users WHERE Username = ?", username).Scan(&userID)
			CheckErr(err, "couldn't get userID to updateSocials")
			/*
				err = database.QueryRow("SELECT ID FROM Comments WHERE User_ID = ?", userID).Scan(&commentID)
				   CheckErr(err, "couldn't get commentID to updateSocials")
				   err = database.QueryRow("SELECT ID FROM Posts WHERE User_ID = ?", userID).Scan(&postID)
				   CheckErr(err, "couldn't get postID to updateSocials")
			*/
			// Check if that post had been liked already by user. If yes --> UPDATE (from 0 or -1 to 1), if not --> INSERT
			var alreadyInSocialsTable bool
			rows, err := database.Query("SELECT * FROM Socials WHERE User_ID = ? AND Post_ID = ?", userID, postID)
			CheckErr(err, "select rows in UpdateSocials.go")
			if rows.Next() {
				alreadyInSocialsTable = true
			}
			rows.Close()
			if alreadyInSocialsTable {
				stmt, errdb := database.Prepare("UPDATE Socials SET Socials = 1 WHERE User_ID = ? AND Post_ID = ?")
				CheckErr(errdb, "Requete DB updateSocials : update")
				_, err = stmt.Exec(userID, postID)
				CheckErr(err, "insert like in socials.go")
			} else {
				stmt, errdb := database.Prepare("INSERT INTO Socials (User_ID, Comment_ID, Post_ID, Socials) VALUES (?, ?, ?, ?)")
				CheckErr(errdb, "Requete DB updateSocials : insert")
				_, err = stmt.Exec(userID, -1, postID, 1)
				CheckErr(err, "insert like in socials.go")
			}
		} else {
			//****************************
			// 			UNLIKING
			//****************************
			var userID string
			// var commentID string
			// var postID string
			err := database.QueryRow("SELECT ID FROM Users WHERE Username = ?", username).Scan(&userID)
			CheckErr(err, "couldn't get userID to updateSocials")
			/*
				err = database.QueryRow("SELECT ID FROM Comments WHERE User_ID = ?", userID).Scan(&commentID)
				   CheckErr(err, "couldn't get commentID to updateSocials")
				   err = database.QueryRow("SELECT ID FROM Posts WHERE User_ID = ?", userID).Scan(&postID)
				   CheckErr(err, "couldn't get postID to updateSocials")
			*/
			// Check if that post had been liked already by user. If yes --> UPDATE (from 1 to 0), if not --> INSERT
			var alreadyInSocialsTable bool
			rows, err := database.Query("SELECT * FROM Socials WHERE User_ID = ? AND Post_ID = ?", userID, postID)
			CheckErr(err, "select rows in UpdateSocials.go")
			if rows.Next() {
				alreadyInSocialsTable = true
			}
			rows.Close()
			if alreadyInSocialsTable {
				stmt, errdb := database.Prepare("UPDATE Socials SET Socials = 0 WHERE User_ID = ? AND Post_ID = ?")
				CheckErr(errdb, "Requete DB updateSocials : update")
				_, err = stmt.Exec(userID, postID)
				CheckErr(err, "insert like after update in socials.go")
			} else {
				stmt, errdb := database.Prepare("INSERT INTO Socials (User_ID, Comment_ID, Post_ID, Socials) VALUES (?, ?, ?, ?)")
				CheckErr(errdb, "Requete DB updateSocials : insert")
				_, err = stmt.Exec(userID, -1, postID, 1)
				CheckErr(err, "insert like after insert in socials.go")
			}
		}
	} else if r.PostForm["socials"][0] == "disliked" {
		//===================================
		// 		DISLIKES HANDLING
		//===================================
		if !hasAlreadyDislikedPost(postID, username, database) {
			//****************************
			// 			DISLIKING
			//****************************
			var userID string
			// var commentID string
			// var postID string
			err := database.QueryRow("SELECT ID FROM Users WHERE Username = ?", username).Scan(&userID)
			CheckErr(err, "couldn't get userID to updateSocials")
			/*
				err = database.QueryRow("SELECT ID FROM Comments WHERE User_ID = ?", userID).Scan(&commentID)
				   CheckErr(err, "couldn't get commentID to updateSocials")
				   err = database.QueryRow("SELECT ID FROM Posts WHERE User_ID = ?", userID).Scan(&postID)
				   CheckErr(err, "couldn't get postID to updateSocials")
			*/
			// Check if that post had been disliked already by user. If yes --> UPDATE (from -1 to 0), if not --> INSERT
			var alreadyInSocialsTable bool
			rows, err := database.Query("SELECT * FROM Socials WHERE User_ID = ? AND Post_ID = ?", userID, postID)
			CheckErr(err, "select rows in UpdateSocials.go")
			if rows.Next() {
				alreadyInSocialsTable = true
			}
			rows.Close()
			if alreadyInSocialsTable {
				stmt, errdb := database.Prepare("UPDATE Socials SET Socials = -1 WHERE User_ID = ? AND Post_ID = ?")
				CheckErr(errdb, "Requete DB updateSocials : update")
				_, err = stmt.Exec(userID, postID)
				CheckErr(err, "insert like in socials.go")
			} else {
				stmt, errdb := database.Prepare("INSERT INTO Socials (User_ID, Comment_ID, Post_ID, Socials) VALUES (?, ?, ?, ?)")
				CheckErr(errdb, "Requete DB updateSocials : insert")
				_, err = stmt.Exec(userID, -1, postID, -1)
				CheckErr(err, "insert like in socials.go")
			}
		} else {
			//****************************
			// 			UNDISLIKING
			//****************************
			var userID string
			// var commentID string
			// var postID string
			err := database.QueryRow("SELECT ID FROM Users WHERE Username = ?", username).Scan(&userID)
			CheckErr(err, "couldn't get userID to updateSocials")
			/*
				err = database.QueryRow("SELECT ID FROM Comments WHERE User_ID = ?", userID).Scan(&commentID)
				   CheckErr(err, "couldn't get commentID to updateSocials")
				   err = database.QueryRow("SELECT ID FROM Posts WHERE User_ID = ?", userID).Scan(&postID)
				   CheckErr(err, "couldn't get postID to updateSocials")
			*/
			// Check if that post had been liked already by user. If yes --> UPDATE (from -1 to 0), if not --> INSERT
			var alreadyInSocialsTable bool
			rows, err := database.Query("SELECT * FROM Socials WHERE User_ID = ? AND Post_ID = ?", userID, postID)
			CheckErr(err, "select rows in UpdateSocials.go")
			if rows.Next() {
				alreadyInSocialsTable = true
			}
			rows.Close()
			if alreadyInSocialsTable {
				stmt, errdb := database.Prepare("UPDATE Socials SET Socials = 0 WHERE User_ID = ? AND Post_ID = ?")
				CheckErr(errdb, "Requete DB updateSocials : update")
				_, err = stmt.Exec(userID, postID)
				CheckErr(err, "insert like after update in socials.go")
			} else {
				stmt, errdb := database.Prepare("INSERT INTO Socials (User_ID, Comment_ID, Post_ID, Socials) VALUES (?, ?, ?, ?)")
				CheckErr(errdb, "Requete DB updateSocials : insert")
				_, err = stmt.Exec(userID, -1, postID, -1)
				CheckErr(err, "insert like after insert in socials.go")
			}
		}
	} else {
		fmt.Println("==> wtf man?? <==")
	}
}

func SearchSocial(id string, database *sql.DB) (likes, dislikes int, likeCounter, dislikeCounter []int) {
	stmt, err := database.Prepare("SELECT * FROM Socials WHERE Post_ID = ? AND Socials = 1")
	CheckErr(err, "viewpostshandler")
	res, err := stmt.Query(id)
	CheckErr(err, "viewpostshandler")
	if err != sql.ErrNoRows {
		for res.Next() {
			likes++
		}
		likeCounter = append(likeCounter, likes)
	}
	res.Close()
	stmt, err = database.Prepare("SELECT * FROM Socials WHERE Post_ID = ? AND Socials = -1")
	CheckErr(err, "viewpostshandler")
	res, err = stmt.Query(id)
	CheckErr(err, "viewpostshandler")
	if err != sql.ErrNoRows {
		for res.Next() {
			dislikes++
		}
		dislikeCounter = append(dislikeCounter, dislikes)
	}
	res.Close()

	return likes, dislikes, likeCounter, dislikeCounter
}

// TEST
func hasAlreadyLikedComment(user_id string, commentID int, db *sql.DB) bool {
	var res string
	err := db.QueryRow("SELECT Socials FROM Socials WHERE User_ID = ? AND Comment_ID = ?", user_id, commentID).Scan(&res)
	// CheckErr(err, "select socials in has already likedComment")
	if err == sql.ErrNoRows {
		return false
	} else {
		return true
	}
}

func UpdateSocialsComment(user_id string, liked, commentID int, database *sql.DB) {
	if hasAlreadyLikedComment(user_id, commentID, database) { // update du dislike en like
		stmt, errdb := database.Prepare("UPDATE Socials SET Socials=? WHERE User_ID=? AND Comment_ID=?")
		CheckErr(errdb, "Requete DB update Comment like")
		_, err := stmt.Exec(liked, user_id, commentID)
		CheckErr(err, "Update like comment in socials.go")
	} else { // Insertion du like
		fmt.Println("HERE")
		stmt, errdb := database.Prepare("INSERT INTO Socials (User_ID, Comment_ID, Post_ID, Socials) VALUES (?, ?, ?, ?)")
		CheckErr(errdb, "Requete DB Insert Comment like")
		_, err := stmt.Exec(user_id, commentID, -1, liked)
		CheckErr(err, "insert like comment in socials.go")
	}
}
