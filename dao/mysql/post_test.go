package mysql

//测试mysql下的函数

import (
	"bluebell/models"
	"bluebell/settings"
	"testing"
)

// 初始化mysql连接, 防止下面在测试调用时数据库时，使用的是空的db
func init() {
	dbCfg := settings.MysqlConfig{
		Host: "127.0.0.1",
		Port: 3306,
		User: "root",
		Password: "mysql123",
		DB: "sql_demo",
		MaxOpenConns: 10,
		MaxIdleConns: 10,
	}
	if err := Init(&dbCfg); err != nil {
		panic(err)
	}
}

func TestCreatePost(t *testing.T){
	// 测试 在mysql数据库中创建帖子

	post := models.Post{
		ID: 10,
		AuthorID: 1,
		CommunityID: 1,
		Title: "测试标题",
		Content: "测试内容",
	}
	err := CreatePost(&post)
	if err != nil {
		t.Fatalf("create post failed, err:%v", err)
	}
	t.Logf("create post success, post id:%d", post.ID)
}
