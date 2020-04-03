package msg

import (
	c "github.com/Biubiubiuuuu/warehouse/server/common/tips/code"
)

var MsgFlags = map[int]string{
	c.SUCCESS:          "SUCCESS",
	c.ERROR:            "ERROR",
	c.NOTFOUND:         "404 Not Found",
	c.TOKEN_ERROR:      "token生成失败",
	c.TOKEN_AUTH_ERROR: "token错误",
	c.TOKEN_TIMEOUT:    "token已过期",
	c.AUTH_NOT_BEARER:  "Query not 'token' param OR header Authorization has not Bearer token",
	c.USERNAME_EXIST:   "用户名已存在",
	c.PARAMETER_ERROR:  "参数错误",
	c.NOTNULL:          "用户名或密码不能为空",
	c.NOTADMINISTRATOR: "没有权限操作",
	c.NOTMORE:          "没有更多了~",
}

// get return msg
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[c.ERROR]
}
