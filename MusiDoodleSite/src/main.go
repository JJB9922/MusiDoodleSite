package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Println("HTTP Server running on 8080!")

	//Serve images from the assets folder.
	http.Handle("/assets/", http.StripPrefix("/assets/",
		http.FileServer(http.Dir("../assets"))))

	//Declare a function which deals with http responses.
	//It parses the index.html file.
	index := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("../views/index.html",
			"../views/partials/navbar.tmpl"))
		tmpl.Execute(w, nil)
	}

	screenshots := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("../views/screenshots.html",
			"../views/partials/navbar.tmpl"))

		tmpl.Execute(w, nil)
	}

	//Handlers
	http.HandleFunc("/", index)
	http.HandleFunc("/screenshots", screenshots)

	//Listen for traffic on 8080, fatal log if there's a problem.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
