package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	data "forum/internal"

	"github.com/gorilla/websocket"
)

type ClientInfo struct {
	Conn     *websocket.Conn
	Username string
}

var ClientsInfo []ClientInfo

var (
	broadcast = make(chan data.ChatMessage)
	upgrader  = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	clientInfo := ClientInfo{
		Conn: conn,
	}

	// fmt.Println("New Client joined the hub!")
	ClientsInfo = append(ClientsInfo, clientInfo)

	for {
		var msg data.ChatMessage
		err := conn.ReadJSON(&msg)
		usernameClientByConn(msg.Username, conn)
		msg.ListUserConnected = listUserConnected()
		msg.ListUser, msg.ListUserNewMsg = FullListUser(msg.Username)

		if err != nil {
			fmt.Println("client leave ", err, ClientsInfo)
			for _, client := range ClientsInfo {
				client.Conn.WriteJSON(msg)
			}
			deleteClientByConn(conn)
			return
		}

		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast
		// fmt.Println("msg : ", msg)
		if !msg.TypingProgress {
			writeMsgDB(msg.Username, msg.UserPrivateMessage, msg.Horaire, msg.Message)
		}
		msg.ListUserConnected = listUserConnected()
		updateOnly := msg.MajlistConnected

		for _, client := range ClientsInfo {
			tmpListUser, tmpListUserNewMsg := FullListUser(client.Username)
			// fmt.Println("ListUser : ", tmpListUser)
			// fmt.Println("ListUserNewMsg : ", tmpListUserNewMsg)
			msg.ListUser, msg.ListUserNewMsg = sortLastMsg(tmpListUser, tmpListUserNewMsg)
			// Filtrage du client pour la privatisation des messages
			if (msg.Message == "-->New_User_In_Chat<--" || msg.Message == "-->User_Leave_Chat<--") || updateOnly || msg.TypingProgress {
				msg.MajlistConnected = true
			}
			err := client.Conn.WriteJSON(msg)
			if err != nil {
				fmt.Println("Client leave the hub!", err, ClientsInfo)
				deleteClientByConn(client.Conn)
			}
			msg.MajlistConnected = false
		}
		updateOnly = false
	}
}

func sortLastMsg(listUserSort []string, listUserNewMsg []bool) ([]string, []bool) {
	var sortListUserNewMsg []bool
	for userNewMsg := range listUserNewMsg {
		if listUserNewMsg[userNewMsg] == true {
			sortListUserNewMsg = append(sortListUserNewMsg, listUserNewMsg[userNewMsg])
			// fmt.Println("tabUsers", sortListUserNewMsg)
		}
	}
	return listUserSort, listUserNewMsg
}

func listUserConnected() (listUser []string) {
	for _, client := range ClientsInfo {
		listUser = append(listUser, client.Username)
	}
	return listUser
}

func FullListUser(usernameConnected string) (listUser []string, ListUserNewMsg []bool) {
	database, err := sql.Open("sqlite3", "./internal/forum.db")
	CheckErr(err, "open db in websocket FullListUser")
	defer database.Close()

	var NewUserMsg string
	ListUserNewMsgInDB, errdb := database.Prepare("SELECT NewUserMsg FROM Users WHERE Username = ?")
	CheckErr(errdb, "Requete DB")
	ListUserNewMsgInDB.QueryRow(usernameConnected).Scan(&NewUserMsg)
	ArrayListUserNewMsg := strings.Split(NewUserMsg, ":,:")

	listUserInDB, errdb := database.Prepare("SELECT Username FROM Users ORDER BY Username")
	CheckErr(errdb, "Requete DB")
	rows, err := listUserInDB.Query()
	CheckErr(err, "db query")
	var username string
	for rows.Next() {
		err = rows.Scan(&username)
		CheckErr(err, "db rows.scan websocket FullListUser")
		listUser = append(listUser, username)
		if checkNewMsg(username, ArrayListUserNewMsg) {
			ListUserNewMsg = append(ListUserNewMsg, true)
		} else {
			ListUserNewMsg = append(ListUserNewMsg, false)
		}

	}
	return listUser[1:], ListUserNewMsg[1:] // user_deleted et le user 0
}

func checkNewMsg(user string, ArrayListUserNewMsg []string) bool {
	for _, element := range ArrayListUserNewMsg {
		if element == user {
			return true
		}
	}
	return false
}

// Fonction pour supprimer un client de l'array en utilisant la connexion WebSocket
func deleteClientByConn(conn *websocket.Conn) {
	for i, client := range ClientsInfo {
		if client.Conn == conn {
			client.Conn.Close()
			ClientsInfo = append(ClientsInfo[:i], ClientsInfo[i+1:]...)
		}
	}
}

func usernameClientByConn(username string, conn *websocket.Conn) {
	for i, client := range ClientsInfo {
		if client.Conn == conn {
			ClientsInfo[i].Username = username
			break
		}
	}
}

