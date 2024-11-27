package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error){
	// 1. 生成postID（雪花算法）
	p.ID = snowflake.GenID()
	// 2. 数据库保存post
	mysql.CreatePost(p)
	// 3. 将数据同步到redis中
	redis.CreatePost(p.ID,p.CommunityID)
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

	// mysql数据库查完了，轮到redis数据库查询投票数了
	voteData:= redis.GetPostVoteDataByID(id)
	
	data = &models.ApiPostDetail{
		Post: post,
		VoteNum: voteData,
		AuthorName: user.Username,
		CommunityDetail: community, 
	}
	// 查询并返回数据
	return 
}

// // 获取帖子列表（ 没用了 ）
// func GetPostList(page,size int64)(data []*models.ApiPostDetail, err error){
// 	posts, err := mysql.GetPostList(page,size)
// 	// 将结果追加到data中：（其实还是一个帖子列表，只是每个帖子都携带作者信息和社区信息等）
// 	data = make([]*models.ApiPostDetail, 0, len(posts))
// 	// 遍历每个帖子，补全信息
// 	for _, post := range posts {
// 		// TODO: 发现：在嵌套结构体中，外层结构体的实例可以直接访问内层结构体的变量，所以下面可以写成：
// 		// user, err := GetUserById(post.AuthorID)
// 		user, err := mysql.GetUserById(post.Post.AuthorID)
// 		if err != nil {
// 			zap.L().Error("GetUserById() failed", zap.Int64("author_id",post.Post.AuthorID),zap.Error(err))
// 			return nil, err
// 		}
// 		community, err := mysql.GetCommunityDetailByID(post.Post.CommunityID)
// 		if err != nil {
// 			zap.L().Error("GetCommunityDetailByID() failed", zap.Error(err))
// 			return nil, err
// 		}
// 		// 包装每一个帖子细节：
// 		postDetail := &models.ApiPostDetail{
// 			AuthorName: user.Username,
// 			CommunityDetail: community,
// 			Post: post.Post,
// 		}
// 		data = append(data, postDetail)
// 	}
// 	return

// }

// 获取帖子列表(可以获得投票数)
func GetPostList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	// 去redis中查询id列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		zap.L().Error("GetPostList2 GetPostIDsInOrder failed", zap.Error(err))
		return
	}
	// 根据id在mysql中查询帖子详情
	if len(ids) == 0 {
		zap.L().Warn("GetPostList2 ids is empty")
		return
	}
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		zap.L().Error("GetPostList2 GetPostListByIDs failed", zap.Error(err))
		return
	}

	// 提前获取每篇帖子的投票数据
	voteData,err := redis.GetPostVoteData(ids)
	if err != nil {
		zap.L().Error("GetPostList2 GetPostVoteData failed", zap.Error(err))
		return
	}
	 

	data = make([]*models.ApiPostDetail, 0, len(posts))
	// 遍历每个帖子，补全信息
	for idx, post := range posts {
		// 发现：在嵌套结构体中，外层结构体的实例可以直接访问内层结构体的变量，所以下面可以写成：
		// user, err := GetUserById(post.AuthorID)
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("GetUserById() failed", zap.Int64("author_id",post.AuthorID),zap.Error(err))
			return nil, err
		}
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("GetCommunityDetailByID() failed", zap.Error(err))
			return nil, err
		}
		// 包装每一个帖子细节：
		postDetail := &models.ApiPostDetail{
			AuthorName: user.Username,
			VoteNum: voteData[idx],
			CommunityDetail: community,
			Post: post,
		}
		data = append(data, postDetail)
	}
	return

}


func GetCommunityPostList(p *models.ParamPostList)(data []*models.ApiPostDetail, err error){
	// 去redis中查询id列表
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		zap.L().Error("GetPostList2 GetPostIDsInOrder failed", zap.Error(err))
		return
	}
	// 根据id在mysql中查询帖子详情
	if len(ids) == 0 {
		zap.L().Warn("GetPostList2 ids is empty")
		return
	}
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		zap.L().Error("GetPostList2 GetPostListByIDs failed", zap.Error(err))
		return
	}

	// 提前获取每篇帖子的投票数据
	voteData,err := redis.GetPostVoteData(ids)
	if err != nil {
		zap.L().Error("GetPostList2 GetPostVoteData failed", zap.Error(err))
		return
	}
	 
	data = make([]*models.ApiPostDetail, 0, len(posts))
	// 遍历每个帖子，补全信息
	for idx, post := range posts {
		// 发现：在嵌套结构体中，外层结构体的实例可以直接访问内层结构体的变量，所以下面可以写成：
		// user, err := GetUserById(post.AuthorID)
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("GetUserById() failed", zap.Int64("author_id",post.AuthorID),zap.Error(err))
			return nil, err
		}
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("GetCommunityDetailByID() failed", zap.Error(err))
			return nil, err
		}
		// 包装每一个帖子细节：
		postDetail := &models.ApiPostDetail{
			AuthorName: user.Username,
			VoteNum: voteData[idx],
			CommunityDetail: community,
			Post: post,
		}
		data = append(data, postDetail)
	}
	return

}

// 根据请求参数是否有communityID去相应的进行查询, 合并两个查询帖子的接口
func GetPostListSelect(p *models.ParamPostList)(data []*models.ApiPostDetail, err error){
	if p.CommunityID == 0 {
		// 查所有帖子
		data,err = GetPostList2(p)
	}else{
		// 根据社区查询帖子
		data,err = GetCommunityPostList(p)
	}
	if err != nil {
		zap.L().Error("GetPostListSelect failed", zap.Error(err))
		return nil,err
	}
	return 
}
