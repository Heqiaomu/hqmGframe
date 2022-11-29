package class

import (
	"context"
	"github.com/Heqiaomu/hqmGframe/model/entity"
	"github.com/Heqiaomu/hqmGframe/model/request"
)

type Service struct {
	
}

func (s Service) NewClass(ctx context.Context, class request.Class) (interface{} , error) {
	// 入库操作
	modelClass := entity.Class{
		ClassName:   class.ClassName,
		Description: class.Description,
		ClassLeader: class.ClassLeader,
	}
	err := entity.New(ctx).CreateClass(modelClass)
	return modelClass,err
}