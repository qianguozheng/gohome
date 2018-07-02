package main

import (
	"flag"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	portPtr := flag.String("port", ":80", "port to serve the service")

	flag.Parse()

	fmt.Println("port=", *portPtr)
	fmt.Println("goHome standalone web server")
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	homeCtx := NewHomeCtx()

	e.GET("/", homeCtx.Handle)
	e.GET("/home", homeCtx.Handle)
	//e.POST("/wx", wechatCtx.HandlePost)
	e.Logger.Fatal(e.Start(*portPtr))
}

type HomeCtx struct{}

func NewHomeCtx() *HomeCtx {
	home := HomeCtx{}
	return &home
}

func (home HomeCtx) Handle(c echo.Context) error {
	//return c.String(http.StatusOK, "")
	return c.File("html/index.html")
}
