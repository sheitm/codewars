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

// func Test_board_initialize(t *testing.T) {
// 	b := board{}
// 	b.initialize(4)
// 	c := b.cells[2][2]
// 	if c.val != 0 {
// 		t.Errorf("Expected 0, got %d", c)
// 	}
// }

func Test_matrixToCellsAndBackAgain(t *testing.T) {
	sourceMatrix := [][]int{
		{2, 1, 4, 3},
		{3, 4, 1, 2},
		{4, 2, 3, 1},
		{1, 3, 2, 4}}
	cells := matrixToCells(sourceMatrix)
	matrix := cellsToMatrix(cells)

	if !reflect.DeepEqual(matrix, sourceMatrix) {
		t.Errorf("matrix conversion failed. got %v. want %v", matrix, sourceMatrix)
	}
}

func Test_vectorConversionAndBackAgain(t *testing.T) {
	source := []int{2, 1, 4, 3}
	mv := convertIntVector(source)
	v := convertCellVector(mv)
	if !reflect.DeepEqual(source, v) {
		t.Errorf("vector conversion faile. got %v. expected %v", v, source)
	}
}

func Test_getVectorForIndex(t *testing.T) {
	matrix := [][]int{
		{2, 1, 4, 3},
		{3, 4, 1, 2},
		{4, 2, 3, 1},
		{1, 3, 2, 4}}
	cells := matrixToCells(matrix)
	type args struct {
		cells [][]*tCell
		index int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "0", args: args{cells: cells, index: 0}, want: []int{2, 3, 4, 1}},
		{name: "1", args: args{cells: cells, index: 1}, want: []int{1, 4, 2, 3}},
		{name: "2", args: args{cells: cells, index: 2}, want: []int{4, 1, 3, 2}},
		{name: "3", args: args{cells: cells, index: 3}, want: []int{3, 2, 1, 4}},

		{name: "4", args: args{cells: cells, index: 4}, want: []int{3, 4, 1, 2}},
		{name: "5", args: args{cells: cells, index: 5}, want: []int{2, 1, 4, 3}},
		{name: "6", args: args{cells: cells, index: 6}, want: []int{1, 3, 2, 4}},
		{name: "7", args: args{cells: cells, index: 7}, want: []int{4, 2, 3, 1}},

		{name: "8", args: args{cells: cells, index: 8}, want: []int{4, 1, 2, 3}},
		{name: "9", args: args{cells: cells, index: 9}, want: []int{2, 3, 1, 4}},
		{name: "10", args: args{cells: cells, index: 10}, want: []int{3, 2, 4, 1}},
		{name: "11", args: args{cells: cells, index: 11}, want: []int{1, 4, 3, 2}},

		{name: "12", args: args{cells: cells, index: 12}, want: []int{1, 3, 2, 4}},
		{name: "13", args: args{cells: cells, index: 13}, want: []int{4, 2, 3, 1}},
		{name: "14", args: args{cells: cells, index: 14}, want: []int{3, 4, 1, 2}},
		{name: "15", args: args{cells: cells, index: 15}, want: []int{2, 1, 4, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gCells := getVectorForIndex(tt.args.cells, tt.args.index)
			got := convertCellVector(gCells)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getVectorForIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
