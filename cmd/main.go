package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name      string
	age       uint16
	money     int
	ball      float64
	happiness float64
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go is Super easy!")
}

func docPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Read doc")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/doc/", docPage)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()

}
