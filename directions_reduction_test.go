package main

import (
	"reflect"
	"testing"
)

func TestDirReduc(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "1", args: args{arr: []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}}, want: []string{"WEST"}},
		{name: "2", args: args{arr: []string{"NORTH", "WEST", "SOUTH", "EAST"}}, want: []string{}},
		{name: "3", args: args{arr: []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"}}, want: []string{"NORTH"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DirReduc(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DirReduc() = %v, want %v", got, tt.want)
			}
		})
	}
}
