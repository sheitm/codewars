package main

// https://www.codewars.com/kata/5541f58a944b85ce6d00006a/train/go

// ProductFib returns the product fib
func ProductFib(prod uint64) [3]uint64 {
	lim := 20000
	seq := fibSequence(lim)
	for i := 0; i < lim; i++ {
		p := seq[i] * seq[i+1]
		if uint64(p) == prod {
			return [3]uint64{uint64(seq[i]), uint64(seq[i+1]), 1}
		}
		if uint64(p) > prod {
			return [3]uint64{uint64(seq[i]), uint64(seq[i+1]), 0}
		}
	}

	return [3]uint64{1, 1, 1}
}

func fibSequence(upperBound int) []int {
	if upperBound == 0 {
		return []int{0}
	}

	res := []int{0, 1}
	if upperBound == 1 {
		return res
	}

	for i := 1; i < upperBound; i++ {
		l := len(res)
		nn := res[l-2] + res[l-1]
		res = append(res, nn)
	}

	return res
}
