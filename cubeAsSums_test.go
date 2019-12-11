package main

import (
	"reflect"
	"testing"
)

func TestFindSummands(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "1", args: args{n:1}, want: []int{1}},
		{name: "2", args: args{n:2}, want: []int{3, 5}},
		{name: "3", args: args{n:3}, want: []int{7, 9, 11}},
		{name: "4", args: args{n:4}, want: []int{13, 15, 17, 19}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSummands(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindSummands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindSummandsLargeNumbers(t *testing.T) {
	// Arrange
	large := []int{241, 493, 798, 1000}

	// Act
	for _, i := range large {
		summands := FindSummands(i)
		if len(summands) != i {
			t.Errorf("expected %d elements. got %d", i, len(summands))
		}
	}
}