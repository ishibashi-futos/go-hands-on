package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Profile struct {
	Name          string   `json:"name"`
	Age           int      `json:"age"`
	Gender        string   `json:"gender"`
	FavoriteFoods []string `json:"favorite_foods"`
}

var bob Profile = Profile{
	Name:          "Bob",
	Age:           25,
	Gender:        "Man",
	FavoriteFoods: []string{"Hamburger", "Cookie", "Chocolate"},
}

var alice Profile = Profile{
	Name:          "Alice",
	Age:           24,
	Gender:        "Woman",
	FavoriteFoods: []string{"Apple", "Orange", "Melon"},
}

var savedProfiles map[string]Profile = map[string]Profile{
	alice.Name: alice,
	bob.Name:   bob,
}

func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t := template.Must(template.ParseFiles("templates/index.html.tpl"))
	err := t.ExecuteTemplate(w, "index.html.tpl", savedProfiles)
	if err != nil {
		log.Fatal(err)
	}
}

func GetProfile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("name")

	var resProfile Profile
	if name == "Bob" {
		resProfile = bob
	} else if name == "Alice" {
		resProfile = alice
	} else {
		http.Error(w, fmt.Sprintf("%d Not Found", http.StatusNotFound), http.StatusNotFound)
		return
	}

	t := template.Must(template.ParseFiles("templates/profile.html.tpl"))
	err := t.ExecuteTemplate(w, "profile.html.tpl", resProfile)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/profile/:name", GetProfile)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
