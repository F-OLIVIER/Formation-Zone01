package pkg

import (
	"database/sql"
	"fmt"
)

// optionType = 'follow', 'mp', 'post', 'comment', 'groupInvite', 'event'
func InsertNotif(ID, UserID_Followers int, date, optionType string, db *sql.DB) {
	var request string
	switch optionType {
	case "follow":
		request = "INSERT INTO NOTIFICATIONS(IDfollow,UserID_Receiver,Date) Values(?,?,?)"
	case "mp":
		request = "INSERT INTO NOTIFICATIONS(IDPrivateMessage,UserID_Receiver,Date) Values(?,?,?)"
	case "post":
		request = "INSERT INTO NOTIFICATIONS(IDPost,UserID_Receiver,Date) Values(?,?,?)"
	case "comment":
		request = "INSERT INTO NOTIFICATIONS(IDComment,UserID_Receiver,Date) Values(?,?,?)"
	case "groupInvite":
		request = "INSERT INTO NOTIFICATIONS(IDGroup,UserID_Receiver,Date) Values(?,?,?)"
	case "event":
		request = "INSERT INTO NOTIFICATIONS(IDEvent,UserID_Receiver,Date) Values(?,?,?)"
	case "groupmsg":
		request = "INSERT INTO NOTIFICATIONS(IDPrivateGroupMessage,UserID_Receiver,Date) Values(?,?,?)"
	default:
		fmt.Println("\nY A DE LA MERDE ICI : vérifie t'on optionType à la fonction InsertNotif\n Option correct : 'follow', 'mp', 'post', 'comment', 'group', 'event'")
		return
	}

	stmt, err := db.Prepare(request)
	CheckErr(err, "InsertNotif Prepare db")
	_, err = stmt.Exec(ID, UserID_Followers, date)
	CheckErr(err, "InsertNotif db Exec")
}

// optionType = 'follow', 'mp', 'post', 'comment', 'group', 'event'
func DeleteNotif(ID int, optionType string, db *sql.DB) {
	var request string
	switch optionType {
	case "follow":
		request = "DELETE FROM NOTIFICATIONS WHERE IDFollow = ?"
	case "mp":
		request = "DELETE FROM NOTIFICATIONS WHERE IDPrivateMessage = ?"
	case "post":
		request = "DELETE FROM NOTIFICATIONS WHERE IDPost = ?"
	case "comment":
		request = "DELETE FROM NOTIFICATIONS WHERE IDComment = ?"
	case "group":
		request = "DELETE FROM NOTIFICATIONS WHERE IDGroup = ?"
	case "inviteGroup":
		request = "DELETE FROM NOTIFICATIONS WHERE IDNotif = ?"
	case "event":
		request = "DELETE FROM NOTIFICATIONS WHERE IDEvent = ?"
	case "groupmsg":
		request = "DELETE FROM NOTIFICATIONS WHERE IDPrivateGroupMessage = ?"
	case "askgroup":
		request = "DELETE FROM NOTIFICATIONS WHERE IDAsking = ?"
	default:
		fmt.Println("\nY A DE LA MERDE ICI : vérifie t'on optionType à la fonction DeleteNotif\n Option correct : 'follow', 'mp', 'post', 'comment', 'group', 'event'")
		return
	}

	stmt, err := db.Prepare(request)
	CheckErr(err, "DeleteNotif Prepare db")
	_, err = stmt.Exec(ID)
	CheckErr(err, "DeleteNotif db Exec")
}

