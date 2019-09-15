package main

import "testing"

func Test_getFloat(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		args   args
		wantF  float64
		wantIs bool
	}{
		{"valid float", args{s: "1.0"}, 1.0, true},
		{"invalid float", args{s: "foo"}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotF, gotIs := getFloat(tt.args.s)
			if gotF != tt.wantF {
				t.Errorf("isFloat() gotF = %v, want %v", gotF, tt.wantF)
			}
			if gotIs != tt.wantIs {
				t.Errorf("isFloat() gotIs = %v, want %v", gotIs, tt.wantIs)
			}
		})
	}
}
