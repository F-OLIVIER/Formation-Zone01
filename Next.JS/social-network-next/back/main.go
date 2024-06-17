package main

import (
	"social/backend"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	backend.StartServer()
}

/*func main() {
	// Ouvrir la connexion à la base de données SQLite
	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
		return
	}
	defer db.Close()

	// Exécuter la requête SQL pour récupérer tous les enregistrements de la table USERS
	rows, err := db.Query("SELECT * FROM USERS")
	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la requête:", err)
		return
	}
	defer rows.Close()

	// Parcourir les résultats et afficher les données
	for rows.Next() {
		var id int
		var email string
		var password string
		var firstName string
		var lastName string
		var dateOfBirth time.Time
		var avatar []byte
		var username string
		var aboutMe sql.NullString
		var privateProfile sql.NullInt64

		err = rows.Scan(&id, &email, &password, &firstName, &lastName, &dateOfBirth, &avatar, &username, &aboutMe, &privateProfile)
		if err != nil {
			fmt.Println("Erreur lors de la lecture des données de la ligne:", err)
			return
		}

		// Gérer le cas où PrivateProfile est NULL
		var privateProfileValue int
		if privateProfile.Valid {
			privateProfileValue = int(privateProfile.Int64)
		} else {
			privateProfileValue = 0 // Ou une autre valeur par défaut appropriée
		}

		// Gérer le cas où AboutMe est NULL
		var aboutMeValue string
		if aboutMe.Valid {
			aboutMeValue = aboutMe.String
		} else {
			aboutMeValue = "NULL"
		}

		fmt.Printf("ID: %d, Email: %s, Password: %s, FirstName: %s, LastName: %s, DateOfBirth: %s, Avatar: %s, Username: %s, AboutMe: %s, PrivateProfile: %d\n", id, email, password, firstName, lastName, dateOfBirth, avatar, username, aboutMeValue, privateProfileValue)
	}

	// Vérifier s'il y a eu des erreurs pendant l'itération des résultats
	err = rows.Err()
	if err != nil {
		fmt.Println("Erreur pendant l'itération des résultats:", err)
		return
	}
}*/