func GetNotif(UserID int, db *sql.DB) (ListFollowers, listMP, listPost, listComment, listGroup, listEvent, listGroupMsg, listAskGroup [][]interface{}) {
	// récupération des notifications follow
	ListFollowers = make([][]interface{}, 0)
	stmtfollow, err := db.Prepare(`SELECT USERS.ID, USERS.FirstName, USERS.LastName, USERS.Avatar, USERS.Nickname, FOLLOWERS.ID, FOLLOWERS.ValidateFollow, FOLLOWERS.DateFollow
										FROM USERS
										INNER JOIN FOLLOWERS ON FOLLOWERS.UserID_Follower = USERS.ID
										INNER JOIN NOTIFICATIONS ON NOTIFICATIONS.IDFollow = FOLLOWERS.ID
										WHERE NOTIFICATIONS.UserID_Receiver = ?;`)
	CheckErr(err, "GetNotif ListFollowers, db prepare")
	rowsfollow, err := stmtfollow.Query(UserID)
	CheckErr(err, "GetNotif ListFollowers, db query")
	for rowsfollow.Next() {
		var currentUser User
		var currentFollow Followers
		var currentInformation []interface{}
		err = rowsfollow.Scan(&currentUser.Id, &currentUser.Firstname, &currentUser.Lastname, &currentUser.Avatar, &currentUser.Nickname, &currentFollow.Id, &currentFollow.ValidateFollow, &currentFollow.DateFollow)
		CheckErr(err, "GetNotif ListFollowers, db rowsfollow.Next scan")
		var category = "Follow"
		currentUser.Category = category
		currentInformation = append(currentInformation, currentUser)
		currentInformation = append(currentInformation, currentFollow)
		ListFollowers = append(ListFollowers, currentInformation) // information sur l'utilisateur
	}

	// Récupération des notifications MP
	listMP = make([][]interface{}, 0) // Initialisation de la tranche vide
	stmtMp, err := db.Prepare(`SELECT USERS.ID, USERS.FirstName, USERS.LastName, USERS.Avatar, USERS.Nickname, messages.id, messages.content, messages.date
									FROM USERS
									INNER JOIN messages ON messages.sender_id = USERS.ID
									INNER JOIN NOTIFICATIONS ON NOTIFICATIONS.IDPrivateMessage = messages.id
									WHERE NOTIFICATIONS.UserID_Receiver = ?`)
	CheckErr(err, "GetNotif listMP, db prepare")
	rowsmp, err := stmtMp.Query(UserID)
	CheckErr(err, "GetNotif listMP, db query")
	for rowsmp.Next() {
		var currentUser User
		var currentMessage Message
		var currentInformation []interface{}
		err = rowsmp.Scan(&currentUser.Id, &currentUser.Firstname, &currentUser.Lastname, &currentUser.Avatar, &currentUser.Nickname, &currentMessage.Id, &currentMessage.Content, &currentMessage.Date)
		CheckErr(err, "GetNotif listMP, db rowsmp.Next scan")
		var category = "MP"
		currentUser.Category = category
		currentInformation = append(currentInformation, currentUser)
		currentInformation = append(currentInformation, currentMessage)
		listMP = append(listMP, currentInformation) // information sur l'utilisateur
	}

	// récupération des notifications commentaire
	listComment = make([][]interface{}, 0) // Initialisation de la tranche vide

	stmtComment, err := db.Prepare(`SELECT DISTINCT USERS.ID, USERS.FirstName, USERS.LastName, USERS.Avatar, USERS.Nickname,
	                                        COMMENT.IDPost, COMMENT.CommentContent, COMMENT.Date, COMMENT.Image
	                                        FROM USERS
	                                        INNER JOIN COMMENT ON COMMENT.UserID = USERS.ID 
											INNER JOIN Post ON COMMENT.IDPost = POST.IDPost
	                                        INNER JOIN NOTIFICATIONS ON NOTIFICATIONS.IDComment = COMMENT.IDComment
	                                        WHERE NOTIFICATIONS.UserID_Receiver = ?`)
	CheckErr(err, "GetNotif listComment, db prepare")

	rowsComment, err := stmtComment.Query(UserID)
	CheckErr(err, "GetNotif listComment, db query")

	for rowsComment.Next() {
		var currentUser User
		var currentComment Comment
		var currentInformation []interface{}

		err = rowsComment.Scan(&currentUser.Id, &currentUser.Firstname, &currentUser.Lastname, &currentUser.Avatar, &currentUser.Nickname, &currentComment.Post_id, &currentComment.Content, &currentComment.Date, &currentComment.Image)
		CheckErr(err, "listComment, db rowsComment.Next scan")

		// Maintenant, pour chaque commentaire, nous avons les informations sur l'utilisateur et le commentaire lui-même.
		// Vous pouvez ensuite utiliser currentUser et currentComment pour construire vos données de notification.

		var category = "Comment"
		currentUser.Category = category
		currentInformation = append(currentInformation, currentUser)    // Informations sur l'utilisateur
		currentInformation = append(currentInformation, currentComment) // Informations sur le commentaire

		listComment = append(listComment, currentInformation) // Ajouter à la liste des commentaires
	}

	// récupération des notifications post
	listPost = make([][]interface{}, 0)
	stmtpost, err := db.Prepare(`SELECT USERS.FirstName, USERS.LastName, USERS.Avatar, USERS.Nickname,
										POST.IDPost, POST.UserID, POST.Title, POST.PostContent, POST.Date, POST.Image, POST.Private, POST.Likes, POST.NbComments
										FROM USERS
										INNER JOIN POST ON POST.UserID = USERS.ID
										INNER JOIN NOTIFICATIONS ON NOTIFICATIONS.IDPost = POST.IDPost
										WHERE NOTIFICATIONS.UserID_Receiver = ? AND NOTIFICATIONS.IDPost = POST.IDPost;`)
	CheckErr(err, "GetNotif listPost, db prepare")
	rowspost, err := stmtpost.Query(UserID)
	CheckErr(err, "GetNotif listPost, db query")
	for rowspost.Next() {
		var currentUser User
		var currentPost Post
		var currentInformation []interface{}
		err = rowspost.Scan(&currentUser.Firstname, &currentUser.Lastname, &currentUser.Avatar, &currentUser.Nickname, &currentPost.Id, &currentPost.User_id, &currentPost.Title, &currentPost.Content, &currentPost.Date, &currentPost.Image, &currentPost.Private, &currentPost.Likes, &currentPost.NbComments)
		CheckErr(err, "listPost, db rowspost.Next scan")
		currentInformation = append(currentInformation, currentUser) // information sur l'utilisateur
		currentInformation = append(currentInformation, currentPost) // information sur le post
		var category = "Post"
		currentUser.Category = category
		listPost = append(listPost, currentInformation)
	}
	// Récupération des notifications de groupe
	listGroup = make([][]interface{}, 0)
	stmtGroup, err := db.Prepare(`SELECT USERS.FirstName, USERS.LastName, USERS.Avatar, USERS.Nickname,
										LISTGROUPS.IDGroup, LISTGROUPS.NameGroup, LISTGROUPS.AboutUs, LISTGROUPS.Image
										FROM USERS
										INNER JOIN LISTGROUPS ON LISTGROUPS.UserID_Creator = USERS.ID
										INNER JOIN NOTIFICATIONS ON NOTIFICATIONS.IDGroup = LISTGROUPS.IDGroup
										WHERE NOTIFICATIONS.UserID_Receiver = ? AND NOTIFICATIONS.IDGroup = LISTGROUPS.IDGroup;`)
	CheckErr(err, "GetNotif listGroup, db prepare")
	rowsGroup, err := stmtGroup.Query(UserID)
	CheckErr(err, "GetNotif listGroup, db query")
	for rowsGroup.Next() {
		var currentUser User
		var currentGroup Group
		var currentInformation []interface{}
		err = rowsGroup.Scan(&currentUser.Firstname, &currentUser.Lastname, &currentUser.Avatar, &currentUser.Nickname, &currentGroup.IdGroup, &currentGroup.Title, &currentGroup.AboutGroup, &currentGroup.Image)
		CheckErr(err, "listComment, db rowsComment.Next scan")
		currentInformation = append(currentInformation, currentUser)  // information sur l'utilisateur
		currentInformation = append(currentInformation, currentGroup) // information sur le commentaire
		var category = "Group"
		currentUser.Category = category
		currentInformation = append(currentInformation, currentUser)
		listGroup = append(listGroup, currentInformation)
	}

	// Récupération des notifications de messages de groupe
	listGroupMsg = make([][]interface{}, 0)
	stmtGroupMsg, err := db.Prepare(`SELECT USERS.ID, USERS.FirstName, USERS.LastName, USERS.Avatar, USERS.Nickname, groupmessages.id, groupmessages.content, groupmessages.date, LISTGROUPS.IDGroup, LISTGROUPS.NameGroup, LISTGROUPS.AboutUs, LISTGROUPS.Image
									FROM USERS
									INNER JOIN groupmessages ON groupmessages.sender_id = USERS.ID
									INNER JOIN NOTIFICATIONS ON NOTIFICATIONS.IDPrivateGroupMessage = groupmessages.id
									INNER JOIN LISTGROUPS ON LISTGROUPS.IDGroup = groupmessages.group_id
									WHERE NOTIFICATIONS.UserID_Receiver = ?`)
	CheckErr(err, "GetNotif listGroupMsg, db prepare")
	rowsGroupMsg, err := stmtGroupMsg.Query(UserID)
	CheckErr(err, "GetNotif listGroupMsg, db query")
	for rowsGroupMsg.Next() {
		var currentUser User
		var currentGroupMessage struct {
			ID      int
			Content string
			Date    string
		}
		var currentGroup Group
		var currentInformation []interface{}
		err = rowsGroupMsg.Scan(&currentUser.Id, &currentUser.Firstname, &currentUser.Lastname, &currentUser.Avatar, &currentUser.Nickname, &currentGroupMessage.ID, &currentGroupMessage.Content, &currentGroupMessage.Date, &currentGroup.IdGroup, &currentGroup.Title, &currentGroup.AboutGroup, &currentGroup.Image)
		CheckErr(err, "listGroupMsg, db rowsGroupMsg.Next scan")
		currentInformation = append(currentInformation, currentUser)         // information sur l'utilisateur
		currentInformation = append(currentInformation, currentGroupMessage) // information sur le message de groupe
		currentInformation = append(currentInformation, currentGroup)
		var category = "GroupMessage"
		currentUser.Category = category
		listGroupMsg = append(listGroupMsg, currentInformation)
	}

	listEvent = make([][]interface{}, 0)
	stmtEvent, err := db.Prepare(`SELECT EVENTGROUPS.IDEvent, EVENTGROUPS.Title, EVENTGROUPS.Date,
										USERS.ID, USERS.FirstName, USERS.LastName, USERS.Avatar, USERS.Nickname,
										LISTGROUPS.IDGroup, LISTGROUPS.NameGroup, LISTGROUPS.AboutUs, LISTGROUPS.Image
										FROM EVENTGROUPS
										INNER JOIN USERS ON EVENTGROUPS.UserID_Sender = USERS.ID
										INNER JOIN LISTGROUPS ON EVENTGROUPS.IDGroup = LISTGROUPS.IDGroup
										INNER JOIN NOTIFICATIONS ON NOTIFICATIONS.IDEvent = EVENTGROUPS.IDEvent
										WHERE NOTIFICATIONS.UserID_Receiver = ?`)
	CheckErr(err, "GetNotif listEvent, db prepare")
	rowsEvent, err := stmtEvent.Query(UserID)
	CheckErr(err, "GetNotif listEvent, db query")
	for rowsEvent.Next() {
		var currentEvent Event
		var currentUser User
		var currentGroup Group
		var currentInformation []interface{}
		err = rowsEvent.Scan(&currentEvent.IdEvent, &currentEvent.Title, &currentEvent.Date,
			&currentUser.Id, &currentUser.Firstname, &currentUser.Lastname, &currentUser.Avatar, &currentUser.Nickname,
			&currentGroup.IdGroup, &currentGroup.Title, &currentGroup.AboutGroup, &currentGroup.Image)
		CheckErr(err, "listEvent, db rowsEvent.Next scan")
		currentInformation = append(currentInformation, currentUser)  // information sur l'utilisateur
		currentInformation = append(currentInformation, currentGroup) // information sur le groupe
		currentInformation = append(currentInformation, currentEvent) // information sur l'event
		var category = "Event"
		currentUser.Category = category
		listEvent = append(listEvent, currentInformation)
	}

	// Récupération des notifications des demandes de groupe
	listAskGroup = make([][]interface{}, 0)
	stmtAskGroup, err := db.Prepare(`SELECT USERS_SENDER.ID, USERS_SENDER.FirstName, USERS_SENDER.LastName, USERS_SENDER.Avatar, USERS_SENDER.Nickname,
										LISTGROUPS.IDGroup, LISTGROUPS.NameGroup, LISTGROUPS.AboutUs, LISTGROUPS.Image
										FROM USERS AS USERS_SENDER
										INNER JOIN NOTIFICATIONS ON NOTIFICATIONS.UserID_sender = USERS_SENDER.ID
										INNER JOIN LISTGROUPS ON LISTGROUPS.IDGroup = NOTIFICATIONS.IDAsking
										WHERE NOTIFICATIONS.UserID_Receiver = ?`)
	CheckErr(err, "GetNotif listAskGroup, db prepare")
	rowsAskGroup, err := stmtAskGroup.Query(UserID)
	CheckErr(err, "GetNotif listAskGroup, db query")
	for rowsAskGroup.Next() {
		var currentUser User
		var currentGroup Group
		var currentInformation []interface{}
		err = rowsAskGroup.Scan(&currentUser.Id, &currentUser.Firstname, &currentUser.Lastname, &currentUser.Avatar, &currentUser.Nickname, &currentGroup.IdGroup, &currentGroup.Title, &currentGroup.AboutGroup, &currentGroup.Image)
		CheckErr(err, "listAskGroup, db rowsAskGroup.Next scan")
		currentInformation = append(currentInformation, currentUser)  // information sur l'utilisateur
		currentInformation = append(currentInformation, currentGroup) // information sur le groupe
		var category = "AskGroup"
		currentUser.Category = category
		listAskGroup = append(listAskGroup, currentInformation)
	}

	return ListFollowers, listMP, listPost, listComment, listGroup, listEvent, listGroupMsg, listAskGroup
}

