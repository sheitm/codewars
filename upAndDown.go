package main

import "strings"

// https://www.codewars.com/kata/56cac350145912e68b0006f0/train/go

// Arrange do it!
func Arrange(s string) string {
	if s == "" {
		return ""
	}
	arr := strings.Split(s, " ")
	hasSwaps := false
	for {
		arr, hasSwaps = passOver(arr)
		if !hasSwaps {
			return toUpAndDownString(arr)
		}
	}
}

func passOver(arr []string) ([]string, bool) {
	swapped := map[int]int{}
	for i := 0; i < len(arr)-1; i++ {
		if i%2 == 0 && len(arr[i]) > len(arr[i+1]) {
			swapped[i] = i + 1
			swapped[i+1] = i
			break
		}
		if i%2 != 0 && len(arr[i]) < len(arr[i+1]) {
			swapped[i] = i + 1
			swapped[i+1] = i
			break
		}
	}
	if len(swapped) == 0 {
		return arr, false
	}
	res := []string{}
	for i := 0; i < len(arr); i++ {
		if sw, ok := swapped[i]; ok {
			res = append(res, arr[sw])
			continue
		}
		res = append(res, arr[i])
	}
	return res, true
}

func toUpAndDownString(arr []string) string {
	s := ""
	started := false
	for i := 0; i < len(arr); i++ {
		if started {
			s += " "
		}
		started = true
		if i%2 == 0 {
			s += strings.ToLower(arr[i])
		} else {
			s += strings.ToUpper(arr[i])
		}
	}
	return s
}
