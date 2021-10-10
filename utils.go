package main

import (
	"time"
)

func getYear() int {
	var date = time.Date(time.Now().Year(), time.December, 25, 0, 0, 0, 0, time.UTC)

	if time.Until(date) > 0 {
		return time.Now().Year()
	}

	return time.Now().Year() + 1
}

func getHours(h float64) float64 {
	for h > 24 {
		h -= 24
	}

	return h
}

func getSixties(t float64) float64 {
	for t > 60 {
		t -= 60
	}

	return t
}
