package simple_factory

import (
	"testing"
)

func TestType1(t *testing.T) {
	//调用者需要知道1类型将被实例化成什么
	api := NewApi(1)
	api.Say("ttlin")
}

func TestType2(t *testing.T) {
	api := NewApi(2)
	api.Say("ttlin")
}
