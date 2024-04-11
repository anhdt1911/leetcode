package main

import "fmt"

func main() {
	arr := []int{13, 8, 1, 5, 2, 5, 9, 7, 6, 12}
	fmt.Println(PreviouslyGreaterElement(arr))
}

func PreviouslyGreaterElement(arr []int) []int {
	result := make([]int, len(arr))
	stack := []int{}
	// another way to solve this is to loop from the end of stack and apply same logic
	for i, v := range arr {
		for len(stack) > 0 && arr[stack[len(stack)-1]] <= v {
			stack = stack[:len(stack)-1]
		}
		// after the loop, only the index of element that is greater than current i is in the stack
		if len(stack) > 0 {
			result[i] = arr[stack[len(stack)-1]]
		}
		stack = append(stack, i)
	}
	return result
}
