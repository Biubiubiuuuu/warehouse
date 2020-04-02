package models

import (
	"errors"

	"github.com/Biubiubiuuuu/warehouse/server/common/tips/code"
	"github.com/Biubiubiuuuu/warehouse/server/common/tips/msg"
	"github.com/Biubiubiuuuu/warehouse/server/dbs/mysql"
)

// Admin model
type Admin struct {
	Model
	Username      string `gorm:"not null;unique;size:255" json:"username"`          //用户名
	Password      string `gorm:"not null;size:255" json:"-"`                        //密码
	IP            string `gorm:"size:30" json:"ip"`                                 //登录IP
	Token         string `gorm:"size:255" json:"token"`                             // 授权令牌
	Administrator string `gorm:"not null;default:'N';size:10" json:"administrator"` // 超级管理员 Y | N
}

// 登录
func (a *Admin) Login() error {
	db := mysql.GetDB()
	if err := db.Where("username = ? AND password = ?", a.Username, a.Password).First(&a).Error; err != nil {
		return err
	}
	return nil
}

// 注册
func (a *Admin) Register() error {
	db := mysql.GetDB()
	if db.NewRecord(a.Username) {
		return db.Create(&a).Error
	}
	return errors.New(msg.GetMsg(code.USERNAME_EXIST))
}

// 修改账号信息
func (a *Admin) UpdataInfo(arsg map[string]interface{}) error {
	db := mysql.GetDB()
	if err := db.Model(&a).Update(arsg).Error; err != nil {
		return err
	}
	return nil
}

// 查询账号
func (a *Admin) QueryByUsername() bool {
	db := mysql.GetDB()
	if err := db.Where("username = ?", a.Username).First(&a).Error; err != nil {
		return false
	}
	return true
}

// 检查用户权限
func (a *Admin) CheckAdministrator() bool {
	db := mysql.GetDB()
	if err := db.Where("username = ? AND administrator = Y", a.Username).First(&a).Error; err != nil {
		return false
	}
	return true
}

// 删除（可批量）
func (a *Admin) Delete(ids []int64) error {
	db := mysql.GetDB()
	tx := db.Begin()
	for _, id := range ids {
		if id == 0 {
			return errors.New("id is not 0")
		}
		a.ID = id
		if err := tx.Delete(&a).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

// 查询用户（分页查询）
func QueryByLimitOffset(pageSize int, page int) (admins []Admin) {
	db := mysql.GetDB()
	db.Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&admins)
	return
}

// 总记录数
func QueryCount() (count int) {
	db := mysql.GetDB()
	db.Model(&Admin{}).Count(&count)
	return count
}
