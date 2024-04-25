package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 8, 0, 1, 2}, 0))
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

	pivotIndex := low

	// not rotated
	if pivotIndex == 0 {
		low, high = 0, len(nums)-1
		// target at right side of pivot
	} else if pivotIndex != 0 && target < nums[0] {
		low, high = pivotIndex, len(nums)-1
	} else {
		low, high = 0, pivotIndex-1
	}

	for low <= high {
		mid := (low + high) / 2
		if target > nums[mid] {
			low = mid + 1
		} else if target < nums[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
