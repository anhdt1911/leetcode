package main

import "fmt"

func main() {
	fmt.Println(characterReplacement("ABAB", 2))
}

func characterReplacement(s string, k int) int {
	charMap := make(map[rune]int)
	res := 0
	for i, v := range s {
		charMap[v] = i
	}
	return res
}
