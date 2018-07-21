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
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{3.0, 4.0}
		got := rectangle.Area()
		want := 12.0

		assertResult(t, got, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{3.0}
		got := circle.Area()
		want := 28.274333882308139

		assertResult(t, got, want)
	})

}
