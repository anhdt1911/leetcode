package main

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	for i, v := range s {
		var runeArr []rune
		runeArr = append(runeArr, v)
	loop:
		for _, val := range s[i+1:] {
			for _, value := range runeArr {
				if value == val {
					break loop
				}
			}
			runeArr = append(runeArr, val)
		}
		if len(runeArr) > maxLen {
			maxLen = len(runeArr)
		}
	}
	return maxLen
}
