package services

import "errors"

type DivService struct{}

type DivRequest struct {
	A, B int
}

type DivRespon struct {
	Ret int
}

func (DivService) Div(param DivRequest, result *DivRespon) error {
	if param.B == 0 {
		return errors.New("除数不能为0")
	}
	result.Ret = param.A / param.B
	return nil
}
