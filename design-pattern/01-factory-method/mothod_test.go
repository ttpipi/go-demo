package factory_method

import (
	"fmt"
	"testing"
)

func TestOperatorAdd(t *testing.T) {
	op := OperatorAddFactory{}.Create()
	i := op.Result(1, 2)
	fmt.Println(i)
}

func TestOperatorMinus(t *testing.T) {
	op := OperatorMinusFactory{}.Create()
	i := op.Result(1, 2)
	fmt.Println(i)
}
