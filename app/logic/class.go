package logic

import (
	"context"
	"github.com/Heqiaomu/hqmGframe/model/request"
	"github.com/Heqiaomu/hqmGframe/utils/gerror"
	"github.com/pkg/errors"
)

func (l *Logic) NewClass(ctx context.Context,class request.Class) (interface{} , error) {
	if class.ClassLeader != 0 {
		// 开始检验当前用户是否存在
		user, err := userSvc.GetUserByID(ctx, class.ClassLeader)
		if err != nil {
			return nil , errors.Wrap(err, "用户查询失败")
		}
		if user.ID == 0 {
			return nil , errors.Wrap(gerror.ErrorFound, "当前用户未找到")
		}
	}
	// 开始进行入库操作
	return classSvc.NewClass(ctx, class)
}
