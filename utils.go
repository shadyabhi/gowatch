package main

import (
	"strconv"
)

func getFloat(s string) (f float64, is bool) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		// Last chance? Hex?
		uint, err := strconv.ParseUint(s, 16, 64)
		if err != nil {
			return f, false
		}
		f = float64(uint)
	}
	return f, true
}
