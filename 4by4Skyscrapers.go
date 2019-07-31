package main

// https://www.codewars.com/kata/5671d975d81d6c1c87000022/train/go

// SolvePuzzle solves the puzzlr
func SolvePuzzle(clues []int) [][]int {
	b := board{}
	b.initialize(4)

	return b.getIntMatrix()
}

type tCell struct {
	val int
}

type board struct {
	cells [][]*tCell
}

func (b *board) initialize(dim int) {
	cls := make([][]*tCell, dim)
	for y := 0; y < dim; y++ {
		cls[y] = make([]*tCell, dim)
	}

	b.cells = cls
}

func (b board) getIntMatrix() [][]int {
	cells := [][]int{}
	for y := 0; y < 4; y++ {
		row := []int{}
		for x := 0; x < 4; x++ {
			row = append(row, b.cells[y][x].val)
		}
		cells = append(cells, row)
	}
	return cells
}

func getVectorForIndex(cells [][]*tCell, index int) []*tCell {
	if index <= 3 {
		v := []*tCell{}
		for i := 0; i < 4; i++ {
			v = append(v, cells[i][index])
		}
		return v
	}
	if index <= 7 {
		v := []*tCell{}
		for i := 0; i < 4; i++ {
			v = append(v, cells[index-4][3-i])
		}
		return v
	}
	if index <= 11 {
		v := []*tCell{}
		for i := 0; i < 4; i++ {
			v = append(v, cells[3-i][11-index])
		}
		return v
	}

	return cells[15-index]
}

func cellsToMatrix(cells [][]*tCell) [][]int {
	matrix := [][]int{}
	for y := 0; y < len(cells); y++ {
		row := []int{}
		for x := 0; x < len(cells); x++ {
			row = append(row, cells[y][x].val)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func matrixToCells(matrix [][]int) [][]*tCell {
	cells := [][]*tCell{}
	for y := 0; y < len(matrix); y++ {
		row := []*tCell{}
		for x := 0; x < len(matrix); x++ {
			c := tCell{val: matrix[y][x]}
			row = append(row, &c)
		}
		cells = append(cells, row)
	}
	return cells
}

func convertCellVector(cells []*tCell) []int {
	v := []int{}
	for i := 0; i < len(cells); i++ {
		v = append(v, cells[i].val)
	}
	return v
}

func convertIntVector(ints []int) []*tCell {
	v := []*tCell{}
	for i := 0; i < len(ints); i++ {
		c := tCell{val: ints[i]}
		v = append(v, &c)
	}
	return v
}
