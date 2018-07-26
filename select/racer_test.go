package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func delayedServer(delay time.Duration) *httptest.Server {
	delayedFunc := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}

	handlerFunc := http.HandlerFunc(delayedFunc)

	return httptest.NewServer(handlerFunc)
}

func TestRacer(t *testing.T) {
	slowServer := delayedServer(20 * time.Millisecond)
	fastServer := delayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
