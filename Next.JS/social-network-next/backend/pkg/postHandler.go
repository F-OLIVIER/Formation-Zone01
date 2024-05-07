package pkg

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Si GET on rafraichit les posts disponibles (getPosts dans index.js).
	switch r.Method {
	case "GET":
		var posts []Post
		var err error

		param := r.URL.Query().Get("id")
		fmt.Println(param)

		posts, err = FindPostByParam(param)
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
		var newPost Post

		err := r.ParseMultipartForm(10 << 20) // 10MB max size
		if err != nil {

			http.Error(w, "400 bad request: "+err.Error(), http.StatusBadRequest)
			return
		}

		newPost.Title = r.FormValue("title")
		newPost.Content = r.FormValue("content")
		privateStr := r.FormValue("privacy")
		privateInt, err := strconv.Atoi(privateStr)
		if err != nil {
			http.Error(w, "400 bad request: invalid privacy value", http.StatusBadRequest)
			return
		}

		var listuserautorized []string
		if privateInt == 2 {
			listuserautorizedString := r.FormValue("selectedFollowers")
			listuserautorized = strings.Split(listuserautorizedString, ",")
		}

		newPost.Private = privateInt
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
		err = NewPost(newPost, curr, listuserautorized)
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
func NewPost(p Post, u User, listuserautorized []string) error {
	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
	}
	defer db.Close()

	dt := time.Now().Format("01-02-2006 15:04:05")

	postInsert, err := db.Exec(`INSERT INTO POST(UserID, Title, PostContent, Date, Image, Private, Likes, NbComments) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, u.Id, p.Title, p.Content, dt, p.Image, p.Private, p.Likes, p.NbComments)
	if err != nil {
		fmt.Println("error exec", err)
		return err
	}

	if len(listuserautorized) > 0 {
		// Récupérer l'ID généré lors de l'insertion du post
		lastInsertId, err := postInsert.LastInsertId()
		if err != nil {
			fmt.Println("error getting last insert id")
			return err
		}

		for i := 0; i < len(listuserautorized); i++ {
			_, err = db.Exec(`INSERT INTO POSTSELECTUSERS (IDPost, UserID) VALUES (?, ?)`, lastInsertId, listuserautorized[i])
			if err != nil {
				fmt.Println("error exec listuserautorized")
				return err
			}
		}
	}

	return nil
}

// Récupération des posts en fonction d'un paramétre, exemple filtre d'une catégorie.
func FindPostByParam(data string) ([]Post, error) {
	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM POST ORDER BY IDPost DESC`)
	if err != nil {
		return []Post{}, errors.New("failed to find posts")
	}
	var posts []Post
	if data != "" {

		i, err := strconv.Atoi(data)
		if err != nil {
			return []Post{}, errors.New("id must be an integer")
		}
		posts, err = ConvertRowsToPost(db, rows, i)
		if err != nil {
			return []Post{}, errors.New("failed to convert")
		}
	} else {
		posts, err = ConvertRowToPost(rows)
		if err != nil {
			return []Post{}, errors.New("failed to convert")
		}
	}

	return posts, nil
}

// Mise en forme des rows en une array de structures Post.
func ConvertRowToPost(rows *sql.Rows) ([]Post, error) {
	var posts []Post
	for rows.Next() {
		var p Post
		err := rows.Scan(&p.Id, &p.User_id, &p.Title, &p.Content, &p.Date, &p.Image, &p.Private, &p.Likes, &p.NbComments)
		if err != nil {
			break
		}
		posts = append(posts, p)
	}
	return posts, nil
}

// Mise en forme des rows en une array de structures Post.
func ConvertRowsToPost(db *sql.DB, rows *sql.Rows, i int) ([]Post, error) {
	var posts []Post
	listOfFollow := ListFollow(i, "followings", 0)

	for rows.Next() {
		var p Post
		err := rows.Scan(&p.Id, &p.User_id, &p.Title, &p.Content, &p.Date, &p.Image, &p.Private, &p.Likes, &p.NbComments)
		if err != nil {
			fmt.Println("err scan ConvertRowsToPost", err)
			break
		}
		if p.User_id == i || p.Private == 1 {
			posts = append(posts, p)
		} else if p.Private == 0 {
			for _, followedUser := range listOfFollow {
				if p.User_id == followedUser.Id {
					posts = append(posts, p)
					break
				}
			}
		} else if p.Private == 2 {
			autorized := 0
			stmt, errdb := db.Prepare("SELECT ID FROM POSTSELECTUSERS WHERE IDPost = ? AND UserID = ?")
			CheckErr(errdb, "Requete DB ConvertRowsToPost")
			stmt.QueryRow(p.Id, i).Scan(&autorized)
			if autorized > 0 {
				posts = append(posts, p)
			}
		} else {
			fmt.Println("PROBLEME ConvertRowsToPost")
		}
	}
	return posts, nil
}
