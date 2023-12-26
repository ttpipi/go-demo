/*
	按指定权重出现1,2,3,4
*/

package main

import (
	"fmt"
	"math"
	"math/rand"
)

var weight = map[int]int{
	1: 50,
	2: 20,
	3: 30,
	4: 50,
}

var counts = make(map[int]int)

func getTotalWeight() int {
	ret := 0
	for _, v := range weight {
		ret += v
	}
	return ret
}

func getNumber() int {
	x := rand.Intn(math.MaxInt16)
	curPR := 0
	totalPR := getTotalWeight()
	for k, v := range weight {
		curPR += v
		if x <= math.MaxInt16*curPR/totalPR {
			return k
		}
	}
	return -1
}

func main() {
	n := 1000000
	for i := 0; i < n; i++ {
		x := getNumber()
		counts[x]++
	}

	for k, v := range counts {
		fmt.Printf("%d的概率为: %.4f\n", k, float32(v)/float32(n))
	}
}
