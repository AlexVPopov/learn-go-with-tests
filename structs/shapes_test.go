package main

import "testing"

func assertResult(t *testing.T, got, want float64) {
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{3.0, 4.0})
	want := 14.0

	assertResult(t, got, want)
}

func TestArea(t *testing.T) {
	got := Area(Rectangle{3.0, 4.0})
	want := 12.0

	assertResult(t, got, want)
}
