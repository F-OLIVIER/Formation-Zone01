CREATE TABLE IF NOT EXISTS LISTGROUPS
(
    IDGroup
    INTEGER
    PRIMARY
    KEY,
    NameGroup
    VARCHAR
(
    150
),
    UserID_Creator INTEGER,
    AboutUs VARCHAR
(
    500
),
    Image BLOB,
    FOREIGN KEY
(
    UserID_Creator
) REFERENCES USERS
(
    ID
)
    );

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

CREATE TABLE IF NOT EXISTS EVENTGROUPS
(
    IDEvent
    INTEGER
    PRIMARY
    KEY,
    IDGroup
    INTEGER,
    Date
    TEXT,
    Title
    VARCHAR
(
    150
),
    Description VARCHAR
(
    500
),
    UserID_Sender INTEGER,
    FOREIGN KEY
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
