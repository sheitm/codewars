package main

import "testing"

func TestJosephusSurvivor(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{n: 7, k: 3}, want: 4},
		{name: "2", args: args{n: 40, k: 3}, want: 28},
		{name: "3", args: args{n: 11, k: 19}, want: 10},
		{name: "4", args: args{n: 14, k: 2}, want: 13},
		{name: "5", args: args{n: 100, k: 1}, want: 100},
		{name: "6", args: args{n: 1, k: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JosephusSurvivor(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("JosephusSurvivor() = %v, want %v", got, tt.want)
			}
		})
	}
}

// dotest(11, 19, 10)
// dotest(40, 3, 28)
// dotest(14, 2, 13)
// dotest(100, 1, 100)
