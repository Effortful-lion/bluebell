package mysql

import (
	"bluebell/models"
	"database/sql"
	"strings"

	"github.com/jmoiron/sqlx"
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

// 根据帖子ID查询单个帖子详情
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
	sqlStr := `select post_id, author_id, community_id, title, content, create_time
	from post
	order by create_time
	desc
	limit ?,?`
	// 查到所有的帖子
	err = db.Select(&posts,sqlStr,(page-1)*size,size)  // 偏移量 = (页码 - 1) * 每页显示的条数[因为页码从1开始]
	return
}

// 根据给定id列表查询帖子数据
func GetPostListByIDs(ids []string)(postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	where post_id in (?)
	order by FIND_IN_SET(post_id, ?)`  // FIND_IN_SET可以返回指定变量的位置，配合order by，可以根据find_in_set的返回值，按顺序返回数据
	// 动态填充id,使用args解决
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return
	}

	// 使用db.Rebind解决参数替换 (重新绑定参数为一个sql语句)
	query = db.Rebind(query)

	err = db.Select(&postList, query, args...)
	return
}