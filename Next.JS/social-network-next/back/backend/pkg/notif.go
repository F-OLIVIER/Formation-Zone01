package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type Receiver struct {
	UserIDReceiver   int    `json:"user_id_receiver"`
	UserID_Followers int    `json:"user_id_followers"`
	IDnotif          int    `json:"id_notif"`
	TypeID           string `json:"Type_id"`
	TypeNotif        string `json:"type_notif"`
	AddOrDelete      bool   `json:"add_or_delete"`
}

func NotifHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data Receiver
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if data.TypeNotif == "" {
		data.TypeNotif = "get"
	}

	fmt.Println("Data :", data)

	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		http.Error(w, "Erreur lors de l'ouverture de la base de données", http.StatusInternalServerError)
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
		return
	}
	defer db.Close()

	var jsonResponse map[string]interface{}

	switch data.TypeNotif {
	case "get": // Récupération des notifications
		ListFollowers, listMP, listPost, listComment, listGroup, listEvent, listGroupMp, listAskGroup := GetNotif(data.UserIDReceiver, db)

		jsonResponse = map[string]interface{}{
			"success":       true,
			"listFollowers": ListFollowers,
			"listMP":        listMP,
			"listPost":      listPost,
			"listComment":   listComment,
			"listGroup":     listGroup,
			"listEvent":     listEvent,
			"listGroupMp":   listGroupMp,
			"listAskGroup":  listAskGroup,
		}

	case "post": // modification d'une notification car accepté ou refusé
		if data.AddOrDelete { // si true, acceptation de la notification
			switch data.TypeID {
			case "IDFollow":
				validateFollow(data.UserIDReceiver, data.UserID_Followers, db)
			case "IDPrivateMessage":
				// ne rien faire car le message doit rester pour l'historique
			case "IDPost":
				// ne rien faire car le post doit rester
			case "IDComment":
				// ne rien faire car le comentaire doit rester
			case "IDGroup":
				stmt, err := db.Prepare("UPDATE MEMBERSGROUPS SET ValidationInvite = true WHERE IDgroup = ? AND UserID = ?")
				CheckErr(err, "NotifHandler AddOrDelete IDGroup Prepare db")
				_, err = stmt.Exec(data.IDnotif, data.UserIDReceiver)
				CheckErr(err, "NotifHandler AddOrDelete IDGroup db Exec")
			case "IDEvent":
				stmt, err := db.Prepare("UPDATE RESPONSEEVENTGROUPS SET option = 1 WHERE IDEvent = ? AND UserID = ?")
				CheckErr(err, "NotifHandler AddOrDelete IDEvent Prepare db")
				_, err = stmt.Exec(data.IDnotif, data.UserIDReceiver)
				CheckErr(err, "NotifHandler AddOrDelete IDEvent db Exec")
			}

		} else { // si false, refus de la notification
			switch data.TypeID {
			case "IDFollow":
				stmt, err := db.Prepare("DELETE FROM FOLLOWERS WHERE UserID_Following = ? AND UserID_Follower = ?")
				CheckErr(err, "NotifHandler !AddOrDelete IDFollow Prepare db")
				_, err = stmt.Exec(data.UserIDReceiver, data.UserID_Followers)
				CheckErr(err, "NotifHandler !AddOrDelete IDFollow db Exec")
			case "IDPrivateMessage":
				// ne rien faire car le message doit rester pour l'historique
			case "IDPost":
				// ne rien faire car le post doit rester
			case "IDComment":
				// ne rien faire car le comentaire doit rester
			case "IDGroup":
				stmt, err := db.Prepare("DELETE FROM MEMBERSGROUPS WHERE IDgroup = ? AND UserID = ?")
				CheckErr(err, "NotifHandler !AddOrDelete IDGroup  Prepare db")
				_, err = stmt.Exec(data.IDnotif, data.UserIDReceiver)
				CheckErr(err, "NotifHandler !AddOrDelete IDGroup  db Exec")
			case "IDEvent":
				stmt, err := db.Prepare("DELETE FROM  RESPONSEEVENTGROUPS WHERE IDEvent = ? AND UserID = ?")
				CheckErr(err, "NotifHandler !AddOrDelete IDEvent  Prepare db")
				_, err = stmt.Exec(data.IDnotif, data.UserIDReceiver)
				CheckErr(err, "NotifHandler !AddOrDelete IDEvent  db Exec")
			}
		}

		// si accepter ou refuser, supression de la notifcation de la table NOTIFICATIONS
		stmt, err := db.Prepare("DELETE FROM NOTIFICATIONS WHERE ? = ?")
		CheckErr(err, "NotifHandler DELETE NOTIFICATIONS Prepare db")
		_, err = stmt.Exec(data.TypeID, data.IDnotif)
		CheckErr(err, "NotifHandler DELETE NOTIFICATIONS db Exec")

	}

	if err := json.NewEncoder(w).Encode(jsonResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