func id_Users_Chat(sender, receiver string, database *sql.DB, newMessage bool) (ID_sender, ID_receiver int) {
	// récupération des id utilisateur
	user, errdb := database.Prepare(`SELECT ID FROM Users WHERE Username = ?`)
	CheckErr(errdb, "1- Requete DB id_Users_Chat")
	user.QueryRow(sender).Scan(&ID_sender)
	user.QueryRow(receiver).Scan(&ID_receiver)

	// insertion d'un nouveau message dans NewUserMsg
	if newMessage && receiver != "" {
		var NewUserMsg string
		stmt1, errdb := database.Prepare(`SELECT NewUserMsg FROM Users WHERE ID = ?`)
		CheckErr(errdb, "2- Requete DB NewUserMsg id_Users_Chat")
		stmt1.QueryRow(ID_receiver).Scan(&NewUserMsg)

		New_NewUserMsg := NewUserMsg
		if NewUserMsg == "" {
			New_NewUserMsg = sender
		} else {
			if !checkNewMsg(sender, strings.Split(NewUserMsg, ":,:")) {
				New_NewUserMsg += ":,:" + sender
			}
		}

		db, err := database.Prepare("UPDATE Users SET NewUserMsg = ? WHERE ID = ?")
		CheckErr(err, "Prepare db")
		_, err = db.Exec(New_NewUserMsg, ID_receiver)
		CheckErr(err, "db Exec")
	}

	return ID_sender, ID_receiver
}

func writeMsgDB(sender, receiver, timestamp, content string) {
	// ouverture database
	database, err := sql.Open("sqlite3", "./internal/forum.db")
	CheckErr(err, "open db in viewpostshandler")
	defer database.Close()

	ID_sender, ID_receiver := id_Users_Chat(sender, receiver, database, true)
	if ID_sender > 0 && ID_receiver > 0 {
		stmt, err := database.Prepare("INSERT INTO Chat(User_ID_sender,User_ID_receiver,Time_stamp,Msg) Values(?,?,?,?)")
		CheckErr(err, "1- writeMsgDB")
		_, err = stmt.Exec(ID_sender, ID_receiver, timestamp, content)
		CheckErr(err, "2- writeMsgDB")
	}
}

func SearchHistory(r *http.Request, database *sql.DB) (historyMsg data.HistoryMsg) {
	// Récupération de l'id du post à afficher
	var dataJS data.HistoryMsg
	err := json.NewDecoder(r.Body).Decode(&dataJS)
	CheckErr(err, "Erreur de décodage JSON SearchHistory")

	ID_sender, ID_receiver := id_Users_Chat(dataJS.Sender, dataJS.Receiver, database, false)

	// Retire le message "nouveau message" de la db
	var NewUserMsg string
	stmt1, errdb := database.Prepare(`SELECT NewUserMsg FROM Users WHERE ID = ?`)
	CheckErr(errdb, "2- Requete DB NewUserMsg SearchHistory")
	stmt1.QueryRow(ID_sender).Scan(&NewUserMsg)

	ArrayListUserNewMsg := strings.Split(NewUserMsg, ":,:")
	new_NewUserMsg := ""
	for i := 0; i < len(ArrayListUserNewMsg); i++ {
		if ArrayListUserNewMsg[i] != dataJS.Receiver {
			new_NewUserMsg += ArrayListUserNewMsg[i]
			if i != len(ArrayListUserNewMsg)-1 ||
				(i != len(ArrayListUserNewMsg)-2 && ArrayListUserNewMsg[len(ArrayListUserNewMsg)-1] == dataJS.Sender) {
				new_NewUserMsg += ":,:"
			}
		}
	}

	db, err := database.Prepare("UPDATE Users SET NewUserMsg = ? WHERE ID = ?")
	CheckErr(err, "Prepare db")
	_, err = db.Exec(new_NewUserMsg, ID_sender)
	CheckErr(err, "db Exec")

	// cherche l'historique
	var listMsg []*data.ChatMsg
	if ID_sender > 0 && ID_receiver > 0 {
		stmt, errdb := database.Prepare(`
			SELECT User_ID_sender, Time_stamp, Msg
			FROM Chat
			WHERE (User_ID_sender = ? and User_ID_receiver = ?) or (User_ID_sender = ? and User_ID_receiver = ?)
		`)
		CheckErr(errdb, "2- Requete DB SearchHistory")
		rows, err := stmt.Query(ID_sender, ID_receiver, ID_receiver, ID_sender)
		CheckErr(err, "db query")
		if err == nil {
			for rows.Next() {
				var msg data.ChatMsg
				var senderDB int
				rows.Scan(&senderDB, &msg.Timestamp, &msg.Content)
				if senderDB == ID_sender {
					msg.SenderMsg = dataJS.Sender
				} else {
					msg.SenderMsg = dataJS.Receiver
				}
				listMsg = append(listMsg, &msg)
			}
		}
	}
	historyMsg.ListMsg = listMsg
	return historyMsg
}
