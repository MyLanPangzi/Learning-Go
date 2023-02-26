package main

import (
	"html/template"
	"os"
)

func Exec(t *template.Template) error {
	m := map[string]Product{}
	for _, p := range Products {
		m[p.Name] = p
	}
	return t.Execute(os.Stdout, &m)
}
func main() {
	all, err := template.ParseFiles("templates/var.go.txt")
	if err != nil {
		panic(err)
	}
	err = Exec(all.Lookup("main"))
	if err != nil {
		panic(err)
	}
}
