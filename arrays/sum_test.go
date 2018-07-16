package main

import (
	"reflect"
	"testing"
)

func assertSum(t *testing.T, got, want int, numbers []int) {
	if got != want {
		t.Errorf("got %d expected %d, given %v", got, want, numbers)
	}
}

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertSum(t, got, want, numbers)
	})
}

func assertSums(t *testing.T, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{4, 6, 9})
	want := []int{3, 19}

	assertSums(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	t.Run("summing non-empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{4, 6, 9})
		want := []int{2, 15}

		assertSums(t, got, want)
	})

	t.Run("summing empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{4, 6, 9})
		want := []int{0, 15}

		assertSums(t, got, want)
	})
}
