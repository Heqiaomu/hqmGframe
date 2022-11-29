package user

import (
	"github.com/Heqiaomu/hqmGframe/model/request"
	"github.com/Heqiaomu/hqmGframe/utils/gresp"
	"github.com/gin-gonic/gin"
)

type Api struct {
}

func (api *Api) Register(ctx *gin.Context) {
	register := request.Register{}
	//err := gvalidator.ShouldBindFormDataToModel(ctx, &register)
	err := ctx.ShouldBind(&register)
	if err != nil {
		gresp.ErrorParam(ctx, register)
		return
	}
}
