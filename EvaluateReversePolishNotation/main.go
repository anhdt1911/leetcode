package main

import (
	"strconv"
)

func main() {
	evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
}

func evalRPN(tokens []string) int {
	operator := map[string]struct{}{
		"+": {},
		"-": {},
		"*": {},
		"/": {},
	}
	var outStack []int
	for _, v := range tokens {
		if _, ok := operator[v]; !ok {
			numV, _ := strconv.Atoi(v)
			outStack = append(outStack, numV)
		} else {
			latterOp := outStack[len(outStack)-1]
			outStack = outStack[:len(outStack)-1]
			formerOp := outStack[len(outStack)-1]
			outStack = outStack[:len(outStack)-1]
			switch v {
			case "+":
				outStack = append(outStack, (formerOp + latterOp))
			case "-":
				outStack = append(outStack, (formerOp - latterOp))
			case "*":
				outStack = append(outStack, (formerOp * latterOp))
			case "/":
				outStack = append(outStack, (formerOp / latterOp))
			}
		}
	}
	return outStack[0]
}
