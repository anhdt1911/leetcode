package main

import "fmt"

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 12))
}

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			high = mid - 1
		} else if target > nums[mid] {
			low = mid + 1
		}
	}
	return -1
}
