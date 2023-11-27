package adder

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("setup for tests")
	test_result := m.Run()
	fmt.Println("cleanup after tests")
	os.Exit(test_result)
}

func Test_1(t *testing.T) {
	fmt.Println("Running test 1")
}

func Test_2(t *testing.T) {
	fmt.Println("Running test 2")
}

func Test_Add(t *testing.T) {

	result := add(1, 2)
	if result != 3 {
		t.Errorf("Expected 3 but got %d", result)
	}
}
