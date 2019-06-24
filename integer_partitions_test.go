package main

import (
	"testing"
)

func TestPart(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{n: 1}, want: "Range: 0 Average: 1.00 Median: 1.00"},
		{name: "2", args: args{n: 2}, want: "Range: 1 Average: 1.50 Median: 1.50"},
		{name: "3", args: args{n: 3}, want: "Range: 2 Average: 2.00 Median: 2.00"},
		{name: "4", args: args{n: 4}, want: "Range: 3 Average: 2.50 Median: 2.50"},
		{name: "5", args: args{n: 5}, want: "Range: 5 Average: 3.50 Median: 3.50"},
		{name: "6", args: args{n: 6}, want: "Range: 8 Average: 4.75 Median: 4.50"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part(tt.args.n); got != tt.want {
				t.Errorf("Part() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partition_set(t *testing.T) {
	// Arrange
	d := []int{3, 4, 1, 1, 2, 0, 0, 0, 0, 0, 0}
	p := partition{}
	p.set(d)

	// Assert
	if len(p.diagram) != 11 {
		t.Errorf("unexpected diagram length. expected 11. got %d", len(p.diagram))
	}
	if p.sum != 11 {
		t.Errorf("expected sum to be 11. got %d", p.sum)
	}
	if p.key != "00000011234" {
		t.Errorf("unexpected key. got %s", p.key)
	}
}

func Test_makePartitions(t *testing.T) {
	partitions := makePartitions(2)
	_ = partitions
}

func Test_partition_make(t *testing.T) {
	// Arrange
	n := 4
	p := partition{}

	// Act
	p.make(n)

	// Assert
	if len(p.diagram) != 4 {
		t.Errorf("unexpected diagram length. expected 4. got %d", len(p.diagram))
	}
	if p.sum != 4 {
		t.Errorf("expected sum to be 4. got %d", p.sum)
	}
	if p.key != "1111" {
		t.Errorf("unexpected key. got %s", p.key)
	}
}
