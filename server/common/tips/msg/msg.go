package msg

import (
	c "github.com/Biubiubiuuuu/warehouse/server/common/tips/code"
)

var MsgFlags = map[int]string{
	c.SUCCESS:                       "SUCCESS",
	c.ERROR:                         "ERROR",
	c.NOTFOUND:                      "404 Not Found",
	c.TOKEN_ERROR:                   "token生成失败",
	c.TOKEN_AUTH_ERROR:              "token错误",
	c.TOKEN_TIMEOUT:                 "token已过期",
	c.AUTH_NOT_BEARER:               "Query not 'token' param OR header Authorization has not Bearer token",
	c.LOGIN_SUCCESS:                 "登录成功",
	c.LOGIN_ERROR:                   "登录失败",
	c.ADD_SUCCESS:                   "添加成功",
	c.ADD_ERROR:                     "添加失败",
	c.QUERY_SUCCESS:                 "查询成功",
	c.QUERY_ERROR:                   "查询失败",
	c.DELETE_SUCCESS:                "删除成功",
	c.DELETE_ERROR:                  "删除失败",
	c.UPDATE_SUCCESS:                "修改成功",
	c.UPDATE_ERROR:                  "修改失败",
	c.NOTMORE:                       "没有更多了~",
	c.USERNAME_EXIST:                "用户名已存在",
	c.USERNAME_OR_PASSWORD_ERROR:    "用户名或密码错误",
	c.PARAMETER_ERROR:               "参数错误",
	c.USERNAME_OR_PASSWORD_NOT_NULL: "用户名或密码不能为空",
	c.NOT_ADMINISTRATOR:             "没有权限操作",
	c.EXIST:                         "已存在",
	c.NOTNULL:                       "参数不能为空",
	c.STOCK_EXIST:                   "该商品已创建库存，请直接修改库存信息",
	c.OLD_ERROR:                     "旧密码错误",
	c.ORDER_SUCCESS:                 "下单成功",
	c.ORDER_ERROR:                   "下单失败",
}

// get return msg
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[c.ERROR]
}
