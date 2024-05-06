package utils

import (
	"database/sql"
	data "forum/internal"
	"html"
	"log"
	"net/http"
	"strconv"
)

func DisplayForum(w http.ResponseWriter, r *http.Request, catFilter []string, database *sql.DB, user_id, username, postPublish, postLike string) []*data.Post {
	var id, title, date, userDuPost string
	var categ1, categ2, categ3 string
	var likes, dislikes int
	var categories []string
	rows, err := database.Query(`SELECT Posts.ID, Titre, Date, Users.Username, Category_ID1 ,Category_ID2 ,Category_ID3 FROM Posts INNER JOIN Users ON Users.ID=Posts.User_ID`)
	var postTab []*data.Post
	if err != nil {
		log.Fatal(err)
	} else {
		for rows.Next() {
			rows.Scan(&id, &title, &date, &userDuPost, &categ1, &categ2, &categ3)

			if catFilter == nil && postPublish == "" && postLike == "" { // si absence de filtre, on charge tout les posts
				categories = Categories(w, r, database, categ1, categ2, categ3)

				likes, dislikes, _, _ = SearchSocial(id, database)
				idInt, err := strconv.Atoi(id)
				CheckErr(err, "Atoi")
				var PostsData = &data.Post{
					ID:        idInt,
					Title:     html.EscapeString(title),
					Date:      date,
					Author:    html.EscapeString(userDuPost),
					Categorie: categories,
					Likes:     likes,
					Dislikes:  dislikes,
				}
				postTab = append(postTab, PostsData)

			} else if postPublish == "postPublish" { // l'utilisateur demande l'affichage de ces posts publié
				categories = Categories(w, r, database, categ1, categ2, categ3)

				if username == userDuPost {

					likes, dislikes, _, _ = SearchSocial(id, database)
					idInt, err := strconv.Atoi(id)
					CheckErr(err, "Atoi")
					var PostsData = &data.Post{
						ID:        idInt,
						Title:     html.EscapeString(title),
						Date:      date,
						Author:    html.EscapeString(userDuPost),
						Categorie: categories,
						Likes:     likes,
						Dislikes:  dislikes,
					}
					postTab = append(postTab, PostsData)
				}
			} else if postLike == "postLike" { // l'utilisateur demande l'affichage de ces post aimé
				categories = Categories(w, r, database, categ1, categ2, categ3)

				stmt, err := database.Prepare("SELECT Post_ID FROM Socials WHERE user_id = ? AND Socials = 1")
				CheckErr(err, "postLiker")
				res, err := stmt.Query(user_id)
				CheckErr(err, "postLiker")
				var listPost []string
				for res.Next() {
					var postID string
					res.Scan(&postID)
					listPost = append(listPost, postID)
				}

				if checkListPost(listPost, id) {

					likes, dislikes, _, _ = SearchSocial(id, database)
					idInt, err := strconv.Atoi(id)
					CheckErr(err, "Atoi")
					var PostsData = &data.Post{
						ID:        idInt,
						Title:     html.EscapeString(title),
						Date:      date,
						Author:    html.EscapeString(userDuPost),
						Categorie: categories,
						Likes:     likes,
						Dislikes:  dislikes,
					}
					postTab = append(postTab, PostsData)
				}
			} else { // presence d'un filtre pour ne pas afficher tout les posts
				if checkCategId(catFilter, categ1, categ2, categ3) {
					categories = append(categories, categ1)
					categories = append(categories, categ2)
					categories = append(categories, categ3)

					likes, dislikes, _, _ = SearchSocial(id, database)

					idInt, err := strconv.Atoi(id)
					CheckErr(err, "Atoi")
					var PostsData = &data.Post{
						ID:        idInt,
						Title:     html.EscapeString(title),
						Date:      date,
						Author:    html.EscapeString(userDuPost),
						Categorie: categories,
						Likes:     likes,
						Dislikes:  dislikes,
					}
					postTab = append(postTab, PostsData)
				}
			}

			categories = nil
			categ3 = ""
			categ2 = ""
			categ1 = ""
			likes = 0
			dislikes = 0
		}
	}
	return postTab
}

func Categories(w http.ResponseWriter, r *http.Request, database *sql.DB, categ1, categ2, categ3 string) (categories []string) {
	var cat1, cat2, cat3 string

	// fmt.Printf("id : %v, categ1 : %v, categ2 : %v, categ3 : %v\n", id, categ1, categ2, categ3)
	req, err := database.Prepare(`SELECT Name FROM Category WHERE ID = ?`)
	CheckErr(err, "db prepare categorie")

	err = req.QueryRow(categ1).Scan(&cat1)
	CheckErr(err, "db Exet categorie 1")
	req.QueryRow(categ2).Scan(&cat2)
	req.QueryRow(categ3).Scan(&cat3)

	categories = append(categories, cat1)
	categories = append(categories, cat2)
	categories = append(categories, cat3)
	return
}

func checkCategId(catFilter []string, categ1, categ2, categ3 string) bool {
	// fmt.Println("val : ", catFilter)
	for _, val := range catFilter {
		// fmt.Printf("val : %v\t-1 : %v\t-2 : %v\t-3: %v\n", val, categ1, categ2, categ3)
		if val == categ1 || val == categ2 || val == categ3 {
			return true
		}
	}
	return false
}

func checkListPost(listPost []string, ID string) bool {
	for _, val := range listPost {
		// fmt.Printf("val : %v\t-1 : %v\t-2 : %v\t-3: %v\n", val, categ1, categ2, categ3)
		if string(val) == ID {
			return true
		}
	}
	return false
}
