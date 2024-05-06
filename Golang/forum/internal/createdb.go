package data

import (
	"database/sql"
	"fmt"
)

func Createdb() {

	database, err := sql.Open("sqlite3", "./internal/forum.db")
	if err != nil {
		fmt.Println("Err open db in Createdb : ", err)
	}

	// Création des tables si elle n'existe pas
	createBaseTable := `CREATE TABLE IF NOT EXISTS Users (
		ID INTEGER PRIMARY KEY,
		uuid INTEGER, 
		Photo TEXT,
		Username VARCHAR(150) NOT NULL,
		Email VARCHAR(150) NOT NULL, 
		Password VARCHAR(150) NOT NULL
	);
	
	CREATE TABLE IF NOT EXISTS Category (
		ID INTEGER PRIMARY KEY, 
		Name VARCHAR(150)
	);
	
	CREATE TABLE IF NOT EXISTS Posts (
		ID INTEGER PRIMARY KEY,
		Category_ID1 INTEGER, 
		Category_ID2 INTEGER, 
		Category_ID3 INTEGER, 
		User_ID INTEGER, 
		Date TEXT, 
		Titre TEXT,
		Content TEXT, 
		LienImage TEXT,
		FichierJoint TEXT,
		FOREIGN KEY(Category_ID1) REFERENCES Category(ID),
		FOREIGN KEY(Category_ID2) REFERENCES Category(ID),
		FOREIGN KEY(Category_ID3) REFERENCES Category(ID),
		FOREIGN KEY(User_ID) REFERENCES Users(ID)
	);
	
	CREATE TABLE IF NOT EXISTS Comments (
		ID INTEGER PRIMARY KEY, 
		Post_ID INTEGER, 
		User_ID INTEGER, 
		Date TEXT, 
		Content TEXT, 
		FOREIGN KEY(Post_ID) REFERENCES Posts(ID), 
		FOREIGN KEY(User_ID) REFERENCES Users(ID)
	);
	
	CREATE TABLE IF NOT EXISTS Socials (
		ID INTEGER PRIMARY KEY, 
		User_ID INTEGER, 
		Comment_ID INTEGER, 
		Post_ID INTEGER,
		Socials INTEGER, 
		FOREIGN KEY(User_ID) REFERENCES Users(ID), 
		FOREIGN KEY(Comment_ID) REFERENCES Comments(ID)
		FOREIGN KEY(Post_ID) REFERENCES Posts(ID)
	);`

	_, err = database.Exec(createBaseTable)
	if err != nil {
		fmt.Println("Err create db in Createdb : ", err)
	}

	// Insertion des éléments de base dans la table si elle n'existe pas
	idUsernameDelete := database.QueryRow("SELECT ID FROM Users WHERE Username = 'USER_DELETE'").Scan()
	if idUsernameDelete == sql.ErrNoRows {
		insertUser_Deleted := `INSERT INTO Users (Photo, Username, Email, Password) VALUES ('user_delete.jpg','USER_DELETE','','Fauxmotdepassepoureviteruntrucvidepourqu''unutilisateurmalveillantnepuissepasseconnecteraussisimplementquecela');`
		_, err = database.Exec(insertUser_Deleted)
		if err != nil {
			fmt.Println("Err insert db 1 in Createdb :  ", err)
		}
	}

	idCategorie := database.QueryRow("SELECT ID FROM Category WHERE Name = 'Autres'").Scan()
	if idCategorie == sql.ErrNoRows {
		insertBaseCat := `INSERT INTO Category (Name) VALUES ('Boissons');
		INSERT INTO Category (Name) VALUES ('Sauces');
		INSERT INTO Category (Name) VALUES ('Entrées');
		INSERT INTO Category (Name) VALUES ('Légumes');
		INSERT INTO Category (Name) VALUES ('Poissons');
		INSERT INTO Category (Name) VALUES ('Viandes');
		INSERT INTO Category (Name) VALUES ('Desserts');
		INSERT INTO Category (Name) VALUES ('Autres');`
		_, err = database.Exec(insertBaseCat)
		if err != nil {
			fmt.Println("Err insert db 2 in Createdb : ", err)
		}
	}

	database.Close()
}
