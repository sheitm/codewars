package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestBuddy(t *testing.T) {
	type args struct {
		start int
		limit int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "0", args: args{start: 10, limit: 50}, want: []int{48, 75}},
		{name: "1", args: args{start: 48, limit: 50}, want: []int{48, 75}},
		{name: "2", args: args{start: 1071625, limit: 1103735}, want: []int{1081184, 1331967}},
		{name: "3", args: args{start: 57345, limit: 90061}, want: []int{62744, 75495}},
		{name: "4", args: args{start: 2693, limit: 7098}, want: []int{5775, 6128}},
		{name: "5", args: args{start: 6379, limit: 8275}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Buddy(tt.args.start, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Buddy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDivisors(t *testing.T) {
	type args struct {
		d int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "48", args: args{d: 48}, want: []int{1, 2, 3, 4, 6, 8, 12, 16, 24}},
		{name: "75", args: args{d: 75}, want: []int{1, 3, 5, 15, 25}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDivisors(tt.args.d)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDivisorSum(t *testing.T) {
	type args struct {
		d int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "48", args: args{d: 48}, want: 76},
		{name: "76", args: args{d: 75}, want: 49},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDivisorSum(tt.args.d); got != tt.want {
				t.Errorf("getDivisorSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
