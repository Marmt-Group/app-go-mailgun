package main

import (
	"fmt"
	"github.com/labstack/echo"
	"html/template"
	"io"
	"net/http"
	"os"
)

// Template html/template renderer
type Template struct {
	templates *template.Template
}

// Render renders a template document
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// POST Route function
func upload(c echo.Context) error {
	// Read form fields
	// key := c.FormValue("mailgunKey")

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	//TODO: need to create batch process function to pass csv data to

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully.</p>", file.Filename))
}

// GET Route function
func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "World")
}

func main() {
	// mailGunKey := os.Getenv("MAILGUN_KEY")
	port := os.Args[1]

	e := echo.New()
	e.Static("/static", "assets")
	
	t := &Template{
		templates: template.Must(template.ParseGlob("./views/*.html")),
	}
	e.Renderer = t

	// Routes
	e.GET("/", index)
	e.POST("/upload", upload)

	if port == "dev" {
		e.Logger.Fatal(e.Start(":8000"))
	} else {
		e.Logger.Fatal(e.Start(":8080"))
	}

}
