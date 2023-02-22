package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
//func inspectTags(s any, tagName string) {
//	st := reflect.TypeOf(s)
//	v := s.(reflect.Value)
//	fmt.Println(st, v.Type(), v.Elem().Type())
//	for i := 0; i < st.NumField(); i++ {
//		f := st.Field(i)
//		get := f.Tag.Get(tagName)
//		lookup, ok := f.Tag.Lookup(tagName)
//		fmt.Printf(
//			"field %v tag %v tagName %v get [%v] lookup [%v] ok %v\n",
//			f.Name,
//			f.Tag,
//			tagName,
//			get,
//			lookup,
//			ok,
//		)
//	}
//}
//
//type Person struct {
//	Name    string `alias:"id"`
//	City    string `alias:""`
//	Country string
//}
//
//func main() {
//	//inspectTags(Person{}, "alias")
//	//	stringType := reflect.TypeOf("this is a string")
//	//    structType := reflect.StructOf([] reflect.StructField {
//	//        { Name: "Name", Type: stringType, Tag: `alias:"id"` },
//	//        { Name: "City", Type: stringType,Tag: `alias:""`},
//	//        { Name: "Country", Type: stringType },
//	//    })
//	//    inspectTags(reflect.New(structType), "alias")
//
//	stringType := reflect.TypeOf("")
//	t := reflect.StructOf([]reflect.StructField{
//		{Name: "Name", Type: stringType, Tag: `alias:"id"`},
//		{Name: "City", Type: stringType, Tag: `alias:""`},
//		{Name: "Country", Type: stringType},
//	})
//	inspectTags(reflect.New(t), "alias")
//}
//
////func inspectTags(s interface{}, tagName string) {
////    structType := reflect.TypeOf(s)
////    for i := 0; i < structType.NumField(); i++ {
////        field := structType.Field(i)
////        tag := field.Tag
////        valGet := tag.Get(tagName)
////        valLookup, ok := tag.Lookup(tagName)
////        Printfln("Field: %v, Tag %v: %v", field.Name, tagName, valGet)
////        Printfln("Field: %v, Tag %v: %v, Set: %v",
////            field.Name, tagName, valLookup, ok)
////    }
////}
////type Person struct {
////    Name string `alias:"id"`
////    City string `alias:""`
////    Country string
////}
////func main() {
////    inspectTags(Person{}, "alias")
////}
