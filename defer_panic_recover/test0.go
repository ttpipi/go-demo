/**
	1.recover要在defer中使用才能发挥效果
	2.panic(value) -> recover() -> value
	3.defer在压栈的时候确定参数
	4.不在循环中使用defer
 */

package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f.", r)
		}
	}()

	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g.", i)
	fmt.Println("Printing in g.", i)
	g(i+1)
}