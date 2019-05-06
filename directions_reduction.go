package main

// https://www.codewars.com/kata/directions-reduction/train/go

// DirReduc asdsda
func DirReduc(arr []string) []string {
	if len(arr) < 2 {
		return arr
	}
	p := point{x: 0, y: 0}
	for i := 0; i < len(arr); i++ {
		p.move(arr[i])
	}
	return p.getDirections()
}

type point struct {
	x int
	y int
}

func (p *point) move(direction string) {
	switch direction {
	case "NORTH":
		p.y = p.y + 1
	case "SOUTH":
		p.y = p.y - 1
	case "EAST":
		p.x = p.x + 1
	case "WEST":
		p.x = p.x - 1
	}
}

func (p point) getDirections() []string {
	dirs := []string{}
	if p.y > 0 {
		dirs = appendDirs("NORTH", p.y, dirs)
	} else {
		dirs = appendDirs("SOUTH", p.y*-1, dirs)
	}
	if p.x > 0 {
		dirs = appendDirs("EAST", p.x, dirs)
	} else {
		dirs = appendDirs("WEST", p.x*-1, dirs)
	}
	return dirs
}

func appendDirs(dir string, c int, arr []string) []string {
	if c == 0 {
		return arr
	}
	for i := 0; i < c; i++ {
		arr = append(arr, dir)
	}
	return arr
}

//////////////////////////////////
///
/// Second solution
///
//////////////////////////////////

var opposites = map[string]string{
	"NORTH": "SOUTH",
	"SOUTH": "NORTH",
	"EAST":  "WEST",
	"WEST":  "EAST",
}

var dirIndexes = map[string]int{
	"NORTH": 1,
	"SOUTH": 2,
	"WEST":  3,
	"EAST":  4,
}

// DirReduc2 asdsda
func DirReduc2(arr []string) []string {
	// sort.Slice(arr, func(i, j int) bool { return dirIndexes[arr[i]] < dirIndexes[arr[j]] })
	return reduceRecursive(arr)
}

func reduceRecursive(arr []string) []string {
	if len(arr) < 2 {
		return arr
	}
	next := reduceOnce(arr)
	if len(next) == len(arr) {
		return next
	}
	return reduceRecursive(next)
}

func reduceOnce(arr []string) []string {
	res := []string{}
	lim := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		if i == lim {
			res = append(res, arr[i])
			continue
		}
		if opposites[arr[i]] == arr[i+1] {
			i = i + 2
			continue
		}
		res = append(res, arr[i])
	}
	return res
}
