-- database: forum.db
-- Initialisation de la db :
-- 1: sqlite3 ./internal/forum.db
-- 2: .databases
-- 3: .quit
-- run query

CREATE TABLE Users(
    ID INTEGER PRIMARY KEY,
    uuid INTEGER, 
    Photo TEXT,
    Username VARCHAR(150) NOT NULL,
    Email VARCHAR(150) NOT NULL, 
    Password VARCHAR(150) NOT NULL
);

CREATE TABLE Category(
    ID INTEGER PRIMARY KEY, 
    Name VARCHAR(150)
);

CREATE TABLE Posts(
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

CREATE TABLE Comments(
    ID INTEGER PRIMARY KEY, 
    Post_ID INTEGER, 
    User_ID INTEGER, 
    Date TEXT, 
    Content TEXT, 
    FOREIGN KEY(Post_ID) REFERENCES Posts(ID), 
    FOREIGN KEY(User_ID) REFERENCES Users(ID)
);

CREATE TABLE Socials(
    ID INTEGER PRIMARY KEY, 
    User_ID INTEGER, 
    Comment_ID INTEGER, 
    Post_ID INTEGER,
    Socials INTEGER, 
    FOREIGN KEY(User_ID) REFERENCES Users(ID), 
    FOREIGN KEY(Comment_ID) REFERENCES Comments(ID)
    FOREIGN KEY(Post_ID) REFERENCES Posts(ID)
);

INSERT INTO Users (Photo, Username, Email, Password) VALUES ('default/user_delete.jpg','USER_DELETE','','Fauxmotdepassepoureviteruntrucvidepourqu''unutilisateurmalveillantnepuissepasseconnecteraussisimplementquecela');

INSERT INTO Category (Name) VALUES ('Boissons');
INSERT INTO Category (Name) VALUES ('Sauces');
INSERT INTO Category (Name) VALUES ('Entrées');
INSERT INTO Category (Name) VALUES ('Légumes');
INSERT INTO Category (Name) VALUES ('Poissons');
INSERT INTO Category (Name) VALUES ('Viandes');
INSERT INTO Category (Name) VALUES ('Desserts');
-- INSERT INTO Posts (Section,User_ID,Titre,Content) VALUES ('Test',0,'Recette Test','Sert à testé la database');

