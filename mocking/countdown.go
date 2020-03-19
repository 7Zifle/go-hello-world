package main

import (
	"fmt"
	"io"
	"time"
)

const (
	Count        = 3
	FinalWord    = "Go!"
	WriteMessage = "write"
	SleepMessage = "sleep"
)

type ConfigurableSleeper struct {
	Duration      time.Duration
	SleepFunction func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.SleepFunction(c.Duration)
}

type Sleeper interface {
	Sleep()
}

type SpyTime struct {
	DurationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.DurationSlept = duration
}

type CountdownSpy struct {
	Calls []string
}

func (c *CountdownSpy) Sleep() {
	c.Calls = append(c.Calls, SleepMessage)
}

func (c *CountdownSpy) Write(p []byte) (n int, e error) {
	c.Calls = append(c.Calls, WriteMessage)
	return
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := Count; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintf(out, "%d\n", i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, FinalWord)
}
