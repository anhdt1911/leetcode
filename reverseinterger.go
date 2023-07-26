package main

import "math"

func reverse(x int) int {
	revInt := 0

	for x != 0 {
		revInt = revInt*10 + x%10
		x /= 10
	}

	if revInt < math.MinInt32 || revInt > math.MaxInt32 {
		return 0
	}

	return revInt
}
