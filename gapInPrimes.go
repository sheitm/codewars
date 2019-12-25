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
		lower := m
		upper := m+g
		p := big.NewInt(int64(lower)).ProbablyPrime(0)
		gp := big.NewInt(int64(upper)).ProbablyPrime(0)
		m = m+2
		if  !(p&&gp)  {
			continue
		}
		c := m
		for {
			if c >= upper {
				return []int{lower, upper}
			}
			if big.NewInt(int64(c)).ProbablyPrime(0) {
				break
			}
			c = c+2
		}
	}
}

