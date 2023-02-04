package solver

import (
	"context"
	"errors"
	"io"
)

type Processor struct {
	Solver MathSolver
}

type MathSolver interface {
	Resolve(ctx context.Context, expression string) (float64, error)
}

func (p Processor) ProcessExpression(ctx context.Context, r io.Reader) (float64, error) {
	expr, err := readToLine(r)
	if err != nil {
		return 0, err
	}
	if len(expr) == 0 {
		return 0, errors.New("no expression to read")
	}
	result, err := p.Solver.Resolve(ctx, expr)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func readToLine(r io.Reader) (string, error) {
	var out []byte
	b := make([]byte, 1)
	for {
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				return string(out), nil
			}
		}
		if b[0] == '\n' {
			break
		}
		out = append(out, b[0])
	}
	return string(out), nil
}

/*

func readToNewLine(r io.Reader) (string, error) {
	var out []byte
	b := make([]byte, 1)
	for {
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				return string(out), nil
			}
		}
		if b[0] == '\n' {
			break
		}
		out = append(out, b[0])
	}
	return string(out), nil
}

*/
