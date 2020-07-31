package proxy

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {
	sub := ProxySubject{}
	str := sub.Do()
	fmt.Println(str)
}
