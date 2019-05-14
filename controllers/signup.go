package controllers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/proelbtn/kosen-isec-lab-vulnerable-chat-app/models"
	"net/http"
)

func SignupGetHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "signup", nil)
}

func SignupPostHandler(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}

	params, err := c.FormParams()
	if err != nil {
		return err
	}

	id, pass, name := params.Get("id"), params.Get("pass"), params.Get("name")

	// TODO: raw password
	err = models.CreateUser(id, pass, name)

	// TODO: better error handling
	if err != nil {
		println(err.Error())
		return c.Render(http.StatusNotAcceptable, "signup", nil)
	}

	sess.Values["id"] = id
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusSeeOther, "/")
}
