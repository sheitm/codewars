package main

import "fmt"

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
	return ""
}
