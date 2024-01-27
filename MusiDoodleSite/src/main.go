package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Println("HTTP Server running on 8080!")

	index := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./index.html"))
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
