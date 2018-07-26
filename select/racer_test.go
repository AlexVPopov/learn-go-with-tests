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
	t.Run("returns the url, which responded faster", func(t *testing.T) {
		slowServer := delayedServer(20 * time.Millisecond)
		fastServer := delayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got, _ := Racer(slowServer.URL, fastServer.URL)

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("returns an error if a server doesn't repond within 10s", func(t *testing.T) {
		serverA := delayedServer(11 * time.Second)
		serverB := delayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected an eror but didn't get one")
		}
	})
}
