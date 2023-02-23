package main

//
//import (
//	"fmt"
//	"reflect"
//	"time"
//)
//
//func main() {
//	ch1 := createAndSend([]int{1, 2, 3}).(chan int)
//	ch2 := createAndSend([]string{"a", "b", "c"}).(chan string)
//	read(ch1, ch2)
//	done := make(chan struct{})
//	a := make(chan string)
//	time.AfterFunc(time.Second*2, func() {
//		close(done)
//	})
//	time.AfterFunc(time.Second*3, func() {
//		close(a)
//	})
//	go sendA(done, a)
//	for s := range a {
//		fmt.Println(s)
//	}
//}
//
//func sendA(channels ...any) {
//
//	cases := make([]reflect.SelectCase, 0, len(channels))
//	sv := reflect.ValueOf(channels)
//	cases = append(cases, reflect.SelectCase{
//		Dir:  reflect.SelectRecv,
//		Chan: sv.Index(0).Elem(),
//	})
//	for i := 1; i < sv.Len(); i++ {
//		cases = append(cases, reflect.SelectCase{
//			Dir:  reflect.SelectSend,
//			Chan: sv.Index(i).Elem(),
//			Send: reflect.ValueOf("A"),
//		})
//	}
//
//	for {
//		chosen, _, _ := reflect.Select(cases)
//		if chosen == 0 {
//			fmt.Println("done")
//			return
//		}
//		time.Sleep(time.Millisecond * 500)
//	}
//
//}
//func createAndSend(s any) any {
//	sv := reflect.ValueOf(s)
//	cv := reflect.MakeChan(reflect.ChanOf(reflect.BothDir, sv.Type().Elem()), 1)
//	go func() {
//		defer cv.Close()
//		for i := 0; i < sv.Len(); i++ {
//			cv.Send(sv.Index(i))
//		}
//	}()
//	return cv.Interface()
//}
//
//func read(channels ...any) {
//	cases := make([]reflect.SelectCase, 0, len(channels))
//	sv := reflect.ValueOf(channels)
//	for i := 0; i < sv.Len(); i++ {
//		fmt.Println(sv.Index(i).Kind(), sv.Index(i).Elem().Type())
//		cases = append(cases, reflect.SelectCase{
//			Dir:  reflect.SelectRecv,
//			Chan: sv.Index(i).Elem(),
//		})
//	}
//
//	for {
//		chosen, recv, ok := reflect.Select(cases)
//		if ok {
//			fmt.Printf("index %v val %v type %v\n", chosen, recv, recv.Type())
//			continue
//		}
//		if len(cases) == 1 {
//			fmt.Println("all channel closed")
//			return
//		}
//		cases = append(cases[:chosen], cases[chosen+1:]...)
//	}
//}
