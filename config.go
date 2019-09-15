package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type config struct {
	Interval     int64
	Count        int
	ShowOutputs  bool
	WordBoundary bool
	ShowRate     bool
	Cmd          string
}

func (c *config) ParseConfig() error {
	flag.Int64Var(&c.Interval, "d", 1, "Repeat every 'd' seconds.")
	flag.IntVar(&c.Count, "c", 0, "Stop after 'c' executions.")
	flag.BoolVar(&c.ShowOutputs, "o", false, "Show previous, current and diff outputs")
	flag.BoolVar(&c.ShowRate, "r", false, "Show difference from previous output for int/floats")
	flag.BoolVar(&c.WordBoundary, "w", false, "Parse wordwise, not charwise")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), `gowatch is a tool like 'watch' but provides additional features like
seeing difference from previous output for numeric words.

Following command runs the command 'cmd' every second, forever and lists number difference
insread of just the new string: gowatch -r 'cmd'`+"\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Arguments:-\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	c.Cmd = strings.Join(flag.Args(), " ")

	// Everything akay?
	if c.ShowRate {
		c.WordBoundary = true
	}
	if len(c.Cmd) < 1 {
		return fmt.Errorf("error: command to execute not provided")
	}
	return nil
}
