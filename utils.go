package main

import (
	"time"
)

type Until struct {
	Days    int
	Hours   int
	Minutes int
	Seconds int
}

func getYear() int {
	var date = time.Date(time.Now().Year(), time.December, 25, 0, 0, 0, 0, time.UTC)

	if time.Until(date) > 0 {
		return time.Now().Year()
	}

	return time.Now().Year() + 1
}

func getHours(h int) int {
	for h > 24 {
		h -= 24
	}

	return h - 1
}

func getSixties(t int) int {
	for t > 60 {
		t -= 60
	}

	if t == 60 {
		return 0
	}

	return t
}

func IfThenElse[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

func GetUntil(t time.Time) Until {
	until := time.Until(t)

	return Until{
		Days:    int(until.Hours() / 24),
		Hours:   getHours(int(until.Hours())),
		Minutes: getSixties(int(until.Minutes())),
		Seconds: getSixties(int(until.Seconds())),
	}
}
