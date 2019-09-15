package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
	"time"
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
			outs.cur, err = getCMDOutput(c.Cmd)
			if err != nil {
				done <- true
				return fmt.Errorf("error executing command, got error: %s", err)
			}

			watchSummary(c, outs)
			outs.prev = outs.cur

		}
	}
}

// getCMDOutput is a helper function to execute command
func getCMDOutput(cmd string) (out string, err error) {
	outBytes, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return out, fmt.Errorf("error running cmd: %s, got err: %s", cmd, err)
	}
	out = string(outBytes)
	return
}
