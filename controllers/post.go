package controllers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/yone-lab/isec-vulnerable-bbs/models"
	"net/http"
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

	id, ok := sess.Values["id"]
	if !ok {
		return err
	}
	if id == nil {
		return c.Redirect(http.StatusSeeOther, "/")
	}

	content := params.Get("content")

	err = models.CreatePost(id.(string), content)

	// TODO: better error handling
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/")
	}

	return c.Redirect(http.StatusSeeOther, "/")

}
