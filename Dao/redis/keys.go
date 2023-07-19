package redis

// redis key
// redis key 尽量使用命名空间的方式进行区分不同的key
const (
	Prefix            = "bluebell:"
	KeyPostTimeZSet   = "post:time " //帖子及发帖时间
	KeyPostScoreZSet  = "post:score" //帖子及投票的分数
	KeyPostVotedZSetP = "post:voted" // KeyPostVotedZSetP 需要的参数是post_id
	KeyCommunitySetP  = "community:" // KeyCommunitySetP 需要的参数是community_id
)

func getRedisKey(Key string) string {
	return Prefix + Key
}
