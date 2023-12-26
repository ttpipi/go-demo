package main

import "fmt"

var m = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	num := 0
	pre := m[s[0]]
	for i := 1; i < len(s); i++ {
		cur := m[s[i]]
		if pre >= cur {
			num += pre
		} else {
			num -= pre
		}
		pre = cur
	}
	return num + pre
}

func main() {
	fmt.Println(romanToInt("LVIII"))   //58
	fmt.Println(romanToInt("MCMXCIV")) //1994
}
