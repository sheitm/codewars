package main

import "testing"

func TestCircleOfNumbers(t *testing.T) {
	type args struct {
		n           int
		firstNumber int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{n: 10, firstNumber: 2}, want: 7},
		{name: "2", args: args{n: 10, firstNumber: 7}, want: 2},
		{name: "3", args: args{n: 4, firstNumber: 1}, want: 3},
		{name: "4", args: args{n: 6, firstNumber: 3}, want: 0},
		{name: "5", args: args{n: 20, firstNumber: 0}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CircleOfNumbers(tt.args.n, tt.args.firstNumber); got != tt.want {
				t.Errorf("CircleOfNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}