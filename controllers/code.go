package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidPram
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeNeedLogin
	CodeInvalidToken
	CodeVoteTimeExpire
	CodeVoteRepeated
)

var codemsgmap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidPram:     "请求参数错误",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",
	CodeNeedLogin:       "需要登入",
	CodeInvalidToken:    "无效的token",
	CodeVoteTimeExpire:  "投票时间已过",
	CodeVoteRepeated:    "不允许重复投票",
}

func (c ResCode) Msg() string {
	msg, ok := codemsgmap[c]
	if !ok {
		msg = codemsgmap[CodeServerBusy]
	}
	return msg
}
