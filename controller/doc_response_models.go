package controller

import "bluebell/models"

// swagger文档的参数

// GET
// _ResponsePostList 帖子列表接口响应数据
type _ResponsePostList struct {
	Code    ResCode                 `json:"code"`    // 业务响应状态码
	Message string                  `json:"message"` // 提示信息
	Data    []*models.ApiPostDetail `json:"data"`    // 数据
}

type _ResponsePostDetail struct {
	Code    ResCode              `json:"code"`    // 业务响应状态码
	Message string               `json:"message"` // 提示信息
	Data    *models.ApiPostDetail `json:"data"`    // 数据
}

// TODO:帖子创建的swagger参数
type _ResponsePostCreate struct {
	Code    ResCode              `json:"code"`    // 业务响应状态码
	Message string               `json:"message"` // 提示信息
	Data    any 				 `json:"data"`    // 数据
}

type _ResponseCommunityList struct {
	Code ResCode `json:"code"`
	Message string `json:"message"`
	Data []*models.Community `json:"data"`
}

type _ResponseCommunityDetail struct {
	Code ResCode `json:"code"`
	Message string `json:"message"`
	Data *models.CommunityDetail `json:"data"`
}

// POST
type _RequestCreatePost struct {
	Title string `json:"title"`	// 帖子标题
	Content string `json:"content"`	// 帖子内容
	CommunityID int64 `json:"community_id"`	// 帖子所属社区ID
}

type _RequestVotePost struct {
	PostID string `json:"post_id"`
	Direction int8 `json:"direction,string""`
}