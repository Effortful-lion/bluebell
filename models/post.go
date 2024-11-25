package models

import "time"

// 有关帖子的请求参数 和 响应参数 定义

// 内存对齐 ：使得结构体的内存占用尽量小
type Post struct {
	Title string `json:"title" db:"title" binding:"required"`
	Content string `json:"content" db:"content" binding:"required"`
    ID int64 `json:"id" db:"post_id"`
	AuthorID int64 `json:"author_id" db:"author_id"`
	CommunityID int64 `json:"community_id" db:"community_id" binding:"required"`
	Status int32 `json:"status" db:"status"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
} 

// 这个结构体是为了满足：当接口得到数据后发现：得到的作者是作者id，社区是社区id，需要将id查询到作者和社区信息
// ApiPostDetail 帖子接口的结构体
type ApiPostDetail struct {
	AuthorName string `json:"author_name"`  // 发帖作者的名称
	*Post		// 嵌入Post帖子结构体
	*CommunityDetail  // 嵌入Community社区结构体
}