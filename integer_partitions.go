package main

// https://www.codewars.com/kata/getting-along-with-integer-partitions/train/go

import (
	"fmt"
	"sort"
)

var mem map[int][][]int

// Part solve kata
func Part(n int) string {
	if n < 1 || n > 50 {
		return ""
	}
	mem = map[int][][]int{}
	enum := getEnum(n)
	product := getProduct(enum)
	r := getRange(product)
	avg := average(product)
	median := median(product)
	return fmt.Sprintf("Range: %d Average: %.2f Median: %.2f", r, avg, median)
}

func getEnum(n int) [][]int {
	if e, ok := mem[n]; ok {
		return e
	}
	if n == 1 {
		e := [][]int{{1}}
		mem[n] = e
		return e
	}
	e := getEnum(n - 1)
	ee := [][]int{{n}}
	for _, a := range e {
		a = append(a, 1)
		ee = append(ee, a)
	}
	mem[n] = ee
	return ee
}

func getProduct(enum [][]int) []int {
	m := map[int]int{}
	res := []int{}
	for _, arr := range enum {
		p := 1
		for i := range arr {
			p = p * arr[i]
		}
		if _, ok := m[p]; !ok {
			m[p] = p
			res = append(res, p)
		}
	}
	return res
}

func median(arr []int) float64 {
	if len(arr) == 1 {
		return float64(arr[0])
	}
	sort.Ints(arr)
	if len(arr)%2 != 0 {
		return float64(arr[int(len(arr)/2)])
	}
	m := len(arr) / 2
	return (float64(arr[m-1]) + float64(arr[m])) / 2
}

func average(arr []int) float64 {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	return float64(sum) / float64(len(arr))
}

func getRange(arr []int) int {

	min := 999999999999999999
	max := -1
	for i := 0; i < len(arr); i++ {
		n := arr[i]
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return max - min
}
