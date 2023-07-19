package controllers

import (
	"bluebell/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	strconv2 "strconv"
)

//跟社区相关

func CommunityHandler(c *gin.Context) {
	// 查询到所有社区(community_id,community_name)以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("login.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	//1.获取社区id
	idStr := c.Param("id")
	id, err := strconv2.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidPram)
		return
	}
	// 查询到所有社区(community_id,community_name)以列表的形式返回
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("login.GetCommunityDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
