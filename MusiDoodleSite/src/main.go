package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	fmt.Println("HTTP Server running on 10000!")

	http.Handle("/assets/", http.StripPrefix("/assets/",
		http.FileServer(http.Dir("../assets"))))

	http.Handle("/", r)

	// Define routes
	r.HandleFunc("/", index)
	r.HandleFunc("/technical", technical)
	r.HandleFunc("/versions", versions)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/download", download)
	r.HandleFunc("/submit-contact", contactFormSubmitted).Methods("POST")

	http.ListenAndServe(":10000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../views/index.html",
		"../views/partials/navbar.tmpl", "../views/partials/version.tmpl"))
	tmpl.Execute(w, nil)
}

func technical(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../views/technical.html",
		"../views/partials/navbar.tmpl", "../views/partials/version.tmpl"))
	tmpl.Execute(w, nil)
}

func versions(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../views/versions.html",
		"../views/partials/navbar.tmpl", "../views/partials/version.tmpl"))
	tmpl.Execute(w, nil)
}

func download(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../views/download.html",
		"../views/partials/navbar.tmpl", "../views/partials/version.tmpl"))
	tmpl.Execute(w, nil)
}
