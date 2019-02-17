package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	method        = flag.String("X", "GET", "method type")
	name          = flag.String("name", "", "[Required]type your name")
	age           = flag.Int("age", -1, "[Required]type your age")
	gender        = flag.String("gender", "", "[Required]type your gender")
	favoriteFoods = flag.String("favorite-foods", "", "type your favorite foods(Separate by space) e.g. apple orange")
)

const url string = "http://localhost:8080/"

func GetProfile() {
	response, err := http.Get(fmt.Sprintf(url+"Profile/%s", *name))
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
}

type Profile struct {
	Name          string   `json:"name"`
	Age           int      `json:"age"`
	Gender        string   `json:"gender"`
	FavoriteFoods []string `json:"favorite_foods"`
}

func PostProfile() {
	profile := ArgsToProfile()
	json, err := json.Marshal(profile)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", url+"Profile/", bytes.NewBuffer([]byte(json)))
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}

func ArgsToProfile() Profile {
	return Profile{
		Name:          *name,
		Age:           *age,
		Gender:        *gender,
		FavoriteFoods: strings.Split(*favoriteFoods, " "),
	}
}

func main() {
	flag.Parse()
	switch *method {
	case "GET":
		GetProfile()
	case "POST":
		PostProfile()
	default:
		fmt.Printf("Select your method")
	}
}
