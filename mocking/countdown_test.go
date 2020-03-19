package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("Prints 3 -> Go!", func(t *testing.T) {
		spySleeperPrinter := &CountdownSpy{}
		buffer := &bytes.Buffer{}
		Countdown(buffer, spySleeperPrinter)

		got := buffer.String()
		want := "3\n2\n1\nGo!"
		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("Sleep before every print", func(t *testing.T) {
		spySleeperPrinter := &CountdownSpy{}
		Countdown(spySleeperPrinter, spySleeperPrinter)

		expected := []string{
			SleepMessage,
			WriteMessage,
			SleepMessage,
			WriteMessage,
			SleepMessage,
			WriteMessage,
			SleepMessage,
			WriteMessage,
		}

		if !reflect.DeepEqual(expected, spySleeperPrinter.Calls) {
			t.Errorf("wanted calls %v but got %v", expected,
				spySleeperPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := &ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.DurationSlept != sleepTime {
		t.Errorf("Should have slept for %v but slept for %v", sleepTime,
			spyTime.DurationSlept)
	}
}
