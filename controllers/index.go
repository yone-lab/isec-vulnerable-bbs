package controllers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/proelbtn/vulnerable-bbs/models"
	"net/http"
)

type indexGetHandlerParams struct {
	User  *models.User
	Posts *[]models.Post
}

/*
 If the user logged in, they are navigated to bss page.
 Otherwise, they are navigated to top page.
*/
func IndexGetHandler(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}

	if id, ok := sess.Values["id"]; ok {
		user, _ := models.SearchUser(id.(string))
		posts, _ := models.GetPosts()
		return c.Render(http.StatusOK, "bbs", indexGetHandlerParams{user, posts})
	}

	return c.Render(http.StatusOK, "index", nil)
}
