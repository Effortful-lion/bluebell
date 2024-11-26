package redis

// 设计redis的key


// redis key 注意使用命名空间的方式，方便查询和拆分
var (
	KeyPrefix = "bluebell:"				// key的声明前缀
	KeyPostTimeZSet = "post:time"  		// 按发帖时间排序的投票结果（ post_id : time）  redis.time = mysql.create_time
	KeyPostScoreZSet = "post:score"		// 按投票分数排序的投票结果,参数是帖子id（ post_id : score ）
	KeyPostVotedZSetPF = "post:voted:"		// 记录用户投票状态和投票的类型，1表示赞成，-1表示反对,参数是post_id
	KeyCommunitySetPF = "community:" // set;保存每个分区下帖子的id
	// TODO:这里记录投票状态的设计，不理解
)

// 返回加上前缀的rediskey
func getRedisKey(key string) string {
	return KeyPrefix + key
}