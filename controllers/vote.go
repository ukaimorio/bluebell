package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func PostVoteController(c *gin.Context) {
	//参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			ResponseError(c, CodeInvalidPram)
			return
		}
		errData := removeTopStruct(errs.Translate(trans)) //翻译并去除错误提示的结构体
		ResponseErrorWithMsg(c, CodeInvalidPram, errData)
		return
	}
	//获取当前请求的用户id
	userID, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost() failed", zap.Error(err))

		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
