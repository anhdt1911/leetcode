package main

import "fmt"

func main() {
	for _, v := range generateParenthesis(3) {
		fmt.Println(v)
	}
}
func generateParenthesis(n int) []string {
	out := &[]string{}
	current := make([]byte, n*2)
	backtracking(out, n, 0, 0, current)
	return *out
}

func backtracking(out *[]string, n, leftBracket, rightBracket int, current []byte) {
	if leftBracket+rightBracket == n*2 {
		*out = append(*out, string(current))
	}

	if leftBracket < n {
		current[leftBracket+rightBracket] = '('
		backtracking(out, n, leftBracket+1, rightBracket, current)
	}

	if rightBracket < leftBracket {
		current[leftBracket+rightBracket] = ')'
		backtracking(out, n, leftBracket, rightBracket+1, current)
	}
}
