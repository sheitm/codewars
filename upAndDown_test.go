package main

import "testing"

func TestArrange(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{s: ""}, want: ""},
		{name: "2", args: args{s: "who retaining hit that a the we taken"}, want: "who RETAINING hit THAT a THE we TAKEN"},
		{name: "3", args: args{s: "who hit retaining The That a we taken"}, want: "who RETAINING hit THAT a THE we TAKEN"},
		{name: "4", args: args{s: "on I came up were so grandmothers"}, want: "i CAME on WERE up GRANDMOTHERS so"},
		{name: "5", args: args{s: "way the my wall them him"}, want: "way THE my WALL him THEM"},
		{name: "6", args: args{s: "YO"}, want: "yo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Arrange(tt.args.s); got != tt.want {
				t.Errorf("Arrange() = %v, want %v", got, tt.want)
			}
		})
	}
}
