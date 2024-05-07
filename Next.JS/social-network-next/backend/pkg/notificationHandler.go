package pkg

import (
	"database/sql"
	"fmt"
)

// optionType = 'follow', 'mp', 'post', 'comment', 'group', 'event'
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
	case "group":
		request = "INSERT INTO NOTIFICATIONS(IDGroup,UserID_Receiver,Date) Values(?,?,?)"
	case "event":
		request = "INSERT INTO NOTIFICATIONS(IDEvent,UserID_Receiver,Date) Values(?,?,?)"
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
	case "event":
		request = "DELETE FROM NOTIFICATIONS WHERE IDEvent = ?"
	default:
		fmt.Println("\nY A DE LA MERDE ICI : vérifie t'on optionType à la fonction DeleteNotif\n Option correct : 'follow', 'mp', 'post', 'comment', 'group', 'event'")
		return
	}

	stmt, err := db.Prepare(request)
	CheckErr(err, "DeleteNotif Prepare db")
	_, err = stmt.Exec(ID)
	CheckErr(err, "DeleteNotif db Exec")
}

func GetNotif(UserID int, db *sql.DB) (ListFollowers []User, listMP, listPost, listComment, listGroup, listEvent [][]interface{}) {
	// récupération des notifications follow
	stmtfollow, err := db.Prepare(`SELECT USERS.ID, USERS.FirstName, USERS.LastName, USERS.Avatar, USERS.Nickname
										FROM USERS
										INNER JOIN FOLLOWERS ON FOLLOWERS.UserID_Follower = USERS.ID
										INNER JOIN NOTIFICATIONS ON NOTIFICATIONS.IDFollow = FOLLOWERS.ID
										WHERE FOLLOWERS.ValidateFollow = '1' AND NOTIFICATIONS.UserID_Receiver = ?;`)
	CheckErr(err, "GetNotif ListFollowers, db prepare")
	rowsfollow, err := stmtfollow.Query(UserID)
	CheckErr(err, "GetNotif ListFollowers, db query")
	for rowsfollow.Next() {
		var currentUser User
		err = rowsfollow.Scan(&currentUser.Id, &currentUser.Firstname, &currentUser.Lastname, &currentUser.Avatar, &currentUser.Nickname)
		CheckErr(err, "GetNotif ListFollowers, db rowsfollow.Next scan")
		ListFollowers = append(ListFollowers, currentUser) // information sur l'utilisateur
	}

	// récupération des notifications message privée

	// récupération des notifications post
	listPost = make([][]interface{}, 2)
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
		listPost = append(listPost, currentInformation)
	}

	// récupération des notifications commentaire
	listComment = make([][]interface{}, 3)
	stmtComment, err := db.Prepare(`SELECT USERS.FirstName, USERS.LastName, USERS.Avatar, USERS.Nickname,
										COMMENT.IDPost, COMMENT.CommentContent, COMMENT.Date, COMMENT.Image
										FROM USERS
										INNER JOIN COMMENT ON COMMENT.UserID = USERS.ID
										INNER JOIN Post ON COMMENT.IDPost = POST.IDPost
										INNER JOIN NOTIFICATIONS ON NOTIFICATIONS.IDComment = COMMENT.IDComment
										WHERE NOTIFICATIONS.UserID_Receiver = ? AND NOTIFICATIONS.IDComment = COMMENT.IDComment;`)
	CheckErr(err, "GetNotif listComment, db prepare")
	rowsComment, err := stmtComment.Query(UserID)
	CheckErr(err, "GetNotif listComment, db query")
	for rowsComment.Next() {
		var currentUser User
		var currentComment Comment
		var currentInformation []interface{}
		err = rowsComment.Scan(&currentUser.Firstname, &currentUser.Lastname, &currentUser.Avatar, &currentUser.Nickname, &currentComment.Post_id, &currentComment.Content, &currentComment.Date, &currentComment.Image)
		CheckErr(err, "listComment, db rowsComment.Next scan")
		currentInformation = append(currentInformation, currentUser)    // information sur l'utilisateur
		currentInformation = append(currentInformation, currentComment) // information sur le commentaire
		listPost = append(listPost, currentInformation)
	}

	// récupération des notifications de group
	listGroup = make([][]interface{}, 2)
	stmtGroup, err := db.Prepare(`SELECT USERS.FirstName, USERS.LastName, USERS.Avatar, USERS.Nickname,
										LISTGROUPS.NameGroup, LISTGROUPS.Image, LISTGROUPS.AboutUs
										FROM USERS
										INNER JOIN LISTGROUPS ON LISTGROUPS.UserID_Creator = USERS.ID
										WHERE NOTIFICATIONS.UserID_Receiver = ? AND NOTIFICATIONS.IDGroup = LISTGROUPS.IDGroup;`)
	CheckErr(err, "GetNotif listGroup, db prepare")
	rowsGroup, err := stmtGroup.Query(UserID)
	CheckErr(err, "GetNotif listGroup, db query")
	for rowsGroup.Next() {
		var currentUser User
		var currentGroup Group
		var currentInformation []interface{}
		err = rowsComment.Scan(&currentUser.Firstname, &currentUser.Lastname, &currentUser.Avatar, &currentUser.Nickname, &currentGroup.NameGroup, &currentGroup.Image, &currentGroup.Description)
		CheckErr(err, "listComment, db rowsComment.Next scan")
		currentInformation = append(currentInformation, currentUser)  // information sur l'utilisateur
		currentInformation = append(currentInformation, currentGroup) // information sur le groupe
		listGroup = append(listGroup, currentInformation)
	}

	// récupération des notifications d'event de groupe
	listEvent = make([][]interface{}, 3)
	stmtevent, err := db.Prepare(`SELECT EVENTGROUPS.Date, EVENTGROUPS.Title, EVENTGROUPS.UserID_Sender,
										LISTGROUPS.NameGroup, LISTGROUPS.Image, LISTGROUPS.AboutUs
										FROM EVENTGROUPS
										INNER JOIN EVENTGROUPS ON EVENTGROUPS.IDgroup = LISTGROUPS.IDgroup
										INNER JOIN NOTIFICATIONS ON NOTIFICATIONS.IDEvent = EVENTGROUPS.IDEvent
										WHERE NOTIFICATIONS.UserID_Receiver = ? AND NOTIFICATIONS.IDEvent = EVENTGROUPS.IDEvent;`)
	CheckErr(err, "GetNotif listEvent, db prepare")
	rowsevent, err := stmtevent.Query(UserID)
	CheckErr(err, "GetNotif listEvent, db query")
	for rowsevent.Next() {
		var currentGroup Group
		var currentEvent EventGroup
		var currentInformation []interface{}
		err = rowsComment.Scan(&currentEvent.Date, &currentEvent.Title, &currentEvent.UserIDCreatorEvent, &currentGroup.NameGroup, &currentGroup.Image, &currentGroup.Description)
		CheckErr(err, "listComment, db rowsComment.Next scan")
		currentInformation = append(currentInformation, currentGroup) // information sur le groupe
		currentInformation = append(currentInformation, currentEvent) // information sur l'event
		listPost = append(listPost, currentInformation)
	}

	return ListFollowers, listMP, listPost, listComment, listGroup, listEvent
}
