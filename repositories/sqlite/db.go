package sqlite_repositories

import (
	"database/sql"
	"log"
)

func NewSqliteDB() *sql.DB {
	database := ConnectSQLiteDB("sqlite.db")

	return database
}

func ConnectSQLiteDB(path string) *sql.DB {
	database, err := sql.Open("sqlite3", path)

	if err != nil {
		log.Fatal(err)
	}

	return database
}
