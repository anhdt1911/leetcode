package main

// https://leetcode.com/discuss/study-guide/2347639/A-comprehensive-guide-and-template-for-monotonic-stack-based-problems ->
// this is big help to understand monotonic stack
func main() {
	largestRectangleArea([]int{1, 1})
}

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	heights = append([]int{0}, heights...)
	nextSmaller := make([]int, len(heights))
	previousSmaller := make([]int, len(heights))
	stack := []int{}
	for i, v := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > v {
			nextSmaller[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			previousSmaller[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	area := 0
	for i, height := range heights {
		width := nextSmaller[i] - previousSmaller[i] - 1
		if height*width > area {
			area = height * width
		}
	}
	return area
}
