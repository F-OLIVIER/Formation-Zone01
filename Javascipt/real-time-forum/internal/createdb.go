package data

import (
	"database/sql"
	"fmt"
	"os"
)

func Createdb() {
	database, err := sql.Open("sqlite3", "./internal/forum.db")
	CheckErr(err, "1- Err open db in Createdb")

	scriptSQL, err := os.ReadFile("./internal/script.sql")
	CheckErr(err, "open file script.sql")
	_, err = database.Exec(string(scriptSQL))
	CheckErr(err, "2- Err open db in Createdb")

	// Insertion des éléments de base dans la table si elle n'existe pas
	idUsernameDelete := database.QueryRow("SELECT ID FROM Users WHERE Username = 'USER_DELETE'").Scan()
	if idUsernameDelete == sql.ErrNoRows {
		insertUser_Deleted := `INSERT INTO Users (Photo, Username, Email, Password) VALUES ('user_delete.jpg','USER_DELETE','','Fauxmotdepassepoureviteruntrucvidepourqu''unutilisateurmalveillantnepuissepasseconnecteraussisimplementquecela');`
		_, err = database.Exec(insertUser_Deleted)
		CheckErr(err, "3- Err insert db 1 in Createdb")
	}

	idCategorie := database.QueryRow("SELECT ID FROM Category WHERE Name = 'Autres'").Scan()
	if idCategorie == sql.ErrNoRows {
		insertBaseCat := `
		INSERT INTO Category (Name) VALUES ('Boissons');
		INSERT INTO Category (Name) VALUES ('Sauces');
		INSERT INTO Category (Name) VALUES ('Entrées');
		INSERT INTO Category (Name) VALUES ('Légumes');
		INSERT INTO Category (Name) VALUES ('Poissons');
		INSERT INTO Category (Name) VALUES ('Viandes');
		INSERT INTO Category (Name) VALUES ('Desserts');
		INSERT INTO Category (Name) VALUES ('Autres');`
		_, err = database.Exec(insertBaseCat)
		CheckErr(err, "4- Err insert db 1 in Createdb")
	}

	database.Close()
}

func CheckErr(err error, str string) {
	if err != nil {
		fmt.Printf("\n___________________\nERROR : %v\n%v\n", str, err)
	}
}
