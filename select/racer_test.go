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
		got, err := Racer(slowServer.URL, fastServer.URL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("returns an error if a server doesn't repond within 10s", func(t *testing.T) {
		server := delayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an eror but didn't get one")
		}
	})
}
