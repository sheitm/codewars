package main

import (
	"testing"
)

func Test_stringToBin(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name          string
		args          args
		wantBinString string
	}{
		{name: "1", args: args{s: "test"}, wantBinString: "1110100110010111100111110100"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBinString := stringToBin(tt.args.s); gotBinString != tt.wantBinString {
				t.Errorf("stringToBin() = %v, want %v", gotBinString, tt.wantBinString)
			}
		})
	}
}

func Test_bitStringToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "0", args: args{s: "0"}, want: 0},
		{name: "3", args: args{s: "11"}, want: 3},
		{name: "65", args: args{s: "1000001"}, want: 65},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitStringToInt(tt.args.s); got != tt.want {
				t.Errorf("bitStringToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeToString(t *testing.T) {
	type args struct {
		d []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: args{d: []byte("test")}, want: "fPNKd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeToString(tt.args.d); got != tt.want {
				t.Errorf("EncodeToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
