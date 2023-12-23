package slice_test

import (
	"fmt"
	"testing"
)

func autoSlice(min, hight int) []int {
	var b = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 99: 100}
	fmt.Println(b)
	fmt.Printf("slice b: len(%d), cap(%d)\n", len(b), cap(b))

	return b[min:hight]
}

func TestSlice(t *testing.T) {
	b1 := autoSlice(3, 7)
	t.Logf("slice b1: len(%d), cap(%d), elements(%v)\n", len(b1), cap(b1), b1)
}

func TestSlice1(t *testing.T) {
	var arr = []int{1, 10: 100}
	t.Log(arr)
}

func TestSlice2(t *testing.T) {
	var b = []int{1, 2, 3, 4}
	t.Logf("slice b: len(%d), cap(%d), elements(%v)\n", len(b), cap(b), b)

	b1 := b[:2]
	t.Logf("slice b1: len(%d), cap(%d), elements(%v)\n", len(b1), cap(b1), b1)

	t.Log("\nappend 11 to b1:")
	b1 = append(b1, 11)
	t.Logf("slice b1: len(%d), cap(%d), elements(%v)\n", len(b1), cap(b1), b1)
	t.Logf("slice b: len(%d), cap(%d), elements(%v)\n", len(b), cap(b), b)

	t.Log("\nappend 22 to b1:")
	b1 = append(b1, 22)
	t.Logf("slice b1: len(%d), cap(%d), elements(%v)\n", len(b1), cap(b1), b1)
	t.Logf("slice b: len(%d), cap(%d), elements(%v)\n", len(b), cap(b), b)

	t.Log("\nappend 33 to b1:")
	b1 = append(b1, 33)
	t.Logf("slice b1: len(%d), cap(%d), elements(%v)\n", len(b1), cap(b1), b1)
	t.Logf("slice b: len(%d), cap(%d), elements(%v)\n", len(b), cap(b), b)

	b1[0] *= 100
	t.Log("\nb1[0] multiply 100:")
	t.Logf("slice b1: len(%d), cap(%d), elements(%v)\n", len(b1), cap(b1), b1)
	t.Logf("slice b: len(%d), cap(%d), elements(%v)\n", len(b), cap(b), b)
}
