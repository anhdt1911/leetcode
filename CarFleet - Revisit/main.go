package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(carFleet(25, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))
}

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	cars := make([][2]int, n)
	for i := range cars {
		cars[i] = [2]int{position[i], speed[i]}
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] > cars[j][0]
	})

	fleet := 0
	time := 0.0
	for i := range cars {
		duration := (float64(target) - float64(cars[i][0])) / float64(cars[i][1])
		if duration > time {
			time = duration
			fleet++
		}
	}
	return fleet
}
