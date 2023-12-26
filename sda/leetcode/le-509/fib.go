package fib

func fib(n int) int {
	if n < 0 {
		return -1
	}

	if n == 0 || n == 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
