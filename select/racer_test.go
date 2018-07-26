package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func fastReponse(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func slowResonse(w http.ResponseWriter, r *http.Request) {
	time.Sleep(20 * time.Millisecond)
	fastReponse(w, r)
}

func TestRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(slowResonse))
	fastServer := httptest.NewServer(http.HandlerFunc(fastReponse))

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}

	slowServer.Close()
	fastServer.Close()
}
