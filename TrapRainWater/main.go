package main

import "fmt"

func main() {
	height := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(height))
}
func trap(height []int) int {
	length := len(height)
	var vol int
	for i := 1; i <= length-2; i++ {
		leftCol, rightCol := 0, 0
		for left := i - 1; left >= 0; left-- {
			if height[left] > leftCol {
				leftCol = height[left]
			}
		}
		for right := i + 1; right <= length-1; right++ {
			if height[right] > rightCol {
				rightCol = height[right]
			}
		}
		if leftCol <= rightCol && leftCol > height[i] {
			vol += leftCol - height[i]
		} else if rightCol < leftCol && rightCol > height[i] {
			vol += rightCol - height[i]
		}
	}
	return vol
}
func trapV2(height []int) int {
	left, right := 0, len(height)-1
	vol, leftCol, rightCol := 0, 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] > leftCol {
				leftCol = height[left]
			} else {
				vol += leftCol - height[left]
			}
			left++
		} else {
			if height[right] > rightCol {
				rightCol = height[right]
			} else {
				vol += rightCol - height[right]
			}
			right--
		}
	}
	return vol
}
