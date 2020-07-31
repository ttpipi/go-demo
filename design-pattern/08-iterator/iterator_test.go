package iterator

import "testing"

func TestIterator(t *testing.T) {
	var aggreagte Aggregate
	aggreagte = NewNumbers(1, 10)
	IteratorPrint(aggreagte.Iterator())
}
