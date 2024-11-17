package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) {
	// 判断用户名是否存在
	mysql.QueryUserByUsername()
	// 生成ID
	snowflake.GenID()
	//密码加密

	//保存进数据库
	mysql.InsertUser()

}