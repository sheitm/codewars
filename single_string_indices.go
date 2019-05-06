package main

import "errors"

// Solution solve kata
func Solution(str string, idx uint) (uint, error) {
	r := string(str[idx])
	if r != "(" {
		return 0, errors.New("Not an opening bracket")
	}
	depth := 0
	var res uint
	for i := int(idx) + 1; i < len(str); i++ {
		s := string(str[i])
		if s == ")" {
			if depth == 0 {
				res = uint(i)
				break
			}
			depth--
		}
		if s == "(" {
			depth++
		}
	}
	return res, nil
}
