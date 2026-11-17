package main

import "fmt"

func longestValidParentheses(s string) int {
	stack := []int{-1}
	ans := 0

	for i := 0; i < len(s); i++ {

		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			length := i - stack[len(stack)-1]

			if length > ans {
				ans = length
			}
		}
	}

	return ans
}

func main() {
	var s string

	fmt.Print("Enter parentheses string: ")
	fmt.Scanln(&s)

	result := longestValidParentheses(s)

	fmt.Println("Longest Valid Parentheses Length:", result)
}
