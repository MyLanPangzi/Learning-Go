package main

import (
	"encoding/binary"
	"fmt"
	"math/bits"
	"unsafe"
)

func main() {
	d := Data{
		Value:  8675309,
		Active: true,
	}
	copy(d.Label[:], "Phone")
	fmt.Println(d, unsafe.Alignof(d), unsafe.Alignof(d.Value), unsafe.Alignof(d.Label), unsafe.Alignof(d.Active))
	b := [16]byte{0, 132, 95, 237, 80, 104, 111, 110, 101, 0, 0, 0, 0, 0, 1, 0}
	fmt.Println(b)
	d1 := DataFromBytes(b)
	d2 := DataFromBytesUnsafe(b)
	if d1 != d2 {
		panic(fmt.Sprintf("%v %v", d1, d2))
	}
	fmt.Printf("%+v\n", d1)
	b1 := BytesFromData(d)
	b2 := BytesFromDataUnsafe(d)
	if b1 != b2 {
		panic(fmt.Sprintf("%v %v", b1, b2))
	}
	fmt.Println(b1)
}

func BytesFromDataUnsafe(d Data) [16]byte {
	if isLE {
		d.Value = bits.ReverseBytes32(d.Value)
	}
	return *(*[16]byte)(unsafe.Pointer(&d))
}

func BytesFromData(d Data) [16]byte {
	bytes := [16]byte{}
	binary.BigEndian.PutUint32(bytes[:4], d.Value)
	copy(bytes[4:14], d.Label[:])
	if d.Active {
		bytes[14] = 1
	}
	return bytes
}

func DataFromBytesUnsafe(b [16]byte) Data {
	r := *(*Data)(unsafe.Pointer(&b))
	if isLE {
		r.Value = bits.ReverseBytes32(r.Value)
	}
	return r
}

func DataFromBytes(b [16]byte) Data {
	d := Data{}
	d.Value = binary.BigEndian.Uint32(b[:4])
	copy(d.Label[:], b[4:14])
	d.Active = b[14] != 0
	return d
}

type Data struct {
	Value  uint32
	Label  [10]byte
	Active bool
}

var isLE bool

func init() {
	var x uint16 = 0xFF00
	bytes := *(*[2]byte)(unsafe.Pointer(&x))
	isLE = bytes[0] == 0x00
}
