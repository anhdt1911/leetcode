package main

import "fmt"

func main() {
	// add 2 number
	// l1 := &ListNode{1, nil}
	// l2 := &ListNode{1, nil}
	// addTwoNumbers(l1, l2)

	// longest sub string
	// s := "pwwkew&^"
	// maxLen := lengthOfLongestSubstring(s)
	// fmt.Println(maxLen)

	// median of 2 sorted array
	arr1 := []int{1, 3}
	arr2 := []int{2}
	median := findMedianSortedArrays(arr1, arr2)
	fmt.Println(median)
}
