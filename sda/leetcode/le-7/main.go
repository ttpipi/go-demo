package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	num := 0
	for x != 0 {
		y := x % 10
		x = x / 10
		num = num*10 + y
		if num > math.MaxInt32 || num < math.MinInt32 {
			return 0
		}
	}
	return num
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-1000))
}
