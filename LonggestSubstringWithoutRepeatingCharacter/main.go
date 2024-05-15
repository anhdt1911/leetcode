package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
}

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[rune]int)
	start := 0
	longestLen := 0
	for i, v := range s {
		if val, ok := charMap[v]; ok && val >= start {
			start = val + 1
		}
		if i-start+1 > longestLen {
			longestLen = i - start + 1
		}
		charMap[v] = i
	}
	return longestLen
}
