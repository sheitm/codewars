package main

import (
	"fmt"
	"math/big"
)

// https://www.codewars.com/kata/5a045fee46d843effa000070/train/go

const biggestPrime = 1212121

// Decomp decompose into prime factors
func Decomp(n int) string {
	f := factorial(n)
	p := decomposeIntoPrimes(f)
	return toFactorialString(p)
}

func factorial(n int) int64 {
	if n == 0 {
		return 1
	}
	f := 1
	for i := 1; i <= n; i++ {
		f = f * i
	}
	return int64(f)
}

func decomposeIntoPrimes(n int64) []int64 {
	factors := []int64{}
	for {
		p := getSmallestPrimeFactor(n)
		factors = append(factors, p)
		if p == n {
			return factors
		}
		n = n / p
	}
}

func getSmallestPrimeFactor(n int64) int64 {
	p := 2
	for {
		attempts := 0
		if p > biggestPrime {
			attempts = 20
		}
		if big.NewInt(int64(p)).ProbablyPrime(attempts) {
			if n%int64(p) == 0 {
				return int64(p)
			}
		}
		p++
	}
}

func toFactorialString(arr []int64) string {
	if len(arr) == 1 {
		return fmt.Sprintf("%v", arr[0])
	}
	n := arr[0]
	c := 1
	sarr := []string{}
	for i := 1; i < len(arr); i++ {
		if arr[i] == n {
			c++
			continue
		}
		if c == 1 {
			sarr = append(sarr, fmt.Sprintf("%v", n))
		} else {
			sarr = append(sarr, fmt.Sprintf("%v^%v", n, c))
		}
		c = 1
		n = arr[i]
	}
	if c == 1 {
		sarr = append(sarr, fmt.Sprintf("%v", n))
	} else {
		sarr = append(sarr, fmt.Sprintf("%v^%v", n, c))
	}
	s := sarr[0]
	for i := 1; i < len(sarr); i++ {
		s = fmt.Sprintf("%s * %s", s, sarr[i])
	}
	return s
}
