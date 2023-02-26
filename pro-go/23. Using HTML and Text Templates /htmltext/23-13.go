package main

//
//import (
//	"html/template"
//	"os"
//)
//
//func main() {
//	all, err := template.ParseGlob("templates/*.html")
//	if err != nil {
//		panic(err)
//	}
//	t := all.Lookup("range.go.html")
//	err = Exec(t)
//	if err != nil {
//		panic(err)
//	}
//}
//
//func Exec(t *template.Template) error {
//	return t.Execute(os.Stdout, Products)
//}