// optionType = 'post', 'comment', 'group', 'event'
// optionID = mettre l'id du comment ou du groupe ou de l'event
func WhoDisplayNotif(UserID, optionID int, optionType string, db *sql.DB) (listUser []User) {

	switch optionType {
	case "post": // Uniquement si le post est public pour les autres option l'utilisateur choisie lors de la création du post
		listUser = ListFollow(UserID, "followers", 0)
	case "comment":
		// Récupére la privacité du post associé au commentaire 'optionID'
		stmtPrivacitePost, err := db.Prepare(`SELECT POST.Private FROM POST
												INNER JOIN COMMENT ON COMMENT.IDPost = POST.IDPost
												WHERE COMMENT.IDComment = ?;`)
		CheckErr(err, "WhoDisplayNotif stmtPrivacitePost comment, db prepare")
		var private int
		stmtPrivacitePost.QueryRow(optionID).Scan(&private)

		// Si post public
		if private == 1 { // Retourne la liste des followers de l'utilisateur
			listUser = ListFollow(UserID, "followers", 0)
		}
		// Si post semi-privée
		if private == 2 { // retourne la liste des utilisateurs correspondant dans la table 'POSTSELECTUSERS'
			stmtComment, err := db.Prepare(`SELECT USERS.ID, USERS.FirstName, USERS.LastName, USERS.Avatar, USERS.Nickname FROM USERS
												INNER JOIN POSTSELECTUSERS ON POSTSELECTUSERS.UserID = USERS.ID
												WHERE POSTSELECTUSERS.IDPost = ?;`)
			CheckErr(err, "WhoDisplayNotif comment, db prepare")
			rowspost, err := stmtComment.Query(optionID)
			CheckErr(err, "WhoDisplayNotif comment, db query")
			for rowspost.Next() {
				var currentUser User
				err = rowspost.Scan(&currentUser.Id, &currentUser.Firstname, &currentUser.Lastname, &currentUser.Avatar, &currentUser.Nickname)
				CheckErr(err, "WhoDisplayNotif comment, db rowspost.Next scan")
				listUser = append(listUser, currentUser)
			}
		}
		// Si privée, Ne rien faire et retourner listUser vide

	case "group": // Retourne la liste des membres validé du groupe 'optionID'
		stmtGroup, err := db.Prepare(`SELECT USERS.ID, USERS.FirstName, USERS.LastName, USERS.Avatar, USERS.Nickname FROM USERS
										INNER JOIN MEMBERSGROUPS ON MEMBERSGROUPS.UserID = USERS.ID
										WHERE MEMBERSGROUPS.IDgroup = ? AND MEMBERSGROUPS.ValidationInvite = true;`)
		CheckErr(err, "WhoDisplayNotif group, db prepare")
		rowspost, err := stmtGroup.Query(optionID)
		CheckErr(err, "WhoDisplayNotif group, db query")
		for rowspost.Next() {
			var currentUser User
			err = rowspost.Scan(&currentUser.Id, &currentUser.Firstname, &currentUser.Lastname, &currentUser.Avatar, &currentUser.Nickname)
			CheckErr(err, "WhoDisplayNotif group, db rowspost.Next scan")
			listUser = append(listUser, currentUser)
		}

	case "event": // Retourne la liste des membres qui ont validé l'event 'optionID'
		stmtEvent, err := db.Prepare(`SELECT USERS.ID, USERS.FirstName, USERS.LastName, USERS.Avatar, USERS.Nickname FROM USERS
										INNER JOIN RESPONSEEVENTGROUPS ON RESPONSEEVENTGROUPS.UserID = USERS.ID
										WHERE RESPONSEEVENTGROUPS.IDEvent = ? AND RESPONSEEVENTGROUPS.Option = 1;`)
		CheckErr(err, "WhoDisplayNotif event, db prepare")
		rowspost, err := stmtEvent.Query(optionID)
		CheckErr(err, "WhoDisplayNotif event, db query")
		for rowspost.Next() {
			var currentUser User
			err = rowspost.Scan(&currentUser.Id, &currentUser.Firstname, &currentUser.Lastname, &currentUser.Avatar, &currentUser.Nickname)
			CheckErr(err, "WhoDisplayNotif event, db rowspost.Next scan")
			listUser = append(listUser, currentUser)
		}

	case "groupmessage": // Retourne la liste des membres du groupe 'optionID' sauf le créateur
		stmtGroupMessage, err := db.Prepare(`SELECT USERS.ID, USERS.FirstName, USERS.LastName, USERS.Avatar, USERS.Nickname
											 FROM USERS
											 INNER JOIN MEMBERSGROUPS ON MEMBERSGROUPS.UserID = USERS.ID
											 INNER JOIN LISTGROUPS ON LISTGROUPS.IDGroup = MEMBERSGROUPS.IDGroup
											 WHERE MEMBERSGROUPS.IDGroup = ? AND MEMBERSGROUPS.UserID != LISTGROUPS.UserID_Creator;`)
		CheckErr(err, "WhoDisplayNotif groupmessage, db prepare")
		rowspost, err := stmtGroupMessage.Query(optionID)
		CheckErr(err, "WhoDisplayNotif groupmessage, db query")
		for rowspost.Next() {
			var currentUser User
			err = rowspost.Scan(&currentUser.Id, &currentUser.Firstname, &currentUser.Lastname, &currentUser.Avatar, &currentUser.Nickname)
			CheckErr(err, "WhoDisplayNotif groupmessage, db rowspost.Next scan")
			listUser = append(listUser, currentUser)
		}
	}

	return listUser
}
