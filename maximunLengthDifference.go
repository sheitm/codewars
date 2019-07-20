package main

// https://www.codewars.com/kata/5663f5305102699bad000056/train/go

// MxDifLg solve it
func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	min1, max1 := getMinMax(a1)
	min2, max2 := getMinMax(a2)
	if (max1 - min2) > (max2 - min1) {
		return max1 - min2
	}
	return max2 - min1
}

func getMinMax(a []string) (int, int) {
	min := 4294967295
	max := -1
	for _, s := range a {
		l := len(s)
		if l < min {
			min = l
		}
		if l > max {
			max = l
		}
	}
	return min, max
}
