package main

import (
	"strconv"
	"strings"
)

func getFloat(isHex bool, s string) (f float64, is bool) {
	if isHex {
		if !strings.HasPrefix(s, "0x") {
			s = "0x" + s
		}
		uint, err := strconv.ParseUint(s, 0, 64)
		if err != nil {
			return f, false
		}
		return float64(uint), true
	}

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return f, false
	}
	return f, true
}
