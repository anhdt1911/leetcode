package main

import "fmt"

func main() {
	arr := []int{13, 8, 1, 5, 2, 5, 9, 7, 6, 12}
	fmt.Println(NextGreaterElement(arr))
}

func NextGreaterElement(arr []int) []int {
	result := make([]int, len(arr))
	stack := []int{}
	for i, v := range arr {
		for len(stack) > 0 && arr[stack[len(stack)-1]] < v {
			result[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return result
}
