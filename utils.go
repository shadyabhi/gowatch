package main

import "strconv"

func getFloat(s string) (f float64, is bool) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return f, false
	}
	return f, true
}
