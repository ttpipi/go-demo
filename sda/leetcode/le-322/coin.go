package coin

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	memo := make(map[int]int)
	return dp(coins, amount, memo)
}

func dp(coins []int, amount int, memo map[int]int) int {
	if _, ok := memo[amount]; ok {
		return memo[amount]
	}

	if amount < 0 {
		return -1
	}

	if amount == 0 {
		return 0
	}

	var res = math.MaxInt64
	for _, coin := range coins {
		var sub int
		if sub = dp(coins, amount-coin, memo); sub == -1 {
			continue
		}

		if sub+1 < res {
			res = sub + 1
		}
	}

	if res == math.MaxInt64 {
		return -1
	}

	memo[amount] = res
	return res
}
