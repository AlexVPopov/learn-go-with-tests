package main

import "testing"

func assertResult(t *testing.T, got, want float64) {
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{3.0, 4.0}
	got := rectangle.Perimeter()
	want := 14.0

	assertResult(t, got, want)
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		assertResult(t, got, want)
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{3.0, 4.0}
		checkArea(t, rectangle, 12.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{3.0}
		checkArea(t, circle, 28.274333882308139)
	})
}
