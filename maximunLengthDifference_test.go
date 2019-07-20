package main

import "testing"

func TestMxDifLg(t *testing.T) {
	s1 := []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
	s2 := []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}
	ss1 := []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}
	ss2 := []string{"bbbaaayddqbbrrrv"}
	type args struct {
		a1 []string
		a2 []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "0", args: args{a1: s1, a2: s2}, want: 13},
		{name: "1", args: args{a1: ss1, a2: ss2}, want: 10},
		{name: "2", args: args{a1: []string{}, a2: ss2}, want: -1},
		{name: "3", args: args{a1: ss1, a2: []string{}}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MxDifLg(tt.args.a1, tt.args.a2); got != tt.want {
				t.Errorf("MxDifLg() = %v, want %v", got, tt.want)
			}
		})
	}
}
