package myrpc

import (
	"context"
	"crawler/comm"

	"github.com/olivere/elastic/v7"
)

type SaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *SaverService) Save(item comm.Profile, respon *string) error {
	_, err := s.Client.Index().
		Index(s.Index).
		Id(item.Id).
		BodyJson(item.Data).
		Do(context.Background())

	if err != nil {
		return err
	}

	*respon = "ok"
	return nil
}
