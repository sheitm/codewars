package main

// BouncingBall solve the kata
func BouncingBall(h, bounce, window float64) int {
	if bounce <= 0 || bounce >= 1 || window >= h {
		return -1
	}
	nh := bounce * h
	if nh < window {
		return 1
	}
	return BouncingBall(nh, bounce, window) + 2
}
