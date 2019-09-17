package main

import (
	"html/template"
	"io"
	"net/http"
	"github.com/labstack/echo"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type Template struct {
    templates *template.Template
}

// Render renders a template document
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

// Route function
func Hello(c echo.Context) error {
    return c.Render(http.StatusOK, "hello", "World")
}

func main() {
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("./public/views/*.html")),
	}
	e.Renderer = t

	// Routes
	e.GET("/", Hello)

	e.Logger.Fatal(e.Start(":8080"))
}