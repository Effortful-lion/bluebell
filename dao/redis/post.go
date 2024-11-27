package redis

import (
	"bluebell/models"
	"strconv"
	"time"
	"github.com/go-redis/redis"
)

// TODO: 写一个定时任务，使得失效的key自动删除（举例创建时间一周后的帖子）



func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// 从redis中获取id
	// 默认 时间
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	//确定查询的起始点
	return getIDsFromKey(key, p.Page, p.Size)
}


// TODO:利用投票的相关算法将投票数据(voteNum)转换为分数
// 获取投票数据
func GetPostVoteData(ids []string)(data []int64,err error){
	// 遍历ids的每个id，获取每篇帖子
	// 获取每篇帖子的投票数据
	// for _, id := range ids {
	// 	key := getRedisKey(KeyPostVotedZSetPF+id)
	// 	// 查找 key 中分数为1的元素数量
	// 	v := rdb.ZCount(key, "1", "1").Val()
	// 	data = append(data, v)
	// }

	// 为了提高请求响应速度（假如有100条记录就要查100次，但是使用pipeline只需要一次），使用pipeline，减少RTT（网络往返时间）做一次请求

	pipeline := rdb.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPF+id)
		pipeline.ZCount(key, "1", "1")
	}
	// 得到一个切片
	cmders, err := pipeline.Exec()
	if err != nil{
		return nil, err
	}
	data = make([]int64, 0,len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	
	return
}

// 按社区查询帖子id，根据创建时间或分数//TODO:理解
func GetCommunityPostIDsInOrder(p *models.ParamPostList)([]string, error){
	// 使用zinterstore 把分区的帖子set与帖子分数的zset进行聚合，获取对应的分区的分数，生成一个新的zset
	// 针对这个新的 zset 按之前的逻辑取数据
	// 从redis中获取id
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}

	// 利用缓存key减少zinterstore执行的次数
	key := orderKey + strconv.Itoa(int(p.CommunityID))
	// 社区的key
	ckey := getRedisKey(KeyCommunitySetPF+strconv.Itoa(int(p.CommunityID)))
	if rdb.Exists(key).Val() < 1 {
		// 不存在 交集的key ，需要创建
		pipeline := rdb.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, ckey, orderKey)
		// 拼接 超时时间：相当于给这个临时的查询结果限定时间// TODO:思考为什么？
		pipeline.Expire(key, 60*time.Second)
		_, err := pipeline.Exec()
		if err != nil{
			return nil, err
		}
	}
	// 存在key，根据key查询 
	// 确定查询的起始点
	return getIDsFromKey(key, p.Page, p.Size)
}

// 提取公共部分：根据查询参数确定查询的起始点
func getIDsFromKey(key string,page,size int64)([]string,error){
	// 确定查询的索引起始点
	start := (page - 1) * size
	end := start + size - 1
	// zrevrange 返回的是按照分数从大到小排序的元素
	return rdb.ZRevRange(key,start,end).Result()
}

// 根据id查询redis中的投票数
func GetPostVoteDataByID(id int64)(int64){
	// 拼接key
	zkey := getRedisKey(KeyPostVotedZSetPF+strconv.Itoa(int(id)))
	res := rdb.ZCount(zkey, "1", "1").Val()
	return res
}
