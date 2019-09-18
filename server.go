package main

import (
	"html/template"
	"io"
	"os"
	"net/http"
	"github.com/labstack/echo"
)

// Template html/template renderer
type Template struct {
    templates *template.Template
}

// Render renders a template document
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

// Route function
func Index(c echo.Context) error {
    return c.Render(http.StatusOK, "index", "World")
}

func main() {
	// mailGunKey := os.Getenv("MAILGUN_KEY")
	port := os.Args[1]

	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("./public/views/*.html")),
	}
	e.Renderer = t

	// Routes
	e.GET("/", Index)

	
	if port == "dev" {
		e.Logger.Fatal(e.Start(":8000"))
	} else {
		e.Logger.Fatal(e.Start(":8080"))
	}
	
}