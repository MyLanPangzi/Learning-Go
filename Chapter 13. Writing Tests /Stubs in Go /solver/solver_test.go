package solver

import (
	"context"
	"errors"
	"strings"
	"testing"
)

type MathSolverStub struct {
}

func (m MathSolverStub) Resolve(ctx context.Context, expression string) (float64, error) {
	switch expression {
	case "2 + 2 * 10":
		return 22, nil
	case "( 2 + 2 ) * 10":
		return 40, nil
	case "(2 + 2 * 10":
		return 0, errors.New("invalid expression: " + expression)
	}
	return 0, nil
}

func TestProcessorProcessExpression(t *testing.T) {
	p := Processor{MathSolverStub{}}
	in := strings.NewReader(`2 + 2 * 10
( 2 + 2 ) * 10
( 2 + 2 * 10 `)
	data := []float64{22, 40, 0}
	hadErr := []bool{false, false, true}
	for i, d := range data {
		result, err := p.ProcessExpression(context.Background(), in)
		if err != nil && !hadErr[i] {
			t.Error(err)
		}
		if result != d {
			t.Errorf("Expected result %f, got %f", d, result)
		}

	}
}
