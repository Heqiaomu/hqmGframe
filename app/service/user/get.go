package user

import (
	"context"
	dao "github.com/Heqiaomu/hqmGframe/model/entity/userdao"
)

type Service struct {
}

func (s *Service) GetUserByPhone(ctx context.Context, phoneNum int64) (*dao.User, error) {
	user, err := dao.NewUserDAO(ctx).GetUserByPhone(phoneNum)
	if err != nil {
		return nil, err
	}
	return user, nil
}
