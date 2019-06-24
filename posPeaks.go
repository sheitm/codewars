package main

// https://www.codewars.com/kata/pick-peaks/train/go

// PosPeaks contains peak nfo
type PosPeaks struct {
	Pos   []int
	Peaks []int
}

// PickPeaks finds all peaks
func PickPeaks(array []int) PosPeaks {
	if len(array) == 0 {
		return PosPeaks{}
	}

	pos := []int{}
	peaks := []int{}
	index := 0
	for {
		po, pe, nextIndex := searchPeak(index, array)
		if po > 0 {
			pos = append(pos, po)
			peaks = append(peaks, pe)
		}
		if nextIndex < 0 {
			break
		}
		index = nextIndex
	}

	return PosPeaks{
		Pos:   pos,
		Peaks: peaks,
	}
}

func searchPeak(start int, array []int) (int, int, int) {
	index := start
	pos := 0
	peak := 0
	for {
		if index == len(array)-1 {
			return 0, 0, -1
		}
		if array[index+1] <= array[index] {
			ffd := fastForwardDownSlope(index+1, array)
			if ffd < 0 {
				return 0, 0, -1
			}
			return pos, peak, ffd
		}
		pos = index + 1
		peak = array[pos]
		index++
	}
}

func fastForwardDownSlope(index int, array []int) int {
	for {
		if index == len(array)-1 {
			return -1
		}
		if array[index+1] > array[index] {
			return index
		}
		index++
	}
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
