package user

import (
	"github.com/Heqiaomu/hqmGframe/model/request"
	"github.com/Heqiaomu/hqmGframe/utils/gresp"
	"github.com/Heqiaomu/hqmGframe/utils/gvalidator"
	"github.com/gin-gonic/gin"
)

type Api struct {
}

func (api *Api) Register(ctx *gin.Context) {
	register := request.Register{}
	err := gvalidator.ShouldBindFormDataToModel(ctx, &register)
	if err != nil {
		gresp.ErrorParam(ctx, register)
		return
	}
}
