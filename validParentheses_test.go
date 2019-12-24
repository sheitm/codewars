package main

import "testing"

func TestValidParentheses(t *testing.T) {
	type args struct {
		parens string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{parens:"()"}, want: true},
		{name: "2", args: args{parens:")(()))"}, want: false},
		{name: "3", args: args{parens:"("}, want: false},
		{name: "4", args: args{parens:"(())((()())())"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidParentheses(tt.args.parens); got != tt.want {
				t.Errorf("ValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}