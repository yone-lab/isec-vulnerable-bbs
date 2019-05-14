package models

import (
	"database/sql"
)

func getDatabase() (*sql.DB, error) {
	return sql.Open("sqlite3", "file:database.sqlite3")
}
