CREATE TABLE IF NOT EXISTS NOTIFICATIONS
(
    IDNotif
    INTEGER
    PRIMARY
    KEY,
    IDPrivateMessage
    INTEGER
    DEFAULT
    0,
    IDPost
    INTEGER
    DEFAULT
    0,
    IDComment
    INTEGER
    DEFAULT
    0,
    IDGroup
    INTEGER
    DEFAULT
    0,
    IDEvent
    INTEGER
    DEFAULT
    0,
    IDFollow
    INTEGER
    DEFAULT
    0,
    UserID_Receiver
    INTEGER,
    Date
    TEXT,
    FOREIGN
    KEY
(
    IDPrivateMessage
) REFERENCES PRIVATEMESSAGE
(
    IDMessage
),
    FOREIGN KEY
(
    IDPost
) REFERENCES POST
(
    IDPost
),
    FOREIGN KEY
(
    IDComment
) REFERENCES COMMENT
(
    IDComment
),
    FOREIGN KEY
(
    IDGroup
) REFERENCES LISTGROUPS
(
    IDGroup
),
    FOREIGN KEY
(
    IDEvent
) REFERENCES EVENTGROUPS
(
    IDEvent
),
    FOREIGN KEY
(
    UserID_Receiver
) REFERENCES USERS
(
    ID
)
    );