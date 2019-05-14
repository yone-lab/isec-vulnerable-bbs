package controllers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"net/http"
)

func IndexGetHandler(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}

	if id, ok := sess.Values["id"].(string); ok {
		return c.Render(http.StatusOK, "bbs", id)
	}

	return c.Render(http.StatusOK, "index", nil)
}
