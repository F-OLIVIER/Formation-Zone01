package pkg

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
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

		// Retrieve user ID from session cookie
		cookie, err := r.Cookie("session")
		if err != nil {
			return
		}

		foundVal := cookie.Value
		curr, err := CurrentUser(foundVal)
		if err != nil {
			DeleteCookie(w)
			return
		}
		var newEvent EventGroup

		err = r.ParseMultipartForm(10 << 20) // 10MB max size
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

		// Call the function to create a new post
		err = NewEvent(newEvent, curr)
		if err != nil {
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}

		ListUser, err := HowIsInGroup(newEvent, curr)
		if err != nil {
			fmt.Println("error exec", err)
			return
		}

		// Send response indicating success
		msg := Resp{Msg: "New EventGroup added", Target: ListUser}
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

func HowIsInGroup(p EventGroup, u User) ([]int, error) {
	GroupId := p.IDGroup
	OwnerUserId := u.Id

	// Connexion à la base de données
	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		return nil, fmt.Errorf("Erreur lors de l'ouverture de la base de données: %v", err)
	}
	defer db.Close()

	// Préparation de la requête SQL pour récupérer les UserID
	stmt, err := db.Prepare(`SELECT UserID FROM MEMBERSGROUPS WHERE IDGroup = ? AND UserID != ?`)
	if err != nil {
		return nil, fmt.Errorf("Erreur lors de la préparation de la requête SQL: %v", err)
	}
	defer stmt.Close()

	// Exécution de la requête
	rows, err := stmt.Query(GroupId, OwnerUserId)
	if err != nil {
		return nil, fmt.Errorf("Erreur lors de l'exécution de la requête SQL: %v", err)
	}
	defer rows.Close()

	// Parcourir les résultats et les ajouter à un tableau
	var userIds []int
	for rows.Next() {
		var userId int
		err = rows.Scan(&userId)
		if err != nil {
			return nil, fmt.Errorf("Erreur lors de la lecture des résultats: %v", err)
		}
		userIds = append(userIds, userId)
	}

	// Vérification des erreurs lors de la itération
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Erreur lors de l'itération des résultats: %v", err)
	}

	return userIds, nil
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

	ListUser, err := HowIsInGroup(p, u)
	if err != nil {
		fmt.Println("error exec", err)
		return err
	}
	date := time.Now().Format("01-02-2006 15:04:05")

	// convert int64 to int
	lastInsertId, _ := postInsert.LastInsertId()

	for _, user := range ListUser {
		InsertNotif(int(lastInsertId), user, date, "event", db)
	}

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
