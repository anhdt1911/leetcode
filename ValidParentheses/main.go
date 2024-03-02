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
