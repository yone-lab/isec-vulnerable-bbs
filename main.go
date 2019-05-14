package main

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"html/template"
	"io"
	"net/http"
)

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

		if val := sess.Values["id"]; val != nil {
			c.Render(http.StatusOK, "bulletin", nil)
		}

		return c.Render(http.StatusOK, "index", nil)
	})

	e.GET("/login", func(c echo.Context) error {
		return c.Render(http.StatusOK, "login", nil)
	})

	e.GET("/signup", func(c echo.Context) error {
		return c.Render(http.StatusOK, "signup", nil)
	})

	e.POST("/signup", func(c echo.Context) error {
		_, err := c.FormParams()
		if err != nil {
			return err
		}

		//id, pass, name := params.Get("id"), params.Get("pass"), params.Get("name")

		return nil
	})

	e.Logger.Fatal(e.Start(":8080"))
}
