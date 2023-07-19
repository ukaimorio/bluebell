package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	strconv2 "strconv"
)

// CreatePostHandler 创建帖子的处理函数
func CreatePostHandler(c *gin.Context) {
	// 1.获取参数以及参数的校验

	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) err", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		ResponseError(c, CodeInvalidPram)
	}
	//从c取道当前发送请求的用户id
	userID, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	// 2.创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 获取帖子详情的处理函数
func GetPostDetailHandler(c *gin.Context) {
	// 1.获取参数
	pidStr := c.Param("id")
	pid, err := strconv2.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidPram)
		return
	}
	// 2.根数id查询帖子数据
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostByID(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, data)
}

func GetPostListHandler(c *gin.Context) {
	// 获取数据
	page, size := getPageInfo(c)
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// GetPostListHandler2 (按分数按创建时间排序)根据前端传来的参数动态的去获取帖子列表
// 1.获取参数
// 2.去redis查询id列表
// 3.根据id去数据库查询帖子详细信息
func GetPostListHandler2(c *gin.Context) {
	//GET请求参数：/api/v1/posts2?page=1&size=10&order=time
	// 获取数据
	// c.ShouldBindJSON() 只有请求中携带的是json数据才使用这个方法去获取数据
	// c.ShouldBind() 根据请求中的数据类型选择相应的方法去获取数据
	// 初始化结构体是指定初始参数
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidPram)
		return
	}
	data,err:=logic.GetPostListNew(p) //跟新:合二为一
	if err != nil {
		zap.L().Error("logic.GetPostList2() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

/*func GetCommunityPostListHandler(c *gin.Context) {
	p := &models.ParamCommunityPostList{
		ParamPostList: &models.ParamPostList{
			Page:  1,
			Size:  10,
			Order: models.OrderTime,
		},
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetCommunityPostListHandler with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidPram)
		return
	}
	data, err := logic.GetCommunityPostList(p)
	if err != nil {
		zap.L().Error("logic.GetCommunityPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}*/
