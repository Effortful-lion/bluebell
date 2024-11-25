package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
	"fmt"

	"go.uber.org/zap"
)

// 用户登录
func Login(p *models.ParamLogin) (token string, err error) {
	// 唯一性登录：用户id和用户名
	// 1. 规定：（用用户名登录）根据用户名查询用户是否存在
	user, err := mysql.GetUserByName(p.Username)
	if user == nil {
		return "", err
	}
	fmt.Println(user)
	// 2. 用户存在，检查对应密码是否正确
	// 密码加密后（在dao层进行密码处理）比较
	PostPassword := mysql.EncryptPassword(p.Password)
	if user.Password != PostPassword {
		return "", err
	}
	// 3. 密码正确，生成token
	token, err = jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return "", err
	}
	// 生成token后，将token存储在redis中
	ID,err := mysql.GetIDByUserName(user.Username)
	if err != nil {
		zap.L().Error("mysql.GetIDByUserName() failed") 
		return "", err
	}
	redis.SetUserToken(token, ID)

	// 作出响应
	return token, nil
}

// 用户注册
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