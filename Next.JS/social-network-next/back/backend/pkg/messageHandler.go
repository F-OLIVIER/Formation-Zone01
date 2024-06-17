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

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/message" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	// Si GET on récupère l'historique des messages d'un chat.
	switch r.Method {
	case "GET":
		// On récupère le sender s = utilisateur connecté via sa session cookie.
		cookie, err := r.Cookie("session")
		if err != nil {
			return
		}
		foundVal := cookie.Value
		curr, err := CurrentUser(foundVal)
		if err != nil {
			return
		}
		s := strconv.Itoa(curr.Id)

		// On récupère le receiver r = id dans l'url quand on a cliqué sur son username.
		firstId, _ := strconv.Atoi(r.URL.Query().Get("firstId"))
		r := r.URL.Query().Get("receiver")

		// Récupération de tous les messages de la database en commun entre les deux users.
		messages, err := FindChatMessages(s, r, firstId)
		if err != nil {
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}

		// On retourne l'ensemble des messages.
		resp, err := json.Marshal(messages)
		if err != nil {
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	// Si POST, on envoit un message (écrit en fait dans la database, et il sera affiché par le GET au-dessus)
	case "POST":

		var newMessage Message
		// Décodage de la requête dans newMessage.
		err := json.NewDecoder(r.Body).Decode(&newMessage)
		if err != nil {
			http.Error(w, "400 bad request.", http.StatusBadRequest)
			return
		}

		// Appel de la fonction NewMessage (ci-dessous)
		err = NewMessage(newMessage)
		if err != nil {
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// Création d'un nouveau message.
func NewMessage(m Message) error {
	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
	}
	defer db.Close()

	result, err := db.Exec(`INSERT INTO messages(sender_id, receiver_id, content, date) values(?, ?, ?, ?)`, m.Sender_id, m.Receiver_id, m.Content, m.Date)
	if err != nil {
		return err
	}

	err = UpdateChatTime(m.Sender_id, m.Receiver_id, db)
	if err != nil {
		return err
	}

	dt := time.Now().Format("01-02-2006 15:04:05")

	idMessage, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Insertion du nouveau commentaire dans notifs
	InsertNotif(int(idMessage), m.Receiver_id, dt, "mp", db)

	return nil
}

// Récupération de tous les messages de la database en commun entre 2 users donnés.
func FindChatMessages(sender, receiver string, firstId int) ([]Message, error) {
	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
	}
	defer db.Close()
	// On converti les id de sender et receiver en INT.
	s, err := strconv.Atoi(sender)
	if err != nil {
		return []Message{}, errors.New("sender id must be an integer")
	}
	r, err := strconv.Atoi(receiver)
	if err != nil {
		return []Message{}, errors.New("receiver id must be an integer")
	}

	q, err := db.Query(`SELECT * FROM messages WHERE ((sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)) AND ( id <= ? ) ORDER BY id DESC LIMIT 10`, s, r, r, s, firstId)
	if err != nil {
		return []Message{}, errors.New("could not find chat messages")
	}

	messages, err := ConvertRowToMessage(q)
	if err != nil {
		return []Message{}, errors.New("failed to convert")
	}

	return messages, nil
}

// find the last message between two users
func FindLastMessage(sender, receiver string) (Message, error) {
	// Opens the database
	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
	}
	defer db.Close()

	// Converts sender and receiver ids to integers
	s, err := strconv.Atoi(sender)
	if err != nil {
		return Message{}, errors.New("sender id must be an integer")
	}

	r, err := strconv.Atoi(receiver)
	if err != nil {
		return Message{}, errors.New("receiver id must be an integer")
	}

	// search database for last message between the two users
	q, err := db.Query(`SELECT * FROM messages WHERE ((sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)) ORDER BY id DESC LIMIT 1`, r, r, r, s)
	if err != nil {
		return Message{}, errors.New("could not find chat messages")
	}

	// Converts rows to an array of message structs
	messages, err := ConvertRowToMessage(q)
	if err != nil {
		return Message{}, errors.New("failed to convert")
	}

	return messages[0], nil
}

// Mise en forme des rows en une array de structures Message.
func ConvertRowToMessage(rows *sql.Rows) ([]Message, error) {
	var messages []Message
	for rows.Next() {
		var m Message
		err := rows.Scan(&m.Id, &m.Sender_id, &m.Receiver_id, &m.Content, &m.Date)
		if err != nil {
			break
		}

		messages = append(messages, m)
	}
	return messages, nil
}
