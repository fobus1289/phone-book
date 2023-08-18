package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func MustOpen(name string) *sql.DB {
	db, err := Open(name)

	if err != nil {
		panic(err)
	}

	return db
}

func Open(name string) (*sql.DB, error) {
	return sql.Open("sqlite3", name)
}
