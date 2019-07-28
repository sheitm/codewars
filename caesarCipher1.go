package main

import "fmt"

const alphabet = "abcdefghijklmnopqrstuvwxz"

// big: 65 - 90
// small: 97 - 122

// MovingShift encrypt
func MovingShift(s string, shift int) []string {
	accum := ""
	for _, l := range s {
		ll := encrypt(int(l), shift)
		accum = fmt.Sprintf("%s%s", accum, string(ll))
		shift++
		if shift > 25 {
			shift = 0
		}
	}
	return toStringArr(accum)
}

// DemovingShift decrypt
func DemovingShift(arr []string, shift int) string {
	s := ""
	for _, a := range arr
	return ""
}

func encrypt(letter, shift int) int {
	if !isLetter(letter) {
		return letter
	}
	r := letter + shift
	if isLowercase(letter) {
		if r <= 122 {
			return r
		}
		return 97 + ((r - 122) - 1)
	}
	if r <= 90 {
		return r
	}
	return 65 + ((r - 90) - 1)
}

func isLowercase(l int) bool {
	return l >= 97 && l <= 122
}

func isLetter(l int) bool {
	if l >= 65 && l <= 90 {
		return true
	}
	return l >= 97 && l <= 122
}

func toStringArr(s string) []string {
	lengths := getStringArrLengths(len(s))
	arr := []string{}
	a := 0
	for _, i := range lengths {
		arr = append(arr, s[a:a+i])
		a = a + i
	}
	return arr
}

func getStringArrLengths(l int) []int {
	d := l / 5
	m := l % 5

	if m == 0 {
		return []int{d, d, d, d, d}
	}

	d++
	s := 0
	arr := []int{}
	for i := 0; i < 5; i++ {
		if s == l {
			arr = append(arr, 0)
			continue
		}
		if (s + d) > l {
			arr = append(arr, l-s)
			s = l
			continue
		}
		arr = append(arr, d)
		s += d
	}
	return arr
}

func toStringArrCore(s string, substringCount int) []string {
	d := len(s) / substringCount
	res := []string{}
	for i := 0; i < substringCount; i++ {
		ii := i * d
		res = append(res, s[ii:ii+d])
	}
	return res
}

func getPerfectMod(s string, substringCount int) (int, int) {
	l := len(s)
	sc := substringCount
	for i := 0; i < len(s); i++ {
		if l%sc == 0 {
			return l, sc
		}
		l--
		sc--
	}
	return 0, 0
}
