package main

import (
	"math/rand"
	"testing"
	"time"
)

var ints = make([]int, 1000)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		ints[i] = i
	}
	m.Run()
}
func Benchmark_contains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		contains(ints, rand.Intn(1000))
	}
}
func Benchmark_contains2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		contains2(ints, rand.Intn(1000))
	}
}
