package fib

func fib3(n int) int {
	if n < 0 {
		return -1
	}

	if n == 0 || n == 1 {
		return n
	}

	pre, cur := 0, 1
	value := 0
	for i := 2; i <= n; i++ {
		value = pre + cur
		pre, cur = cur, value
	}
	return value
}
