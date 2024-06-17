package pkg

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func EventHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/event" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Si GET on rafraichit les posts disponibles (getPosts dans index.js).
	switch r.Method {
	case "GET":
		var events []EventGroup
		var err error

		param := r.URL.Query().Get("id")
		fmt.Println(param)

		events, err = FindEventByParam(param)
		if err != nil {
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}

		// On renvoit l'array de structures Post.
		resp, err := json.Marshal(events)
		if err != nil {
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resp)

		// Si POST on créé un nouveau post (postData(url:post)).
	case "POST":
		var newEvent EventGroup

		err := r.ParseMultipartForm(10 << 20) // 10MB max size
		if err != nil {

			http.Error(w, "400 bad request: "+err.Error(), http.StatusBadRequest)
			return
		}

		param := r.URL.Query().Get("id")
		fmt.Println("event_id", param)
		eventid, err := strconv.Atoi(param)

		newEvent.Title = r.FormValue("title")
		newEvent.Description = r.FormValue("content")
		newEvent.Date = r.FormValue("date")

		newEvent.IDGroup = eventid

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
		err = NewEvent(newEvent, curr)
		if err != nil {
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}

		// Send response indicating success
		msg := Resp{Msg: "New EventGroup added"}
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
func NewEvent(p EventGroup, u User) error {
	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
	}
	defer db.Close()
	p.Coming = 0
	p.NotComing = 0
	postInsert, err := db.Exec(`INSERT INTO EVENTGROUPS(IDGroup, Date, Title, Come, NotCome, Description, UserID_Sender) VALUES (?, ?, ?, ?, ?, ?, ?)`, p.IDGroup, p.Date, p.Title, p.Coming, p.NotComing, p.Description, u.Id)
	if err != nil {
		fmt.Println("error exec", err)
		return err
	}
	fmt.Println(postInsert)

	return nil
}

func FindEventByParam(data string) ([]EventGroup, error) {
	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
		return nil, err
	}
	defer db.Close()

	var rows *sql.Rows
	if data != "" {
		rows, err = db.Query(`SELECT * FROM EVENTGROUPS WHERE IDGroup = ? ORDER BY IDEvent DESC`, data)
	} else {
		rows, err = db.Query(`SELECT * FROM EVENTGROUPS ORDER BY IDEvent DESC`)
	}

	if err != nil {
		return []EventGroup{}, errors.New("failed to find events")
	}
	defer rows.Close()

	events, err := ConvertRowsToEvent(rows)
	if err != nil {
		return []EventGroup{}, errors.New("failed to convert rows to events")
	}

	return events, nil
}

// Mise en forme des rows en une array de structures Event.
func ConvertRowsToEvent(rows *sql.Rows) ([]EventGroup, error) {
	var events []EventGroup

	for rows.Next() {
		var p EventGroup
		err := rows.Scan(&p.IDEvent, &p.IDGroup, &p.Date, &p.Title, &p.Coming, &p.NotComing, &p.Description, &p.UserIDCreatorEvent)
		if err != nil {
			fmt.Println("err scan ConvertRowsToPost", err)
			break
		}
		events = append(events, p)
	}
	return events, nil
}
