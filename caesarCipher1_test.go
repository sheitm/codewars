package main

import (
	"reflect"
	"testing"
)

func TestMovingShift(t *testing.T) {
	type args struct {
		s     string
		shift int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MovingShift(tt.args.s, tt.args.shift); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovingShift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toStringArr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "1", args: args{s: "a"}, want: []string{"a", "", "", "", ""}},
		{name: "11", args: args{s: "abcdefghijk"}, want: []string{"abc", "def", "ghi", "jk", ""}},
		{name: "16", args: args{s: "abcdefghijklmnop"}, want: []string{"abcd", "efgh", "ijkl", "mnop", ""}},
		{name: "17", args: args{s: "abcdefghijklmnopq"}, want: []string{"abcd", "efgh", "ijkl", "mnop", "q"}},
		{name: "20", args: args{s: "abcdefghijklmnopqrst"}, want: []string{"abcd", "efgh", "ijkl", "mnop", "qrst"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toStringArr(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toStringArr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getStringArrLengths(t *testing.T) {
	type args struct {
		l int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "1", args: args{l: 1}, want: []int{1, 0, 0, 0, 0}},
		{name: "11", args: args{l: 11}, want: []int{3, 3, 3, 2, 0}},
		{name: "17", args: args{l: 17}, want: []int{4, 4, 4, 4, 1}},
		{name: "20", args: args{l: 20}, want: []int{4, 4, 4, 4, 4}},
		{name: "21", args: args{l: 21}, want: []int{5, 5, 5, 5, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStringArrLengths(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getStringArrLengths() = %v, want %v", got, tt.want)
			}
		})
	}
}
