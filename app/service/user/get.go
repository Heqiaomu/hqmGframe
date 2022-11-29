package user

import (
	"context"
	"github.com/Heqiaomu/hqmGframe/model/entity"
)

type Service struct {
}

func (s *Service) GetUserByPhone(ctx context.Context, phoneNum int64) (*entity.User, error) {
	user, err := entity.New(ctx).GetUserByPhone(phoneNum)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) GetUserByID(ctx context.Context, id int) (*entity.User, error) {
	user, err := entity.New(ctx).GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}