package main

// https://www.codewars.com/kata/pick-peaks/train/go

// PosPeaks contains peak info
type PosPeaks struct {
	Pos   []int
	Peaks []int
}

// PickPeaks finds all peaks
func PickPeaks(array []int) PosPeaks {
	if len(array) == 0 {
		return PosPeaks{
			Pos:   []int{},
			Peaks: []int{},
		}
	}
	slopes := toSlopeArray(array)
	pos := []int{}
	peaks := []int{}
	for i := 0; i < len(array)-1; i++ {
		if slopes[i] != ">" {
			continue
		}
		if slopes[i+1] == "<" {
			pos = append(pos, i)
			peaks = append(peaks, array[i])
			continue
		}
		if slopes[i+1] == "=" && isPlateu(i, slopes) {
			pos = append(pos, i)
			peaks = append(peaks, array[i])
		}
	}
	return PosPeaks{
		Pos:   pos,
		Peaks: peaks,
	}
}

func isPlateu(index int, slopes []string) bool {
	for i := index; i < len(slopes)-1; i++ {
		if slopes[i+1] == "<" {
			return true
		}

		if slopes[i+1] == ">" {
			return false
		}
	}
	return false
}

func toSlopeArray(array []int) []string {
	res := []string{}
	for i := 0; i < len(array); i++ {
		if i == 0 {
			res = append(res, "0")
			continue
		}
		if array[i] > array[i-1] {
			res = append(res, ">")
			continue
		}
		if array[i] < array[i-1] {
			res = append(res, "<")
			continue
		}
		res = append(res, "=")
	}
	return res
}
