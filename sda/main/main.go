package main

import (
	"fmt"
)

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	w, n := sum/2, len(nums)

	dp := make([][]bool, w+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	for j := range dp[0] {
		dp[0][j] = true
	}

	for i := 1; i <= w; i++ {
		for j := 1; j <= n; j++ {
			if nums[j-1] > i {
				dp[i][j] = dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1] || dp[i-nums[j-1]][j-1]
			}
		}
	}
	return dp[w][n]
}

func main() {
	fmt.Println(canPartition([]int{1, 2, 5}))
}
