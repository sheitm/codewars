package main

import "fmt"

// https://www.codewars.com/kata/5671d975d81d6c1c87000022/train/go

// SolvePuzzle solves the puzzlr
func SolvePuzzle(clues []int) [][]int {
	p := permutations{}
	p.init()

	mh0 := getMustHaves(0, clues)
	mh1 := getMustHaves(1, clues)
	mh2 := getMustHaves(2, clues)
	mh3 := getMustHaves(3, clues)

	for _, p0 := range p.getValidPermutations(mh0) {
		for _, p1 := range p.getValidPermutations(mh1) {
			for _, p2 := range p.getValidPermutations(mh2) {
				for _, p3 := range p.getValidPermutations(mh3) {
					m := [][]int{p0, p1, p2, p3}
					if mIsValidWithClues(m, clues) {
						return m
					}
				}
			}
		}
	}

	return [][]int{}
}

type permutations struct {
	perms [][]int
}

func (p *permutations) init() {
	p.perms = [][]int{}
	p.generateAllPermutations([]int{1, 2, 3, 4}, 4)
}

func (p *permutations) generateAllPermutations(a []int, size int) {
	if size == 1 {
		ac := []int{}
		for _, av := range a {
			ac = append(ac, av)
		}
		p.perms = append(p.perms, ac)
	}

	for i := 0; i < size; i++ {
		p.generateAllPermutations(a, size-1)

		if size%2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
}

func (p permutations) getValidPermutations(mh map[int]int) [][]int {
	res := [][]int{}
	for _, prm := range p.perms {
		ok := true
		for k, v := range mh {
			if prm[k] != v {
				ok = false
			}
		}
		if ok {
			res = append(res, prm)
		}
	}
	return res
}

func getMustHaves(rowIndex int, clues []int) map[int]int {
	res := map[int]int{}
	if rowIndex == 0 {
		if clues[0] == 1 {
			res[0] = 4
		}
		if clues[1] == 1 {
			res[1] = 4
		}
		if clues[2] == 1 {
			res[2] = 4
		}
		if clues[3] == 1 {
			res[3] = 4
		}
		if clues[4] == 1 {
			res[3] = 4
		}
		if clues[15] == 1 {
			res[0] = 4
		}
		return res
	}
	if rowIndex == 1 {
		if clues[5] == 1 {
			res[3] = 4
		}
		if clues[14] == 1 {
			res[0] = 4
		}
		return res
	}
	if rowIndex == 2 {
		if clues[6] == 1 {
			res[3] = 4
		}
		if clues[13] == 1 {
			res[0] = 4
		}
		return res
	}

	if clues[7] == 1 {
		res[3] = 4
	}
	if clues[8] == 1 {
		res[3] = 4
	}
	if clues[9] == 1 {
		res[2] = 4
	}
	if clues[10] == 1 {
		res[1] = 4
	}
	if clues[11] == 1 {
		res[0] = 4
	}
	if clues[12] == 1 {
		res[0] = 4
	}
	return res
}

func vIsValid(r []int) bool {
	if len(r) != 4 {
		return false
	}

	am := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	}
	for _, v := range r {
		if _, ok := am[v]; !ok {
			return false
		}
		delete(am, v)
	}
	return len(am) == 0
}

func mIsValidWithClues(m [][]int, clues []int) bool {
	if !mIsValid(m) {
		return false
	}
	for i, clue := range clues {
		if clue == 0 {
			continue
		}
		v := getVectorForIndex(m, i)
		if countVisibleTowers(v) != clue {
			return false
		}
	}
	return true
}

func mIsValid(m [][]int) bool {
	pos := map[string]string{}
	for i := 0; i < len(m); i++ {
		v := m[i]
		if !vIsValid(v) {
			return false
		}
		for j := 0; j < len(v); j++ {
			k := fmt.Sprintf("%d-%d", j, v[j])
			if _, exists := pos[k]; exists {
				return false
			}
			pos[k] = k
		}
	}
	return true
}

func countVisibleTowers(v []int) int {
	c := 0
	tallest := 0
	for _, t := range v {
		if t > tallest {
			c++
			tallest = t
		}
	}
	return c
}

func getVectorForIndex(cells [][]int, index int) []int {
	if index <= 3 {
		v := []int{}
		for i := 0; i < 4; i++ {
			v = append(v, cells[i][index])
		}
		return v
	}
	if index <= 7 {
		v := []int{}
		for i := 0; i < 4; i++ {
			v = append(v, cells[index-4][3-i])
		}
		return v
	}
	if index <= 11 {
		v := []int{}
		for i := 0; i < 4; i++ {
			v = append(v, cells[3-i][11-index])
		}
		return v
	}

	return cells[15-index]
}
