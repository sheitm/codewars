package main

// Interpreter do it
func Interpreter(code, tape string) string {
	cells := []cell{}
	for i := 0; i < len(tape); i++ {
		c := cell{index: i, val: int(tape[i]) == 49}
		cells = append(cells, c)
	}

	pos := 0
	anchorIndexes := getAnchorIndexes(code)
	fastforward := false
	bail := false
	backcount := 0
	for i := 0; i < len(code); i++ {
		if bail || backcount >= 300 {
			break
		}

		current := toString(cells)
		_ = current

		command := string(code[i])
		switch command {
		case "*":
			if !fastforward {
				cells[pos].flip()
			}
		case ">":
			if !fastforward {
				pos++
				if pos >= len(tape) {
					bail = true
					break
				}
			}
		case "<":
			if !fastforward {
				pos--
				if pos < 0 {
					bail = true
					break
				}
			}
		case "[":
			if !fastforward {
				if !cells[pos].val {
					fastforward = true
				}
			}
		case "]":
			backcount++
			if fastforward {
				fastforward = false
				break
			}

			if cells[pos].val {
				i = getMatchingAnchorIndex(i, anchorIndexes) - 1
			}
		}
	}
	return toString(cells)
}

type cell struct {
	val   bool
	index int
}

func (c *cell) flip() {
	c.val = !c.val
}

func getMatchingAnchorIndex(index int, anchors []int) int {
	anchor := -1
	for i := 0; i < len(anchors); i++ {
		a := anchors[i]
		if a > index {
			return anchor
		}
		anchor = a
	}
	return anchor
}

func getAnchorIndexes(code string) []int {
	anchors := []int{}
	for i := 0; i < len(code); i++ {
		if string(code[i]) == "[" {
			anchors = append(anchors, i)
		}
	}
	return anchors
}

func toString(cells []cell) string {
	res := ""
	for i := 0; i < len(cells); i++ {
		if cells[i].val {
			res += "1"
		} else {
			res += "0"
		}
	}
	return res
}
