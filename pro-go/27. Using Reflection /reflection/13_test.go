package main

//
//import (
//	"math/rand"
//	"os"
//	"testing"
//	"time"
//)
//
//var names = []string{"hello", "world", "go"}
//var n string
//
//func TestMain(m *testing.M) {
//	rand.Seed(time.Now().UnixNano())
//	code := m.Run()
//	os.Exit(code)
//}
//
//func Benchmark_selectValue(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		n = selectValue(names, rand.Intn(3)).(string)
//	}
//}
//func Benchmark_selectValue2(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		n = selectValue2(names, rand.Intn(3))
//	}
//}
