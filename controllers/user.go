package controllers

import (
	"bluebell/Dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	//1.参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误，直接返回响应
		zap.L().Error("SingUp with invalid param", zap.Error(err))
		//判断err是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidPram)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidPram, removeTopStruct(errs.Translate(trans)))
		return
	}
	//手动对请求参数进行详细的业务规校验
	/*if len(p.Username)==0||len(p.Password)==0||len(p.RePassword)==0||p.RePassword!=p.Password{
		zap.L().Error("SingUp with invalid param")
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}*/
	//2.业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	//1.参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		//判断err是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidPram)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidPram, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2.业务处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	//3.返回响应
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d",user.UserID),
		"user_name": user.Username,
		"token":     user.Token,
	})
}
