package controllers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/yone-lab/isec-vulnerable-bbs/models"
	"net/http"
)

func LoginGetHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "login", nil)
}

func LoginPostHandler(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}

	params, err := c.FormParams()
	if err != nil {
		return err
	}

	id, pass := params.Get("id"), params.Get("pass")
	user, err := models.SearchUser(id)

	// TODO: better error handling
	if err != nil {
		return c.Render(http.StatusNotAcceptable, "login", nil)
	}

	println(user.Pass, pass)

	if user.Pass == pass {
		sess.Values["id"] = user.ID
		sess.Save(c.Request(), c.Response())
		return c.Redirect(http.StatusSeeOther, "/")
	}

	return c.Render(http.StatusNotAcceptable, "login", nil)
}
