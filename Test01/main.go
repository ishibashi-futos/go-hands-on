package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func FizzBuzz(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	num, err := strconv.Atoi(p.ByName("num"))
	if err != nil || num < 1 {
		http.Error(w, fmt.Sprintf("%d Bad Request", http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var str string
	for i := 1; i <= num; i++ {
		if i%15 == 0 {
			str += fmt.Sprintf("%d: %s\n", i, "FizzBuzz!")
		} else if i%5 == 0 {
			str += fmt.Sprintf("%d: %s\n", i, "Buzz")
		} else if i%3 == 0 {
			str += fmt.Sprintf("%d: %s\n", i, "Fizz")
		} else {
			str += fmt.Sprintf("%d:\n", i)
		}
	}
	fmt.Fprintf(w, str)
}

func main() {
	router := httprouter.New()
	router.GET("/FizzBuzz/:num", FizzBuzz)
	router.GET("/Profile/:name", GetProfile)
	router.POST("/Profile", PostProfile)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
