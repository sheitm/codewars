package main

import "testing"

func TestEncode(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "empty", args: args{s: "", n: 3}, want: ""},
		{name: "1", args: args{s: "WEAREDISCOVEREDFLEEATONCE", n: 3}, want: "WECRLTEERDSOEEFEAOCAIVDEN"},
		{name: "2", args: args{s: "Hello, World!", n: 3}, want: "Hoo!el,Wrdl l"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
