package main

// https://www.codewars.com/kata/51b6249c4612257ac0000005/train/go

var valueMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// RomanNumeralsDecode solve it!
func RomanNumeralsDecode(roman string) int {
	v := 0
	for i := 0; i < len(roman); i++ {
		cv := valueMap[string(roman[i])]
		if i < len(roman)-1 {
			ccv := valueMap[string(roman[i+1])]
			if cv < ccv {
				v += ccv - cv
				i++
				continue
			}
		}
		v += cv
	}
	return v
}
