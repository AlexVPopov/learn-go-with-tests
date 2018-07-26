package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func delayedResponse(w http.ResponseWriter, r *http.Request, delay time.Duration) {
	time.Sleep(delay)
	w.WriteHeader(http.StatusOK)
}

func fastReponse(w http.ResponseWriter, r *http.Request) {
	delayedResponse(w, r, 0)
}

func slowResonse(w http.ResponseWriter, r *http.Request) {
	delayedResponse(w, r, 20)
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
