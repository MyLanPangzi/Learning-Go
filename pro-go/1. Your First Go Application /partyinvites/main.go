package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"
)

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}
type formData struct {
	*Rsvp
	Errors []string
}

var responses = make([]*Rsvp, 0, 10)

func main() {
	loadTemplates()

	mux := http.NewServeMux()
	mux.HandleFunc("/", welcomeHandler)
	mux.HandleFunc("/list", listHandler)
	mux.HandleFunc("/form", formHandler)
	server := http.Server{
		Addr:              ":9000",
		Handler:           mux,
		ReadTimeout:       3 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       20 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := templates["form"].Execute(w, formData{&Rsvp{}, []string{}})
		if err != nil {
			panic(err)
		}
		return
	}
	if r.Method == http.MethodPost {
		rsvp := Rsvp{
			Name:       r.FormValue("name"),
			Email:      r.FormValue("email"),
			Phone:      r.FormValue("phone"),
			WillAttend: r.FormValue("willattend") == "true",
		}
		var errors []string
		if strings.TrimSpace(rsvp.Name) == "" {
			errors = append(errors, "Please enter your name")
		}
		if strings.TrimSpace(rsvp.Email) == "" {
			errors = append(errors, "Please enter your email")
		}
		if strings.TrimSpace(rsvp.Phone) == "" {
			errors = append(errors, "Please enter your phone")
		}
		if len(errors) > 0 {
			err := templates["form"].Execute(w, formData{&rsvp, errors})
			if err != nil {
				panic(err)
			}
			return
		}

		responses = append(responses, &rsvp)
		if rsvp.WillAttend {
			err := templates["thanks"].Execute(w, rsvp.Name)
			if err != nil {
				panic(err)
			}
			return
		}
		err := templates["sorry"].Execute(w, rsvp.Name)
		if err != nil {
			panic(err)
		}

	}

}

func listHandler(w http.ResponseWriter, r *http.Request) {
	err := templates["list"].Execute(w, responses)
	if err != nil {
		panic(err)
	}
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	err := templates["welcome"].Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

var templates = make(map[string]*template.Template, 3)

func loadTemplates() {
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}
	for i, name := range templateNames {
		t, err := template.ParseFiles("layout.go.html", name+".go.html")
		if err != nil {
			panic(err)
		}
		templates[name] = t
		fmt.Println("Loaded template", i, name)
	}
}
