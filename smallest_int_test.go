package main

import (
	"reflect"
	"testing"
)

func TestSmallest(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{name: "1", args: args{n: 261235}, want: []int64{126235, 2, 0}},
		{name: "2", args: args{n: 209917}, want: []int64{29917, 0, 1}},
		{name: "3", args: args{n: 132}, want: []int64{123, 1, 2}},
		{name: "3", args: args{n: 111111111}, want: []int64{111111111, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Smallest(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Smallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
