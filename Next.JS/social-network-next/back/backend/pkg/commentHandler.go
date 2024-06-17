package pkg

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func CommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/comment" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Si GET on rafraichit les commentaires disponibles (getComments dans index.js).
	switch r.Method {
	case "GET":

		// Il faut forcément un param car un commentaire seul n'existe pas.
		param := r.URL.Query().Get("param")
		data := r.URL.Query().Get("data")

		if param == "" || data == "" {
			http.Error(w, "400 bad request", http.StatusBadRequest)
			return
		}

		// On recherche en fonction du param les commentaires associés
		comments, err := FindCommentByParam(param, data)
		if err != nil {
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}

		// On renvoit l'array de structures Comment.
		resp, err := json.Marshal(comments)
		if err != nil {
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resp)

	// Si POST on créé un nouveau commentaire (postData(url:comment))
	case "POST":
		var newComment Comment
		err := r.ParseMultipartForm(10 << 20) // 10MB max size
		if err != nil {
			http.Error(w, "400 bad request: "+err.Error(), http.StatusBadRequest)
			return
		}
		newComment.Content = r.FormValue("content")
		privateStr := r.FormValue("post_id")
		privateInt, err := strconv.Atoi(privateStr)
		if err != nil {
			http.Error(w, "400 bad request: invalid privacy value", http.StatusBadRequest)
			return
		}
		newComment.Post_id = privateInt

		_, imageHeader, err := r.FormFile("image")
		if err != nil && err != http.ErrMissingFile {
			http.Error(w, "Error processing image", http.StatusInternalServerError)
			return
		}

		if imageHeader != nil {
			fmt.Println("i have an image")
			file, _, err := r.FormFile("image")
			if err != nil {
				fmt.Println("here")
				http.Error(w, "400 bad request: "+err.Error(), http.StatusBadRequest)
				return
			}
			defer file.Close()

			// Read the file data
			fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				http.Error(w, "500 internal server error: "+err.Error(), http.StatusInternalServerError)
				return
			}

			newComment.Image = fileBytes
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

		newComment.User_id = curr.Id

		// Appel de la fonction NewComment (ci-dessous)
		err = NewComment(newComment)
		if err != nil {
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}

		var userIDFromPost int
		db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
		if err != nil {
			fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
		}
		defer db.Close()
		// Récupérer l'USER ID de la table POST en utilisant l'IDPost de la table COMMENT
		err = db.QueryRow(`SELECT UserID FROM POST WHERE IDPost = ?`, newComment.Post_id).Scan(&userIDFromPost)
		if err != nil {
			fmt.Println("Erreur lors de la récupération de l'USER ID depuis POST:", err)
			return
		}
		// On reformate JSON pour envoyer l'information que notre commentaire est créé.
		msg := Resp{Msg: "Sent comment", Receiver: userIDFromPost}
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

// Création d'un nouveau commentaire.
func NewComment(c Comment) error {
	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
	}
	defer db.Close()

	dt := time.Now().Format("01-02-2006 15:04:05")

	// Récupérer l'USER ID de la table POST en utilisant l'IDPost de la table COMMENT
	var userIDFromPost int
	err = db.QueryRow(`SELECT UserID FROM POST WHERE IDPost = ?`, c.Post_id).Scan(&userIDFromPost)
	if err != nil {
		fmt.Println("Erreur lors de la récupération de l'USER ID depuis POST:", err)
		return err
	}

	result, err := db.Exec(`INSERT INTO COMMENT(IDPost, UserID, CommentContent, Date, Image) values(?, ?, ?, ?, ?)`, c.Post_id, c.User_id, c.Content, dt, c.Image)
	if err != nil {
		return err
	}

	idComment, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Insertion du nouveau commentaire dans notifs
	InsertNotif(int(idComment), userIDFromPost, dt, "comment", db)

	_, err = db.Exec("UPDATE POST SET NbComments = NbComments + 1 WHERE IDPost = ?", c.Post_id)
	if err != nil {
		return err
	}

	return nil
}

// Récupération des commentaires en fonction d'un paramétre, id_post ici.
func FindCommentByParam(param, data string) ([]Comment, error) {
	var q *sql.Rows
	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
	}
	defer db.Close()

	i, err := strconv.Atoi(data)
	if err != nil {
		return []Comment{}, errors.New("not an integer")
	}
	switch param {
	case "id":
		q, err = db.Query(`SELECT * FROM COMMENT WHERE IDComment = ?`, i)
		if err != nil {
			return []Comment{}, errors.New("could not find id")
		}
	case "post_id":
		q, err = db.Query(`SELECT * FROM COMMENT WHERE IDPost = ?`, i)
		if err != nil {
			return []Comment{}, errors.New("could not find post_id")
		}
	case "user_id":
		q, err = db.Query(`SELECT * FROM COMMENT WHERE UserID = ?`, i)
		if err != nil {
			return []Comment{}, errors.New("could not find user_id")
		}
	default:
		return []Comment{}, errors.New("cannot search by that parameter")
	}

	comments, err := ConvertRowToComment(q)
	if err != nil {
		return []Comment{}, errors.New("failed to convert")
	}

	return comments, nil
}

// Converts comment table query results to an array of comment structs
func ConvertRowToComment(rows *sql.Rows) ([]Comment, error) {
	var comments []Comment
	for rows.Next() {
		var c Comment
		err := rows.Scan(&c.Id, &c.Post_id, &c.User_id, &c.Content, &c.Date, &c.Image)
		if err != nil {
			break
		}
		comments = append(comments, c)
	}
	return comments, nil
}
