package main

type A struct {
	a string
}

func (a A) doSomething() string {
	return ""
}

func main() {
	var a A
	a.doSomething()
}
