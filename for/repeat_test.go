package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertCorrect := func(t *testing.T, repeated, expected string) {
		if repeated != expected {
			t.Errorf("Expected %s got %s", expected, repeated)
		}
	}

	t.Run("repeating 4 times", func(t *testing.T) {
		repeated := Repeat("a", 4)
		expected := "aaaa"
		assertCorrect(t, expected, repeated)
	})

	t.Run("repeating 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrect(t, expected, repeated)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repated := Repeat("a", 3)
	fmt.Println(repated)
	// Output: aaa
}
