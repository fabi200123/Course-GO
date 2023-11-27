package bench

import (
	"fmt"
	"testing"
)

var blackhole int

func Test_GetFileLen(t *testing.T) {

	result, err := GetFileLen("testdata/data.txt", 1)
	expected := 1335
	if err != nil {
		t.Fatalf("Expected error to be nil, but got %s", err.Error())
	}

	if result != expected {
		t.Fatalf("Expected result to be %d, but got %d", expected, result)
	}

}

func Benchmark_GetFileLen(b *testing.B) {
	for _, bufferSize := range []int{1, 10, 100, 1000, 10000} {
		b.Run(fmt.Sprintf("bufferSize=%d", bufferSize), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, err := GetFileLen("testdata/data.txt", 1)
				if err != nil {
					b.Fatal(err)
				}
				blackhole = result
			}
		})
	}
}
