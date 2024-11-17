package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error){
	// 判断用户名是否存在
	if err = mysql.CheckUserExit(p.Username);err != nil {
		return err
	}
	// 生成ID
	userID := snowflake.GenID()
	//构造一个用户实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	//保存进数据库
	return mysql.InsertUser(user)

}