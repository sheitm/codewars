package main

import (
	"sort"
)

// https://www.codewars.com/kata/sudoku-solution-validator/train/go

// ValidateSolution validate the Sudoko board
func ValidateSolution(m [][]int) bool {
	if len(m) != 9 {
		return false
	}
	for i := 0; i < len(m); i++ {
		v := getHVector(i, m)
		if !vectorIsValid(v) {
			return false
		}
		v = getVVector(i, m)
		if !vectorIsValid(v) {
			return false
		}
		v = getSVector(i, m)
		if !vectorIsValid(v) {
			return false
		}
	}
	return true
}

func getHVector(index int, m [][]int) []int {
	v := []int{}
	for i := 0; i < len(m[index]); i++ {
		v = append(v, m[index][i])
	}
	return v
}

func getVVector(index int, m [][]int) []int {
	v := []int{}
	for i := 0; i < len(m[index]); i++ {
		v = append(v, m[i][index])
	}
	return v
}

func getSVector(index int, m [][]int) []int {
	v := []int{}
	for i := 0; i < len(m[index]); i++ {
		hi, vi := getSVectorIndices(index, i)
		v = append(v, m[hi][vi])
	}
	return v
}

func getSVectorIndices(sectorIndex, index int) (int, int) {
	mv := (sectorIndex / 3) * 3
	v := mv + (index / 3)

	mh := (sectorIndex % 3) * 3
	h := mh + (index % 3)

	return h, v
}

func vectorIsValid(v []int) bool {
	if len(v) != 9 {
		return false
	}
	sort.Ints(v)
	for i := 0; i < len(v); i++ {
		if v[i] != i+1 {
			return false
		}
	}
	return true
}
