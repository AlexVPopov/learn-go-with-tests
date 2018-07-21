package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"

	assertString(t, got, want, "test")
}

func assertString(t *testing.T, got, want, given string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s', want '%s', given '%s'", got, want, given)
	}
}
