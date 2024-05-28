package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"text/template"

	data "forum/internal"
	utils "forum/middleware"

	_ "github.com/mattn/go-sqlite3"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./index.html")
	if err != nil {
		log.Fatal(err)
	}
	ts.Execute(w, nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// lecture du cookie
	cookie, err1 := r.Cookie("user_token")
	// ouverture de la database
	database, err := sql.Open("sqlite3", "./internal/forum.db")
	utils.CheckErr(err, "open db in homehandler")
	defer database.Close()

	// déclaration des informations de base
	userData := &data.User{
		StructCategories: utils.AllColumOfCategory(database),
		Logged:           false,
		Redirect:         "/",
	}
	var globalData data.GlobalData

	if err1 != http.ErrNoCookie && r.URL.Path == "/api/home" { // si presence d'un cookie et en home
		// -----------------------------------------------------------
		// ---------------------- USER CONNECTED ---------------------
		// -----------------------------------------------------------

		// récupération des informations de l'utilisateur
		_, username, email, _ := utils.UserInfo(cookie.Value, database)
		// vérification de la validité du cookie avec la map
		if !utils.CheckToken(utils.Sessions, cookie, email, database) {
			utils.Logout(w, r, database)
		} else {
			globalData.PostListData = utils.DisplayForum(w, r, database, "", username)
			userData.Username = html.EscapeString(username)
			userData.Logged = true
		}

		globalData.UserData = *userData
		jsonData, err := json.Marshal(globalData)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)

	} else if err1 != http.ErrNoCookie { // si presence d'un cookies
		// -----------------------------------------------------------
		// ---------------------- USER CONNECTED ---------------------
		// -----------------------------------------------------------

		if r.Method == "POST" {
			switch r.URL.Path {
			//     ╭──────────────────────────────────────────────────────────────╮
			//     │                          SEPARATEUR                          │
			//     ╰──────────────────────────────────────────────────────────────╯
			case "/api/PostEditor":
				// récupération des informations utilisateur
				user_id, username, email, _ := utils.UserInfo(cookie.Value, database)

				// vérification de la validité du cookie avec la map
				if !utils.CheckToken(utils.Sessions, cookie, email, database) {
					utils.Logout(w, r, database)
				} else {
					msgUser := utils.PostsEditor(w, r, user_id, username, database)
					userData.Msgerr = msgUser
					userData.Username = html.EscapeString(username)
					userData.Logged = true
					if msgUser == "" {
						userData.Redirect = "/home"
					}
				}
				break
				// ╭──────────────────────────────────────────────────────────────╮
				// │                          SEPARATEUR                          │
				// ╰──────────────────────────────────────────────────────────────╯
			case "/api/ViewPost":
				// récupération des informations utilisateur
				_, username, email, _ := utils.UserInfo(cookie.Value, database)
				// vérification de la validité du cookie avec la map
				if !utils.CheckToken(utils.Sessions, cookie, email, database) {
					utils.Logout(w, r, database)
				} else {
					userData.Username = html.EscapeString(username)
					userData.Logged = true
					userData.Redirect = ""

					// Récupération de l'id du post à afficher
					var viewPost data.ViewPost
					err = json.NewDecoder(r.Body).Decode(&viewPost)
					utils.CheckErr(err, "Erreur de décodage JSON viewpost")
					postID, _ := strconv.Atoi(viewPost.PostId)

					// Récupération des informations du post
					var postAuthor, postPhoto, postContent, postDate, postTitre, contentPostPhoto, contentPostFile, name_Category_1, id_Category_ID2, id_Category_ID3 string
					var postCategory []string
					requete := `SELECT 
											Users.Username, Users.Photo, 
											Posts.Content, Posts.Date, Posts.Titre, Posts.Category_ID2, Posts.Category_ID3, 
											Category.Name 
											FROM Posts 
											INNER JOIN Users ON Users.ID=Posts.User_ID 
											INNER JOIN Category ON Posts.Category_ID1=Category.ID 
											WHERE Posts.ID=?`
					stmtPost, err := database.Prepare(requete)
					utils.CheckErr(err, "INNER JOIN viewPost")
					stmtPost.QueryRow(postID).Scan(&postAuthor, &postPhoto, &postContent, &postDate, &postTitre, &id_Category_ID2, &id_Category_ID3, &name_Category_1)

					// Gestion des catégories du post à afficher
					postCategory = append(postCategory, name_Category_1)
					if id_Category_ID3 != "" {
						stmt, err := database.Prepare("SELECT Name FROM Category WHERE (ID=? OR ID=?)")
						utils.CheckErr(err, "1- Request category ID 2 viewPost")
						rows, err := stmt.Query(id_Category_ID2, id_Category_ID3)
						utils.CheckErr(err, "2- Request category ID 2 viewPost")
						for rows.Next() {
							var categ string
							err = rows.Scan(&categ)
							utils.CheckErr(err, "3- Request category ID 2 viewPost")
							postCategory = append(postCategory, categ)
						}
					} else if id_Category_ID2 != "" {
						var name_Category_2 string
						stmt, err := database.Prepare("SELECT Name FROM Category WHERE ID = ?")
						stmt.QueryRow(id_Category_ID2).Scan(&name_Category_2)
						utils.CheckErr(err, "Request category ID 2 viewPost")
						postCategory = append(postCategory, name_Category_2)
					}

					post := &data.Post{
						ID:           postID,
						Title:        html.EscapeString(postTitre),
						Content:      postContent,
						Date:         postDate,
						Author:       html.EscapeString(postAuthor),
						PhotoAuthor:  "/assets/photoCompte/" + postPhoto,
						Categorie:    postCategory,
						ContentPhoto: contentPostPhoto,
						ContentFile:  contentPostFile,

						Comments: utils.SearchComment(postID, database),
					}
					if post != nil {
						globalData.PostData = *post
					}
				}
				break
				// ╭──────────────────────────────────────────────────────────────╮
				// │                          SEPARATEUR                          │
				// ╰──────────────────────────────────────────────────────────────╯
			case "/api/NewComment":
				// récupération des informations utilisateur
				user_id, username, email, _ := utils.UserInfo(cookie.Value, database)
				// vérification de la validité du cookie avec la map
				if !utils.CheckToken(utils.Sessions, cookie, email, database) {
					utils.Logout(w, r, database)
				} else {
					PostID := utils.PostsComment(r, user_id, database)
					userData.Redirect = "/ViewPost/?post=" + PostID
					userData.Username = html.EscapeString(username)
					userData.Logged = true
					userData.Redirect = ""
				}
				break
				// ╭──────────────────────────────────────────────────────────────╮
				// │                          SEPARATEUR                          │
				// ╰──────────────────────────────────────────────────────────────╯
			case "/api/historyMsg":
				globalData.History = utils.SearchHistory(r, database)
				userData.Logged = true
				userData.Redirect = ""
				break
				// ╭──────────────────────────────────────────────────────────────╮
				// │                          SEPARATEUR                          │
				// ╰──────────────────────────────────────────────────────────────╯
			case "/api/Compte":
				// récupération des informations utilisateur
				_, username, email, linkPhoto, Age, Gender, FirstName, LastName := utils.UserAccount(cookie.Value, database)
				// vérification de la validité du cookie avec la map
				if !utils.CheckToken(utils.Sessions, cookie, email, database) {
					utils.Logout(w, r, database)
				} else {
					// mise des informations dans UserData pour l'envoi json
					userData.Photo = linkPhoto
					userData.Email = html.EscapeString(email)
					userData.Username = html.EscapeString(username)
					userData.Logged = true
					userData.Age = Age
					userData.Gender = Gender
					userData.FirstName = html.EscapeString(FirstName)
					userData.LastName = html.EscapeString(LastName)
					userData.Redirect = ""
				}
			}
			// ╭──────────────────────────────────────────────────────────────╮
			// │                          SEPARATEUR                          │
			// ╰──────────────────────────────────────────────────────────────╯
		}

		globalData.UserData = *userData
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(globalData)

	} else { // si absence de cookies
		// -----------------------------------------------------------
		// ------------------- USER NOT CONNECTED --------------------
		// -----------------------------------------------------------

		if r.Method == "POST" {
			switch r.URL.Path {
			//     ╭──────────────────────────────────────────────────────────────╮
			//     │                          SEPARATEUR                          │
			//     ╰──────────────────────────────────────────────────────────────╯
			case "/api/Register":
				if utils.RegisterUser(w, r, database) {
					userData.Redirect = "/Login"
				}
				break
				// ╭──────────────────────────────────────────────────────────────╮
				// │                          SEPARATEUR                          │
				// ╰──────────────────────────────────────────────────────────────╯
			case "/api/Login":
				valid, username := utils.Login(w, r, database)
				if valid {
					userData.Logged = true
					userData.Redirect = "/home"
					userData.Username = username
				} else {
					userData.Msgerr = "Mauvais e-mail ou mot de passe."
					userData.Redirect = ""
				}
				break
				// ╭──────────────────────────────────────────────────────────────╮
				// │                          SEPARATEUR                          │
				// ╰──────────────────────────────────────────────────────────────╯
			}
		}

		globalData.UserData = *userData
		jsonData, err := json.Marshal(globalData)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}
