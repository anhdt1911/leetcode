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

// bpfbhmipx
func lengthOfLongestSubstringUseingSlidingWindow(s string) int {

	if len(s) <= 1 {
		return len(s)
	}

	runeMap, result, start := make(map[rune]int), 0, 0

	for i, v := range s {
		if _, ok := runeMap[v]; !ok {
			runeMap[v] = i
		} else {
			start = runeMap[v] + 1

			for j := 0; j < start; j++ {
				char := rune(s[j])
				if runeMap[char] >= start {
					continue
				}
				delete(runeMap, char)
			}

			runeMap[v] = i
		}

		if len(runeMap) > result {
			result = len(runeMap)
		}
	}

	return result
}
