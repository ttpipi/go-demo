package main

import "fmt"

func isValid(s string) bool {
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack []rune
	for _, b := range s {
		if b == '(' || b == '[' || b == '{' {
			stack = append(stack, b)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != m[b] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("}}}"))
}
