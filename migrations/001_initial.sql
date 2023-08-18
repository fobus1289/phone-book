CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY,
	login TEXT UNIQUE NOT NULL CHECK(length(login) >= 4),
	password TEXT CHECK(length(password) = 60),
	name TEXT,
	age INTEGER
);

CREATE TABLE IF NOT EXISTS phones (
	id INTEGER PRIMARY KEY,
	phone TEXT UNIQUE NOT NULL CHECK(length(phone) = 12),
	is_fax BOOLEAN,
	description TEXT,
	user_id INTEGER,
	FOREIGN KEY (user_id) REFERENCES users(id)
);
