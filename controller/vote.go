package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// TODO: post类的接口都无法显示
// @Summary 帖子投票接口
// @Description 帖子投票接口
// @Tags 帖子投票接口
// @Accept application/json
// @Produce application/json
// @Param req body models.ParamVoteData true 接受帖子投票接口参数
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /vote [post]
// 投票
func PostVoteHandler(c *gin.Context) {
	// 获取参数和参数校验：用户 + 帖子 + 投票选项
	p := new(models.ParamVoteData)
	// 未获得参数:因为在请求结构体中，本身是string的变量在json块中又被声明为string了
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok{
			ResponseError(c, CodeInvalidParam)
			return
		}
		// 是validator.ValidationErrors类型：validator校验器的错误信息是英文，转为中文并去掉结构体标识
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	userID, err := getCurrentUserId(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	
	err = logic.VoteForPost(userID,p)
	if err != nil {
		errData := fmt.Sprintf("vote for post failed, err:%v", err) 
		ResponseErrorWithMsg(c,CodeServerBusy,errData)
		return
	}
	ResponseSuccess(c, nil)
}