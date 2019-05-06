package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		str string
		idx uint
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		{name: "1", args: args{str: "((1)23(45))(aB)", idx: 0}, want: 10},
		{name: "2", args: args{str: "((1)23(45))(aB)", idx: 1}, want: 3},
		{name: "3", args: args{str: "((1)23(45))(aB)", idx: 2}, wantErr: true},
		{name: "4", args: args{str: "((1)23(45))(aB)", idx: 6}, want: 9},
		{name: "5", args: args{str: "((1)23(45))(aB)", idx: 11}, want: 14},
		{name: "6", args: args{str: "(g(At)IO(f)(tM(qk)YF(n)Nr(E)))", idx: 11}, want: 28},
		{name: "7", args: args{str: "(g(At)IO(f)(tM(qk)YF(n)Nr(E)))", idx: 0}, want: 29},
		{name: "8", args: args{str: "(>_(va)`?(h)C(as)(x(hD)P|(fg)))", idx: 19}, want: 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Solution(tt.args.str, tt.args.idx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Solution() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
