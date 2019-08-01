package main

// https://www.codewars.com/kata/build-a-pile-of-cubes/train/go

// FindNb finds the number
func FindNb(m int) int {
	c := 2
	for {
		ok, attempts := tryCubeRoot(c, m)
		if ok {
			return c
		}
		if attempts <= 2 {
			return -1
		}
		c++
	}
}

func tryCubeRoot(c, goal int) (bool, int) {
	s := 0
	for i := 0; i < c; i++ {
		cc := c - i
		s = s + (cc * cc * cc)
		if s == goal {
			return true, i + 1
		}
		if s > goal {
			return false, i + 1
		}
	}
	return false, 100
}
