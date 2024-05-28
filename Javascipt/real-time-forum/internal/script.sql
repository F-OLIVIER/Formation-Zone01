-- database: forum.db
-- Initialisation de la db :
-- 1: sqlite3 ./internal/forum.db
-- 2: .databases
-- 3: .quit
-- run query
CREATE TABLE IF NOT EXISTS Users (
    ID INTEGER PRIMARY KEY,
    uuid INTEGER,
    Email VARCHAR(150) NOT NULL,
    Password VARCHAR(150) NOT NULL,
    Username VARCHAR(150) NOT NULL,
    Age INTEGER,
    Gender VARCHAR(150),
    FirstName VARCHAR(150),
    LastName VARCHAR(150),
    Photo TEXT,
    NewUserMsg TEXT DEFAULT ""
);

CREATE TABLE IF NOT EXISTS Category (ID INTEGER PRIMARY KEY, Name VARCHAR(150));

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
    FOREIGN KEY (Category_ID1) REFERENCES Category (ID),
    FOREIGN KEY (Category_ID2) REFERENCES Category (ID),
    FOREIGN KEY (Category_ID3) REFERENCES Category (ID),
    FOREIGN KEY (User_ID) REFERENCES Users (ID)
);

CREATE TABLE IF NOT EXISTS Comments (
    ID INTEGER PRIMARY KEY,
    Post_ID INTEGER,
    User_ID INTEGER,
    Date TEXT,
    Content TEXT,
    FOREIGN KEY (Post_ID) REFERENCES Posts (ID),
    FOREIGN KEY (User_ID) REFERENCES Users (ID)
);

CREATE TABLE IF NOT EXISTS Socials (
    ID INTEGER PRIMARY KEY,
    User_ID INTEGER,
    Comment_ID INTEGER,
    Post_ID INTEGER,
    Socials INTEGER,
    FOREIGN KEY (User_ID) REFERENCES Users (ID),
    FOREIGN KEY (Comment_ID) REFERENCES Comments (ID) FOREIGN KEY (Post_ID) REFERENCES Posts (ID)
);

CREATE TABLE IF NOT EXISTS Chat (
    ID INTEGER PRIMARY KEY,
    User_ID_sender INTEGER,
    User_ID_receiver INTEGER,
    Time_stamp VARCHAR(50),
    Msg TEXT,
    FOREIGN KEY (User_ID_sender) REFERENCES Users (ID),
    FOREIGN KEY (User_ID_receiver) REFERENCES Users (ID)
);