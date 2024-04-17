package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
}

func minEatingSpeed(piles []int, h int) int {
	max := 0
	for _, v := range piles {
		if v > max {
			max = v
		}
	}

	low, high := 1, max
	for low < high {
		mid := (low + high) / 2
		hours := 0
		for _, v := range piles {
			hours += int(math.Ceil(float64(v) / float64(mid)))
		}
		if hours <= h {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return high
}
