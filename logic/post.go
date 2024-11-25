package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error){
	// 1. 生成postID（雪花算法）
	p.ID = snowflake.GenID()
	// 2. 数据库保存post
	mysql.CreatePost(p)
	return 
}

func GetPostDetail(id int64) (data *models.ApiPostDetail, err error) {
	// 查询并组合我们接口想用的数据
	post, err := mysql.GetPostById(id)
	if err != nil {
		zap.L().Error("mysql.GetPostById() failed", zap.Error(err))
		return
	}
	
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserById() failed",zap.Int64("author_id",post.AuthorID), zap.Error(err))
		return
	}
	
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID() failed", zap.Error(err))
		return
	}
	
	data = &models.ApiPostDetail{
		Post: post,
		AuthorName: user.Username,
		CommunityDetail: community, 
	}
	// 查询并返回数据
	return 
}

func GetPostList(page,size int64)(data []*models.ApiPostDetail, err error){
	// 获取帖子列表
	posts, err := mysql.GetPostList(page,size)
	// 将结果追加到data中：（其实还是一个帖子列表，只是每个帖子都携带作者信息和社区信息等）
	data = make([]*models.ApiPostDetail, 0, len(posts))
	// 遍历每个帖子，补全信息
	for _, post := range posts {
		// TODO: 发现：在嵌套结构体中，外层结构体的实例可以直接访问内层结构体的变量，所以下面可以写成：
		// user, err := GetUserById(post.AuthorID)
		user, err := mysql.GetUserById(post.Post.AuthorID)
		if err != nil {
			zap.L().Error("GetUserById() failed", zap.Int64("author_id",post.Post.AuthorID),zap.Error(err))
			return nil, err
		}
		community, err := mysql.GetCommunityDetailByID(post.Post.CommunityID)
		if err != nil {
			zap.L().Error("GetCommunityDetailByID() failed", zap.Error(err))
			return nil, err
		}
		// 包装每一个帖子细节：
		postDetail := &models.ApiPostDetail{
			AuthorName: user.Username,
			CommunityDetail: community,
			Post: post.Post,
		}
		data = append(data, postDetail)
	}
	return

}
