package models

import (
	"time"
)

type Post struct {
	ID        int
	UID       string
	Content   string
	CreatedAt time.Time
}

func CreatePost(uid, content string) error {
	db, err := getDatabase()
	if err != nil {
		return err
	}

	now := time.Now()

	// TODO: raw password
	_, err = db.Exec("INSERT INTO posts (uid, content, created_at) VALUES (?, ?, ?)", uid, content, now)

	return err
}

func GetPosts() (*[]Post, error) {
	var (
		post  Post
		posts []Post
	)

	db, err := getDatabase()
	if err != nil {
		return nil, err
	}

	// TODO: raw password
	rows, err := db.Query("SELECT id, uid, content, created_at FROM posts")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&post.ID, &post.UID, &post.Content, &post.CreatedAt)
		posts = append(posts, post)
	}

	return &posts, nil
}
