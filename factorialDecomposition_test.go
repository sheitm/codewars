package main

import (
	"reflect"
	"testing"
)

func TestDecomp(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "5", args: args{n: 5}, want: "2^3 * 3 * 5"},
		{name: "14", args: args{n: 14}, want: "2^19 * 3^9 * 5^4 * 7^3 * 11^2 * 13 * 17 * 19"},
		{name: "17", args: args{n: 17}, want: "2^15 * 3^6 * 5^3 * 7^2 * 11 * 13 * 17"},
		// {name: "22", args: args{n: 22}, want: "2^19 * 3^9 * 5^4 * 7^3 * 11^2 * 13 * 17 * 19"},
		// {name: "25", args: args{n: 25}, want: "2^22 * 3^10 * 5^6 * 7^3 * 11^2 * 13 * 17 * 19 * 23"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decomp(tt.args.n); got != tt.want {
				t.Errorf("Decomp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSmallestPrimeFactor(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "2", args: args{n: 2}, want: 2},
		{name: "3", args: args{n: 3}, want: 3},
		{name: "25", args: args{n: 25}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSmallestPrimeFactor(tt.args.n); got != tt.want {
				t.Errorf("getSmallestPrimeFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decomposeIntoPrimes(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{name: "2", args: args{n: 2}, want: []int64{2}},
		{name: "3", args: args{n: 3}, want: []int64{3}},
		{name: "8", args: args{n: 8}, want: []int64{2, 2, 2}},
		{name: "25", args: args{n: 25}, want: []int64{5, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decomposeIntoPrimes(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decomposeIntoPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toFactorialString(t *testing.T) {
	type args struct {
		arr []int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0", args: args{arr: []int64{2, 3}}, want: "2 * 3"},
		{name: "2", args: args{arr: []int64{2, 2, 2, 2}}, want: "2^4"},
		{name: "3", args: args{arr: []int64{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 5, 5, 5, 7, 7, 7, 11, 11, 13, 17, 19}}, want: "2^19 * 3^9 * 5^4 * 7^3 * 11^2 * 13 * 17 * 19"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toFactorialString(tt.args.arr); got != tt.want {
				t.Errorf("toFactorialString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_factorial(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "0", args: args{n: 0}, want: 1},
		{name: "1", args: args{n: 1}, want: 1},
		{name: "5", args: args{n: 5}, want: 120},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorial(tt.args.n); got != tt.want {
				t.Errorf("factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
