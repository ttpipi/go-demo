package main

import "fmt"

func removeDuplicates(nums []int) int {
	i, j := 0, 1
	for j < len(nums) {
		if nums[i] == nums[j] {
			j++
		} else {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func main() {
	arr := []int{1, 1, 2}
	fmt.Println(removeDuplicates(arr)) //2

	arr = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(arr)) //5

	arr = []int{0, 1}
	fmt.Println(removeDuplicates(arr))
}
