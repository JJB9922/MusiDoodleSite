package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	fmt.Println("HTTP Server running on 8080!")

	http.Handle("/assets/", http.StripPrefix("/assets/",
		http.FileServer(http.Dir("../assets"))))

	http.Handle("/", r)

	// Define routes
	r.HandleFunc("/", index)
	r.HandleFunc("/screenshots", screenshots)
	r.HandleFunc("/technical", technical)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/download", download)
	r.HandleFunc("/submit-contact", contactFormSubmitted).Methods("POST")

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../views/index.html",
		"../views/partials/navbar.tmpl"))
	tmpl.Execute(w, nil)
}

func screenshots(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../views/screenshots.html",
		"../views/partials/navbar.tmpl"))
	tmpl.Execute(w, nil)
}

func technical(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../views/technical.html",
		"../views/partials/navbar.tmpl"))
	tmpl.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../views/contact.html",
		"../views/partials/navbar.tmpl"))
	tmpl.Execute(w, nil)
}

func download(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../views/download.html",
		"../views/partials/navbar.tmpl"))
	tmpl.Execute(w, nil)
}

func contactFormSubmitted(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		fmt.Println("Received contact form submission:")
		fmt.Println("Email:", r.FormValue("email"))
		fmt.Println("Subject:", r.FormValue("subject"))
		fmt.Println("Message:", r.FormValue("message"))

		w.Write([]byte("Form submitted successfully!"))
		return
	}

	tmpl := template.Must(template.ParseFiles("../views/screenshots.html",
		"../views/partials/navbar.tmpl"))
	tmpl.Execute(w, nil)
}
