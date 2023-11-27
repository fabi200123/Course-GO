package compare

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_CreatePerson(t *testing.T) {

	expected := Person{
		Name: "John",
		Age:  30,
	}

	result := CreatePerson("John", 30)

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Fatalf("Expected result to match expected result, but got diff %s", diff)
	}

}

func Test_CreatePersonIgnoreInsertDate(t *testing.T) {

	expected := Person{
		Name: "John",
		Age:  30,
	}

	result := CreatePerson("John", 30)

	comp := cmp.Comparer(func(x, y Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})

	if diff := cmp.Diff(expected, result, comp); diff != "" {
		t.Fatalf("Expected result to match expected result, but got diff %s", diff)
	}

	if result.InsertDate.IsZero() {
		t.Errorf("Expected InsertDate to be set, but got zero value")
	}
}
