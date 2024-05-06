package handlers

import (
	"database/sql"
	"fmt"
	data "forum/internal"
	utils "forum/middleware"
	"html"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// Page d'accueil "/"
// /////////////// USER CONNECTED AND DISCONNECTED \\\\\\\\\\\\\\\\\
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err1 := r.Cookie("user_token")
	database, err := sql.Open("sqlite3", "./internal/forum.db")
	utils.CheckErr(err, "open db in homehandler")
	defer database.Close()

	var user_id, username string
	if err1 == nil {
		var email string
		stmt, err := database.Prepare("SELECT Email, ID, Username FROM Users WHERE uuid = ?")
		stmt.QueryRow(cookie.Value).Scan(&email, &user_id, &username)
		utils.CheckErr(err, "homehandler")
		if !utils.CheckToken(utils.Sessions, cookie, email, database) {
			utils.Logout(w, r, database)
		}
	}
	if r.URL.Path == "/" {
		var catFilter []string
		if r.Method == "POST" {
			r.ParseForm()
			if r.FormValue("logout") == "logout" && err1 == nil {
				utils.Logout(w, r, database)
			}
			catFilter = r.Form["categorieFilter"]
			for _, v := range catFilter {
				if v, err := strconv.Atoi(v); v > 8 || v < 1 && err != nil {
					w.WriteHeader(500)
					http.Redirect(w, r, "/", 303)
				}
			}
		}
		// fmt.Println("catFilter : ", catFilter)

		keys, ok := r.URL.Query()["session"]
		if !ok && err1 == http.ErrNoCookie || len(keys) < 1 && err1 == http.ErrNoCookie {
			utils.HomeNotConnected(w, r, catFilter, database)
			return
		} else if !ok && err1 != http.ErrNoCookie || len(keys) < 1 && err1 != http.ErrNoCookie {

			postData := utils.DisplayForum(w, r, catFilter, database, "", username, "", "")
			userData := &data.User{
				StructCategories: utils.AllColumOfCategory(database),
				Username:         html.EscapeString(username),
				Logged:           true,
			}
			data := &data.GlobalData{
				UserData:     *userData,
				PostListData: postData,
			}
			utils.UserRenderTemplate(w, "Home", data)
			return
		}
		key := string(keys[0])
		if key == "connected" && err1 != http.ErrNoCookie {
			postData := utils.DisplayForum(w, r, catFilter, database, "", username, "", "")

			UserData := &data.User{
				StructCategories: utils.AllColumOfCategory(database),
				Username:         html.EscapeString(username),
				Logged:           true,
			}
			data := &data.GlobalData{
				UserData:     *UserData,
				PostListData: postData,
			}
			utils.UserRenderTemplate(w, "Home", data)
		} else if key == "connected" && err1 == http.ErrNoCookie {
			utils.HomeNotConnected(w, r, catFilter, database)
		}
		if key == "disconnected" || (key == "expired" && err1 != http.ErrNoCookie) {
			utils.HomeNotConnected(w, r, catFilter, database)
		}

		if key == "postPublish" {
			var catFilter []string
			postData := utils.DisplayForum(w, r, catFilter, database, user_id, username, "postPublish", "")
			userData := &data.User{
				StructCategories: utils.AllColumOfCategory(database),
				Username:         html.EscapeString(username),
				Logged:           true,
			}
			data := &data.GlobalData{
				UserData:     *userData,
				PostListData: postData,
			}

			utils.UserRenderTemplate(w, "Home", data)
		} else if key == "postLike" {
			var catFilter []string
			postData := utils.DisplayForum(w, r, catFilter, database, user_id, username, "", "postLike")
			userData := &data.User{
				StructCategories: utils.AllColumOfCategory(database),
				Username:         html.EscapeString(username),
				Logged:           true,
			}
			data := &data.GlobalData{
				UserData:     *userData,
				PostListData: postData,
			}

			utils.UserRenderTemplate(w, "Home", data)
		}
	} else if r.URL.Path != "/" {
		w.WriteHeader(404)
		utils.RenderTemplate(w, "404", &data.DataTemplate{})
		// w.Write([]byte("--------------- sa passe ici --------------------"))
	} else {
		w.WriteHeader(400)
	}
}

