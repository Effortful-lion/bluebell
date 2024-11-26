package redis

import (
	"bluebell/models"
	"errors"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated   = errors.New("不允许重复投票")
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 // 每一票值多少分
	// 这里是为了做到：每投一票，就可以为这个帖子“续命”432秒，200票就是一天：86400秒  （时间戳增大）
)

// 创建帖子的时候，将帖子的信息缓存在redis中
func CreatePost(postID, communityID int64) error {
	pipeline := rdb.TxPipeline()
	// 帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	// 更新：把帖子id加到社区的set
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityID)))
	pipeline.SAdd(cKey, postID)
	_, err := pipeline.Exec()
	return err
}

func VoteForPost(userID string, p *models.ParamVoteData)(err error) {
	// 0. 判断投票的限制(帖子可投？(时间限制一周))
	// 1. 判断用户是否投过票(用户可投？)
	if err = CheckAuthToVote(userID, p.PostID,p.Direction);err != nil{
		// 不可以投票：帖子过期/用户已经投过了
		if err == ErrVoteTimeExpire{
			// 帖子过期
			return 
		}else{
			// 用户已经投过票
			return 
		} 
	}
	// 2. 判断投票类型
	// 3. 更新帖子的分数
	// 4. 更新帖子投票的用户
	UpdateVote(userID, p.PostID,p.Direction)
	return 
}

func CheckAuthToVote(userID string, postID string,direction int8) error{
	// 检查帖子的发帖时间到目前是否超过一周
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	// 检查用户是否已经投过票
	// 先查当前用户给当前帖子的投票记录:如果这一次投票的值和之前保存的值一致，就提示不允许重复投票 
	ov := rdb.ZScore(getRedisKey(KeyPostVotedZSetPF+postID), userID).Val()
	if ov == float64(direction) {
		return ErrVoteRepeated
	}
	// 返回nil，表示可以投票
	return nil
}

func UpdateVote(userID string,postID string, direction int8) error{
	// 获取帖子投票结果的key
	postKey := getRedisKey(KeyPostScoreZSet)
	old_score := rdb.ZScore(postKey,postID).Val()
	// 根据投票类型，更新帖子的分数
	// 用户之前的投票记录
	ov := rdb.ZScore(getRedisKey(KeyPostVotedZSetPF+postID), userID).Val()
	var newScore float64
	// 现在相对于之前来说，是做增加操作还是减少操作
	diff := float64(direction) - ov
	newScore = old_score + diff * 432
	// 更新分数
	rdb.ZAdd(getRedisKey(KeyPostScoreZSet),redis.Z{
		Score:newScore,
		Member:postID,
	})
	// 更新帖子投票记录
	rdb.ZAdd(getRedisKey(KeyPostVotedZSetPF+postID),redis.Z{
		Score:float64(direction),
		Member:userID,
	})
	return nil
}