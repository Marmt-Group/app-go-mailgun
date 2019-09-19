package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/labstack/echo"
	"github.com/mailgun/mailgun-go/v3"
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

// Render renders a template document
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Mailgun batch send function
func SendTemplateMessage(file string) (string, error) {
	mg := mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_KEY"))
	m := mg.NewMessage(
		"Excited User>"+"<"+os.Getenv("SENDER_EMAIL")+">",
		"Hey %recipient.first%",
		"If you wish to unsubscribe, click http://mailgun/unsubscribe/%recipient.id%",
	) // IMPORTANT: No To:-field recipients!

	// Open the file
	csvfile, err := os.Open(file)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	// Parse the file
	r := csv.NewReader(csvfile)

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
		fmt.Printf("Question: %s Answer %s\n", record[0], record[1])
	}

	m.AddRecipientAndVariables("bob@example.com", map[string]interface{}{
		"first": "bob",
		"id":    1,
	})

	m.AddRecipientAndVariables("alice@example.com", map[string]interface{}{
		"first": "alice",
		"id":    2,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, id, err := mg.Send(ctx, m)
	return id, err
}

// POST Route function
func upload(c echo.Context) error {
	// Read form fields

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	// SendTemplateMessage(file)
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	return c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully.", file.Filename))
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
