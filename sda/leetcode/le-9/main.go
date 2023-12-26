package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	num := 0
	temp := x
	for x != 0 {
		y := x % 10
		x = x / 10
		num = num*10 + y
	}

	return num == temp
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
}
