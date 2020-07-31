package factory_method

type Operator interface {
	Result(a, b int) int
}

type OperatorFactory interface {
	Create() Operator
}

//一个工厂只能创建一种产品对象, 添加新类, 该类需要同时实现两个接口
//OperatorAdd
type OperatorAdd struct{}
type OperatorAddFactory struct{}

func (o OperatorAddFactory) Create() Operator {
	return &OperatorAdd{}
}

func (o OperatorAdd) Result(a, b int) int {
	return a + b
}

//OperatorMinus
type OperatorMinus struct{}
type OperatorMinusFactory struct{}

func (o OperatorMinusFactory) Create() Operator {
	return &OperatorMinus{}
}

func (o OperatorMinus) Result(a, b int) int {
	return a - b
}
