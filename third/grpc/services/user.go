package services

import (
	context "context"
	"io"
	"log"
	"time"
)

type UserService struct{}

func (*UserService) GetUserScore(_ context.Context, in *UserScoreRequest) (*UserScoreResponse, error) {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for _, user := range in.Users {
		user.Score = score
		score++
		users = append(users, user)
	}
	return &UserScoreResponse{Users: users}, nil
}

func (*UserService) GetUserScoreStream(in *UserScoreRequest, stream UserService_GetUserScoreStreamServer) error {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for index, user := range in.Users {
		user.Score = score
		score++
		users = append(users, user)

		//每隔2条分批发送
		if (index+1)%2 == 0 && index > 0 {
			err := stream.Send(&UserScoreResponse{Users: users})
			if err != nil {
				return err
			}

			//清空切片
			users = (users)[0:0]
		}

		//模拟耗时处理情景
		time.Sleep(time.Second * 2)
	}

	if len(users) > 0 {
		err := stream.Send(&UserScoreResponse{Users: users})
		if err != nil {
			return err
		}
	}
	return nil
}

func (*UserService) GetUserScoreByClientStream(stream UserService_GetUserScoreByClientStreamServer) error {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&UserScoreResponse{Users: users})
			}
			return err
		}

		for _, user := range req.Users {
			user.Score = score
			score++
			users = append(users, user)
		}
	}
}

func (*UserService) GetUserScoreByTWS(stream UserService_GetUserScoreByTWSServer) error {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		for _, user := range req.Users {
			user.Score = score
			score++
			users = append(users, user)
		}

		if err = stream.Send(&UserScoreResponse{Users: users}); err != nil {
			log.Println(err)
		}

		users = (users)[0:0]
	}
}
