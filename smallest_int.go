package main

import (
	"fmt"
	"strconv"
)

// https://www.codewars.com/kata/find-the-smallest/train/go

// Smallest do it
func Smallest(n int64) []int64 {
	arr := getArr(n)
	candidates := map[int64][]permutation{}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i != j {
				p := permutation{sourceIndex: i}
				p.make(arr, j)
				if _, ok := candidates[p.num]; !ok {
					candidates[p.num] = []permutation{}
				}
				candidates[p.num] = append(candidates[p.num], p)
			}
		}
	}
	smallest := int64(9223372036854775807)
	var perms []permutation
	for c, p := range candidates {
		if c < smallest {
			smallest = c
			perms = p
		}
	}
	return perms[0].getVector(n)
}

type permutation struct {
	num         int64
	sourceIndex int
	targetIndex int
}

func (p *permutation) make(arr []int, target int) {
	p.targetIndex = target
	if p.sourceIndex > p.targetIndex {
		p.makeLeft(arr)
		return
	}
	p.makeRight(arr)
}

func (p *permutation) makeRight(arr []int) {
	s := ""
	for i := 0; i < len(arr); i++ {
		if i == p.sourceIndex {
			continue
		}
		s = fmt.Sprintf("%s%d", s, arr[i])
		if i == p.targetIndex {
			s = fmt.Sprintf("%s%d", s, arr[p.sourceIndex])
		}
	}
	num, _ := strconv.Atoi(s)
	p.num = int64(num)
}

func (p *permutation) makeLeft(arr []int) {
	s := ""
	for i := 0; i < len(arr); i++ {
		if i == p.sourceIndex {
			continue
		}
		if i == p.targetIndex {
			s = fmt.Sprintf("%s%d%d", s, arr[p.sourceIndex], arr[i])
			continue
		}
		s = fmt.Sprintf("%s%d", s, arr[i])
	}
	num, _ := strconv.Atoi(s)
	p.num = int64(num)
}

func (p permutation) getVector(original int64) []int64 {
	if p.num == original {
		return []int64{p.num, int64(0), int64(0)}
	}
	return []int64{p.num, int64(p.sourceIndex), int64(p.targetIndex)}
}

func getArr(n int64) []int {
	s := fmt.Sprintf("%d", n)
	res := []int{}
	for i := 0; i < len(s); i++ {
		b := string(s[i])
		n, _ := strconv.Atoi(b)
		res = append(res, n)
	}
	return res
}

// // Smallest2 x
// func Smallest2(n int64) []int64 {
// 	arr := getArr(n)
// 	smallestIndex := getIndexOfSmallest(arr, 10)
// 	v := getNewNumber(arr, int(smallestIndex), 0)

// 	return []int64{v, smallestIndex, 0}
// }

// func getNewNumber(arr []int64, fromIndex, toIndex int) int64 {
// 	s := ""
// 	for i := 0; i < len(arr); i++ {
// 		if i == fromIndex {
// 			continue
// 		}
// 		if i == toIndex {
// 			s = fmt.Sprintf("%s%d%d", s, arr[fromIndex], arr[i])
// 			continue
// 		}
// 		s = fmt.Sprintf("%s%d", s, arr[i])
// 	}
// 	r, _ := strconv.Atoi(s)
// 	return int64(r)
// }

// func getIndexOfSmallest(arr []int64, currentSmallest int64) int64 {
// 	smallest := currentSmallest
// 	smallestIndex := -1
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] < smallest {
// 			smallest = arr[i]
// 			smallestIndex = i
// 		}
// 	}
// 	return int64(smallestIndex)
// }

// func getArr(n int64) []int64 {
// 	s := fmt.Sprintf("%d", n)
// 	res := []int64{}
// 	for i := 0; i < len(s); i++ {
// 		b := string(s[i])
// 		n, _ := strconv.Atoi(b)
// 		res = append(res, int64(n))
// 	}
// 	return res
// }
