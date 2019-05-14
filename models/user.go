package models

import (
	"errors"
)

type User struct {
	ID   string
	Pass string
	Name string
}

func validateUser(id, pass, name string) error {
	var flag = true

	flag = flag && id != ""
	flag = flag && pass != ""
	flag = flag && name != ""

	if !flag {
		return errors.New("validate is failed")
	}

	return nil
}

func CreateUser(id, pass, name string) error {
	db, err := getDatabase()
	if err != nil {
		return err
	}

	err = validateUser(id, pass, name)
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

	if id == "" {
		return nil, errors.New("no id specified")
	}

	rows, err := db.Query("SELECT id, pass, display_name FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, errors.New("no users found")
	}
	rows.Scan(&user.ID, &user.Pass, &user.Name)

	if rows.Next() {
		return nil, errors.New("multiple users found")
	}

	return &user, nil
}
