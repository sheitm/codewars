package main

import (
	"fmt"
	"sort"
)

// https://www.codewars.com/kata/54d496788776e49e6b00052f/train/go

func SumOfDivided(lst []int) string {
	m := map[int][]int{}
	keys := []int{}
	for _, i := range lst {
		factors := getDistinctPrimeFactors(i)
		for _, f := range factors {
			if l, ok := m[f]; ok {
				m[f] = append(l, i)
			} else {
				m[f] = []int{i}
				keys = append(keys, f)
			}
		}
	}
	sort.Ints(keys)
	res := ""
	for _, key := range keys {
		res = res + fmt.Sprintf("(%d %d)", key, sumSlice(m[key]))
	}
	return res
}

func sumSlice(s []int) int {
	sum := 0
	for _, i := range s {
		sum += i
	}
	return sum
}

func getDistinctPrimeFactors(n int) []int {
	m := map[int]bool{}
	distinct := []int{}
	for n%2 == 0 {
		if _, ok := m[2]; !ok {
			distinct = append(distinct, 2)
			m[2] = true
		}

		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			if _, ok := m[i]; !ok {
				distinct = append(distinct, i)
				m[i] = true
			}

			n = n / i
		}
	}

	if n > 2 {
		distinct = append(distinct, n)
	}

	return distinct
}
