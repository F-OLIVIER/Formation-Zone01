package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// DISCLAIMER : Ce code n'est pas uniquement de moi, aide sur git, GPT et tutoriels.

// Génère la liste d'users disponible pour le chat.
func ChatHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/chat" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// On ne veut que une requête GET.
	// updateUsers dans index.js
	if r.Method != "GET" {
		http.Error(w, "405 method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	// On récupére l'id de l'user connecté.
	user_id := r.URL.Query().Get("user_id")
	if user_id == "" {
		http.Error(w, "400 bad request", http.StatusBadRequest)
		return
	}

	uid, err := strconv.Atoi(user_id)
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}

	// On recherche tous les chats disponibles pour l'id.
	users, err := FindUserChats(uid)
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}

	// On check sa position pour savoir si on ouvre le chat 1/2 ou 2/1. ("receiver ou sender")
	var chatUsers []int
	for _, u := range users {
		if u.User_one == uid {
			chatUsers = append(chatUsers, u.User_two)
		} else {
			chatUsers = append(chatUsers, u.User_one)
		}
	}

	// On retourne le bon "chat" et l'event type ""
	msg := OnlineUsers{
		UserIds:  chatUsers,
		Msg_type: "",
	}
	resp, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// Lors d'un envoi de message dans le chat, actualise l'heure (appelé par message.go)
func UpdateChatTime(user1, user2 int, db *sql.DB) error {
	now := time.Now()

	chats, err := FindChatsBetween(user1, user2, db)
	if err != nil {
		return err
	}
	// pas de chat existant, on insert.
	if len(chats) == 0 {
		_, err = db.Exec(`INSERT INTO chats(id_one, id_two, time) values(? ,?, ?)`, user1, user2, now.UnixMilli())
		if err != nil {
			return err
		}
		// chat existant, on update.
	} else {
		_, err = db.Exec(`UPDATE chats SET time = ? WHERE id_one = ? AND id_two = ?`, now.UnixMilli(), chats[0].User_one, chats[0].User_two)
		if err != nil {
			return err
		}
	}
	return nil
}

// Recherche de tous les chats en commun avec un id donné.
func FindUserChats(uid int) ([]Chat, error) {
	var q *sql.Rows

	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
	}
	defer db.Close()

	defer db.Close()

	q, err = db.Query(`SELECT * FROM chats WHERE id_one = ? OR id_two = ? ORDER BY time DESC`, uid, uid)
	if err != nil {
		return []Chat{}, err
	}

	users, err := ConvertRowToChat(q)
	if err != nil {
		return []Chat{}, err
	}

	return users, nil
}

// On recherche les chats en commun pour 2 users donnés.
func FindChatsBetween(u1, u2 int, db *sql.DB) ([]Chat, error) {
	var q *sql.Rows
	q, err := db.Query(`SELECT * FROM chats WHERE id_one = ? AND id_two = ? OR id_one = ? AND id_two = ?`, u1, u2, u2, u1)
	fmt.Print(q)
	if err != nil {
		return []Chat{}, err
	}

	users, err := ConvertRowToChat(q)
	if err != nil {
		return []Chat{}, err
	}

	return users, nil
}

// Mise en forme des rows en une array de structures Chat.
func ConvertRowToChat(rows *sql.Rows) ([]Chat, error) {
	var chats []Chat
	defer rows.Close()
	for rows.Next() {
		var c Chat
		err := rows.Scan(&c.User_one, &c.User_two, &c.Time)
		if err != nil {
			break
		}
		chats = append(chats, c)
	}
	return chats, nil
}
