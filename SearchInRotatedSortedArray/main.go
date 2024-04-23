package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 8, 0, 1, 2}, 4))
}

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := (low + high) / 2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}
