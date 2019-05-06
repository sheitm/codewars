package main

import (
	"testing"
)

func TestInterpreter(t *testing.T) {
	type args struct {
		code string
		tape string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{code: "*", tape: "00101100"}, want: "10101100"},
		{name: "2", args: args{code: ">*>*", tape: "00101100"}, want: "01001100"},
		{name: "3", args: args{code: "*>*>*>*>*>*>*>*", tape: "00101100"}, want: "11010011"},
		{name: "4", args: args{code: "*>*>>*>>>*>*", tape: "00101100"}, want: "11111111"},
		{name: "5", args: args{code: ">>>>>*<*<<*", tape: "00101100"}, want: "00000000"},
		{name: "6", args: args{code: "[<<<]*", tape: "0000"}, want: "1000"},
		{name: "7", args: args{code: "*>*>>>*>*>>>>>*>[>*]", tape: "0000000000000000000000000000000000000000000"}, want: "1100110000100000000000000000000000000000000"},
		{name: "8", args: args{code: "[*]", tape: "1"}, want: "0"},
		{name: "9", args: args{code: "[*]", tape: "0"}, want: "0"},
		{name: "10", args: args{code: "[>*]", tape: "1"}, want: "1"},
		{name: "11", args: args{code: "*[>[*]]", tape: "1111111111"}, want: "0111111111"},
		{name: "12", args: args{code: "[[]*>*>*]", tape: "000"}, want: "000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Interpreter(tt.args.code, tt.args.tape); got != tt.want {
				t.Errorf("Interpreter() = %v, want %v", got, tt.want)
			}
		})
	}
}
