package logic

// TODO:完善思路
// 投票功能：
// 1. 用户投票后，需要把投票数据保存到redis中
// 2.

// 基于用户投票的相关算法：http://www.ruanyifeng.com/blog/algorithm/

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"strconv"

	"go.uber.org/zap"
)

// 投票思路

// 投票的几种情况:
/*
direction = 1:
	1. 之前没有投过票，直接投赞成票
	2. 之前投反对票，改为投赞成票，并且更新分数
direction = -1:
	1. 之前没有投过票，直接投反对票
	2. 之前投赞成票，改为投反对票，并且更新分数
direction = 0:
	1. 之前投过赞成票，取消投票，并且更新分数
	2. 之前投反对票，取消投票，并且更新分数
	3. 之前没有投过票，返回错误

投票的限制：
	每个帖子从发表开始，可以投票的时间是一周：
		1. 到期后，将redis中保存的赞成票数和反对票数保存到mysql数据库中
		2. 删除redis中的投票的 KeyPostVotedZSetPF
*/

// 投票
func VoteForPost(userID int64, p *models.ParamVoteData) (err error) {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))
		err = redis.VoteForPost(strconv.Itoa(int(userID)), p)
	return err
}

