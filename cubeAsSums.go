package main

import "sort"

// https://www.codewars.com/kata/58af1bb7ac7e31b192000020/train/go

func FindSummands(n int) []int {
	cube := n*n*n
	low  := cube / n
	hi := low
	res := []int{}
	if low%2 == 0 {
		low--
		hi++
		res = append(res, low)
		res = append(res, hi)
	} else {
		res = append(res, low)
	}

	for {
		if len(res) == n {
			sort.Ints(res)
			return res
		}
		low = low - 2
		hi = hi + 2
		res = append(res, low)
		res = append(res, hi)
	}
}