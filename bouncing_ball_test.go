package main

import "testing"

func TestBouncingBall(t *testing.T) {
	type args struct {
		h      float64
		bounce float64
		window float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{h: 3, bounce: 0.66, window: 1.5}, want: 3},
		{name: "2", args: args{h: 30, bounce: 0.4, window: 10}, want: 3},
		{name: "3", args: args{h: 10, bounce: 0.6, window: 10}, want: -1},
		{name: "4", args: args{h: 40, bounce: 1, window: 10}, want: -1},
		{name: "5", args: args{h: 5, bounce: -1, window: 1.5}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BouncingBall(tt.args.h, tt.args.bounce, tt.args.window); got != tt.want {
				t.Errorf("BouncingBall() = %v, want %v", got, tt.want)
			}
		})
	}
}
