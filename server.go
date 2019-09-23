package main

import (
	"context"
	"encoding/csv"
	// "fmt"
	"github.com/mailgun/mailgun-go/v3"
	"github.com/labstack/echo/v4"
  	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Template html/template renderer
type Template struct {
	templates *template.Template
}

type Item struct {
	email string
	company string
}

type Data struct {
	Record []Item
}

// Render renders a template document
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// POST Route Mailgun batch send function
func sendTemplateMessage(c echo.Context) error {

	mg_domain := os.Getenv("MAILGUN_DOMAIN") 
	mg_key := os.Getenv("MAILGUN_KEY")

	// Source file
	file, err := c.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	// Read the file
	r := csv.NewReader(src)
	index := 0

	// TODO: set up secrets in Cloud Run

	// Setup Mailgun message
	mg := mailgun.NewMailgun(mg_domain, mg_key)
	m := mg.NewMessage(
		"David J. Davis <davidjamesdavis.djd@gmail.com>",
		"Hey %recipient.first%",
		"I would like to discuss and opportunity",
		"admin@marmt.io",
	) // IMPORTANT: No To:-field recipients!

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		m.AddRecipientAndVariables(record[8], map[string]interface{}{ // record indexes according to your csv records
			"first": record[3], // record indexes according to your csv records
			"id":    index,
		})

		index++
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, id, err := mg.Send(ctx, m)

	if err != nil {
		return c.String(500, err.Error())
	} else {
		return c.String(http.StatusOK, id)
	}
}

// GET Route function
func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "")
}

func main() {
	// mailGunKey := os.Getenv("MAILGUN_KEY")

	e := echo.New()
	e.Use(middleware.Logger()) // remove in production
	e.Static("/static", "assets")

	t := &Template{
		templates: template.Must(template.ParseGlob("./views/*.html")),
	}
	e.Renderer = t

	// Routes
	e.GET("/", index)
	e.POST("/upload", sendTemplateMessage)

	e.Logger.Fatal(e.Start(":8080"))

}
