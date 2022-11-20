package user

import (
	"fmt"
	"github.com/Heqiaomu/hqmGframe/app/logic"
	"github.com/Heqiaomu/hqmGframe/model/request"
	"github.com/Heqiaomu/hqmGframe/utils/gconsts"
	"github.com/Heqiaomu/hqmGframe/utils/gresp"
	"github.com/gin-gonic/gin"
)

func (api *Api) Login(ctx *gin.Context) {
	login := request.Login{}

	err := ctx.ShouldBind(&login)
	//err := gvalidator.ShouldBindFormDataToModel(ctx, &login)
	if err != nil {
		gresp.ErrorParam(ctx, login)
		return
	}
	l := &logic.Logic{}
	resp, err := l.Login(ctx, login)
	if err != nil {
		gresp.Fail(ctx, gconsts.CurdLoginFailCode, fmt.Sprintf("%s:%s", gconsts.CurdLoginFailMsg, err.Error()), nil)
		return
	}
	gresp.Success(ctx, "success", resp)
	return
}
