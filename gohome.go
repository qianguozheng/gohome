package main

import (
	"flag"
	"fmt"

	"io"

	"./admin"

	"html/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	portPtr := flag.String("port", ":80", "port to serve the service")

	flag.Parse()

	fmt.Println("port=", *portPtr)
	fmt.Println("goHome standalone web server")
	e := echo.New()

	render := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("html/admin.html")),
	}
	e.Renderer = render

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

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	//	if _, isMap := data.(map[string]interface{}); isMap {
	//		//viewContext["reverse"] = c.Echo().Reverse
	//		fmt.Println("reverse...")
	//	}
	return t.templates.ExecuteTemplate(w, name, data)
}
