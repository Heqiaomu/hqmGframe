package logic

import (
	"github.com/Heqiaomu/hqmGframe/app/service/class"
	"github.com/Heqiaomu/hqmGframe/app/service/user"
)

var (
	userSvc = new(user.Service)
	classSvc = new(class.Service)
)
