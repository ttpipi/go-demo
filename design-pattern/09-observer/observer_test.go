package observer

import "testing"

func TestObserver(t *testing.T) {
	reader1 := NewReader("reader1")
	reader2 := NewReader("reader2")

	subject := NewSubject()
	subject.Attach(reader1, reader2)

	subject.UpdateContext("run")

	subject.UpdateContext("game over")
}
