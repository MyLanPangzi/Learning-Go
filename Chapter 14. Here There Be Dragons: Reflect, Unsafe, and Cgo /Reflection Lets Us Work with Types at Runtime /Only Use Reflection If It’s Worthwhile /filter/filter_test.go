package filter

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

var target []string
var targetInt []int

func startsWithVowel(s string) bool {
	if len(s) == 0 {
		return false
	}
	switch s[0] {
	case 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}
func buildString() string {
	out := make([]byte, 5)
	for i := 0; i < 5; i++ {
		out[i] = byte(rand.Intn(26) + 65)
	}
	return string(out)
}
func setup() []string {
	rand.Seed(1)
	vals := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		vals[i] = buildString()
	}
	return vals
}
func setupInt() []int {
	rand.Seed(2)
	r := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		r[i] = i
	}
	return r

}
func isEven(i int) bool {
	return i%2 == 0
}

func BenchmarkFilterReflectString(b *testing.B) {
	vals := setup()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target = Filter(vals, startsWithVowel).([]string)
	}
}

func BenchmarkFilterString(b *testing.B) {
	vals := setup()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target = FilterString(vals, startsWithVowel)
	}
}

func BenchmarkFilterInt(b *testing.B) {
	vals := setupInt()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		targetInt = FilterInt(vals, isEven)
	}
}

func BenchmarkFilterReflectInt(b *testing.B) {
	vals := setupInt()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		targetInt = Filter(vals, isEven).([]int)
	}
}
func TestFilter(t *testing.T) {
	names := []string{"Andrew", "Bob", "Clara", "Hortense"}
	longNames := Filter(names, func(s string) bool {
		return len(s) > 3
	}).([]string)
	fmt.Println(longNames)
	expected := []string{"Andrew", "Clara", "Hortense"}
	if !reflect.DeepEqual(longNames, expected) {
		t.Errorf("expected %v got %v", expected, longNames)
	}
	ages := []int{20, 50, 13}
	adults := Filter(ages, func(age int) bool {
		return age >= 18
	}).([]int)
	fmt.Println(adults)
	expectedInt := []int{20, 50}
	if !reflect.DeepEqual(expectedInt, adults) {
		t.Errorf("expected %v got %v", expectedInt, adults)
	}
}
