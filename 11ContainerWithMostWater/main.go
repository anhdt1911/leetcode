package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea1stSubmition(height))
}

func maxArea1stSubmition(height []int) int {
	left, right := 0, len(height)-1
	var maxAmount int
	for left < right {
		smaller, compare := smaller(height[left], height[right])
		currentAmount := smaller * (right - left)
		if currentAmount > maxAmount {
			maxAmount = currentAmount
		}
		switch compare {
		case "left":
			left++
		case "right":
			right--
		default:
			left++
			right--
		}
	}
	return maxAmount
}

func smaller(left, right int) (int, string) {
	if left < right {
		return left, "left"
	} else if left == right {
		return left, "equal"
	} else {
		return right, "right"
	}
}
