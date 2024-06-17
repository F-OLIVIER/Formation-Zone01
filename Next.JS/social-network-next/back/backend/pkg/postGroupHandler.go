package pkg

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func PostGroupHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/postgroup" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Si GET on rafraichit les posts disponibles (getPosts dans index.js).
	switch r.Method {
	case "GET":
		var posts []PostGroup
		var err error

		param := r.URL.Query().Get("id")
		fmt.Println(param)

		posts, err = FindPostByParamGroup(param)
		if err != nil {
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}

		// On renvoit l'array de structures Post.
		resp, err := json.Marshal(posts)
		if err != nil {
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resp)

		// Si POST on créé un nouveau post (postData(url:post)).
	case "POST":
		var newPost PostGroup

		err := r.ParseMultipartForm(10 << 20) // 10MB max size
		if err != nil {

			http.Error(w, "400 bad request: "+err.Error(), http.StatusBadRequest)
			return
		}

		param := r.URL.Query().Get("id")
		fmt.Println("group_id", param)
		groupid, err := strconv.Atoi(param)

		newPost.Title = r.FormValue("title")
		newPost.Content = r.FormValue("content")

		newPost.Group_id = groupid
		_, imageHeader, err := r.FormFile("image")
		if err != nil && err != http.ErrMissingFile {
			http.Error(w, "Error processing image", http.StatusInternalServerError)
			return
		}

		if imageHeader != nil {
			file, _, err := r.FormFile("image")
			if err != nil {
				fmt.Println("here")
				http.Error(w, "400 bad request: "+err.Error(), http.StatusBadRequest)
				return
			}
			defer file.Close()

			// Read the file data
			fileBytes, err := io.ReadAll(file)
			if err != nil {
				http.Error(w, "500 internal server error: "+err.Error(), http.StatusInternalServerError)
				return
			}

			newPost.Image = fileBytes
		}

		// Retrieve user ID from session cookie
		cookie, err := r.Cookie("session")
		if err != nil {
			return
		}

		foundVal := cookie.Value
		curr, err := CurrentUser(foundVal)
		if err != nil {
			return
		}

		// Call the function to create a new post
		err = NewPostGroup(newPost, curr)
		if err != nil {
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}

		// Send response indicating success
		msg := Resp{Msg: "New post added"}
		resp, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	default:
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// Création d'un nouveau post.
func NewPostGroup(p PostGroup, u User) error {
	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
	}
	defer db.Close()

	dt := time.Now().Format("01-02-2006 15:04:05")

	postInsert, err := db.Exec(`INSERT INTO POST_GROUP(UserID, Title, PostContent, Date, Image, GroupID, Likes, NbComments) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, u.Id, p.Title, p.Content, dt, p.Image, p.Group_id, p.Likes, p.NbComments)
	if err != nil {
		fmt.Println("error exec", err)
		return err
	}
	fmt.Println(postInsert)

	return nil
}
func ConvertRowsToPostGroup(rows *sql.Rows) ([]PostGroup, error) {
	var posts []PostGroup

	for rows.Next() {
		var p PostGroup
		err := rows.Scan(&p.Id, &p.User_id, &p.Title, &p.Content, &p.Date, &p.Image, &p.Group_id, &p.Likes, &p.NbComments)
		if err != nil {
			fmt.Println("err scan ConvertRowsToPostGroup", err)
			break
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func FindPostByParamGroup(data string) ([]PostGroup, error) {
	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
		return nil, err
	}
	defer db.Close()

	var rows *sql.Rows
	if data != "" {
		rows, err = db.Query(`SELECT * FROM POST_GROUP WHERE GroupID = ? ORDER BY IDPost DESC`, data)
	} else {
		rows, err = db.Query(`SELECT * FROM POST_GROUP ORDER BY IDPost DESC`)
	}

	if err != nil {
		return []PostGroup{}, errors.New("failed to find posts")
	}
	defer rows.Close()

	posts, err := ConvertRowsToPostGroup(rows)
	if err != nil {
		return []PostGroup{}, errors.New("failed to convert rows to posts")
	}

	return posts, nil
}
