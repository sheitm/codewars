package main

import (
	"reflect"
	"testing"
)

func TestSolvePuzzle(t *testing.T) {
	clues := []int{
		0, 0, 1, 2,
		0, 2, 0, 0,
		0, 3, 0, 0,
		0, 1, 0, 0}
	outcome := [][]int{
		{2, 1, 4, 3},
		{3, 4, 1, 2},
		{4, 2, 3, 1},
		{1, 3, 2, 4}}
	type args struct {
		clues []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "0", args: args{clues: clues}, want: outcome},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolvePuzzle(tt.args.clues); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolvePuzzle() = %v, want %v", got, tt.want)
			}
		})
	}
}
