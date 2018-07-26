package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {}

var tenSecondTimeout = 10 * time.Second

// Racer calls ConfigurableRacer with a 10s timeout period
func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

// ConfigurableRacer makes requests to two urls and returns the url, from which it first
// got a response or returns an error if it doesn't receive a reponse after the given
// timeout period
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
