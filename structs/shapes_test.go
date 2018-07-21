package main

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{3.0, 4.0})
	want := 14.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(Rectangle{3.0, 4.0})
	want := 12.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
