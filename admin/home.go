package admin

import (
	"net/http"

	"../model"
	"github.com/labstack/echo"
)

type HomeCtx struct{}

func NewHomeCtx() *HomeCtx {
	home := HomeCtx{}
	return &home
}

func (home HomeCtx) Handle(c echo.Context) error {
	//return c.String(http.StatusOK, "")

	model.InsertCount(model.Database)
	model.UpdateCount(model.Database)
	_, accessNumber := model.QueryCount(model.Database)
	//return c.File("html/index.html")
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"accessNumber": accessNumber,
	})
}
