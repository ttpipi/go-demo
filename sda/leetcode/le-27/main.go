package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 1 && nums[0] != val {
		return 1
	}

	i, j := 0, len(nums)
	for i < j {
		if nums[i] != val {
			i++
		} else {
			for j > i {
				j--
				if nums[j] != val {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
	return i
}

func main() {
	arr := []int{3, 2, 2, 3}
	fmt.Println(removeElement(arr, 3)) //2

	arr = []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(arr, 2)) //5

	arr = []int{3, 3}
	fmt.Println(removeElement(arr, 5))
}
