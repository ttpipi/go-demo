package coin

import "testing"

var tests = []struct {
	coins  []int
	amount int
	result int
}{
	{[]int{1, 2, 5}, 11, 3},
	{[]int{2}, 3, -1},
	{[]int{1}, 2, 2},
	{[]int{186, 419, 83, 408}, 6249, 20},
}

func TestCoinChange2(t *testing.T) {
	for i, tt := range tests {
		if n := coinChange2(tt.coins, tt.amount); n != tt.result {
			t.Errorf("coinChange2(); 第%d组; expected %d; but got %d", i, tt.result, n)
		}
	}
}

func TestCoinChange(t *testing.T) {
	for i, tt := range tests {
		if n := coinChange(tt.coins, tt.amount); n != tt.result {
			t.Errorf("coinChange(); 第%d组; expected %d; but got %d", i, tt.result, n)
		}
	}
}
