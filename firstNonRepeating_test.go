package main

import "testing"

func Test_firstNonRepeating(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{str: "t"}, want: "t"},
		{name: "abba", args: args{str: "abba"}, want: ""},
		{name: "empty", args: args{str: ""}, want: ""},
		{name: "4", args: args{str: "~><#~><"}, want: "#"},
		{name: "5", args: args{str: "hello world, eh?"}, want: "w"},
		{name: "6", args: args{str: "sTreSS"}, want: "T"},
		{name: "7", args: args{str: "Go hang a salami, I'm a lasagna hog!"}, want: ","},
		{name: "8", args: args{str: "sSs"}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstNonRepeating(tt.args.str); got != tt.want {
				t.Errorf("firstNonRepeating() = %v, want %v", got, tt.want)
			}
		})
	}
}