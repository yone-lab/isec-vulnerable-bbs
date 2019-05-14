package controllers

import (
	"database/sql"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"net/http"
	"time"
)

func PostPostHandler(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}

	params, err := c.FormParams()
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", "file:database.sqlite3")
	if err != nil {
		return err
	}

	id, ok := sess.Values["id"]
	if !ok {
		return err
	}

	content := params.Get("content")
	createdAt := time.Now()

	if id == nil {
		return c.Redirect(http.StatusSeeOther, "/")
	}

	_, err = db.Exec("INSERT INTO posts (uid, content, created_at) VALUES (?, ?, ?)", id, content, createdAt)

	// TODO: better error handling
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/")
	}

	return c.Redirect(http.StatusSeeOther, "/")

}
