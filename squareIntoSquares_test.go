package main

import (
	"reflect"
	"testing"
)

func TestDecompose(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{name: "2", args: args{n: 2}, want: []int64{}},
		{name: "5", args: args{n: 5}, want: []int64{3, 4}},
		{name: "7", args: args{n: 7}, want: []int64{2, 3, 6}},
		{name: "12", args: args{n: 12}, want: []int64{1, 2, 3, 7, 9}},
		{name: "50", args: args{n: 50}, want: []int64{1, 3, 5, 8, 49}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decompose(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decompose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_squareSum(t *testing.T) {
	type args struct {
		arr []int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "0", args: args{arr: []int64{}}, want: 0},
		{name: "1", args: args{arr: []int64{1}}, want: 1},
		{name: "2", args: args{arr: []int64{1, 4, 9}}, want: 98},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squareSum(tt.args.arr); got != tt.want {
				t.Errorf("squareSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_decomposeDownFrom(t *testing.T) {
// 	type args struct {
// 		upper     int64
// 		targetSqr int64
// 		path      []int64
// 	}
// 	tests := []struct {
// 		name  string
// 		args  args
// 		want  []int64
// 		want1 bool
// 		want2 bool
// 	}{
// 		{name: "5", args: args{upper: 4, targetSqr: 25, path: []int64{}}, want: []int64{4, 3}, want1: false, want2: true},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, got1, got2 := decomposeDownFrom(tt.args.upper, tt.args.targetSqr, tt.args.path)
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("decomposeDownFrom() got = %v, want %v", got, tt.want)
// 			}
// 			if got1 != tt.want1 {
// 				t.Errorf("decomposeDownFrom() got1 = %v, want %v", got1, tt.want1)
// 			}
// 			if got2 != tt.want2 {
// 				t.Errorf("decomposeDownFrom() got2 = %v, want %v", got2, tt.want2)
// 			}
// 		})
// 	}
// }
