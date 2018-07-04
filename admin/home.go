package admin

import (
	"github.com/labstack/echo"
)

type HomeCtx struct{}

func NewHomeCtx() *HomeCtx {
	home := HomeCtx{}
	return &home
}

func (home HomeCtx) Handle(c echo.Context) error {
	//return c.String(http.StatusOK, "")
	return c.File("html/index.html")
}
