package main

import (
	"reflect"
	"testing"
)

func TestGap(t *testing.T) {
	type args struct {
		g int
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "1", args: args{g: 2, m: 100, n: 110}, want: []int{101, 103}},
		{name: "2", args: args{g: 4, m: 100, n: 110}, want: []int{103, 107}},
		{name: "3", args: args{g: 10, m: 100, n: 110}, want: nil},
		{name: "4", args: args{g: 8, m: 300, n: 400}, want: []int{359, 367}},
		{name: "5", args: args{g: 10, m: 300, n: 400}, want: []int{337, 347}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gap(tt.args.g, tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Gap() = %v, want %v", got, tt.want)
			}
		})
	}
}