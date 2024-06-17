-- Step 1: Create the table
CREATE TABLE IF NOT EXISTS LISTGROUPS (
    IDGroup INTEGER PRIMARY KEY AUTOINCREMENT,
    NameGroup VARCHAR(150),
    UserID_Creator INTEGER,
    AboutUs VARCHAR(500),
    Image BLOB,
    FOREIGN KEY(UserID_Creator) REFERENCES USERS(ID)
);

-- Step 2: Insert a dummy row with IDGroup 10000
INSERT INTO LISTGROUPS (IDGroup, NameGroup, UserID_Creator, AboutUs, Image)
VALUES (10000, 'dummy', 0, 'dummy', NULL);

-- Step 3: Delete the dummy row
DELETE FROM LISTGROUPS WHERE IDGroup = 10000;
    

CREATE TABLE IF NOT EXISTS MEMBERSGROUPS
(
    ID
    INTEGER
    PRIMARY
    KEY,
    IDGroup
    INTEGER,
    UserID
    INTEGER,
    ValidateInvite
    INTEGER, -- 0 True, 1 False
    FOREIGN
    KEY
(
    IDGroup
) REFERENCES LISTGROUPS
(
    IDGroup
),
    FOREIGN KEY
(
    UserID
) REFERENCES USERS
(
    ID
)
    );

CREATE TABLE IF NOT EXISTS CHATSGROUPS
(
    ID
    INTEGER
    PRIMARY
    KEY,
    IDGroup
    INTEGER,
    Date
    TEXT,
    UserID_Sender
    INTEGER,
    ContentMessage
    TEXT,
    FOREIGN
    KEY
(
    IDGroup
) REFERENCES LISTGROUPS
(
    IDGroup
),
    FOREIGN KEY
(
    UserID_Sender
) REFERENCES USERS
(
    ID
)
    );

CREATE TABLE IF NOT EXISTS EVENTGROUPS (
    IDEvent INTEGER PRIMARY KEY,
    IDGroup INTEGER,
    Date TEXT,
    Title VARCHAR(150),
    Come INTEGER,
    NotCome INTEGER,
    Description VARCHAR(500),
    UserID_Sender INTEGER,
    FOREIGN KEY (IDGroup) REFERENCES LISTGROUPS(IDGroup),
    FOREIGN KEY (UserID_Sender) REFERENCES USERS(ID)
);

CREATE TABLE IF NOT EXISTS RESPONSEEVENTGROUPS
(
    ID
    INTEGER
    PRIMARY
    KEY,
    IDEvent
    INTEGER,
    UserID
    INTEGER,
    Option
    INTEGER, -- 0 Not Answered﻿/1 YES/2 NO
    FOREIGN
    KEY
(
    IDEvent
) REFERENCES EVENTGROUPS
(
    IDEvent
),
    FOREIGN KEY
(
    UserID
) REFERENCES USERS
(
    ID
)
    );

CREATE TABLE IF NOT EXISTS coming_event (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    event_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    come BOOL,
    notcome BOOL,
    FOREIGN KEY (event_id) REFERENCES EVENTGROUPS(IDEvent),
    FOREIGN KEY (user_id) REFERENCES USERS(ID)
);