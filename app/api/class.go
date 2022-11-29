package api

import (
	"github.com/Heqiaomu/hqmGframe/model/request"
	"github.com/Heqiaomu/hqmGframe/utils/gconsts"
	"github.com/Heqiaomu/hqmGframe/utils/gresp"
	"github.com/gin-gonic/gin"
)

func (api *Api) NewClass(ctx *gin.Context) {

	class := request.Class{}
	err := ctx.ShouldBind(&class)
	if err != nil {
		gresp.ErrorParam(ctx, class)
		return
	}

	// 进入 logic
	newClass, err := l.NewClass(ctx, class)
	if err != nil {
		gresp.Fail(ctx, gconsts.ServerOccurredErrorCode , gconsts.ServerOccurredErrorMsg+err.Error() , nil)
		return
	}
	gresp.Success(ctx,"ok", newClass)
	return
}
