package main

import (
	"html/template"
	"os"
)

func GetCategories(products []Product) []string {
	var out []string
	m := make(map[string]string, len(products))
	for _, v := range products {
		m[v.Category] = ""
	}
	for k := range m {
		out = append(out, k)
	}
	return out

}

func Exec(t *template.Template) error {
	return t.Execute(os.Stdout, Products)
}
func main() {
	all := template.New("all")
	all.Funcs(map[string]any{
		"getCategories": GetCategories,
	})
	all, err := all.ParseFiles("templates/func.go.html")
	if err != nil {
		panic(err)
	}
	err = Exec(all.Lookup("main"))
	if err != nil {
		panic(err)
	}
}
