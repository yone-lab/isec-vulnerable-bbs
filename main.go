package main

import (
	"database/sql"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"io"
	"net/http"
)

type User struct {
	id   string `db:"id"`
	pass string `db:"pass"`
	name string `db:"display_name"`
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Debug = true
	e.Renderer = t

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.GET("/", func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			return err
		}

		if id := sess.Values["id"]; id != nil {
			return c.Render(http.StatusOK, "bbs", nil)
		}

		return c.Render(http.StatusOK, "index", nil)
	})

	e.GET("/login", func(c echo.Context) error {
		return c.Render(http.StatusOK, "login", nil)
	})

	e.POST("/login", func(c echo.Context) error {
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

		id, pass := params.Get("id"), params.Get("pass")

		rows, err := db.Query("SELECT * FROM users WHERE id = ? AND pass = ?", id, pass)

		// TODO: better error handling
		if err != nil {
			return c.Render(http.StatusNotAcceptable, "login", nil)
		}

		var user User
		if rows.Scan(user) != nil && !rows.Next() {
			sess.Values["id"] = user.id
			sess.Save(c.Request(), c.Response())
			return c.Redirect(http.StatusSeeOther, "/")
		}

		return c.Render(http.StatusNotAcceptable, "login", nil)
	})

	e.GET("/signup", func(c echo.Context) error {
		return c.Render(http.StatusOK, "signup", nil)
	})

	e.POST("/signup", func(c echo.Context) error {
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

		id, pass, name := params.Get("id"), params.Get("pass"), params.Get("name")

		_, err = db.Exec("INSERT INTO users VALUES (?, ?, ?)", id, pass, name)

		// TODO: better error handling
		if err != nil {
			return c.Render(http.StatusNotAcceptable, "signup", nil)
		}

		sess.Values["id"] = id
		sess.Save(c.Request(), c.Response())

		return c.Redirect(http.StatusSeeOther, "/")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
