package main

import "fmt"

func main() {
	fmt.Println(characterReplacement("ABABCD", 2))
}

func characterReplacement(s string, k int) int {
	charMap := make(map[byte]int)
	maxLen := 0
	var maxRepChar byte = 0
	start, end := 0, 0
	for end < len(s) {
		currentChar := s[end]
		charMap[currentChar]++
		count := charMap[currentChar]
		// check the most repetition character
		if maxRepChar == 0 || charMap[maxRepChar] < count {
			maxRepChar = currentChar
		}

		// keep track of the window size if they exceed max replacement K
		if end-start+1-charMap[maxRepChar] > k {
			// reduce char at start index rep number in the map
			charMap[s[start]]--
			// decrease window size
			start++
		}

		// keep track of max rep
		if maxLen < end-start+1 {
			maxLen = end - start + 1
		}
		end++
	}
	return maxLen
}
