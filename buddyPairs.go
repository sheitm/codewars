package main

import "math"

// https://www.codewars.com/kata/buddy-pairs/train/go

// Buddy returns the first buddy pair
func Buddy(start, limit int) []int {
	for i := start; i <= limit; i++ {
		ds := getDivisorSum(i)
		if ds > i {
			ds2 := getDivisorSum(ds - 1)
			if ds2-1 == i {
				return []int{i, ds - 1}
			}
		}
	}
	return []int{}
}

func getDivisorSum(d int) int {
	s := 0
	for _, i := range findDivisors(d) {
		s += i
	}
	return s
}

func findDivisors(d int) []int {
	divisors := []int{}
	lim := int(math.Sqrt(float64(d))) + 1
	for i := 1; i < lim; i++ {
		if d%i == 0 {
			divisors = append(divisors, i)
			if d/i != i && i != 1 {
				divisors = append(divisors, d/i)
			}
		}
	}
	return divisors
}
