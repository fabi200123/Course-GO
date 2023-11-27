package main

import "fmt"

func modify(a int) {
	a = 10
	fmt.Println("In modify(), value changed to", a)
}
func main() {
	x := 5
	fmt.Println("Before modify(), x =", x)
	modify(x)
	fmt.Println("After modify(), x =", x)
}
