package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type GlobalData struct {
	UserData      User                  `json:"userData"`
	Message       []Message             `json:"message"`
	UsersMessages map[int]UsersMessages `json:"usersmessages"`
	Listuser      []User                `json:"listuser"`
	Online        []int                 `json:"online"`
}

type UsersMessages struct {
	User     User
	Messages []Message
}

var ListOnline []int

func CrossPlatformHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("\n--------------- enter crossplatform ---------------")
	// Ouverture de la database
	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	CheckErr(err, "Erreur lors de l'ouverture de la base de données:")
	defer db.Close()

	// Déclaration des informations de base à renvyer à l'application
	var globalData GlobalData

	if r.Method == "POST" {
		switch r.URL.Path {
		case "/crossplatform/login":
			globalData.UserData = crosslogin(w, r, db)
			break

		case "/crossplatform/getallmessage":
			var jsondata Message
			errjson := json.NewDecoder(r.Body).Decode(&jsondata)
			if errjson != nil {
				http.Error(w, "500 internal server error: Failed to connect to database. "+errjson.Error(), http.StatusInternalServerError)
			}
			// fmt.Println("jsondata getallmessage : ", jsondata)
			globalData.UserData = crossGetuser(jsondata.Id, jsondata.Session_uuid, db)
			globalData.UsersMessages = crossGetAllMessage(jsondata, db)
			break

		case "/crossplatform/getuser":
			var jsondata User
			errjson := json.NewDecoder(r.Body).Decode(&jsondata)
			if errjson != nil {
				http.Error(w, "500 internal server error: Failed to connect to database. "+errjson.Error(), http.StatusInternalServerError)
			}
			// fmt.Println("jsondata getuser : ", jsondata)
			globalData.UserData = crossGetuser(jsondata.Id, jsondata.Session_uuid, db)
			break

		case "/crossplatform/listuser":
			var jsondata User
			errjson := json.NewDecoder(r.Body).Decode(&jsondata)
			if errjson != nil {
				http.Error(w, "500 internal server error: Failed to connect to database. "+errjson.Error(), http.StatusInternalServerError)
			}
			// fmt.Println("jsondata listuser : ", jsondata)
			globalData.Listuser = crossListUsers(jsondata.Id, db)
			globalData.Online = ListOnline

			break

		}
	}

	// Envoi de la réponse de serveur à l'application
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(globalData)
}

func crosslogin(w http.ResponseWriter, r *http.Request, db *sql.DB) User {
	var logindata User
	errjson := json.NewDecoder(r.Body).Decode(&logindata)
	if errjson != nil {
		http.Error(w, "500 internal server error: Failed to connect to database. "+errjson.Error(), http.StatusInternalServerError)
	}

	// vérification et récupération des informations de l'utilisateur
	stmt, errdb := db.Prepare("SELECT ID, Email, Password, FirstName, LastName, Nickname FROM USERS WHERE Nickname = ? OR Email = ?")
	CheckErr(errdb, "userInfo db prepare")
	var userinfo User
	stmt.QueryRow(logindata.Email, logindata.Email).Scan(&userinfo.Id, &userinfo.Email, &userinfo.Password, &userinfo.Firstname, &userinfo.Lastname, &userinfo.Nickname)

	err := bcrypt.CompareHashAndPassword([]byte(userinfo.Password), []byte(logindata.Password))

	var baduser User
	if err != nil {
		return baduser
	}
	// Par sécurité, supression du mot de passe blanc
	userinfo.Password = ""

	// Delete de la session-cookie dans la db si existant (session identique à celle du site internet).
	_, err = db.Exec(`DELETE FROM SESSIONS WHERE UserID = ?`, userinfo.Id)
	// _, err = db.Exec(`DELETE FROM SESSIONS WHERE UserID = ? WHERE app = 1`, userinfo.Id)
	if err != nil {
		fmt.Println("error cookie supp")
		http.Error(w, "500 internal server error.", http.StatusInternalServerError)
		return baduser
	}

	// Création de la nouvelle session pour l'apllication
	sessionId, err := uuid.NewV4()
	_, err = db.Exec(`INSERT INTO SESSIONS (SessionToken, UserID) values(?, ?)`, sessionId, userinfo.Id)
	// _, err = db.Exec(`INSERT INTO SESSIONS (SessionToken, UserID, app) values(?, ?, ?)`, sessionId, userinfo.Id, 1)
	if err != nil {
		fmt.Println("error DB cookie")
		http.Error(w, "500 internal server error.", http.StatusInternalServerError)
		return baduser
	}

	userinfo.Session_uuid = sessionId.String()

	return userinfo
}

// Récupération des informations de l'utilisateur demandé
func crossGetuser(userID int, uuid string, db *sql.DB) (user User) {
	user.Session_uuid = uuid
	stmt, errdb := db.Prepare("SELECT ID, FirstName, LastName, Nickname FROM USERS WHERE id = ?")
	CheckErr(errdb, "userInfo db prepare")
	stmt.QueryRow(userID).Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Nickname)
	return user
}

