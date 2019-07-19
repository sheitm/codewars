package main

import "fmt"

// https://www.codewars.com/kata/5547cc7dcad755e480000004/train/go

// RemovNb2 do it!
func RemovNb2(m uint64) [][2]uint64 {
	max := (m * (m - 1)) / 2
	matches := [][2]uint64{}
	for a := uint64(1); a <= m; a++ {
		for b := uint64(1); b <= m; b++ {
			if a == 15 && b == 21 {
				x := 0
				_ = x
			}
			if a*b == max-(a+b) {
				matches = append(matches, [2]uint64{a, b})
			}
		}
	}
	if len(matches) == 0 || len(matches) > 2 {
		return nil
	}
	return matches
}

// RemovNb returns the values
func RemovNb(m uint64) [][2]uint64 {
	sums := map[string]int{}
	for a := 1; a <= int(m); a++ {
		for b := 1; b <= int(m); b++ {
			if a != b {
				k := getSumKey(a, b)
				if _, ok := sums[k]; !ok {
					sums[k] = sumExpect(int(m), a, b)
				}
			}
		}
	}

	matches := [][2]uint64{}
	for a := 1; a <= int(m); a++ {
		for b := 1; b <= int(m); b++ {
			if a != b {
				k := getSumKey(a, b)
				if a*b == sums[k] {
					matches = append(matches, [2]uint64{uint64(a), uint64(b)})
				}
			}
		}
	}
	if len(matches) == 0 || len(matches) > 2 {
		return nil
	}
	return matches
}

func getSumKey(a, b int) string {
	if a < b {
		return fmt.Sprintf("%d_%d", a, b)
	}
	return fmt.Sprintf("%d_%d", b, a)
}

func sumExpect(lim, a, b int) int {
	s := 0
	for i := 1; i <= lim; i++ {
		if i != a && i != b {
			s += i
		}
	}
	return s
}
