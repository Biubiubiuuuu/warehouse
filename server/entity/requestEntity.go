package entity

// 登录
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 注册
type Register struct {
	OnlineUsername string `json:"online_username"` // 当前登录用户名
	Username       string `json:"username"`        // 用户名
	Password       string `json:"password"`        // 密码
	Administrator  string `json:"administrator"`   // 超级管理员 Y | N
}

// 修改密码
type Register struct {
	OldPassword   string `json:"oldPassword"`   // 当前登录用户名
	Username      string `json:"username"`      // 用户名
	Password      string `json:"password"`      // 密码
	Administrator string `json:"administrator"` // 超级管理员 Y | N
}
