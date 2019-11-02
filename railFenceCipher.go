package main

import (
	"fmt"
	"unicode/utf8"
)

// https://www.codewars.com/kata/58c5577d61aefcf3ff000081/train/go

// Encode encode it
func Encode(s string, n int) string {
	if s == "" {
		return ""
	}

	rails := map[int]string{}
	for i := 0; i < n; i++ {
		rails[i] = ""
	}

	railIndex := 0
	down := true
	for i := 0; i < len(s); i++ {
		rails[railIndex] = fmt.Sprintf("%s%s", rails[railIndex], string(s[i]))
		if down {
			railIndex++
			if railIndex == n {
				railIndex = n - 2
				down = false
			}
		} else {
			railIndex--
			if railIndex < 0 {
				railIndex = 1
				down = true
			}
		}
	}

	out := ""
	for i := 0; i < len(rails); i++ {
		out = fmt.Sprintf("%s%s", out, rails[i])
	}
	return out
}

// Decode decode it
func Decode(s string, n int) string {
	cycle := 2 + (n-2)*2
	r := len(s) / cycle
	lengths := map[int]int{}
	lengths[0] = r
	lengths[n-1] = r
	for i := 1; i < n-1; i++ {
		lengths[i] = 2 * r
	}

	m := len(s) % cycle
	railIndex := 0
	down := true
	for i := 0; i < m; i++ {
		lengths[railIndex] = lengths[railIndex] + 1
		if down {
			railIndex++
			if railIndex == n {
				railIndex = n - 2
				down = false
			}
		} else {
			railIndex--
			if railIndex < 0 {
				railIndex = 1
				down = true
			}
		}
	}

	rails := map[int]string{}
	accum := ""
	currentRail := 0
	for i := 0; i < len(s); i++ {
		accum = fmt.Sprintf("%s%s", accum, string(s[i]))
		if len(accum) == lengths[currentRail] {
			rails[currentRail] = accum
			currentRail++
			accum = ""
		}
	}

	railIndex = 0
	accum = ""
	down = true
	for i := 0; i < len(s); i++ {
		c := rails[railIndex][0]
		rails[railIndex] = trimFirstRune(rails[railIndex])
		accum = fmt.Sprintf("%s%s", accum, string(c))
		if down {
			railIndex++
			if railIndex == n {
				railIndex = n - 2
				down = false
			}
		} else {
			railIndex--
			if railIndex < 0 {
				railIndex = 1
				down = true
			}
		}
	}

	return accum
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
