package main

import "fmt"

// break the task
// find all permutation of s1
// traverse through s2 using 2 pointer, but with what condition for window resizing?

func checkInclusion(s1 string, s2 string) bool {
	s1CharMap := map[byte]int{}
	removedCharMap := map[byte]int{}
	s1Length := len(s1)
	start, end := 0, 0
	for _, v := range s1 {
		s1CharMap[byte(v)]++
	}
	for end < len(s2) {
		if _, ok := s1CharMap[s2[end]]; !ok {
			start++
			end++
			continue
		} else {
			end++
		}
	}

	return false
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
}
