package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"text/template"
)

var templates = template.Must(template.ParseFiles("view.html", "edit.html", "list.html"))
var validPath = regexp.MustCompile("^/(view|edit|save)/([a-zA-Z0-9]+)$")

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	return os.WriteFile(fmt.Sprintf("%s.txt", p.Title), p.Body, 0600)
}
func loadPage(title string) (*Page, error) {
	body, err := os.ReadFile(fmt.Sprintf("%s.txt", title))
	if err != nil {
		return nil, err
	}
	return &Page{title, body}, nil
}
func main() {
	p1 := &Page{"hello", []byte("world")}
	err := p1.save()
	if err != nil {
		log.Fatal(err)
	}
	p2, err := loadPage(p1.Title)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(p2.Body))

	http.Handle("/view/", wikiHandler(viewHandler))
	http.Handle("/save/", wikiHandler(saveHandler))
	http.Handle("/edit/", wikiHandler(editHandler))
	http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {
		wikis := findWikis()
		fmt.Println(wikis)
		templates.ExecuteTemplate(writer, "list.html", wikis)
		//if err != nil {
		//	http.Error(writer, "not found wiki list", http.StatusInternalServerError)
		//	return
		//}
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

func findWikis() []string {
	pwd, err := os.Getwd()
	if err != nil {
		return nil
	}
	var wikis []string
	files, err := os.ReadDir(pwd)
	if err != nil {
		return wikis
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".txt") {
			wikis = append(wikis, strings.TrimSuffix(file.Name(), ".txt"))
		}
	}
	return wikis
}

type wikiHandler func(http.ResponseWriter, *http.Request, string)

func (w wikiHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	title, err := getTitle(writer, request)
	if err != nil {
		return
	}
	w(writer, request, title)
}

func viewHandler(writer http.ResponseWriter, request *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(writer, request, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(writer, "view", p)
}

func editHandler(writer http.ResponseWriter, request *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(writer, "edit", p)
}

func renderTemplate(writer http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(writer, tmpl+".html", p)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

}

func saveHandler(writer http.ResponseWriter, request *http.Request, title string) {
	body := request.FormValue("body")
	p := Page{title, []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(writer, request, "/view/"+title, http.StatusFound)
}

func getTitle(writer http.ResponseWriter, request *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(request.URL.Path)
	if m == nil {
		http.NotFound(writer, request)
		return "", errors.New("invalid title")
	}
	return m[2], nil
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hi there, I love %s", request.URL.Path[1:])
}
