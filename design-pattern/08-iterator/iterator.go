package iterator

import "fmt"

type Aggregate interface {
	Iterator() Iterator
}

type Numbers struct {
	start, end int
}

func (n *Numbers) Iterator() Iterator {
	return &NumbersIterator{
		numbers: n,
		next:    n.start,
	}
}

func NewNumbers(start, end int) *Numbers {
	return &Numbers{
		start: start,
		end:   end,
	}
}

type Iterator interface {
	First()
	IsDone() bool
	Next() interface{}
}

func IteratorPrint(i Iterator) {
	for i.First(); !i.IsDone(); {
		c := i.Next()
		fmt.Printf("%#v\n", c)
	}
}

type NumbersIterator struct {
	numbers *Numbers
	next    int
}

func (n *NumbersIterator) First() {
	n.next = n.numbers.start
}

func (n *NumbersIterator) IsDone() bool {
	return n.next > n.numbers.end
}

func (n *NumbersIterator) Next() interface{} {
	if !n.IsDone() {
		next := n.next
		n.next++
		return next
	}
	return nil
}
