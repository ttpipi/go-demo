package fib

func __fib(dp []int, n int) int {

	if n == 0 || n == 1 {
		return n
	}

	if dp[n] != 0 {
		return dp[n]
	}

	dp[n] = __fib(dp, n-1) + __fib(dp, n-2)
	return dp[n]
}

func fib1(n int) int {
	if n < 0 {
		return -1
	}

	dp := make([]int, n+1)
	return __fib(dp, n)
}
