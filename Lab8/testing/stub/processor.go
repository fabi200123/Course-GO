package stub

import (
	"fmt"
	"io"
)

type Processor struct {
	Solver MathSolver
}

type MathSolver interface {
	Resolve(expresson string) (int, error)
}

func read(r io.Reader) (string, error) {
	var out []byte
	buf := make([]byte, 1)
	for {
		_, err := r.Read(buf)
		if err != nil {
			if err == io.EOF {
				return string(out), nil
			}
			return "", fmt.Errorf("error reading from reader: %s", err.Error())
		}
		if buf[0] == '\n' {
			break
		}
		out = append(out, buf[0])
	}
	return string(out), nil
}

func (p Processor) DoExpression(r io.Reader) (int, error) {
	expression, err := read(r)
	if err != nil {
		return 0, err
	}

	if len(expression) == 0 {
		return 0, fmt.Errorf("expression is empty")
	}

	result, err := p.Solver.Resolve(expression)
	return result, err
}

type MathSolverStub struct{}

func (ms MathSolverStub) Resolve(expression string) (int, error) {
	switch expression {
	case "1+2":
		return 3, nil
	case "1-2":
		return -1, nil
	case "1*2*34":
		return 68, nil
	case "(1/2*10":
		return 0, fmt.Errorf("invalid expression")
	}
	return 0, nil
}
