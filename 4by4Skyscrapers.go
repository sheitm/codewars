package main

// https://www.codewars.com/kata/5671d975d81d6c1c87000022/train/go

// SolvePuzzle solves the puzzlr
func SolvePuzzle(clues []int) [][]int {
	b := board{}
	b.initialize(4)
	return b.cells
}

type board struct {
	cells [][]int
}

func (b *board) initialize(dim int) {
	cls := make([][]int, dim)
	for y := 0; y < dim; y++ {
		cls[y] = make([]int, dim)
	}
	b.cells = cls
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
