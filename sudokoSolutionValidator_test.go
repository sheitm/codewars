package main

import (
	"reflect"
	"testing"
)

func TestValidateSolution(t *testing.T) {
	testTrue := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}
	testFalse := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 0, 3, 4, 8},
		{1, 0, 0, 3, 4, 2, 5, 6, 0},
		{8, 5, 9, 7, 6, 1, 0, 2, 0},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 0, 1, 5, 3, 7, 2, 1, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 0, 0, 4, 8, 1, 1, 7, 9},
	}
	testInvalidSector := [][]int{ // bytter 5 og 7 i (0,0) og /1, 1
		{7, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 5, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}
	type args struct {
		m [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "empty", args: args{m: [][]int{}}, want: false},
		{name: "valid", args: args{m: testTrue}, want: true},
		{name: "notValid", args: args{m: testFalse}, want: false},
		{name: "invalidSector", args: args{m: testInvalidSector}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateSolution(tt.args.m); got != tt.want {
				t.Errorf("ValidateSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHVector(t *testing.T) {
	m := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}
	type args struct {
		index int
		m     [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "0", args: args{index: 0, m: m}, want: []int{5, 3, 4, 6, 7, 8, 9, 1, 2}},
		{name: "3", args: args{index: 3, m: m}, want: []int{8, 5, 9, 7, 6, 1, 4, 2, 3}},
		{name: "8", args: args{index: 8, m: m}, want: []int{3, 4, 5, 2, 8, 6, 1, 7, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHVector(tt.args.index, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getVVector(t *testing.T) {
	m := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}
	type args struct {
		index int
		m     [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "0", args: args{index: 0, m: m}, want: []int{5, 6, 1, 8, 4, 7, 9, 2, 3}},
		{name: "3", args: args{index: 3, m: m}, want: []int{6, 1, 3, 7, 8, 9, 5, 4, 2}},
		{name: "8", args: args{index: 8, m: m}, want: []int{2, 8, 7, 3, 1, 6, 4, 5, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getVVector(tt.args.index, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getVVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSVector(t *testing.T) {
	m := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}
	type args struct {
		index int
		m     [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "0", args: args{index: 0, m: m}, want: []int{5, 6, 1, 3, 7, 9, 4, 2, 8}},
		{name: "8", args: args{index: 8, m: m}, want: []int{2, 6, 1, 8, 3, 7, 4, 5, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSVector(tt.args.index, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSVector() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_vectorIsValid(t *testing.T) {
	type args struct {
		v []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "empty", args: args{v: []int{}}, want: false},
		{name: "short", args: args{v: []int{1, 2, 3, 4, 5, 6, 7, 8}}, want: false},
		{name: "long", args: args{v: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}, want: false},
		{name: "invalid1", args: args{v: []int{5, 3, 4, 6, 7, 8, 9, 1, 0}}, want: false},
		{name: "invalid2", args: args{v: []int{2, 10, 7, 4, 1, 9, 6, 3, 5}}, want: false},
		{name: "invalid3", args: args{v: []int{2, 8, 7, 5, 1, 9, 6, 3, 5}}, want: false},
		{name: "valid", args: args{v: []int{2, 8, 7, 4, 1, 9, 6, 3, 5}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vectorIsValid(tt.args.v); got != tt.want {
				t.Errorf("vectorIsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSVectorIndices(t *testing.T) {
	type args struct {
		sectorIndex int
		index       int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{name: "0_0", args: args{sectorIndex: 0, index: 0}, want: 0, want1: 0},
		{name: "0_1", args: args{sectorIndex: 0, index: 1}, want: 1, want1: 0},
		{name: "2_2", args: args{sectorIndex: 2, index: 2}, want: 8, want1: 0},
		{name: "8_8", args: args{sectorIndex: 8, index: 8}, want: 8, want1: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getSVectorIndices(tt.args.sectorIndex, tt.args.index)
			if got != tt.want {
				t.Errorf("getSVectorIndices() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getSVectorIndices() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}