package adder_test

import (
	"pulbic_test/adder"
	"testing"
)

func TestAdd(t *testing.T) {
	result := adder.Add(1, 2)
	if result != 3 {
		t.Errorf("expected 3 got %d", result)
	}
}
