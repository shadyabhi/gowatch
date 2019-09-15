package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	c := config{}
	if err := c.ParseConfig(); err != nil {
		log.Fatalf("Error parsing command-line arguments: err: %s", err)
	}

	if err := run(c); err != nil {
		fmt.Printf("Error running command, err: %s", err)
	}
}

// run function executes the program with c config provided
// as argument
func run(c config) (err error) {
	ticker := time.NewTicker(time.Duration(c.Interval) * time.Second)
	done := make(chan bool, 1)

	var outs outputs
	for {
		select {
		case <-done:
			return

		case <-ticker.C:
			outs.i++
			if c.Count > 0 && outs.i > c.Count {
				ticker.Stop()
				done <- true
				return nil
			}
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
