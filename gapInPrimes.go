package main

// https://www.codewars.com/kata/561e9c843a2ef5a40c0000a4/train/go

import "math/big"

func Gap(g, m, n int) []int {
	if m % 2 == 0 {
		m++
	}
	for {
		if m + g > n {
			return nil
		}
		if big.NewInt(int64(m)).ProbablyPrime(0) && big.NewInt(int64(m+g)).ProbablyPrime(0) {
			return []int{m, m+g}
		}
		m = m+2
	}
}
