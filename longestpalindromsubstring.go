package main

func longestPalindrome(s string) string {

	strLen := len(s)
	var high, low int
	var result string

	for i := 0; i < strLen; i++ {
		high = i
		low = i - 1
		current := ""
		for low >= 0 && high < strLen {
			if s[low] == s[high] {
				current = string(s[low]) + current + string(s[high])
				high += 1
				low -= 1
			} else {
				break
			}
		}
		if len(current) > len(result) {
			result = current
		}

		current = string(s[i])
		high = i + 1
		low = i - 1
		for low >= 0 && high < strLen {
			if s[low] == s[high] {
				current = string(s[low]) + current + string(s[high])
				high += 1
				low -= 1
			} else {
				break
			}
		}
		if len(current) > len(result) {
			result = current
		}

	}

	return result
}
