package main

// https://www.codewars.com/kata/52774a314c2333f0a7000688/solutions/go

var upDown = map[string]func(int)int {
	"(": func(ii int) int {return ii + 1},
	")": func(ii int) int {return ii - 1},
}

func ValidParentheses(parens string) bool {
	c := 0
	for i := 0; i < len(parens); i++  {
		c = upDown[string(parens[i])](c)
		if c < 0 {
			return false
		}
	}
	return c == 0
}
