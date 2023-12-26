package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	return strs[0]
}

func main() {
	s := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(s))

	s = []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(s))
}
