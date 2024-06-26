package utils

import (
	"database/sql"
	"encoding/json"
	data "forum/internal"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// Envoie les données du post dans le db (tab sql)
func PostsEditor(w http.ResponseWriter, r *http.Request, user_id, username string, database *sql.DB) string {
	// Récupération des informations du json
	var createpost data.Createpost
	err := json.NewDecoder(r.Body).Decode(&createpost)
	CheckErr(err, "Erreur de décodage JSON viewpost")
	titre := createpost.NameRecette
	content := createpost.Recette
	categories := createpost.Category

	//récupère plusieurs valeurs html ayant un name en commun
	valided, msgUser := PostChecker(titre, content, categories, database, username)
	if valided {
		var boolFileJoint, boolPictureJoint bool
		var nameImage, nameFile string
		fileImage, headerimage, errFormImage := r.FormFile("UploadImage")
		if errFormImage == nil {
			//Permet de savoir la taille de l'image
			if headerimage.Size < 200000 {
				defer fileImage.Close()
				nameImage = RandomFileName() + headerimage.Filename
				boolPictureJoint = UploadPicture(fileImage, headerimage, "./assets/StockageClients/Images/"+nameImage)
			} else {
				// fmt.Println(headerimage.Size)
				return "L'image dépasse les 20mb\n"
			}
		}
		//idem pour fichier
		filejoint, headerfile, errFormFile := r.FormFile("UploadFile")
		if errFormFile == nil {
			defer filejoint.Close()
			nameFile = RandomFileName() + headerfile.Filename
			UploadFile(filejoint, headerfile, "./assets/StockageClients/FichiersJoints/"+nameFile)
			boolFileJoint = true
		}
		//On évite de se prendre la tête, on fais un stockage ;)
		var ValeurInsert = data.InsertC{
			FichB:      boolFileJoint,
			ImageB:     boolPictureJoint,
			Titre:      titre,
			Content:    content,
			User_id:    user_id,
			Image:      nameImage,
			Fichier:    nameFile,
			Categories: categories,
		}
		//Insert toutes les valeurs dans le tab sql
		INSERTion(w, r, ValeurInsert, database)
		return ""
	} else {
		msgUser = username + msgUser
		return msgUser
	}
}

// Vérifie si il n'y pas de poste similaire ou que les requis du posts sont bien làs
func PostChecker(titre, content string, categories []string, database *sql.DB, username string) (bool, string) {
	if titre == "" || content == "" {
		return false, ", absence de titre et/ou de contenu"
	}
	//Sert à récupèrer la valeur donné dans la catégorie, pour ensuite la comparer pour voir si elle existe
	usedtitre := database.QueryRow("SELECT Titre FROM Posts WHERE Titre = ?", titre).Scan()
	usedcontent := database.QueryRow("SELECT Content FROM Posts WHERE Content = ?", content).Scan()
	userName := database.QueryRow("SELECT Content FROM Posts WHERE UsernameDuPost = ?", username).Scan()
	// meme titre mais pas du meme auteur, pas le meme
	if (usedtitre != sql.ErrNoRows && usedcontent != sql.ErrNoRows) || (userName != sql.ErrNoRows && usedcontent != sql.ErrNoRows) || usedcontent != sql.ErrNoRows {
		return false, ", contenu ou titre déja présent sur le forum"
	} else if usedcontent == nil && usedtitre == nil {
		// fmt.Println("Il faut qu'au moins il y est un titre et un contenu pour pouvoir publier le poste ")
		return false, ", il faut qu'au moins il y est un titre et un contenu pour pouvoir publier le poste (cette erreur s'affiche aussi en cas de mauvaise categorie appliquée) "
	} else {
		if len(categories) > 0 && len(categories) < 4 {
			return true, ""
		} else {
			return false, ", il faut au minimum une categorie et au maximum 3 categories"
		}
	}
}

// Insert les données dans le SQL en fonction de ce que le client nous envoient comme données à traiter
func INSERTion(w http.ResponseWriter, r *http.Request, V data.InsertC, database *sql.DB) {
	date := GetTime()
	Categories := InsertCategories(V.Categories, database)
	if V.FichB == true && V.ImageB == true { // Insertion dans la DB : avec image et fichier joint
		stmt, err := database.Prepare("INSERT INTO Posts(Titre,Content,User_id,LienImage,FichierJoint,Date,Category_ID1,Category_ID2,Category_ID3) Values(?,?,?,?,?,?,?,?,?)")
		CheckErr(err, "Y'a une merde ici: INSERTion img + file")
		stmt.Exec(V.Titre, V.Content, V.User_id, V.Image, V.Fichier, date, Categories[0], Categories[1], Categories[2])
	} else if V.ImageB == true { // Insertion dans la DB : avec seulement une image jointe
		stmt, err := database.Prepare("INSERT INTO Posts(Titre,Content,User_id,LienImage,Date,Category_ID1,Category_ID2,Category_ID3) Values(?,?,?,?,?,?,?,?)")
		CheckErr(err, "Y'a une merde ici: INSERTion img")
		stmt.Exec(V.Titre, V.Content, V.User_id, V.Image, date, Categories[0], Categories[1], Categories[2])
	} else if V.FichB == true { // Insertion dans la DB : avec seulement un fichier joint
		stmt, err := database.Prepare("INSERT INTO Posts(Titre,Content,User_id,FichierJoint,Date,Category_ID1,Category_ID2,Category_ID3) Values(?,?,?,?,?,?,?,?)")
		CheckErr(err, "Y'a une merde ici: INSERTion img + File")
		stmt.Exec(V.Titre, V.Content, V.User_id, V.Fichier, date, Categories[0], Categories[1], Categories[2])
	} else { // Insertion dans la DB : sans image, ni fichier joint
		stmt, err := database.Prepare("INSERT INTO Posts(Titre,Content,User_id,Date,Category_ID1,Category_ID2,Category_ID3) Values(?,?,?,?,?,?,?)")
		CheckErr(err, "Y'a une merde ici: INSERTion simple")
		stmt.Exec(V.Titre, V.Content, V.User_id, date, Categories[0], Categories[1], Categories[2])
	}
}

// Permet de préparer les valeurs en remplissant avec les valeurs et du vides si nécessaires
func InsertCategories(categories []string, database *sql.DB) []string {
	// fmt.Println("categories : ", categories[0], categories[1], categories[2], len(categories))

	CategoSQL := make([]string, 3)
	var index = 0
	if len(categories) == 1 {
		req, err := database.Prepare(`SELECT ID FROM Category WHERE Name = ?`)
		CheckErr(err, "db prepare categorie 1 avant")
		rows, err := req.Query(categories[0])
		CheckErr(err, "db Exet Query categorie 1")
		for rows.Next() {
			var categ string
			err = rows.Scan(&categ)
			CheckErr(err, "row categorie 1")
			CategoSQL[index] = categ
		}
	} else if len(categories) == 2 {
		req, err := database.Prepare(`SELECT ID FROM Category WHERE (Name = ? or Name = ?)`)
		CheckErr(err, "db prepare categorie 2 avant")
		rows, err := req.Query(categories[0], categories[1])
		CheckErr(err, "db Exet Query categorie 2")
		for rows.Next() {
			var categ string
			err = rows.Scan(&categ)
			CheckErr(err, "row categorie 2")
			CategoSQL[index] = categ
			index++
		}
	} else {
		req, err := database.Prepare(`SELECT ID FROM Category WHERE (Name = ? or Name = ? or Name = ?)`)
		CheckErr(err, "db prepare categorie 3 avant")
		rows, err := req.Query(categories[0], categories[1], categories[2])
		CheckErr(err, "db Exet Query categorie 3")
		for rows.Next() {
			var categ string
			err = rows.Scan(&categ)
			CheckErr(err, "row categorie 3")
			CategoSQL[index] = categ
			index++
		}
	}
	return CategoSQL
}

// Envoie les données du post dans le db (tab sql)
func PostsComment(r *http.Request, user_id string, database *sql.DB) string {
	// Récupération des informations du json
	var addComment data.Createpost
	err := json.NewDecoder(r.Body).Decode(&addComment)
	CheckErr(err, "Erreur de décodage JSON viewpost")
	// fmt.Println("addComment : ", addComment)

	stmt, err := database.Prepare("INSERT INTO Comments(Post_ID,User_ID,Date,Content) Values(?,?,?,?)")
	CheckErr(err, "Y'a une merde ici: INSERTion img + file")
	stmt.Exec(addComment.PostID, user_id, GetTime(), addComment.Comment)

	return addComment.PostID
}
