package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	return 0
}

func main() {
	fmt.Println(strStr("hello", "ll"))  //2
	fmt.Println(strStr("aaaaa", "bba")) //-1
	fmt.Println(strStr("hello", ""))    //0
}
