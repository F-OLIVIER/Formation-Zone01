CREATE TABLE IF NOT EXISTS USERS
(
    ID
    INTEGER
    PRIMARY
    KEY,
    Email
    VARCHAR
(
    150
) NOT NULL,
    Password VARCHAR
(
    150
) NOT NULL,
    FirstName VARCHAR
(
    150
) NOT NULL,
    LastName VARCHAR
(
    150
) NOT NULL,
    DateOfBirth DATE NOT NULL,
    Avatar BLOB,
    Nickname VARCHAR
(
    150
),
    AboutMe VARCHAR
(
    500
),
    PrivateProfile INT -- 0 True, 1 False
    );


CREATE TABLE IF NOT EXISTS SESSIONS (
    UserID INT,
    SessionToken TEXT
);

