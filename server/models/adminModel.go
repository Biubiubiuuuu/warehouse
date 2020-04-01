package models

// Admin model
type Admin struct {
	Model
	Username      string `gorm:"not null;unique;size:255" json:"username"`          //用户名
	Password      string `gorm:"not null;size:255" json:"-"`                        //密码
	IP            string `gorm:"size:30" json:"ip"`                                 //登录IP
	Token         string `gorm:"size:255" json:"token"`                             // 授权令牌
	Administrator string `gorm:"not null;default:'N';size:10" json:"administrator"` // 超级管理员 Y | N
}
