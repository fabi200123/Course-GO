package table

import "fmt"

func DoMath(x, y int, op string) (int, error) {
	switch op {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "*":
		return x * y, nil
	case "/":
		if y == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return x / y, nil
	default:
		return 0, fmt.Errorf("invalid operator")
	}
}
