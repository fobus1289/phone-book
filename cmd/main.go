package main

import (
	"github.com/fobus1289/phone-book/internal/common"
	"github.com/fobus1289/phone-book/internal/repository/database"
	"github.com/fobus1289/phone-book/internal/router"
	"github.com/labstack/echo/v4"
)

const INITIAL_SQL = `
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
`

func main() {

	db := database.MustOpen(common.Database())

	server := echo.New()

	if _, err := db.Exec(INITIAL_SQL); err != nil {
		server.Logger.Fatal(err)
	}

	router.RegisterHandler(server, db)

	server.Logger.Fatal(server.Start(":8081"))
}
