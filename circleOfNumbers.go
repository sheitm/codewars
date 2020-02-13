package main

func CircleOfNumbers(n int, firstNumber int) int {
	jumps := n / 2
	result := firstNumber + jumps
	if result < n {
		return result
	}
	return result - n
}