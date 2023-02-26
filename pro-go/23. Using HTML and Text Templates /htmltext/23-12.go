package main

//
//import (
//	"html/template"
//	"os"
//)
//
////	func Exec(t *template.Template) error {
////	   return t.Execute(os.Stdout, &Kayak)
////	}
////
////	func main() {
////	   allTemplates, err := template.ParseGlob("templates/*.html")
////	   if (err == nil) {
////	       selectedTemplated := allTemplates.Lookup("template.html")
////	       err = Exec(selectedTemplated)
////	   }
////	   if (err != nil) {
////	       Printfln("Error: %v %v", err.Error())
////	   }
////	}
//func main() {
//	all, err := template.ParseGlob("templates/*.html")
//	if err != nil {
//		panic(err)
//	}
//	t := all.Lookup("templates.go.html")
//	err = Exec(t)
//	if err != nil {
//		panic(err)
//	}
//}
//
//func Exec(t *template.Template) error {
//	return t.Execute(os.Stdout, &Kayak)
//}
