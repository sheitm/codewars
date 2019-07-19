package main

// https://www.codewars.com/kata/54521e9ec8e60bc4de000d6c/train/go

// MaximumSubarraySum finds the maximun sum
func MaximumSubarraySum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	largest := 0
	for i := 0; i < len(numbers); i++ {
		candidate := getLargestSumFromIndex(i, numbers)
		if candidate > largest {
			largest = candidate
		}
	}
	return largest
}

func getLargestSumFromIndex(index int, numbers []int) int {
	if numbers[index] < 0 {
		return 0
	}
	s := 0
	largest := 0
	for i := index; i < len(numbers); i++ {
		s += numbers[i]
		if s > largest {
			largest = s
		}
	}
	return largest
}
