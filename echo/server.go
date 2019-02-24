package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type Message struct {
	Title string   `json:"title_name"`
	Text  string   `json:"text_message"`
	Texts []string `json:"texts"`
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("./views/*.html")),
	}
	e := echo.New()
	e.Renderer = t
	e.GET("/hello", func(c echo.Context) error {
		message := &Message{
			Title: "Hello, Golang!",
			Text:  "HogeHoge",
		}
		return c.Render(http.StatusOK, "Hello", message)
	})
	e.GET("/api/hello", func(c echo.Context) error {
		message := &Message{
			Title: "Hello, Golang!",
			Text:  "HogeHoge",
			Texts: []string{"hoge", "fuga"},
		}
		return c.JSON(http.StatusOK, message)
	})
	e.GET("/plus", func(c echo.Context) error {
		type Input struct {
			First  int `query:"first" form:"first"`
			Second int `query:"second" form:"second"`
		}
		input := new(Input)
		if err := c.Bind(input); err != nil {
			log.Fatal(err)
		}

		type Result struct {
			Result int `json:"Result"`
		}
		result := &Result{}
		return c.JSON(http.StatusOK, result)
	})
	e.Start(":8080")
}
