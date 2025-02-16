package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
)

var secret = "bluebell"   // TODO:这里的secret的作用是什么？

// 查询指定用户名的用户id
func GetIDByUserName(username string) (ID int64, err error) {
	sqlStr := `select id from user where username = ?`
	err = db.Get(&ID, sqlStr, username)
	return
}

func GetIDByUserID(userID int64) (ID int64, err error) {
	sqlStr := `select id from user where user_id = ?`
	err = db.Get(&ID, sqlStr, userID)
	return
}

// 查询指定用户名的用户
func GetUserByName(username string) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id,password, username from user where username = ?`
	err = db.Get(user, sqlStr, username)
	if err == sql.ErrNoRows {
		return nil,ErrorUserNotExist 
	}
	if err != nil { 
		// 查询出错
		return nil,err
	}
	return user,err 
}

// 查询指定用户名的用户是否存在
func CheckUserExit(username string) (err error) {
	//执行查询语句
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

func InsertUser(user *models.User) (err error){
	//执行插入语句
	user.Password = EncryptPassword(user.Password)
	sqlStr := `insert into user(user_id,username,password) values (?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)// 语句 + 几个参数
	if err != nil {
		fmt.Println(err)
	}
	return
}

// 加密函数
func EncryptPassword(oPwd string) string {
	// 对密码进行加密
	h := md5.New()
	// 将用户密码作为参数传进去
	h.Write([]byte(secret))
	// 将加密结果作为参数传进去 并 转换为字符串
	return hex.EncodeToString(h.Sum([]byte(oPwd)))
}

// 根据用户ID获取用户信息
func GetUserById(id int64) (user *models.User, err error) {
	user = new(models.User)
	// 查询用户
	sqlStr := "select user_id, username from user where user_id = ?"
	err = db.Get(user, sqlStr, id)
	if err == sql.ErrNoRows {
		return nil,ErrorUserNotExist 
	}
	if err != nil { 
		// 查询出错
		return nil,err
	}
	return
}