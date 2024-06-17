CREATE TABLE IF NOT EXISTS POST
(
    IDPost
    INTEGER
    PRIMARY
    KEY AUTOINCREMENT,
    UserID
    INTEGER,
    Title
    VARCHAR
(
    150
),
    PostContent TEXT,
    Date TEXT,
    Image BLOB,
    Private INTEGER, -- 0 Private, 1 Public, 2 Select user
    Likes INTEGER, 
    NbComments INTEGER,
    FOREIGN KEY
(
    UserID
) REFERENCES USERS
(
    ID
)
    );
CREATE TABLE IF NOT EXISTS COMMENT
(
    IDComment
    INTEGER
    PRIMARY
    KEY AUTOINCREMENT,
    IDPost
    INTEGER,
    UserID
    INTEGER,
    CommentContent
    TEXT,
    Date
    TEXT,
    Image
    BLOB,
    FOREIGN
    KEY
(
    IDPost
) REFERENCES POST
(
    IDPost
),
    FOREIGN KEY
(
    UserID
) REFERENCES USERS
(
    ID
)
    );
CREATE TABLE IF NOT EXISTS POSTSELECTUSERS
(
    ID
    INTEGER
    PRIMARY
    KEY,
    IDPost
    INTEGER,
    UserID
    INTEGER,
    FOREIGN
    KEY
(
    IDPost
) REFERENCES POST
(
    IDPost
),
    FOREIGN KEY
(
    UserID
) REFERENCES USERS
(
    ID
)
    );

CREATE TABLE IF NOT EXISTS liked_post (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		post_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		like BOOL,
		FOREIGN KEY(post_id) REFERENCES posts(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	);


CREATE TABLE IF NOT EXISTS POST_GROUP
(
    IDPost
    INTEGER
    PRIMARY
    KEY AUTOINCREMENT,
    UserID
    INTEGER,
    Title
    VARCHAR
(
    150
),
    PostContent TEXT,
    Date TEXT,
    Image BLOB,
    GroupID INTEGER,
    Likes INTEGER, 
    NbComments INTEGER,
    FOREIGN KEY
(
    UserID
) REFERENCES USERS
(
    ID
)
    );
CREATE TABLE IF NOT EXISTS COMMENT_GROUP
(
    IDComment
    INTEGER
    PRIMARY
    KEY AUTOINCREMENT,
    IDPost
    INTEGER,
    UserID
    INTEGER,
    CommentContent
    TEXT,
    Date
    TEXT,
    Image
    BLOB,
    FOREIGN
    KEY
(
    IDPost
) REFERENCES POST
(
    IDPost
),
    FOREIGN KEY
(
    UserID
) REFERENCES USERS
(
    ID
)
    );