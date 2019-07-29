package main

import (
	"reflect"
	"testing"
)

func TestMovingShift(t *testing.T) {
	type args struct {
		s     string
		shift int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "0", args: args{s: "abcde", shift: 1}, want: []string{"b", "d", "f", "h", "j"}},
		{name: "1", args: args{s: "Z", shift: 1}, want: []string{"A", "", "", "", ""}},
		{name: "3", args: args{s: "I should have known that you would have a perfect answer for me!!!", shift: 1}, want: []string{"J vltasl rlhr ", "zdfog odxr ypw", " atasl rlhr p ", "gwkzzyq zntyhv", " lvz wp!!!"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MovingShift(tt.args.s, tt.args.shift); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovingShift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDemovingShift(t *testing.T) {
	type args struct {
		arr   []string
		shift int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0", args: args{arr: []string{"b", "d", "f", "h", "j"}, shift: 1}, want: "abcde"},
		{name: "1", args: args{arr: []string{"A", "", "", "", ""}, shift: 1}, want: "Z"},
		{name: "2", args: args{arr: []string{"J vltasl rlhr ", "zdfog odxr ypw", " atasl rlhr p ", "gwkzzyq zntyhv", " lvz wp!!!"}, shift: 1}, want: "I should have known that you would have a perfect answer for me!!!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DemovingShift(tt.args.arr, tt.args.shift); got != tt.want {
				t.Errorf("DemovingShift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toStringArr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "1", args: args{s: "a"}, want: []string{"a", "", "", "", ""}},
		{name: "11", args: args{s: "abcdefghijk"}, want: []string{"abc", "def", "ghi", "jk", ""}},
		{name: "16", args: args{s: "abcdefghijklmnop"}, want: []string{"abcd", "efgh", "ijkl", "mnop", ""}},
		{name: "17", args: args{s: "abcdefghijklmnopq"}, want: []string{"abcd", "efgh", "ijkl", "mnop", "q"}},
		{name: "20", args: args{s: "abcdefghijklmnopqrst"}, want: []string{"abcd", "efgh", "ijkl", "mnop", "qrst"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toStringArr(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toStringArr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getStringArrLengths(t *testing.T) {
	type args struct {
		l int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "1", args: args{l: 1}, want: []int{1, 0, 0, 0, 0}},
		{name: "11", args: args{l: 11}, want: []int{3, 3, 3, 2, 0}},
		{name: "17", args: args{l: 17}, want: []int{4, 4, 4, 4, 1}},
		{name: "20", args: args{l: 20}, want: []int{4, 4, 4, 4, 4}},
		{name: "21", args: args{l: 21}, want: []int{5, 5, 5, 5, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStringArrLengths(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getStringArrLengths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isLetter(t *testing.T) {
	type args struct {
		l int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "a", args: args{l: 'a'}, want: true},
		{name: "f", args: args{l: 'f'}, want: true},
		{name: "z", args: args{l: 'z'}, want: true},
		{name: "A", args: args{l: 'A'}, want: true},
		{name: "F", args: args{l: 'F'}, want: true},
		{name: "Z", args: args{l: 'Z'}, want: true},
		{name: "1", args: args{l: '1'}, want: false},
		{name: "!", args: args{l: '!'}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLetter(tt.args.l); got != tt.want {
				t.Errorf("isLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decrypt(t *testing.T) {
	type args struct {
		letter int
		shift  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "0", args: args{letter: 66, shift: 1}, want: 65},
		{name: "1", args: args{letter: 63, shift: 1}, want: 63},
		{name: "1", args: args{letter: 65, shift: 1}, want: 90},
		{name: "2", args: args{letter: 122, shift: 5}, want: 117},
		{name: "3", args: args{letter: 97, shift: 2}, want: 121},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decrypt(tt.args.letter, tt.args.shift); got != tt.want {
				t.Errorf("decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
