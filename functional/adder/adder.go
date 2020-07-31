/**
闭包
*/

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	a := adder()
	var n int
	for i := 0; i < 10; i++ {
		n = a(i)
	}
	fmt.Println(n)
}
