package auth

import "database/sql"

var Db *sql.DB

func CreateMoveTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user TEXT NOT NULL,
	email  UNIQUE TEXT NOT NULL,
	password TEXT NOT NULL
	);`

	if _, err := Db.Exec(query); err != nil {
		return err
	}

	return nil
}

func StartDbConn() error {

	var err error

	Db, err := sql.Open("sqlite3", "mydatabase.db")
	if err != nil {
		return err
	}

	if err = Db.Ping(); err != nil {
		return err
	}
	return nil
}
