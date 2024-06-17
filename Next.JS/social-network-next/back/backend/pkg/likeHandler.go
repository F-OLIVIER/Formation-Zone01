package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func LikeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/like" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method == "POST" {

		var like Like
		err := json.NewDecoder(r.Body).Decode(&like)
		if err != nil {
			http.Error(w, "500 internal server error: Failed to connect to database. "+err.Error(), http.StatusInternalServerError)
		}

		fmt.Println("here")

		// Retrieve user ID from session cookie
		cookie, err := r.Cookie("session")
		if err != nil {
			return
		}

		fmt.Println(cookie)

		foundVal := cookie.Value
		curr, err := CurrentUser(foundVal)
		if err != nil {
			return
		}

		like.User_id = curr.Id

		db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
		if err != nil {
			fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
			return
		}
		defer db.Close()

		var id int
		var likeStatus bool
		err = db.QueryRow("SELECT ID, Like FROM liked_post WHERE post_id = ? AND user_id = ?", like.Post_Id, like.User_id).Scan(&id, &likeStatus)

		if err == nil {
			// If the "like" exists, toggle the like status
			_, err = db.Exec("UPDATE liked_post SET Like = CASE WHEN Like THEN false ELSE true END WHERE post_id = ? AND user_id = ?", like.Post_Id, like.User_id)
			if likeStatus { // If likeStatus is true, decrement the like count
				_, err = db.Exec("UPDATE POST SET Likes = Likes - 1 WHERE IDPost = ?", like.Post_Id)
			} else { // If likeStatus is false, increment the like count
				_, err = db.Exec("UPDATE POST SET Likes = Likes + 1 WHERE IDPost = ?", like.Post_Id)
			}
		} else if err == sql.ErrNoRows {
			// No existing "like" found, so the user wants to like the post: insert a new record
			_, err = db.Exec("INSERT INTO liked_post (user_id, post_id, like) VALUES (?, ?, true)", like.User_id, like.Post_Id)
			_, err = db.Exec("UPDATE POST SET Likes = Likes + 1 WHERE IDPost = ?", like.Post_Id)
		} else {
			// Handle other potential errors
			http.Error(w, "500 Internal Server Error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		msg := Resp{Msg: "✅ Like Updated"}
		resp, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "500 internal server error: Failed to marshal response."+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}
