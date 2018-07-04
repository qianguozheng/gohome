package admin

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type LoginCtx struct{}

func NewLoginCtx() *LoginCtx {
	login := LoginCtx{}
	return &login
}

func (login LoginCtx) Handle(c echo.Context) error {
	fmt.Println("admin login page")
	return c.File("html/login.html")
}

func (login LoginCtx) HandlePost(c echo.Context) error {

	//userName := c.QueryParam("username")
	//password := c.QueryParam("password")

	//Read data from Form

	userName := c.FormValue("username")
	password := c.FormValue("password")

	fmt.Printf("username=%s, password=%s\n", userName, password)

	//TODO: store cookie into database
	cookie := new(http.Cookie)
	cookie.Name = "sessionId"
	cookie.Value = userName
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	return c.Redirect(http.StatusFound, "/admin/")
}
