package models

// 用户结构体，这是我们根据用户注册的时候，需要填入的信息，不是全部照抄数据库的所有类型信息
// 这里的tag信息映射的是我们的数据库信息

type User struct {
	UserID int64 `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
}