package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
////	func createChannelAndSend(data interface{}) interface{} {
////	   dataVal := reflect.ValueOf(data)
////	   channelType := reflect.ChanOf(reflect.BothDir, dataVal.Type().Elem())
////	   channel := reflect.MakeChan(channelType, 1)
////	   go func() {
////	       for i := 0; i < dataVal.Len(); i++ {
////	           channel.Send(dataVal.Index(i))
////	       }
////	       channel.Close()
////	   }()
////	   return channel.Interface()
////	}
//func createAndSend(data any) any {
//	dv := reflect.ValueOf(data)
//	cv := reflect.MakeChan(reflect.ChanOf(reflect.BothDir, dv.Type().Elem()), 1)
//	go func() {
//		defer cv.Close()
//		for i := 0; i < dv.Len(); i++ {
//			cv.Send(dv.Index(i))
//		}
//	}()
//	return cv.Interface()
//}
//func main() {
//	c := createAndSend([]string{"a", "b", "c"}).(chan string)
//	for s := range c {
//		fmt.Println(s)
//	}
//}
//
////func main() {
////    values := []string { "Alice", "Bob", "Charlie", "Dora"}
////    channel := createChannelAndSend(values).(chan string)
////    for {
////        if val, open := <- channel; open {
////            Printfln("Received value: %v", val)
////        } else {
////            break
////        }
////    }
////}
