package models

type User struct {
	Id   string `db:"id"`
	Pass string `db:"pass"`
	Name string `db:"display_name"`
}
