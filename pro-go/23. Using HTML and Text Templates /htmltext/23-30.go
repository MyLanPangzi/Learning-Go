package main

//
//import (
//	"html/template"
//	"os"
//)
//
//func main() {
//	all, err := template.ParseFiles("templates/block.go.html", "templates/list.go.html")
//	if err != nil {
//		panic(err)
//	}
//	t := all.Lookup("mainTemplate")
//	err = Exec(t)
//	if err != nil {
//		panic(err)
//	}
//}
//
//func Exec(t *template.Template) error {
//	return t.Execute(os.Stdout, Products)
//}
