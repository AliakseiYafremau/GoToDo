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

	err = CreateTables(database)

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

func CreateTables(database *sql.DB) error {
	sql_statement := `
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			text TEXT
	);
	`

	_, err := database.Exec(sql_statement)

	if err != nil {
		return errors.ErrTablesCreationFailed
	}

	return nil
}
