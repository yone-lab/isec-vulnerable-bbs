package models

type User struct {
	ID   string
	Pass string
	Name string
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

	row := db.QueryRow("SELECT id, pass, display_name FROM users WHERE id = ?", id)
	row.Scan(&user.ID, &user.Pass, &user.Name)

	return &user, nil
}
