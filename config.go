package main

import (
	"flag"
	"fmt"
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
	flag.Parse()

	c.Cmd = strings.Join(flag.Args(), " ")

	// Everything akay?
	if len(c.Cmd) < 1 {
		return fmt.Errorf("error: command to execute not provided")
	}
	return nil
}