// Récupération de la liste des messages en fonction d'une liste utilisateur autorisé à communiqué avec l'utilisateur (si profils public ou si follower ou si following)
func crossGetAllMessage(msg Message, db *sql.DB) map[int]UsersMessages {
	sender := strconv.Itoa(msg.Id)

	listuser := crossListUsers(msg.Id, db)

	// Récupération des messages correspondant à l'utilisateur
	stmt2, errdb := db.Prepare(`SELECT Sender_id, Receiver_id, Date, Content FROM messages WHERE sender_id = ? OR Receiver_id = ? ORDER BY id DESC`)
	CheckErr(errdb, "crossGetAllMessage db prepare")
	rows2, err := stmt2.Query(sender, sender)
	CheckErr(err, "crossGetAllMessage, db query")
	allMessages := make(map[int]UsersMessages)
	for rows2.Next() {
		var currentmessage Message
		err = rows2.Scan(&currentmessage.Sender_id, &currentmessage.Receiver_id, &currentmessage.Date, &currentmessage.Content)
		CheckErr(err, "crossGetAllMessage, db rows.Next scan")

		if msg.Id == currentmessage.Receiver_id { // cas 1 : l'utilisateur a reçu un message
			userMessages, found := allMessages[currentmessage.Sender_id]
			if found { // Ajoutez le message au slice existant
				userMessages.Messages = append(userMessages.Messages, currentmessage)
				allMessages[currentmessage.Sender_id] = userMessages
			} else { // Créez un nouveau slice et ajoutez-le à la map
				var UserMessage UsersMessages
				UserMessage.Messages = append(UserMessage.Messages, currentmessage)

				for i := 0; i < len(listuser); i++ {
					if listuser[i].Id == currentmessage.Sender_id {
						UserMessage.User = listuser[i]
						break
					}
				}
				allMessages[currentmessage.Sender_id] = UserMessage
			}

		} else { // cas 2 : l'utilisateur a envoyé un message
			userMessages, found := allMessages[currentmessage.Receiver_id]
			if found { // Ajoutez le message au slice existant
				userMessages.Messages = append(userMessages.Messages, currentmessage)
				allMessages[currentmessage.Receiver_id] = userMessages
			} else { // Créez un nouveau slice et ajoutez-le à la map
				var UserMessage UsersMessages
				UserMessage.Messages = append(UserMessage.Messages, currentmessage)

				for i := 0; i < len(listuser); i++ {
					if listuser[i].Id == currentmessage.Receiver_id {
						UserMessage.User = listuser[i]
						break
					}
				}
				allMessages[currentmessage.Receiver_id] = UserMessage
			}
		}
	}

	// ajout des utilisateurs sans message absent de la map
	for _, value := range listuser {
		currentData := allMessages[value.Id]
		if allMessages[value.Id].User.Id == 0 {
			currentData.User = value
			allMessages[value.Id] = currentData
		}
	}

	// Printer la map allMessages
	// for key, value := range allMessages {
	// 	fmt.Println("_______________________________________________________")
	// 	fmt.Printf("Clé: %v\nValeur User: %v\nValeur Messages: %v\n\n", key, value.User, value.Messages)
	// 	fmt.Println("_______________________________________________________")
	// }

	return allMessages
}

// Récupération de la liste complétes des utilisateurs present dans la db
func listAllUsers(db *sql.DB) (listuser []User) {
	stmt1, errdb := db.Prepare(`SELECT ID, FirstName, LastName, Nickname, PrivateProfile FROM USERS`)
	CheckErr(errdb, "completeListUsers db prepare")
	rows1, err := stmt1.Query()
	CheckErr(err, "completeListUsers, db query")
	// var listuser []User
	for rows1.Next() {
		var currentUser User
		err = rows1.Scan(&currentUser.Id, &currentUser.Firstname, &currentUser.Lastname, &currentUser.Nickname, &currentUser.PrivateProfile)
		CheckErr(err, "completeListUsers, db rows.Next scan")
		listuser = append(listuser, currentUser)
	}
	return listuser
}

// Récupération de la liste des utilisateurs en fonction d'un utilisateur préci (si profils public ou si follower ou si following)
func crossListUsers(User_ID int, db *sql.DB) (listUser []User) {
	// récupération de la liste compléte des utilisateurs
	completeUserList := listAllUsers(db)
	// récupération des liste followers followings de l'utilisateur
	followers := ListFollow(User_ID, "followers", 0)
	followings := ListFollow(User_ID, "followings", 0)

	for i := 0; i < len(completeUserList); i++ {
		currentUser := completeUserList[i]
		if currentUser.Id != User_ID { // si c'est l'utilisateur qui fais la requete
			if currentUser.PrivateProfile == "0" { // si profils privée
				if inUserList(currentUser.Id, followers) || inUserList(currentUser.Id, followings) { // si dans la liste des followers ou des following
					listUser = append(listUser, currentUser)
				}
			} else { // si profils public
				listUser = append(listUser, currentUser)
			}
		}
	}

	return listUser
}

func inUserList(idsearch int, userList []User) bool {
	for i := 0; i < len(userList); i++ {
		if userList[i].Id == idsearch {
			return true
		}
	}
	return false
}

func ContainsID(id int, slice []int) bool {
	for _, item := range slice {
		if item == id {
			return true
		}
	}
	return false
}

func RemoveSliceInt(slice []int, value int) []int {
	index := -1
	for i, v := range slice {
		if v == value {
			index = i
			break
		}
	}
	if index == -1 {
		// Si l'élément n'est pas trouvé, retourner le slice original
		// fmt.Println("ERROR: no index in online users")
		return slice
	}
	// Créer un nouveau slice sans l'élément
	return append(slice[:index], slice[index+1:]...)
}
