package main

import (
	"os"
	"testing"
)

func Test_config_ParseConfig(t *testing.T) {
	c := &config{}

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{
		"gowatch",
		"-r",
		"echo",
	}
	if err := c.ParseConfig(); err != nil {
		t.Errorf("Unexpected error in parsing, err: %s", err)
	}
	if c.ShowRate != true {
		t.Errorf("Expected ShowRate to be true, got false")
	}
	if c.WordBoundary != true {
		t.Errorf("Expected WordBoundary to be true, got false")
	}
}
