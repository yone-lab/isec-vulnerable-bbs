package main

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/proelbtn/kosen-isec-lab-vulnerable-chat-app/controllers"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// refs: https://stackoverflow.com/questions/18175630/go-template-executetemplate-include-html
func noescape(str string) template.HTML {
	return template.HTML(str)
}

func main() {
	t := &Template{
		templates: template.Must(template.New("main").Funcs(template.FuncMap{"noescape": noescape}).ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Debug = true
	e.Renderer = t

	e.Use(middleware.Logger())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.GET("/", controllers.IndexGetHandler)
	e.GET("/login", controllers.LoginGetHandler)
	e.POST("/login", controllers.LoginPostHandler)
	e.GET("/signup", controllers.SignupGetHandler)
	e.POST("/signup", controllers.SignupPostHandler)
	e.POST("/post", controllers.PostPostHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
