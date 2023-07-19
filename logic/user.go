package logic

import (
	"bluebell/Dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

//存放业务逻辑代码

func SignUp(p *models.ParamSignUp) (err error) {
	//1.判断用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		//数据库查询出错
		return err
	}
	//2.生成UID
	userID := snowflake.GenID()
	//构造一个user实例
	user := models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//3.密码加密

	//4.保存进数据库
	return mysql.InsertUser(&user)
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//传递的是指针，在传入Login后能拿到Userid
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	//生产JWT的token
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
