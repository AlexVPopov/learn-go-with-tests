package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}

const finalWord = "Go!"
const countdownStart = 3

// Sleeper adds a delay to execution
type Sleeper interface {
	Sleep()
}

// DefaultSleeper impelments the sleeper interface
type DefaultSleeper struct{}

// Sleep delays the exection with 1 second
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// ConfigurableSleeper allows for setting a custom sleep duration
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep lets ConfigurableSleep sleep
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// Countdown counts backwords to 0 and prints go
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}
