package stub

import (
	"strings"
	"testing"
)

func Test_Processor(t *testing.T) {
	p := Processor{Solver: MathSolverStub{}}
	input := strings.NewReader(`1+2
1-2`)
	expected := []int{3, -1}
	for i := 0; i < 2; i++ {
		result, err := p.DoExpression(input)
		if err != nil {
			t.Fatalf("Expected error to be nil, but got %s", err.Error())
		}
		if result != expected[i] {
			t.Fatalf("Expected result to be %d, but got %d", expected[i], result)
		}
	}
}
