package main

import "fmt"

// https://www.codewars.com/kata/58c5577d61aefcf3ff000081/train/go

// Encode encode it
func Encode(s string, n int) string {
	if s == "" {
		return ""
	}

	rails := map[int]string{
		0: "",
		1: "",
		2: "",
	}

	mapIndex := 0
	down := true
	for i := 0; i < len(s); i++ {
		rails[mapIndex] = fmt.Sprintf("%s%s", rails[mapIndex], string(s[i]))
		if down {
			mapIndex++
			if mapIndex == n {
				mapIndex = n - 2
				down = false
			}
		} else {
			mapIndex--
			if mapIndex < 0 {
				mapIndex = 1
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
