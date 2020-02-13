package main

import "testing"

func TestSumOfDivided(t *testing.T) {
	type args struct {
		lst []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{lst: []int{12, 15}}, want: "(2 12)(3 27)(5 15)"},
		{name: "2", args: args{lst: []int{15,21,24,30,45}}, want: "(2 54)(3 135)(5 90)(7 21)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfDivided(tt.args.lst); got != tt.want {
				t.Errorf("SumOfDivided() = %v, want %v", got, tt.want)
			}
		})
	}
}