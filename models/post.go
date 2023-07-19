package models

import "time"

// Post 内存对齐概念 就是机组存储有关
type Post struct {
	ID          int64     `json:"id,string" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binging:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binging:"required"`
	Content     string    `json:"content" db:"content" binging:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

// ApiPostDetail 帖子详情接口的结构体

type ApiPostDetail struct {
	AuthorName       string `json:"author_name"`
	VoteNum          int64  `json:"vote_num"`
	*Post                   //嵌入帖子结构体
	*CommunityDetail        //嵌入社区信息
}

// ParamPostList 获取帖子列表参数
type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"` //可以为空
	Page        int64  `json:"page" form:"page"`
	Size        int64  `json:"size" form:"size"`
	Order       string `json:"order" form:"order"`
}

type ParamCommunityPostList struct {
	*ParamPostList
	CommunityID int64 `json:"community_id" form:"community_id"`
}
