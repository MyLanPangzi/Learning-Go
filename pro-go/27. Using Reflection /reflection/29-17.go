package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
////	func sendOverChannel(channel interface{}, data interface{}) {
////	   channelVal := reflect.ValueOf(channel)
////	   dataVal := reflect.ValueOf(data)
////	   if (channelVal.Kind() == reflect.Chan &&
////	           dataVal.Kind() == reflect.Slice &&
////	           channelVal.Type().Elem() == dataVal.Type().Elem()) {
////	       for i := 0; i < dataVal.Len(); i++ {
////	           val := dataVal.Index(i)
////	           channelVal.Send(val)
////	       }
////	       channelVal.Close()
////	   } else {
////	       Printfln("Unexpected types: %v, %v", channelVal.Type(), dataVal.Type())
////	   }
////	}
//func send(c, data any) {
//	cv := reflect.ValueOf(c)
//	dv := reflect.ValueOf(data)
//	if !(cv.Kind() == reflect.Chan &&
//		dv.Kind() == reflect.Slice) || cv.Type().Elem() != dv.Type().Elem() {
//		return
//	}
//	for i := 0; i < dv.Len(); i++ {
//		cv.Send(dv.Index(i))
//	}
//	cv.Close()
//}
//func main() {
//	ch := make(chan string)
//	go send(ch, []string{"a", "b", "c"})
//	for s := range ch {
//		fmt.Println(s)
//	}
//}
//
////func main() {
////    values := []string { "Alice", "Bob", "Charlie", "Dora"}
////    channel := make(chan string)
////    go sendOverChannel(channel, values)
////    for {
////        if val, open := <- channel; open {
////            Printfln("Received value: %v", val)
////        } else {
////            break
////        }
////    }
////}
