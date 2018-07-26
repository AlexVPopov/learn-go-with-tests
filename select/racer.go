package main

import (
	"net/http"
	"time"
)

func main() {}

// Racer makes requests to two urls and returns the url, from which is first go a response
func Racer(a, b string) (winner string) {
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return a
	}

	return b
}
