package person

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestNewPerson(t *testing.T) {
	expected := Person{
		Name: "hello",
		Age:  10,
	}
	person := NewPerson("hello", 10)
	//if diff := cmp.Diff(expected, person); diff != "" {
	//	t.Error(diff)
	//}
	comparer := cmp.Comparer(func(x, y Person) bool {
		return x.Name == y.Name && y.Age == x.Age
	})
	if diff := cmp.Diff(expected, person, comparer); diff != "" {
		t.Error(diff)
	}

}
