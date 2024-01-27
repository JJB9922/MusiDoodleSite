package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

)

func main() {
	fmt.Println("MusiDoodle server is running on port 8080!")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("MusiDoodleSite/index.html"))
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}