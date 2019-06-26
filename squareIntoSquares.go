package main

import "sort"

// https://www.codewars.com/kata/54eb33e5bc1a25440d000891/train/go

// Decompose return squares into squares
func Decompose(n int64) []int64 {
	if n <= 2 {
		return []int64{}
	}

	targetSqr := n * n
	for i := 1; int64(i) < n; i++ {
		arr, ok := decomposeDownFrom(n-int64(i), targetSqr, []int64{})
		if ok {
			sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
			return arr
		}
	}

	return []int64{}
}

func decomposeDownFrom(upper int64, targetSqr int64, path []int64) ([]int64, bool) {
	path = append(path, upper)
	sum := squareSum(path)
	if sum == targetSqr {
		return path, true
	}
	if sum > targetSqr {
		return []int64{}, false
	}

	for i := 1; int64(i) < upper; i++ {
		arr, ok := decomposeDownFrom(upper-int64(i), targetSqr, path)
		if ok {
			return arr, true
		}
	}

	return []int64{}, false
}

func squareSum(arr []int64) int64 {
	var s int64
	for i := 0; i < len(arr); i++ {
		sqr := arr[i] * arr[i]
		s = s + sqr
	}
	return s
}
