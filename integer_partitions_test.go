package main

import "testing"

func TestPart(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{n: 1}, want: "Range: 0 Average: 1.00 Median: 1.00"},
		{name: "2", args: args{n: 2}, want: "Range: 1 Average: 1.50 Median: 1.50"},
		{name: "3", args: args{n: 3}, want: "Range: 2 Average: 2.00 Median: 2.00"},
		{name: "4", args: args{n: 4}, want: "Range: 3 Average: 2.50 Median: 2.50"},
		{name: "5", args: args{n: 5}, want: "Range: 5 Average: 3.50 Median: 3.50"},
		{name: "6", args: args{n: 6}, want: "Range: 8 Average: 4.75 Median: 4.50"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part(tt.args.n); got != tt.want {
				t.Errorf("Part() = %v, want %v", got, tt.want)
			}
		})
	}
}
