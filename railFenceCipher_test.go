package main

import (
	"testing"
)

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
		{name: "3", args: args{s: "0123456789", n: 3}, want: "0481357926"},
		{name: "4", args: args{s: "12", n: 3}, want: "12"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 0481357926

func TestDecode(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0", args: args{s: "0481357926", n: 3}, want: "0123456789"},
		{name: "1", args: args{s: "WECRLTEERDSOEEFEAOCAIVDEN", n: 3}, want: "WEAREDISCOVEREDFLEEATONCE"},
		{name: "2", args: args{s: "Hoo!el,Wrdl l", n: 3}, want: "Hello, World!"},
		{name: "4", args: args{s: "12", n: 3}, want: "12"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
