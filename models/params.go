package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 投票数据
type ParamVoteData struct {
	//UserID从当前请求中获取当前的用户
	PostID    string `json:"post_id" binding:"required"`               //帖子id
	Direction int8   `json:"direction,string" binding:"oneof= 1 0 -1"` // 赞成票/反对票/取消投票/不投票
}
