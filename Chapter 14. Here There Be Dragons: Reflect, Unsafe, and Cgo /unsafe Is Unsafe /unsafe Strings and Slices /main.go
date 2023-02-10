package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func main() {

	hello := "hello"
	bye := []byte("goodbye")
	sHdr := (*reflect.StringHeader)(unsafe.Pointer(&hello))
	bHdr := (*reflect.StringHeader)(unsafe.Pointer(&bye))
	sHdr.Data = bHdr.Data
	sHdr.Len = bHdr.Len
	fmt.Println(hello)
	bye[0] = 'x'
	fmt.Println(hello)

	s := "hello"
	//sHdrData := (*reflect.StringHeader)(unsafe.Pointer((&s))).Data
	sHdrData := unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer((&s))).Data)
	sHdr = (*reflect.StringHeader)(unsafe.Pointer((&s)))
	fmt.Println(sHdr.Len)
	for i := 0; i < sHdr.Len; i++ {
		bp := *(*byte)(unsafe.Pointer(uintptr(sHdrData) + uintptr(i)))
		fmt.Println(string(bp))
		bp = *(*byte)(unsafe.Pointer(sHdr.Data + uintptr(i)))
		fmt.Println(string(bp))
	}
	sHdr.Len = sHdr.Len + 10
	fmt.Println(s)
	bp := (*byte)(unsafe.Pointer(uintptr(sHdrData) + 2))
	//*bp = *bp + 1
	fmt.Println(string(*bp))
	//*bp = *bp + 1
	//fmt.Println(s)
	runtime.KeepAlive(s)

}
