package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	tm "github.com/buger/goterm"
	"github.com/shadyabhi/gowatch/config"
)

type outputs struct {
	// Outputs current and previous
	prev string
	cur  string

	// positions of interest
	prevPos [][]int
	curPos  [][]int

	// Interation
	i int
}

func watchSummary(c config.Cfg, outs outputs) {
	tm.Clear()
	tm.MoveCursor(1, 1)
	tm.Printf("Every: %ds, Iteration: %d, Command: %s\n\n", c.Interval, outs.i, c.Cmd)

	// Actual output
	if c.ShowOutputs {
		tm.Printf("%s\n\n", outs.cur)
		tm.Printf("%s\n\n", outs.prev)
	}

	// Decorated Out
	if c.WordBoundary {
		tm.Printf("%s\n\n", outs.printWordWise(c))
	} else {
		tm.Printf("%s\n\n", outs.printCharWise(c))
	}
	tm.Flush()
}

// printCharWise compares output char-wise
func (o *outputs) printCharWise(_ config.Cfg) (ret string) {
	prevLength := len(o.prev)

	for i := 0; i < len(o.cur); i++ {
		if i < prevLength {
			// If prev string was short,
			// nothing to compare
			if o.cur[i] == o.prev[i] {
				// As is.
				ret += string(o.cur[i])
			} else {
				// Operate
				// TODO: Don't highlight whitespace
				ret += getHighlightedChar(string(o.cur[i]))
			}
		} else {
			ret += string(o.cur[i])
		}
	}
	return
}

// printWordWise compares output word-wise
func (o *outputs) printWordWise(c config.Cfg) (ret string) {
	re := regexp.MustCompile(`\S+`)
	prevWords := re.FindAllStringIndex(o.prev, -1)
	curWords := re.FindAllStringIndex(o.cur, -1)

	for i, w := range curWords {
		// Preserve whitespaces
		if i > 0 {
			ret += o.cur[curWords[i-1][1]:w[0]]
		} else {
			ret += o.cur[0:w[0]]
		}

		var isFloatCur, isFloatPrev bool
		var curFloat, prevFloat float64

		curOutputWord := o.cur[w[0]:w[1]]
		curFloat, isFloatCur = getFloat(c.Hex, curOutputWord)
		// Prev output might be short
		if i < len(prevWords) {
			// Compare same Nth word
			prevOutputWord := o.prev[prevWords[i][0]:prevWords[i][1]]
			prevFloat, isFloatPrev = getFloat(c.Hex, prevOutputWord)
			if reflect.DeepEqual(curOutputWord, prevOutputWord) {
				ret += curOutputWord
				isFloatCur = false
				isFloatPrev = false
				continue
			}
		} // Don't care if prev was longer
		// We had a change

		// Float?
		if c.ShowRate &&
			isFloatCur && isFloatPrev {

			var floatStr string

			diff := curFloat - prevFloat
			isFloat := strings.Contains(curOutputWord, ".")
			if isFloat {
				floatStr = fmt.Sprintf("%.1f", diff)
			} else {
				floatStr = fmt.Sprintf("%.0f", diff)
			}

			nSpaces := len(curOutputWord) - len(floatStr)
			if nSpaces > 0 {
				ret += getHighlightedString(fmt.Sprintf("%s%s", strings.Repeat(" ", nSpaces), floatStr))
			} else {
				ret += getHighlightedString(floatStr)
			}
			// Reset float finds for next iter
			isFloatCur = false
			isFloatPrev = false
			continue
		}

		// String
		if o.i > 1 {
			ret += getHighlightedString(curOutputWord)
		} else {
			// Don't highlight on first run
			ret += curOutputWord
		}
	}
	return
}

func getHighlightedString(str string) (ret string) {
	for i := 0; i < len(str); i++ {
		ret += getHighlightedChar(string(str[i]))
		// ret += "_" + string(str[i]) + "_" // Great to testing and comparing
	}
	return
}

func getHighlightedChar(str string) (ret string) {
	ret = tm.Color(str, tm.BLACK)
	ret = tm.Background(ret, tm.WHITE)
	return
}
