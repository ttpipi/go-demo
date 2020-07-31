package adapter

import (
	"fmt"
	"testing"
)

func TestAdapter(t *testing.T) {
	adaptee := AdapteeImpl{}

	adapter := NewAdapter(adaptee)
	str := adapter.Request()
	fmt.Println(str)
}
