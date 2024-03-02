package main

import "fmt"

func main() {
	s := "(){}}{"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	stack := []rune{30}
	for _, v := range s {
		if string(v) == "(" || string(v) == "[" || string(v) == "{" {
			stack = append(stack, v)
		} else {
			top := string(stack[len(stack)-1])
			switch string(v) {
			case ")":
				if top == "(" {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			case "]":
				if top == "[" {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			case "}":
				if top == "{" {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			default:
				return true
			}
		}
	}
	return len(stack) == 1
}

func isValidV2(s string) bool {
	brackMap := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	stack := []rune{}
	for _, v := range s {
		if _, ok := brackMap[v]; ok {
			stack = append(stack, v)
			// to check if string has none open bracket
		} else if len(stack) == 0 || brackMap[stack[len(stack)-1]] != v {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
