package models

import (
	"github.com/Biubiubiuuuu/warehouse/server/dbs/mysql"
	"github.com/google/uuid"
)

type User struct {
	Model
	Tel      string    `gorm:"not null;unique;size:50" json:"tel"` // 手机号
	Name     string    `json:"name"`                               // 姓名
	Password string    `json:"-"`                                  // 密码
	UUID     uuid.UUID `json:"uuid"`
	Token    string    `gorm:"size:255" json:"token"` // 授权令牌
	IP       string    `gorm:"size:30" json:"ip"`     //登录IP
}

// 登录
func (u *User) LoginUser() error {
	db := mysql.GetDB()
	return db.Where("tel = ? AND password = ?", u.Tel, u.Password).First(&u).Error
}

// 注册
func (u *User) RegisterUser() error {
	db := mysql.GetDB()
	return db.Create(&u).Error
}

// 查询账号是否存在并返回账号信息
func (u *User) QueryUser() bool {
	db := mysql.GetDB()
	if err := db.Where("tel = ? OR (token = ? AND token IS NOT NULL)", u.Tel, u.Token).First(&u).Error; err != nil {
		return false
	}
	return true
}

// 修改账号信息
func (u *User) UpdataUser(args map[string]interface{}) error {
	db := mysql.GetDB()
	return db.Model(&u).Update(args).Error
}

// 查询用户（分页查询）
func QueryUserByLimitOffset(pageSize int, page int) (users []User) {
	db := mysql.GetDB()
	db.Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&users)
	return
}

// 总记录数
func QueryUserCount() (count int) {
	db := mysql.GetDB()
	db.Model(&User{}).Count(&count)
	return count
}
