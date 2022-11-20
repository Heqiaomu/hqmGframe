package logic

import (
	"context"
	"github.com/Heqiaomu/hqmGframe/model/request"
	"github.com/Heqiaomu/hqmGframe/model/response"
	"github.com/Heqiaomu/hqmGframe/utils/gerror"
	"github.com/Heqiaomu/hqmGframe/utils/gjwt"
	"time"
)

// Login 登录的 逻辑实现层
func (l *Logic) Login(ctx context.Context, login request.Login) (interface{}, error) {
	// 从数据库获取当前用的数据，是否一致
	userModel, err := userSvc.GetUserByPhone(context.Background(), login.Phone)
	if err != nil {
		return nil, err
	}

	if userModel.Password != login.Password {
		return nil, gerror.ErrorAuthPassword
	}

	token, err := gjwt.KeyIn().Token(gjwt.CustomClaims{
		UserId:    userModel.UserID,
		Name:      userModel.Username,
		ClientIP:  ctx.Value("ClientIP").(string),
		Timestamp: time.Now().Unix(),
		Expire:    time.Minute * 30,
	})
	if err != nil {
		return nil, gerror.ErrorAuthToken
	}

	return response.LoginResp{
		UserID:    userModel.UserID,
		Username:  userModel.Username,
		Token:     token,
		LoginTime: time.Now(),
		ClientIP:  ctx.Value("ClientIP").(string),
	}, nil
}
