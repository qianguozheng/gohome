package main

import (
	"flag"
	"fmt"

	"./admin"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	portPtr := flag.String("port", ":80", "port to serve the service")

	flag.Parse()

	fmt.Println("port=", *portPtr)
	fmt.Println("goHome standalone web server")
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	//Homepage
	homeCtx := admin.NewHomeCtx()
	e.GET("/", homeCtx.Handle)
	e.GET("/home", homeCtx.Handle)

	//Admin Login
	loginCtx := admin.NewLoginCtx()
	e.GET("/login", loginCtx.Handle)
	e.POST("/login", loginCtx.HandlePost)

	adminCtx := admin.NewAdminCtx()
	e.GET("/admin/", adminCtx.Handle)

	e.Logger.Fatal(e.Start(*portPtr))
}
