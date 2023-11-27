package adder_test

import (
	"testing"

	"wantsome.ro/testing/adder"
)

func Test_Add(t *testing.T) {

	result := adder.Add(1, 2)
	if result != 3 {
		t.Fatalf("Expected 5 but got %d", result)
	}
}
