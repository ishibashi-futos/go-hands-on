package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func GetProfile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("name")
	var responseProfile Profile
	responseProfile, ok := savedProfiles[name]
	if !ok {
		http.Error(w, fmt.Sprintf("%d Not Found", http.StatusNotFound), http.StatusNotFound)
		return
	}
	bytes, err := json.Marshal(responseProfile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Fprintf(w, string(bytes))
}

var savedProfiles map[string]Profile = map[string]Profile{
	alice.Name: alice,
	bob.Name:   bob,
}

func PostProfile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var reqProfile Profile
	err = json.Unmarshal(body, &reqProfile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if reqProfile.Name == "" {
		http.Error(w, "name is required.", http.StatusBadRequest)
		return
	}

	savedProfiles[reqProfile.Name] = reqProfile
	fmt.Fprintf(w, fmt.Sprintf("%d Created", http.StatusCreated))
}
