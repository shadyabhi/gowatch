package main

import (
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
