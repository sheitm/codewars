package main

import "testing"

func TestFindNb(t *testing.T) {
	type args struct {
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "not-ok-1", args: args{m: 15}, want: -1},
		{name: "ok-0", args: args{m: 36}, want: 3},
		{name: "ok-1", args: args{m: 1071225}, want: 45},
		{name: "ok-2", args: args{m: 4183059834009}, want: 2022},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindNb(tt.args.m); got != tt.want {
				t.Errorf("FindNb() = %v, want %v", got, tt.want)
			}
		})
	}
}
