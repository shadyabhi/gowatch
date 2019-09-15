package main

import (
	"flag"
	"fmt"
	"os/exec"
	"reflect"
	"regexp"
	"strings"
	"time"

	tm "github.com/buger/goterm"
)

type config struct {
	Interval     int64
	Count        int
	ShowOutputs  bool
	WordBoundary bool
	ShowRate     bool
	Cmd          string
}

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

func getOutput(cmd string) (out string, err error) {
	outBytes, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return out, fmt.Errorf("error running cmd: %s, got err: %s", cmd, err)
	}
	out = string(outBytes)
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

func run(c config) (err error) {

	ticker := time.NewTicker(time.Duration(c.Interval) * time.Second)
	done := make(chan bool)

	var outs outputs
	for {
		select {
		case <-done:
			fmt.Println("Exiting...")
			return

		case <-ticker.C:
			if c.Count > 0 && outs.i > c.Count {
				done <- true
				return nil
			}
			outs.i++
			outs.cur, err = getOutput(c.Cmd)
			if err != nil {
				done <- true
				return fmt.Errorf("error executing command, got error: %s", err)
			}

			watchSummary(c, outs)
			outs.prev = outs.cur

		}
	}
}

// - Identify all positions, with width
// - Find diff b/w new and old output
// - Replace with diff, maintan width
// - Highlight if it changed
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

func main() {
	// Flags
	c := config{}
	flag.Int64Var(&c.Interval, "d", 1, "Duration to repeat command again")
	flag.IntVar(&c.Count, "c", 0, "Times till which to run the command")
	flag.BoolVar(&c.ShowOutputs, "x", false, "Show current and previous outputs")
	flag.BoolVar(&c.ShowRate, "r", false, "Show diff from previous output")
	flag.BoolVar(&c.WordBoundary, "w", false, "Parse word wise")
	flag.Parse()

	c.Cmd = strings.Join(flag.Args(), " ")
	if err := run(c); err != nil {
		fmt.Printf("Error running command, err: %s", err)
	}
}
