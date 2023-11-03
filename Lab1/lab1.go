package main

import (
	"fmt"
)

func sum(a, b int) (int, error) {
	return a + b, nil
}

func main() {
	i, j := 10, 1000
	var p *int
	p = &i
	*p = 25
	fmt.Println(i)

	*p = *(&j)
	fmt.Println(*p)
}
