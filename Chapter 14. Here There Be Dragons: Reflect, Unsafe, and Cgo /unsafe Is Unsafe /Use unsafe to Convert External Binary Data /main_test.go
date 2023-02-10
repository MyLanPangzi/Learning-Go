package main

import "testing"

var arr [16]byte
var d Data
var input = [16]byte{0, 132, 95, 237, 80, 104, 111, 110, 101, 0, 0, 0, 0, 0, 1, 0}
var inputData = Data{
	Value:  10000,
	Label:  [10]byte{80, 104, 111, 110, 101, 0, 0, 0, 0, 0},
	Active: true,
}

func BenchmarkBytesFromData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr = BytesFromData(inputData)
	}
}
func BenchmarkBytesFromDataUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr = BytesFromDataUnsafe(inputData)
	}
}
func BenchmarkDataFromBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d = DataFromBytes(input)
	}
}
func BenchmarkDataFromBytesUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d = DataFromBytesUnsafe(input)
	}
}
