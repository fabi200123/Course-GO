package table

import "testing"

func Test_doMath(t *testing.T) {
	testCase := []struct {
		name        string
		x           int
		y           int
		op          string
		expected    int
		expectedErr string
	}{
		{"add", 1, 2, "+", 3, ""},
		{"subtract", 1, 2, "-", -1, ""},
		{"multiply", 1, 2, "*", 2, ""},
		{"zero_div", 5, 0, "/", 0, "cannot divide by zero"},
		{"divide", 4, 2, "/", 2, ""},
		{"invalid", 1, 2, "x", 0, "invalid operator"},
	}

	for _, d := range testCase {
		t.Run(d.name, func(t *testing.T) {
			result, err := DoMath(d.x, d.y, d.op)

			if err != nil {
				if err.Error() != d.expectedErr {
					t.Fatalf("Expected error to be %s, but got %s", d.expectedErr, err.Error())
				}
			} else {
				if result != d.expected {
					t.Fatalf("Expected result to be %d, but got %d", d.expected, result)
				}
			}
		})
	}
}
