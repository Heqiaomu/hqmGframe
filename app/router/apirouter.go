package router

import (
	"github.com/Heqiaomu/hqmGframe/app/api/file"
	"github.com/Heqiaomu/hqmGframe/app/api/user"
)

var (
	fileApi = new(file.UploadApi)
	userApi = new(user.Api)
)

var (
	fileUpload = new(fileApiRouter)
)