// ----------------------------------------------------
// ----------------- page utilisateur -----------------
// ----------------------------------------------------

// Page d'enregistrement d'un nouvelle utilisateur '/Register'
// /////////////// USER DISCONNECTED \\\\\\\\\\\\\\\\\
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	_, err1 := r.Cookie("user_token")
	if r.Method == "POST" {
		if r.URL.Path == "/Register" {
			database, _ := sql.Open("sqlite3", "./internal/forum.db")
			defer database.Close()
			if utils.RegisterUser(w, r, database) {
				http.Redirect(w, r, "/?register=success", http.StatusSeeOther)
			}
		}
	} else if r.Method == "GET" && err1 != http.ErrNoCookie {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		utils.RenderTemplate(w, "Register", &data.DataTemplate{})
	}
}

// Page de Login utilisateur enregistré '/Login'
// /////////////// USER DISCONNECTED \\\\\\\\\\\\\\\\\
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	_, err1 := r.Cookie("user_token")
	if r.Method == "POST" {
		if r.URL.Path == "/Login" {
			database, _ := sql.Open("sqlite3", "./internal/forum.db")
			defer database.Close()
			if utils.Login(w, r, database) {
				http.Redirect(w, r, "/?session=connected", http.StatusSeeOther)
			} else {
				UserData := &data.User{
					Msgerr: "Mauvais e-mail ou mot de passe. Nouveau ? <a href=\"/Register\">s'enregistrer</a>",
					Logged: false,
				}
				data := &data.GlobalData{
					UserData: *UserData,
				}
				utils.UserRenderTemplate(w, "Login", data)
				// w.Write([]byte("Wrong Email/Password. New here ? <a href=\"/Register\">Register</a>"))
			}
		} else if r.URL.Path != "/Login" {
			w.WriteHeader(404)
			utils.RenderTemplate(w, "404", &data.DataTemplate{})
		}
	} else if r.Method == "GET" && err1 != http.ErrNoCookie {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		utils.UserRenderTemplate(w, "Login", &data.GlobalData{})
	}
}

// Page de gestion du compte utilisateur '/Compte'
// /////////////// USER CONNECTED \\\\\\\\\\\\\\\\\
func CompteHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err1 := r.Cookie("user_token")

	if err1 != http.ErrNoCookie && r.URL.Path == "/Compte" { // Si Cookie present = user connecté et '/compte'
		database, err := sql.Open("sqlite3", "./internal/forum.db")
		utils.CheckErr(err, "open db in viewpostshandler")
		defer database.Close()

		// récupération des informations utilisateur
		var user_id, username, linkPhoto, email string
		user_id, username, email, linkPhoto = utils.UserInfo(cookie.Value, database)
		if err1 == nil { // si presence d'un cookies
			// vérification de la validité du cookie avec la map
			if !utils.CheckToken(utils.Sessions, cookie, email, database) {
				utils.Logout(w, r, database)
			}
			var msgUser string
			if r.Method == "POST" {
				// Si deconnection demandé par l'utilisateur
				if r.FormValue("logout") == "logout" {
					utils.Logout(w, r, database)
				}
				msgUser, linkPhoto, username = utils.UpdateAccount(w, r, user_id, linkPhoto, username, database)

				if r.FormValue("postLike") == "postLike" {
					http.Redirect(w, r, "/?session=postLike", 303)
					return
				} else if r.FormValue("postPublish") == "postPublish" {
					http.Redirect(w, r, "/?session=postPublish", 303)
					return
				}
			}
			UserData := &data.User{
				Photo:    linkPhoto,
				Email:    html.EscapeString(email),
				Username: html.EscapeString(username),
				Msgerr:   msgUser,
				Logged:   true,
			}
			global := &data.GlobalData{
				UserData: *UserData,
			}
			utils.UserRenderTemplate(w, "Compte", global)
		} else { // Si Cookie Inexistant ou expiré
			utils.RenderTemplate(w, "Home", &data.DataTemplate{})
		}
	} else if r.URL.Path != "/Compte" {
		w.WriteHeader(404)
		utils.RenderTemplate(w, "404", &data.DataTemplate{})
	} else if r.URL.Path == "/Compte" && err1 == http.ErrNoCookie { // si pas de cookie et en '/compte'
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		w.WriteHeader(400)
	}
}

