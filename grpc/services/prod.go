package services

import "context"

type ProdService struct{}

func (*ProdService) GetProdStock(_ context.Context, req *ProdRequest) (*ProdRespon, error) {
	var stock int32
	if req.Area == ProdAreas_A {
		stock = 10
	} else if req.Area == ProdAreas_B {
		stock = 20
	} else {
		stock = 30
	}
	return &ProdRespon{Stock: stock}, nil
}

func (*ProdService) GetProdStocks(_ context.Context, _ *QuerySize) (*ProdResponList, error) {
	stocks := []*ProdRespon{
		&ProdRespon{Stock: 90},
		&ProdRespon{Stock: 91},
		&ProdRespon{Stock: 92},
		&ProdRespon{Stock: 93},
	}
	return &ProdResponList{Stocks: stocks}, nil
}

func (*ProdService) GetProdInfo(context.Context, *ProdRequest) (*ProdModel, error) {
	ret := ProdModel{
		Id:    123,
		Name:  "测试商品",
		Price: 250.9,
	}
	return &ret, nil
}
