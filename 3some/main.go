package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("hello")
}
func threeSum(nums []int) [][]int {
	var out [][]int
	slices.Sort(nums)
	// len(nums) - 2 because need at least two to make a triplet
	for i := 0; i < len(nums)-2; i++ {
		// skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			var arr []int
			sum := nums[left] + nums[right]
			if sum == target {
				arr = append(arr, -target, nums[left], nums[right])
				out = append(out, arr)
				//skip duplicates
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return out
}
