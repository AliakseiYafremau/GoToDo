package sqlite_repositories

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/AliakseiYafremau/GoToDo/errors"
)

func NewSqliteDB() (*sql.DB, error) {
	database, err := ConnectSQLiteDB("sqlite.db")

	if err != nil {
		return nil, err
	}

	return database, nil
}

func ConnectSQLiteDB(path string) (*sql.DB, error) {
	database, err := sql.Open("sqlite3", path)

	if err != nil {
		return nil, errors.ErrConnectionFailed
	}

	return database, nil
}
