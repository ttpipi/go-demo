package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	var result []int
	for i, x := range nums {
		if k, ok := m[target-x]; ok {
			result = append(result, i, k)
			break
		}
		m[x] = i
	}
	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	arr := twoSum(nums, target)
	fmt.Println(arr)
}