// Page permettant à un utilisateur de re-initialisé sont mot de passe '/ForgetPassword'
// /////////////// USER DISCONNECTED \\\\\\\\\\\\\\\\\
func ForgetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	_, err1 := r.Cookie("user_token")
	if r.Method == "POST" {
		var msg string
		var codeinit, codeconfirm bool
		if r.Method == "POST" {
			database, err := sql.Open("sqlite3", "./internal/forum.db")
			utils.CheckErr(err, "open db in viewpostshandler")
			defer database.Close()
			msg, codeinit, codeconfirm = utils.ForgetPassword(r, database)
		}
		data := &data.DataTemplate{
			Msgerr:      msg,
			Codeinit:    codeinit,
			Codeconfirm: codeconfirm,
		}
		utils.RenderTemplate(w, "ForgetPassword", data)
	} else if r.Method == "GET" && err1 != http.ErrNoCookie {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		utils.RenderTemplate(w, "Login", &data.DataTemplate{})
	}
}

// ----------------------------------------------------
// ------------------ page du forum -------------------
// ----------------------------------------------------

// Page permettant de voir les details d'un post en particulier '/ViewPost'
// /////////////// USER CONNECTED AND DISCONNECTED \\\\\\\\\\\\\\\\\
func ViewPostHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err1 := r.Cookie("user_token")
	database, err := sql.Open("sqlite3", "./internal/forum.db")
	utils.CheckErr(err, "open db in viewpostshandler")
	defer database.Close()

	// check de la validité du cookie avec la map de session
	if err1 == nil {
		var email string
		stmt, err := database.Prepare("SELECT Email FROM Users WHERE uuid = ?")
		stmt.QueryRow(cookie.Value).Scan(&email)
		utils.CheckErr(err, "ViewPostHandler")
		if !utils.CheckToken(utils.Sessions, cookie, email, database) {
			utils.Logout(w, r, database)
		}
	}

	if r.URL.Path == "/ViewPost" {
		// Check si l'utillisateur a demandé une deconnexion
		if err1 != http.ErrNoCookie && r.Method == "POST" { // Si Cookie present = user connecté (Method post)
			r.ParseForm()
			if r.FormValue("logout") == "logout" {
				utils.Logout(w, r, database)
			}
		}

		var username, user_id string
		var logged bool
		// Check si l'utilisateur et logger ou non
		if err1 != http.ErrNoCookie {
			stmt, err := database.Prepare("SELECT Username,ID FROM Users WHERE uuid = ?")
			utils.CheckErr(err, "viewpostshandler")
			stmt.QueryRow(cookie.Value).Scan(&username, &user_id)
			utils.CheckErr(err, "viewpostshandler")
			logged = true
		}

		// Récupération de l'ID du post à traité
		keys := r.URL.Query()["id"]
		if len(keys) < 1 { // vérification qu'un key est presente
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		postID, _ := strconv.Atoi(keys[0])

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
		stmt, err := database.Prepare(requete)
		utils.CheckErr(err, "INNER JOIN viewPost")
		stmt.QueryRow(postID).Scan(&postAuthor, &postPhoto, &postContent, &postDate, &postTitre, &id_Category_ID2, &id_Category_ID3, &name_Category_1)

		// Récupération du lien image du post (si existant)
		stmtPhoto, err := database.Prepare("SELECT LienImage FROM Posts WHERE ID = ?")
		utils.CheckErr(err, "ViewPostHandler")
		stmtPhoto.QueryRow(postID).Scan(&contentPostPhoto)

		// Récupération du lien fichier du post (si existant)
		stmtFile, err := database.Prepare("SELECT FichierJoint FROM Posts WHERE ID = ?")
		utils.CheckErr(err, "ViewPostHandler")
		stmtFile.QueryRow(postID).Scan(&contentPostFile)

		// Gestion des catégories du post à afficher
		postCategory = append(postCategory, name_Category_1)
		if id_Category_ID3 != "" {
			stmt, err = database.Prepare("SELECT Name FROM Category WHERE (ID=? OR ID=?)")
			rows, err := stmt.Query(id_Category_ID2, id_Category_ID3)
			utils.CheckErr(err, "Request category ID 2 viewPost")
			for rows.Next() {
				var categ string
				err = rows.Scan(&categ)
				postCategory = append(postCategory, categ)
			}
		} else if id_Category_ID2 != "" {
			var name_Category_2 string
			stmt, err = database.Prepare("SELECT Name FROM Category WHERE ID = ?")
			stmt.QueryRow(id_Category_ID2).Scan(&name_Category_2)
			utils.CheckErr(err, "Request category ID 2 viewPost")
			postCategory = append(postCategory, name_Category_2)
		}

		var likeCounter, dislikeCounter int
		var likes, dislikes []int
		if r.Method == "GET" { // Check en méthod GET si un like ou un dislike et cliquer
			fmt.Println("likeCounter", likeCounter)
			fmt.Println("dislikeCounter ", dislikeCounter)
			likeCounter, dislikeCounter, likes, dislikes = utils.SearchSocial(strconv.Itoa(postID), database)
			fmt.Println("likeCounter", likeCounter)
			fmt.Println("dislikeCounter ", dislikeCounter)
		} else if err1 != http.ErrNoCookie && r.Method == "POST" { // Si Cookie present = user connecté (Method post)
			r.ParseForm()
			likeCounter, dislikeCounter, likes, dislikes = utils.SearchSocial(strconv.Itoa(postID), database)
			socialsLikeComment := r.FormValue("socialsLikeComment")
			socialsDislikeComment := r.FormValue("socialsDislikeComment")

			if r.FormValue("newComment") != "" { // Ajout d'un commentaire
				utils.PostsComment(r, user_id, username, database)
			} else if socialsLikeComment != "" { // Si l'utilisateur like un commentaire
				socialsLikeCommentInt, err := strconv.Atoi(socialsLikeComment)
				utils.CheckErr(err, "Atoi socialsLikeComment")
				utils.UpdateSocialsComment(user_id, 1, socialsLikeCommentInt, database)
				fmt.Println("likeCounter", likeCounter)
				fmt.Println("dislikeCounter ", dislikeCounter)
			} else if socialsDislikeComment != "" { // Si l'utilisateur dislike un commentaire
				socialsDislikeCommentInt, err := strconv.Atoi(socialsDislikeComment)
				fmt.Println("likeCounter", likeCounter)
				fmt.Println("dislikeCounter ", dislikeCounter)
				utils.CheckErr(err, "Atoi socialsLikeComment")
				utils.UpdateSocialsComment(user_id, -1, socialsDislikeCommentInt, database)
				fmt.Println("likeCounter", likeCounter)
				fmt.Println("dislikeCounter ", dislikeCounter)
			} else if r.PostForm["socials"] != nil { // ajout d'un like ou d'un dislike
				if r.PostForm["socials"][0] == "liked" || r.PostForm["socials"][0] == "disliked" {
					// handle social buttons behaviour
					// create UserData to render template accordingly
					keys := r.URL.Query()["id"]
					postID, _ = strconv.Atoi(keys[0])
					utils.UpdateSocials(r, username, database, postID)
					fmt.Println("likeCounter", likeCounter)
					fmt.Println("dislikeCounter ", dislikeCounter)
					likeCounter, dislikeCounter, _, _ = utils.SearchSocial(keys[0], database)
					fmt.Println("likeCounter", likeCounter)
					fmt.Println("dislikeCounter ", dislikeCounter)
				}
			}
		}

		if contentPostPhoto != "" {
			contentPostPhoto = "assets/StockageClients/Images/" + contentPostPhoto
		}
		if contentPostFile != "" {
			contentPostFile = "assets/StockageClients/FichiersJoints/" + contentPostFile
		}
		fmt.Println("likeCounter à envoyer", likeCounter)
		fmt.Println("dislikeCounter à envoyer", dislikeCounter)
		// Information a envoyer en Front
		UserData := &data.User{
			Username: html.EscapeString(username),
			Logged:   logged,
		}
		SocialsData := &data.Post{
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

			LikeCounter:    likes,
			DislikeCounter: dislikes,
			Likes:          likeCounter,
			Dislikes:       dislikeCounter,
		}
		data := &data.GlobalData{
			UserData: *UserData,
			PostData: *SocialsData,
		}
		utils.UserRenderTemplate(w, "ViewPost", data)
	} else if r.URL.Path != "/ViewPost" {
		w.WriteHeader(404)
		utils.RenderTemplate(w, "404", &data.DataTemplate{})
	} else {
		w.WriteHeader(400)
	}
}

