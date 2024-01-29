package main

import (
	"fmt"
	"html/template"
  "net/http"
)

type ContactDetails struct {
  Email string
  Subject string
  Message string
  Honeypot string
}

func contact(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../views/contact.html",
		"../views/partials/navbar.tmpl"))
	tmpl.Execute(w, nil)
}

func contactFormSubmitted(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../views/contact.html", "../views/partials/navbar.tmpl"))
  
  if r.Method != http.MethodPost{
    tmpl.Execute(w, nil)
    return;
  }

  if r.Method == http.MethodPost && len(r.FormValue("honepot")) == 0 {

    details := ContactDetails{
      Email: r.FormValue("email"),
      Subject: r.FormValue("subject"),
      Message: r.FormValue("message"),
    }

    fmt.Println("Received contact form submission:")
		fmt.Println("Email:", details.Email)
		fmt.Println("Subject:", details.Subject)
		fmt.Println("Message:", details.Message)

		tmpl.Execute(w, struct{ Success bool }{true})
	} else {
    fmt.Println("Robot detected!")
  }
}
