package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"strconv"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 帖子相关

// TODO:这个post无法显示出来
// 创建帖子接口
// @Summary 创建帖子接口
// @Description 创建帖子接口 (base or vip)
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param req body models.Post true 接收创建帖子接口参数
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /post [post]
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

// GetPostDetailHandler 获取帖子详情
// @Summary 单个帖子详情接口
// @Description 单个帖子详情接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path int true "帖子id"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostDetail
// @Router /post/{id} [get]
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

// // 获取帖子列表: 分页查询（另外的要求）【没用了，可以作为学习分页查询】
// func GetPostListHandler(c *gin.Context) {
// 	page,size := getPageInfo(c)
// 	// 获取数据
// 	data, err := logic.GetPostList(page, size)
// 	if err != nil {
// 		zap.L().Error("获取帖子列表失败", zap.Error(err))
// 		ResponseError(c, CodeServerBusy)
// 		return
// 	}
// 	// 返回数据
// 	ResponseSuccess(c, data)
// }

// 根据前端传来的参数动态的获取帖子列表
// 按照创建时间排序 或者 按照分数排序

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /post2 [get]
func GetPostListHandler2(c *gin.Context){
	// 1. 获取参数: 初始化结构体时，指定默认参数
	p := &models.ParamPostList{
		Page: models.DefaultPage,
		Size: models.DefaultSize,
		Order: models.OrderTime,
	}

	if err := c.ShouldBindQuery(p); err != nil{
		zap.L().Error("GetPostListHandler2() with invalid params")
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2. 获取数据
	data, err := logic.GetPostListSelect(p)
	if err != nil {
		zap.L().Error("获取帖子列表失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, data)
}
