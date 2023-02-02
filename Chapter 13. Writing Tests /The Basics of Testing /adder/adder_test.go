package adder

import "testing"

func Test_addNumbers(t *testing.T) {
	r := addNumbers(2, 3)
	if r != 5 {
		//t.Error("incorrect result: expected 5, got,", r)
		t.Fatalf("incorrect result: expected 5 got %d", r)
	}

}