// page permettant de crée un post '/PostEditor'
// /////////////// USER CONNECTED \\\\\\\\\\\\\\\\\
func PostEditorHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err1 := r.Cookie("user_token")

	if err1 != http.ErrNoCookie && r.URL.Path == "/PostEditor" {
		database, err := sql.Open("sqlite3", "./internal/forum.db")
		utils.CheckErr(err, "open db in viewpostshandler")
		defer database.Close()
		if err1 == nil { // si presence d'un cookies

			var user_id, username, email, msgUser string
			// récupération des informations utilisateur
			user_id, username, email, _ = utils.UserInfo(cookie.Value, database)
			// vérification de la validité du cookie avec la map
			if !utils.CheckToken(utils.Sessions, cookie, email, database) {
				utils.Logout(w, r, database)
			}

			//Récupère les catégories dans le sql (desserts, boissons, légumes...)
			namesCategories := utils.AllColumOfCategory(database)

			if r.Method == "POST" {
				// Si deconnection demandé par l'utilisateur
				if r.FormValue("logout") == "logout" {
					utils.Logout(w, r, database)
				}
				// va inserer les données du post dans le db
				var namesCat []string
				for _, val := range namesCategories {
					namesCat = append(namesCat, val.Name)
				}
				msgUser = utils.PostsEditor(w, r, user_id, username, database)
				fmt.Println("msgUser : ", msgUser)
				if msgUser == ", votre post à bien été publié" { // si msgUser non vide, un post viens d'être poster, redirection vers la home
					http.Redirect(w, r, "/", http.StatusSeeOther)
					fmt.Println("ici")
					return
				}
			}

			UserData := &data.User{
				Msgerr:           msgUser,
				Username:         html.EscapeString(username),
				Logged:           true,
				StructCategories: utils.AllColumOfCategory(database),
			}
			global := &data.GlobalData{
				UserData: *UserData,
			}
			if msgUser == "" && r.FormValue("texte") != "" {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			utils.UserRenderTemplate(w, "PostEditor", global)
		} else {
			var catFilter []string
			if r.Method == "POST" {
				catFilter = r.Form["categorieFilter"]
				//Checks whether or not the user try to be a sassy boy
				for _, v := range catFilter {
					if v, err := strconv.Atoi(v); v > 8 || v < 1 && err != nil {
						fmt.Println("toto")
						http.Redirect(w, r, "/", 303)
					}
				}
			}
			utils.HomeNotConnected(w, r, catFilter, database)
		}
	} else if r.URL.Path == "/Compte" && err1 == http.ErrNoCookie {
		http.Redirect(w, r, "/", 303)
	} else if r.URL.Path != "/Compte" {
		w.WriteHeader(404)
		utils.RenderTemplate(w, "404", &data.DataTemplate{})
	} else {
		w.WriteHeader(400)
	}
}
