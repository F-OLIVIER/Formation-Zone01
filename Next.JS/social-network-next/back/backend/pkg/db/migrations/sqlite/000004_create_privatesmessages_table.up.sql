	CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		sender_id INTEGER NOT NULL,
		receiver_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		date TEXT NOT NULL,
		FOREIGN KEY(sender_id) REFERENCES users(id),
		FOREIGN KEY(receiver_id) REFERENCES users(id)
	);

	CREATE TABLE IF NOT EXISTS chats (
		id_one INTEGER NOT NULL,
		id_two INTEGER NOT NULL,
		time INTEGER NOT NULL,
		FOREIGN KEY(id_one) REFERENCES users(id),
		FOREIGN KEY(id_two) REFERENCES users(id)
	);

CREATE TABLE IF NOT EXISTS groupmessages (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	sender_id INTEGER NOT NULL,
	group_id INTEGER NOT NULL,
	content TEXT NOT NULL,
	date TEXT NOT NULL,
	FOREIGN KEY(sender_id) REFERENCES users(id)
);