package fib

import "testing"

var tests = []struct {
	i, value int
}{
	{0, 0},
	{1, 1},
	{26, 121393},
	{29, 514229},
	{30, 832040},
}

func TestFib3(t *testing.T) {
	for _, tt := range tests {
		if n := fib3(tt.i); n != tt.value {
			t.Errorf("fib3(%d); got %d; expected %d", tt.i, n, tt.value)
		}
	}
}

func TestFib2(t *testing.T) {
	for _, tt := range tests {
		if n := fib2(tt.i); n != tt.value {
			t.Errorf("fib2(%d); got %d; expected %d", tt.i, n, tt.value)
		}
	}
}

func TestFib1(t *testing.T) {
	for _, tt := range tests {
		if n := fib1(tt.i); n != tt.value {
			t.Errorf("fib1(%d); got %d; expected %d", tt.i, n, tt.value)
		}
	}
}

func TestFib(t *testing.T) {
	for _, tt := range tests {
		if n := fib(tt.i); n != tt.value {
			t.Errorf("fib(%d); got %d; expected %d", tt.i, n, tt.value)
		}
	}
}
