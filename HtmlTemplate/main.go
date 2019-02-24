package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Message struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func (m Message) toJSON() (jsonStr string) {
	bytes, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	jsonStr = string(bytes)
	return
}

func handleHello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	template := template.Must(template.ParseFiles("./Hello.html.tmpl"))
	message := &Message{
		Title: "Hello, Golang!",
		Text:  "hoge fuga",
	}
	if err := template.ExecuteTemplate(w, "Hello.html.tmpl", message); err != nil {
		log.Fatal(err)
	}
}

func handleHelloApi(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	message := &Message{
		Title: "Hello, Golang!",
		Text:  "hoge fuga",
	}
	fmt.Fprintf(w, message.toJSON())
}

func main() {
	router := httprouter.New()
	router.GET("/hello", handleHello)
	router.GET("/api/hello", handleHelloApi)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
