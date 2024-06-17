package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func ComingEventHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/comingevent" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method == "POST" {

		var come Coming
		err := json.NewDecoder(r.Body).Decode(&come)
		if err != nil {
			http.Error(w, "500 internal server error: Failed to connect to database. "+err.Error(), http.StatusInternalServerError)
		}

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

		come.User_id = curr.Id

		db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
		if err != nil {
			fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
			return
		}
		defer db.Close()
		var msg Resp
		var id int
		var comeStatus bool
		var notComeStatus bool
		err = db.QueryRow("SELECT ID, Come, NotCome FROM coming_event WHERE event_id = ? AND user_id = ?", come.Event_id, come.User_id).Scan(&id, &comeStatus, &notComeStatus)

		if !notComeStatus {

			if err == nil {
				// If the "come" exists, toggle the come status
				_, err = db.Exec("UPDATE coming_event SET Come = CASE WHEN Come THEN false ELSE true END WHERE event_id = ? AND user_id = ?", come.Event_id, come.User_id)
				if comeStatus { // If likeStatus is true, decrement the come count
					_, err = db.Exec("UPDATE EVENTGROUPS SET Come = Come - 1 WHERE IDevent = ?", come.Event_id)
				} else { // If likeStatus is false, increment the come count
					_, err = db.Exec("UPDATE EVENTGROUPS SET Come = Come + 1 WHERE IDEvent = ?", come.Event_id)
				}
			} else if err == sql.ErrNoRows {
				// No existing "come" found, so the user wants to come the post: insert a new record
				_, err = db.Exec("INSERT INTO coming_event (user_id, event_id, come, notcome) VALUES (?, ?, true, false)", come.User_id, come.Event_id)
				_, err = db.Exec("UPDATE EVENTGROUPS SET Come = Come + 1 WHERE IDEvent = ?", come.Event_id)
			} else {
				// Handle other potential errors
				http.Error(w, "500 Internal Server Error: "+err.Error(), http.StatusInternalServerError)
				return
			}
		} else if !comeStatus {
			_, err = db.Exec("UPDATE EVENTGROUPS SET NotCome = NotCome - 1 WHERE IDevent = ?", come.Event_id)
			_, err = db.Exec("UPDATE EVENTGROUPS SET Come = Come + 1 WHERE IDEvent = ?", come.Event_id)
			_, err = db.Exec("UPDATE coming_event SET NotCome = CASE WHEN NotCome THEN false ELSE true END WHERE event_id = ? AND user_id = ?", come.Event_id, come.User_id)
			_, err = db.Exec("UPDATE coming_event SET Come = CASE WHEN Come THEN false ELSE true END WHERE event_id = ? AND user_id = ?", come.Event_id, come.User_id)
		}

		msg = Resp{Msg: "✅ Like Updated"}

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

func NotComingEventHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/notcomingevent" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method == "POST" {

		var come Coming
		err := json.NewDecoder(r.Body).Decode(&come)
		if err != nil {
			http.Error(w, "500 internal server error: Failed to connect to database. "+err.Error(), http.StatusInternalServerError)
		}

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

		come.User_id = curr.Id

		db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
		if err != nil {
			fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
			return
		}
		defer db.Close()
		var msg Resp
		var id int
		var comeStatus bool
		var notComeStatus bool
		err = db.QueryRow("SELECT ID, Come, NotCome FROM coming_event WHERE event_id = ? AND user_id = ?", come.Event_id, come.User_id).Scan(&id, &comeStatus, &notComeStatus)

		if !comeStatus {
			if err == nil {
				// If the "come" exists, toggle the come status
				_, err = db.Exec("UPDATE coming_event SET NotCome = CASE WHEN NotCome THEN false ELSE true END WHERE event_id = ? AND user_id = ?", come.Event_id, come.User_id)
				if notComeStatus { // If likeStatus is true, decrement the come count
					_, err = db.Exec("UPDATE EVENTGROUPS SET NotCome = NotCome - 1 WHERE IDevent = ?", come.Event_id)
				} else { // If likeStatus is false, increment the come count
					_, err = db.Exec("UPDATE EVENTGROUPS SET NotCome = NotCome + 1 WHERE IDEvent = ?", come.Event_id)
				}
			} else if err == sql.ErrNoRows {
				// No existing "come" found, so the user wants to come the post: insert a new record
				_, err = db.Exec("INSERT INTO coming_event (user_id, event_id, come, notcome) VALUES (?, ?, true, false)", come.User_id, come.Event_id)
				_, err = db.Exec("UPDATE EVENTGROUPS SET NotCome = NotCome + 1 WHERE IDEvent = ?", come.Event_id)
			} else {
				// Handle other potential errors
				http.Error(w, "500 Internal Server Error: "+err.Error(), http.StatusInternalServerError)
				return
			}
		} else if !notComeStatus {
			_, err = db.Exec("UPDATE EVENTGROUPS SET Come = Come - 1 WHERE IDevent = ?", come.Event_id)
			_, err = db.Exec("UPDATE EVENTGROUPS SET NotCome = NotCome + 1 WHERE IDEvent = ?", come.Event_id)
			_, err = db.Exec("UPDATE coming_event SET NotCome = CASE WHEN NotCome THEN false ELSE true END WHERE event_id = ? AND user_id = ?", come.Event_id, come.User_id)
			_, err = db.Exec("UPDATE coming_event SET Come = CASE WHEN Come THEN false ELSE true END WHERE event_id = ? AND user_id = ?", come.Event_id, come.User_id)
		}

		msg = Resp{Msg: "✅ Like Updated"}
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
