package main

import (
	"reflect"
	"testing"
)

func Test_fibSequence(t *testing.T) {
	type args struct {
		upperBound int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "0", args: args{upperBound: 0}, want: []int{0}},
		{name: "1", args: args{upperBound: 1}, want: []int{0, 1}},
		{name: "2", args: args{upperBound: 0}, want: []int{0, 1, 1}},
		{name: "13", args: args{upperBound: 13}, want: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibSequence(tt.args.upperBound); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fibSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductFib(t *testing.T) {
	type args struct {
		prod uint64
	}
	tests := []struct {
		name string
		args args
		want [3]uint64
	}{
		{name: "714", args: args{prod: 714}, want: [3]uint64{21, 34, 1}},
		{name: "800", args: args{prod: 800}, want: [3]uint64{34, 55, 0}},
		{name: "74049690", args: args{prod: 74049690}, want: [3]uint64{6765, 10946, 1}},
		{name: "84049690", args: args{prod: 84049690}, want: [3]uint64{10946, 17711, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProductFib(tt.args.prod); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductFib() = %v, want %v", got, tt.want)
			}
		})
	}
}
