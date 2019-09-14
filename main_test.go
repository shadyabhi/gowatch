package main

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func Test_getOutput(t *testing.T) {
	type args struct {
		cmd string
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		wantErr bool
	}{
		{
			"valid command ls",
			args{cmd: "echo foobar"},
			"foobar\n",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := getOutput(tt.args.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("getOutput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut != tt.wantOut {
				t.Errorf("getOutput() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_decorate(t *testing.T) {
	type args struct {
		outs outputs
	}
	tests := []struct {
		name    string
		args    args
		wantRet string
	}{
		{
			"simple line with no change",
			args{
				outputs{
					prev: "hello",
					cur:  "hello",
				},
			},
			"hello",
		},
		{
			"simple line with one change",
			args{
				outputs{
					prev: "hello 1",
					cur:  "hello 2",
				},
			},
			fmt.Sprintf("hello %s", getHighlightedChar("2")),
		},
		{
			"simple line with change (first one shorter)",
			args{
				outputs{
					prev: "ab",
					cur:  "cd ef",
				},
			},
			getHighlightedString("cd") + " ef",
		},
		{
			"simple line with change (second one shorter)",
			args{
				outputs{
					prev: "ab cd",
					cur:  "xyz",
				},
			},
			getHighlightedString("xyz"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := decorate(tt.args.outs); gotRet != tt.wantRet {
				spew.Dump(gotRet)
				spew.Dump(tt.wantRet)
				t.Errorf("decorate() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
