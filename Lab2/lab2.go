package main

import (
	"encoding/json"
	"fmt"
)

// type Person struct {
// 	Name   string
// 	Age    int
// 	Height int
// 	// Pet    []Pet
// }

// type Pet struct {
// 	Name  string
// 	Breed string
// }

func parseJson(name string, age int) ([]byte, error) {

	x := struct {
		Name string
		Age  int
	}{
		Name: name,
		Age:  age,
	}
	return json.Marshal(x)
}

func main() {
	// var x []int
	// x = append(x, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// fmt.Printf("Slice address: %p, slice first element: %p\n", &x, &x[0])
	// y := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	// x = append(x, y...)
	// fmt.Printf("Slice address: %p, slice first element: %p", &x, &x[0])

	// x := []int{1, 2, 3, 4, 5}
	// y := x[1:3]
	// y[0] = 100
	// p := x[:]
	// fmt.Println(x, " ", y, " ", p)

	// var m = map[string]int{
	// 	"cheie1": 1,
	// 	"cheie2": 2,
	// 	"cheie3": 3,
	// 	"cheie4": 4,
	// 	"cheie5": 5,
	// }

	// var m = map[string]int{
	// 	"cheie1": 1,
	// 	"cheie2": 2,
	// 	"cheie3": 3,
	// 	"cheie4": 4,
	// 	"cheie5": 5,
	// }

	// m["cheie6"] = 6
	// m["cheie3"] = 2

	// delete(m, "cheie2")

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// fmt.Println(m)

	// var x = Person{
	// 	Name:   "Ion",
	// 	Age:    20,
	// 	Height: 180,
	// }

	// var y = Person{
	// 	Name:   "Maria",
	// 	Age:    20,
	// 	Height: 180,
	// }

	// x := struct {
	// 	Name string
	// 	Age  int
	// }{
	// 	Name: "Ion",
	// 	Age:  20,
	// }

	// fmt.Println(x.Name)
	// fmt.Println(x.Age)
	// fmt.Println(x.Height)
	// fmt.Println(y == x)

	response, _ := parseJson("Ion", 20)
	fmt.Printf("%s", response)

}
