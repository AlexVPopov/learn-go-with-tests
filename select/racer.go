package main

import (
	"net/http"
	"time"
)

func main() {}

// Racer makes requests to two urls and returns the url, from which is first go a response
func Racer(a, b string) (winner string) {
	aDuration := measureReponseTime(a)
	bDuration := measureReponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureReponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
