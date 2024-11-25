package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 帖子相关

// 创建帖子
func CreatePostHandler(c *gin.Context) {
	// 1. 获取参数及参数校验
	p := new (models.Post)
	if err := c.ShouldBindJSON(p);err != nil {
		// 参数获取失败
		zap.L().Error("CreatePostHandler with invalid param",zap.Error(err))
		ResponseError(c,CodeInvalidParam)
		return
	}
	userID,err := getCurrentUserId(c)
	fmt.Println(userID)
	if err != nil {
		// 获取不到当前用户id，就重新登录获取
		ResponseError(c,CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	// 2. 调用logic进行业务处理(创建帖子)
	if err:=logic.CreatePost(p);err != nil{
		zap.L().Error("logic.CreatePost()创建帖子失败",zap.Error(err))
		ResponseError(c,CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c,nil)
}

func GetPostDetailHandler(c *gin.Context){
	// 1.获取参数和校验
	idStr := c.Param("id")
	id,err := strconv.ParseInt(idStr,10,64)
	if err != nil{
		zap.L().Error("获取参数失败",zap.Error(err))
		ResponseError(c,CodeInvalidParam)
		return
	}
	// 2.查询数据库
	data,err := logic.GetPostDetail(id)
	// 3.返回响应
	if err != nil {
		zap.L().Error("获取详情失败",zap.Error(err))
		ResponseError(c,CodeServerBusy)
		return
	}
	ResponseSuccess(c,data)
}

// 获取帖子列表: 分页查询（另外的要求）
func GetPostListHandler(c *gin.Context) {
	page,size := getPageInfo(c)
	// 获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("获取帖子列表失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回数据
	ResponseSuccess(c, data)
}