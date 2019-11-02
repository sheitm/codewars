package main

import "testing"

func TestDecodex(t *testing.T) {
	type args struct {
		roman string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{roman: "XXI"}, want: 21},
		{name: "2", args: args{roman: "I"}, want: 1},
		{name: "3", args: args{roman: "IV"}, want: 4},
		{name: "4", args: args{roman: "MMVIII"}, want: 2008},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RomanNumeralsDecode(tt.args.roman); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
