package main

import (
	"reflect"
	"testing"
)

func TestPickPeaks(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want PosPeaks
	}{
		{name: "1", args: args{array: []int{1, 2, 3, 6, 4, 1, 2, 3, 2, 1}}, want: PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}}},
		{name: "2", args: args{array: []int{}}, want: PosPeaks{Pos: []int{}, Peaks: []int{}}},
		{name: "3", args: args{array: []int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3}}, want: PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}}},
		{name: "4", args: args{array: []int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1}}, want: PosPeaks{Pos: []int{3, 7, 10}, Peaks: []int{6, 3, 2}}},
		{name: "5", args: args{array: []int{2, 1, 3, 1, 2, 2, 2, 2, 1}}, want: PosPeaks{Pos: []int{2, 4}, Peaks: []int{3, 2}}},
		{name: "6", args: args{array: []int{2, 1, 3, 1, 2, 2, 2, 2}}, want: PosPeaks{Pos: []int{2}, Peaks: []int{3}}},
		{name: "7", args: args{array: []int{1, 2, 2, 2, 3, 1}}, want: PosPeaks{Pos: []int{4}, Peaks: []int{3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PickPeaks(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PickPeaks() = %v, want %v", got, tt.want)
			}
		})
	}
}

// otest(
// 	[]int{1, 2, 3, 6, 4, 1, 2, 3, 2, 1},
// 	PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}},

func Test_toSlopeArray(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "1", args: args{array: []int{1, 2, 3, 6, 4, 1, 2, 3, 2, 1}}, want: []string{"0", ">", ">", ">", "<", "<", ">", ">", "<", "<"}},
		{name: "1", args: args{array: []int{1, 2, 2, 2, 3, 1}}, want: []string{"0", ">", "=", "=", ">", "<"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toSlopeArray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toSlopeArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
