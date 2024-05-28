package utils

import (
	"database/sql"
	data "forum/internal"
	"html"
	"log"
	"net/http"
	"strconv"
)

func DisplayForum(w http.ResponseWriter, r *http.Request, database *sql.DB, user_id, username string) []*data.Post {
	var id, title, date, userDuPost string
	var categ1, categ2, categ3 string
	var categories []string
	rows, err := database.Query(`SELECT Posts.ID, Titre, Date, Users.Username, Category_ID1 ,Category_ID2 ,Category_ID3 FROM Posts INNER JOIN Users ON Users.ID=Posts.User_ID`)
	var postTab []*data.Post
	if err != nil {
		log.Fatal(err)
	} else {
		for rows.Next() {
			rows.Scan(&id, &title, &date, &userDuPost, &categ1, &categ2, &categ3)

			categories = Categories(w, r, database, categ1, categ2, categ3)

			idInt, err := strconv.Atoi(id)
			CheckErr(err, "Atoi")
			var PostsData = &data.Post{
				ID:        idInt,
				Title:     html.EscapeString(title),
				Date:      date,
				Author:    html.EscapeString(userDuPost),
				Categorie: categories,
			}
			postTab = append(postTab, PostsData)
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
