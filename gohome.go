package main

import (
	"flag"
	"fmt"

	"io"

	"./admin"

	"html/template"

	"./model"
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
		templates: template.Must(template.ParseGlob("html/*.html")),
	}
	e.Renderer = render

	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Static("/", "static/theme")

	//Init Database
	model.Database = model.InitDB("test.db")
	model.Migrate(model.Database)

	//Homepage
	homeCtx := admin.NewHomeCtx()
	e.GET("/index", homeCtx.Handle)
	e.GET("/", homeCtx.HandleTheme)

	//Admin Login
	loginCtx := admin.NewLoginCtx()
	e.GET("/login", loginCtx.Handle)
	e.POST("/login", loginCtx.HandlePost)

	adminCtx := admin.NewAdminCtx()
	e.GET("/admin/", adminCtx.Handle)

	//Upload file
	e.GET("/admin/upload", adminCtx.HandleUpload)
	e.POST("/admin/upload", adminCtx.HandleUpload)

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
