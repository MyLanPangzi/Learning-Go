package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
////	func inspectChannel(channel interface{}) {
////	   channelType := reflect.TypeOf(channel)
////	   if (channelType.Kind() == reflect.Chan) {
////	       Printfln("Type %v, Direction: %v",
////	           channelType.Elem(), channelType.ChanDir())
////	   }
////	}
////
////	func main() {
////	   var c chan<- string
////	   inspectChannel(c)
////	}
//func inspectChan(c any) {
//	cv := reflect.TypeOf(c)
//	if cv.Kind() == reflect.Chan {
//		fmt.Printf("Type %v, Direction %v\n", cv.Elem(), cv.ChanDir())
//	}
//}
//func main() {
//	var c1 chan<- string
//	inspectChan(c1)
//	var c2 <-chan string
//	inspectChan(c2)
//	var c3 chan string
//	inspectChan(c3)
//
//}
