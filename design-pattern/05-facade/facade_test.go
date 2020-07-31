package facade

import "testing"

func TestFacade(t *testing.T) {
	s := NewSystem()
	s.Run()
}
