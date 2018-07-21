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

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{3.0, 4.0}, 12.0},
		{Circle{3.0}, 28.274333882308139},
		{Triangle{12.0, 6.0}, 36.0},
	}

	for _, tt := range areaTests {
		checkArea(t, tt.shape, tt.want)
	}
}
