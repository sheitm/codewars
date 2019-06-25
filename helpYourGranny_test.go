package main

import (
	"testing"
)

func TestTour(t *testing.T) {
	friends1 := []string{"A1", "A2", "A3", "A4", "A5"}
	fTowns1 := map[string]string{"A1": "X1", "A2": "X2", "A3": "X3", "A4": "X4"}
	dist1 := map[string]float64{"X1": 100.0, "X2": 200.0, "X3": 250.0, "X4": 300.0}

	// friends2 := []string{"B1", "B2"}
	// fTowns2 := map[string]string{"B1": "Y1", "B2": "Y2", "B3": "Y3", "B4": "Y4", "B5": "Y5"}
	// dist2 := map[string]float64{"Y1": 50.0, "Y2": 70.0, "Y3": 90.0, "Y4": 110.0, "Y5": 150.0}

	type args struct {
		arrFriends []string
		ftwns      map[string]string
		h          map[string]float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{arrFriends: friends1, ftwns: fTowns1, h: dist1}, want: 889},
		// {name: "2", args: args{arrFriends: friends2, ftwns: fTowns2, h: dist2}, want: 168},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tour(tt.args.arrFriends, tt.args.ftwns, tt.args.h); got != tt.want {
				t.Errorf("Tour() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDistance(t *testing.T) {
	type args struct {
		v1 float64
		v2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "1", args: args{v1: 5, v2: 3}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDistance(tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("findDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
