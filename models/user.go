package models

import (
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id   string `db:"id"`
	Pass string `db:"pass"`
	Name string `db:"display_name"`
}

func CreateUser(id, pass, name string) error {
	db, err := getDatabase()
	if err != nil {
		return err
	}

	// TODO: raw password
	_, err = db.Exec("INSERT INTO users (id, pass, display_name) VALUES (?, ?, ?)", id, pass, name)

	return err
}

func SearchUser(id string) (*User, error) {
	var user User

	db, err := getDatabase()
	if err != nil {
		return nil, err
	}

	row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	row.Scan(user)

	return &user, nil
}
