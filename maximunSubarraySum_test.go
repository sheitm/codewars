package main

import "testing"

func TestMaximumSubarraySum(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "0", args: args{numbers: []int{}}, want: 0},
		{name: "1", args: args{numbers: []int{-2, -1, -3, -4, -1, -2, -1, -5, -4}}, want: 0},
		{name: "2", args: args{numbers: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, want: 6},
		{name: "3", args: args{numbers: []int{-1, -1, -1, 1, -1, -1}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaximumSubarraySum(tt.args.numbers); got != tt.want {
				t.Errorf("MaximumSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
