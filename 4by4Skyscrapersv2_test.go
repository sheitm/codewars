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

func Test_permutations(t *testing.T) {
	p := permutations{}
	p.init()

	if len(p.perms) != 24 {
		t.Errorf("unexpected perms. got %v", p.perms)
	}
}

func Test_getMustHaves(t *testing.T) {
	clues := []int{
		0, 0, 1, 2,
		0, 2, 0, 0,
		0, 3, 0, 0,
		0, 1, 0, 0}
	type args struct {
		rowIndex int
		clues    []int
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{name: "0", args: args{rowIndex: 0, clues: clues}, want: map[int]int{2: 4}},
		{name: "1", args: args{rowIndex: 1, clues: clues}, want: map[int]int{}},
		{name: "2", args: args{rowIndex: 2, clues: clues}, want: map[int]int{0: 4}},
		{name: "3", args: args{rowIndex: 3, clues: clues}, want: map[int]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMustHaves(tt.args.rowIndex, tt.args.clues); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMustHaves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_vIsValid(t *testing.T) {
	type args struct {
		r []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "0", args: args{r: []int{1, 2, 3, 4}}, want: true},
		{name: "1", args: args{r: []int{4, 1, 3, 2}}, want: true},
		{name: "2", args: args{r: []int{4, 1, 3, 4}}, want: false},
		{name: "3", args: args{r: []int{}}, want: false},
		{name: "4", args: args{r: []int{10, 11, 12, 344}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vIsValid(tt.args.r); got != tt.want {
				t.Errorf("vIsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countVisibleTowers(t *testing.T) {
	type args struct {
		v []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "0", args: args{v: []int{2, 1, 4, 3}}, want: 2},
		{name: "1", args: args{v: []int{1, 2, 3, 4}}, want: 4},
		{name: "2", args: args{v: []int{4, 2, 3, 1}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countVisibleTowers(tt.args.v); got != tt.want {
				t.Errorf("countVisibleTowers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getVectorForIndex(t *testing.T) {
	cells := [][]int{
		{2, 1, 4, 3},
		{3, 4, 1, 2},
		{4, 2, 3, 1},
		{1, 3, 2, 4}}
	type args struct {
		cells [][]int
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
			if got := getVectorForIndex(tt.args.cells, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getVectorForIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mIsValid(t *testing.T) {
	okMatrix := [][]int{
		{2, 1, 4, 3},
		{3, 4, 1, 2},
		{4, 2, 3, 1},
		{1, 3, 2, 4}}
	notOkkMatrix1 := [][]int{
		{2, 1, 4, 3},
		{3, 4, 1, 2},
		{4, 2, 3, 1},
		{3, 1, 2, 4}}
	notOkkMatrix2 := [][]int{
		{2, 1, 4, 3},
		{3, 4, 1, 2},
		{4, 2, 3, 1},
		{1, 3, 2, 33}}
	type args struct {
		m [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "ok", args: args{m: okMatrix}, want: true},
		{name: "not-ok", args: args{m: notOkkMatrix1}, want: false},
		{name: "not-ok-2", args: args{m: notOkkMatrix2}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mIsValid(tt.args.m); got != tt.want {
				t.Errorf("mIsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mIsValidWithClues(t *testing.T) {
	clues := []int{
		0, 0, 1, 2,
		0, 2, 0, 0,
		0, 3, 0, 0,
		0, 1, 0, 0}
	wrongClues := []int{
		0, 0, 1, 2,
		0, 2, 0, 0,
		0, 3, 0, 0,
		0, 2, 0, 0}
	m := [][]int{
		{2, 1, 4, 3},
		{3, 4, 1, 2},
		{4, 2, 3, 1},
		{1, 3, 2, 4}}
	type args struct {
		m     [][]int
		clues []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "ok", args: args{m: m, clues: clues}, want: true},
		{name: "not-ok", args: args{m: m, clues: wrongClues}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mIsValidWithClues(tt.args.m, tt.args.clues); got != tt.want {
				t.Errorf("mIsValidWithClues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_permutations_getValidPermutations(t *testing.T) {
	p := permutations{}
	p.init()

	mAll := map[int]int{0: 1, 1: 2, 2: 3, 3: 4}
	mNone := map[int]int{}
	mOne := map[int]int{0: 4}

	perms := p.getValidPermutations(mAll)
	if len(perms) != 1 {
		t.Errorf("unexpected length. want 1, got %d", len(perms))
	}

	perms = p.getValidPermutations(mNone)
	if len(perms) != 24 {
		t.Errorf("unexpected length. want 24, got %d", len(perms))
	}

	perms = p.getValidPermutations(mOne)
	if len(perms) != 6 {
		t.Errorf("unexpected length. want 24, got %d", len(perms))
	}
}
