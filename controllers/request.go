package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	strconv2 "strconv"
)

const CxtUserIDKey = "UserID"

var ErrorUserNotLogin = errors.New("用户未登入")

// GetCurrentUserID 获取当前用户登入的ID
func GetCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CxtUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// getPageInfo 获取分页信息的要求
func getPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	var (
		page int64
		size int64
		err  error
	)
	page, err = strconv2.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv2.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
