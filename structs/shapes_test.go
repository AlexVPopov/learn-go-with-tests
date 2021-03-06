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
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 3.0, Height: 4.0}, want: 32.0},
		{name: "Circle", shape: Circle{Radius: 3.0}, want: 28.274333882308139},
		{name: "Triangle", shape: Triangle{Base: 12.0, Height: 6.0}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.want)
		})
	}
}
