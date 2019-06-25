package main

import (
	"fmt"
	"math"
	"sort"
)

// https://www.codewars.com/kata/help-your-granny/train/go

// Tour tour the towns
func Tour(arrFriends []string, ftwns map[string]string, h map[string]float64) int {
	keys := getSortedKeys(h)
	distances := map[string]float64{}
	for i := 0; i < len(keys)-1; i++ {
		k := fmt.Sprintf("%s%s", keys[i], keys[i+1])
		d := findDistance(h[keys[i]], h[keys[i+1]])
		distances[k] = d
	}
	return 1
}

func findDistance(v1, v2 float64) float64 {
	var h, a float64
	if v1 > v2 {
		h = v1
		a = v2
	} else {
		h = v2
		a = v1
	}

	l := math.Pow(h, 2) - math.Pow(a, 2)
	return math.Sqrt(l)
}

func getSortedKeys(h map[string]float64) []string {
	keys := []string{}
	for k := range h {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
