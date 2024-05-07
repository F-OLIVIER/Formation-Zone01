CREATE TABLE IF NOT EXISTS FOLLOWERS
(
    ID
    INTEGER
    PRIMARY
    KEY,
    UserID_Following
    INTEGER,
    UserID_Follower
    INTEGER,
    DateFollow
    TEXT,
    ValidateFollow
    INTEGER, -- 0 True, 1 False
    FOREIGN
    KEY
(
    UserID_Following
) REFERENCES USERS
(
    ID
),
    FOREIGN KEY
(
    UserID_Follower
) REFERENCES USERS
(
    ID
)
    );
