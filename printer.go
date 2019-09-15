package main

import (
	"reflect"
	"regexp"

	tm "github.com/buger/goterm"
)

func watchSummary(c config, outs outputs) {
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
		tm.Printf("%s\n\n", outs.printWordWise())
	} else {
		tm.Printf("%s\n\n", outs.printCharWise())
	}
	tm.Flush()
}

// Compares output char-wise
func (o *outputs) printCharWise() (ret string) {
	// For now, I write the regex
	r := regexp.MustCompile(`\d+`)

	o.prevPos = r.FindAllStringIndex(o.prev, -1)
	o.curPos = r.FindAllStringIndex(o.cur, -1)

	prevLength := len(o.prev)

	// var digitsFound int
	for i := 0; i < len(o.cur); i++ {
		if i < prevLength {
			// If prev string was shorted,
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
func (o *outputs) printWordWise() (ret string) {
	re := regexp.MustCompile(`\S+`)
	prevWords := re.FindAllStringIndex(o.prev, -1)
	curWords := re.FindAllStringIndex(o.cur, -1)

	for i, w := range curWords {
		// Add non-word chars if needed
		// var wsJump int
		if i > 0 {
			// wsJump = w[0] - curWords[i-1][1]
			ret += o.cur[curWords[i-1][1]:w[0]]
		} else {
			// wsJump = w[0]
			ret += o.cur[0:w[0]]
		}

		// Prev output might be shorted
		if i < len(prevWords) {
			// Compare same Nth word
			if reflect.DeepEqual(o.cur[w[0]:w[1]], o.prev[prevWords[i][0]:prevWords[i][1]]) {
				ret += o.cur[w[0]:w[1]]
				continue
			}
		} // Don't care if prev was longer
		ret += getHighlightedString(o.cur[w[0]:w[1]])
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
