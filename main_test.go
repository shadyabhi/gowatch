package main

import (
	"fmt"
	"testing"
)

func Test_getCMDOutput(t *testing.T) {
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
			gotOut, err := getCMDOutput(tt.args.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCMDOutput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut != tt.wantOut {
				t.Errorf("getCMDOutput() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
func Test_outputs_printCharWise(t *testing.T) {
	type fields struct {
		prev    string
		cur     string
		prevPos [][]int
		curPos  [][]int
		i       int
	}
	tests := []struct {
		name    string
		fields  fields
		wantRet string
	}{
		{
			"simple line with no change",
			fields{
				prev: "hello",
				cur:  "hello",
			},
			"hello",
		},
		{
			"simple line with one change",
			fields{
				prev: "hello 1",
				cur:  "hello 2",
			},
			fmt.Sprintf("hello %s", getHighlightedChar("2")),
		},
		{
			"simple line with change (first one shorter)",
			fields{
				prev: "ab",
				cur:  "cd ef",
			},
			getHighlightedString("cd") + " ef",
		},
		{
			"simple line with change (second one shorter)",
			fields{
				prev: "ab cd",
				cur:  "xyz",
			},
			getHighlightedString("xyz"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &outputs{
				prev:    tt.fields.prev,
				cur:     tt.fields.cur,
				prevPos: tt.fields.prevPos,
				curPos:  tt.fields.curPos,
				i:       tt.fields.i,
			}
			if gotRet := o.printCharWise(); gotRet != tt.wantRet {
				t.Errorf("outputs.printWordWise() = %#v, want %#v", gotRet, tt.wantRet)
			}
		})
	}
}

func Test_outputs_printWordWise(t *testing.T) {
	type fields struct {
		prev    string
		cur     string
		prevPos [][]int
		curPos  [][]int
		i       int
	}
	tests := []struct {
		name    string
		fields  fields
		wantRet string
	}{
		{
			"same line",
			fields{
				prev: " hello  world",
				cur:  " hello  world",
			},
			" hello  world",
		},
		{
			"difference in start",
			fields{
				prev: " foo  world",
				cur:  " hello  world",
			},
			" " + getHighlightedString("hello") + "  world",
		},
		{
			"difference in middle",
			fields{
				prev: " foo 1 world",
				cur:  " foo 2 world",
			},
			" foo " + getHighlightedString("2") + " world",
		},
		{
			"muliple difference in start, middle, end (final string long)",
			fields{
				prev: " foo 1 world",
				cur:  " hello 2 world 3",
			},
			" " + getHighlightedString("hello") + " " + getHighlightedString("2") + " world " + getHighlightedString("3"),
		},
		{
			"muliple difference in start, middle, end (final string short)",
			fields{
				prev: "verylong 1 world",
				cur:  "hey 1",
			},
			getHighlightedString("hey") + " 1",
		},
		{
			"simple with whitespaces",
			fields{
				prev: "hello\nworld\t\t",
				cur:  "foo\nworld",
			},
			getHighlightedString("foo") + "\nworld",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &outputs{
				prev:    tt.fields.prev,
				cur:     tt.fields.cur,
				prevPos: tt.fields.prevPos,
				curPos:  tt.fields.curPos,
				i:       tt.fields.i,
			}
			if gotRet := o.printWordWise(); gotRet != tt.wantRet {
				t.Errorf("outputs.printWordWise() = %#v, want %#v", gotRet, tt.wantRet)
			}
		})
	}
}
