package main

// https://www.codewars.com/kata/getting-along-with-integer-partitions/train/go

import (
	"fmt"
	"sort"
)

type partition struct {
	key     string
	sum     int
	diagram []int
}

func (p *partition) make(n int) {
	arr := []int{}
	for i := 0; i < n; i++ {
		arr = append(arr, 1)
	}
	p.set(arr)
}

func (p *partition) set(d []int) {
	sort.Ints(d)
	p.diagram = d
	k := ""
	s := 0
	for _, v := range d {
		k = fmt.Sprintf("%s%d", k, v)
		s += v
	}
	p.key = k
	p.sum = s
}

func makePartitions(n int) []partition {
	partitions := map[string]partition{}
	arr := []int{}
	for i := 0; i < n; i++ {
		arr = append(arr, 1)
	}

	current := partition{}
	current.set(arr)
	partitions[current.key] = current

	// for {
	// 	for i := 0; i < len(current.diagram)-1; i++ {
	// 		if current.diagram[i] > 0 {

	// 		}
	// 	}
	// 	// p := partition{}
	// 	// p.set(current)
	// 	// partitions = append(partitions, p)
	// 	// next := []int{}
	// 	// changes := false
	// 	// for i := 0; i < (len(current) - 1); i++ {
	// 	// 	if changes {
	// 	// 		next = append(next, current[i])
	// 	// 		continue
	// 	// 	}
	// 	// 	if current[i] > current[i+1] {
	// 	// 		next = append(next, current[i]-1)
	// 	// 		next = append(next, current[i+1]+1)
	// 	// 		i++
	// 	// 		changes = true
	// 	// 	}
	// 	// }
	// 	// if !changes {
	// 	// 	break
	// 	// }
	// 	// current = next
	// }

	// res := []partition{}
	// for _, p := range partitions {
	// 	res = append(res, p)
	// }
	// return res

	return []partition{}
}

func newArr(arr []int, decr, incr int) partition {
	res := []int{}
	for i, v := range arr {
		if i == decr {
			res = append(res, v-1)
			continue
		}
		if i == incr {
			res = append(res, v+1)
			continue
		}
		res[i] = arr[i]
	}
	p := partition{}
	p.set(res)
	return p
}

// Part solve kata
func Part(n int) string {
	return ""
	// if n < 1 || n > 50 {
	// 	return ""
	// }

	// enum := getEnum(n)
	// product := getProduct(enum)
	// r := getRange(product)
	// avg := average(product)
	// median := median(product)
	// return fmt.Sprintf("Range: %d Average: %.2f Median: %.2f", r, avg, median)
}

// func getEnum(n int) [][]int {
// 	if e, ok := mem[n]; ok {
// 		return e
// 	}
// 	if n == 1 {
// 		e := [][]int{{1}}
// 		mem[n] = e
// 		return e
// 	}
// 	e := getEnum(n - 1)
// 	ee := [][]int{{n}}
// 	for _, a := range e {
// 		a = append(a, 1)
// 		ee = append(ee, a)
// 	}
// 	mem[n] = ee
// 	return ee
// }

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
