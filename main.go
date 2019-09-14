package main

import (
	"flag"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"time"

	tm "github.com/buger/goterm"
)

type config struct {
	Interval    int64
	Count       int
	ShowOutputs bool
	ShowRate    bool
	Cmd         string
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
func decorate(outs outputs) (ret string) {
	// For now, I write the regex
	r := regexp.MustCompile(`\d+`)

	outs.prevPos = r.FindAllStringIndex(outs.prev, -1)
	outs.curPos = r.FindAllStringIndex(outs.cur, -1)

	prevLength := len(outs.prev)

	// var digitsFound int
	for i := 0; i < len(outs.cur); i++ {
		if i < prevLength {
			// If prev string was shorted,
			// nothing to compare
			if outs.cur[i] == outs.prev[i] {
				// As is.
				ret += string(outs.cur[i])
			} else {
				// Operate
				// TODO: Don't highlight whitespace
				// if i == outs.curPos[digitsFound][0] {
				// 	digitsFound++
				// 	prevNum, err := strconv.ParseInt(outs.prev[outs.prevPos[digitsFound][0]:outs.prevPos[digitsFound][1]], 10, 64)
				// 	if err != nil {
				// 		panic(err)
				// 	}
				// 	currNum, err := strconv.ParseInt(outs.cur[outs.curPos[digitsFound][0]:outs.curPos[digitsFound][1]], 10, 64)
				// 	if err != nil {
				// 		panic(err)
				// 	}
				// 	diff := currNum - prevNum
				// 	ret += getHighlightedString(strconv.Itoa(int(diff)))
				// 	continue
				// }
				ret += getHighlightedChar(string(outs.cur[i]))
			}
		} else {
			ret += string(outs.cur[i])
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
	tm.Printf("%s\n\n", decorate(outs))
	tm.Flush()
}

func main() {
	// Flags
	c := config{}
	flag.Int64Var(&c.Interval, "d", 1, "Duration to repeat command again")
	flag.IntVar(&c.Count, "c", 0, "Times till which to run the command")
	flag.BoolVar(&c.ShowOutputs, "x", false, "Show current and previous outputs")
	flag.BoolVar(&c.ShowRate, "r", false, "Show diff from previous output")
	flag.Parse()

	c.Cmd = strings.Join(flag.Args(), " ")
	if err := run(c); err != nil {
		fmt.Printf("Error running command, err: %s", err)
	}
}
