package mysql

import (
	"bluebell/models"
	"database/sql"
	"go.uber.org/zap"
)

// 创建帖子
func CreatePost(p *models.Post)(err error){
    sqlStr := "insert into post(post_id, author_id, community_id, title, content) values (?,?,?,?,?)"
    _, err = db.Exec(sqlStr, p.ID, p.AuthorID, p.CommunityID, p.Title, p.Content)
    if err != nil {
        zap.L().Error("创建帖子失败", zap.Error(err))
        return
    }
    return
}

// 根据帖子ID查询帖子详情
func GetPostById(id int64) (data *models.Post, err error) {
	data = new(models.Post)
	sqlStr := "select post_id, author_id, community_id, title, content, create_time from post where post_id = ?"
	err = db.Get(data, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("该帖子不存在",zap.Error(err))
			err = ErrorInvalidID
			return
		}
		zap.L().Error("查询数据库失败", zap.Error(err))
		return
	}
	return data,nil
} 

// 获取帖子列表
func GetPostList(page,size int64)(posts []*models.ApiPostDetail, err error){
	// 帖子列表
	posts = make([]*models.ApiPostDetail, 0,2)
	// 这里使用 limit 查询前4条，为后面的分页要求做准备
	sqlStr := "select post_id, author_id, community_id, title, content, create_time from post limit ?,?"
	// 查到所有的帖子
	err = db.Select(&posts,sqlStr,(page-1)*size,size)  // 偏移量 = (页码 - 1) * 每页显示的条数[因为页码从1开始]
	return
}
