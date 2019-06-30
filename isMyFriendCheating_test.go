package main

import (
	"reflect"
	"testing"
)

func TestRemovNb2(t *testing.T) {
	type args struct {
		m uint64
	}
	tests := []struct {
		name string
		args args
		want [][2]uint64
	}{
		{name: "26", args: args{m: 26}, want: [][2]uint64{{15, 21}, {21, 15}}},
		// {name: "100", args: args{m: 100}, want: nil},
		// {name: "101", args: args{m: 101}, want: [][2]uint64{{55, 91}, {91, 55}}},
		// {name: "102", args: args{m: 102}, want: [][2]uint64{{70, 73}, {73, 70}}},
		// {name: "110", args: args{m: 110}, want: [][2]uint64{{70, 85}, {85, 70}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemovNb2(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemovNb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemovNb(t *testing.T) {
	type args struct {
		m uint64
	}
	tests := []struct {
		name string
		args args
		want [][2]uint64
	}{
		{name: "26", args: args{m: 26}, want: [][2]uint64{{15, 21}, {21, 15}}},
		{name: "100", args: args{m: 100}, want: nil},
		{name: "101", args: args{m: 101}, want: [][2]uint64{{55, 91}, {91, 55}}},
		{name: "102", args: args{m: 102}, want: [][2]uint64{{70, 73}, {73, 70}}},
		{name: "110", args: args{m: 110}, want: [][2]uint64{{70, 85}, {85, 70}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemovNb(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemovNb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumExpect(t *testing.T) {
	type args struct {
		lim int
		a   int
		b   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "0", args: args{lim: 4, a: 2, b: 3}, want: 5},
		{name: "1", args: args{lim: 101, a: 55, b: 91}, want: 5005},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumExpect(tt.args.lim, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("sumExpect() = %v, want %v", got, tt.want)
			}
		})
	}
}
