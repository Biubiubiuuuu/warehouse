package models

import (
	"errors"

	"github.com/Biubiubiuuuu/warehouse/server/dbs/mysql"
)

type UserInfo struct {
	Model
	Provice     string `json:"provice"`       // 省
	City        string `json:"city"`          // 城市
	ShopAddress string `json:"shop_address1"` // 门店详细地址
	UserID      int64  `json:"user_id"`       // 用户ID
}

// 添加用户地址信息
func (u *UserInfo) AddUserInfo() error {
	db := mysql.GetDB()
	return db.Create(&u).Error
}

// 修改用户地址信息
func (u *UserInfo) UpdateUserInfo(args map[string]interface{}) error {
	db := mysql.GetDB()
	return db.Model(&u).Updates(args).Error
}

// 查询用户地址信息详情
func (u *UserInfo) QueryUserInfoByID() error {
	db := mysql.GetDB()
	return db.First(&u, u.ID).Error
}

// 删除用户地址信息（支持批量）
func (u *UserInfo) DeleteUserInfo(ids []int64) error {
	db := mysql.GetDB()
	tx := db.Begin()
	for _, id := range ids {
		if id == 0 {
			return errors.New("id is not 0")
		}
		u.ID = id
		if err := tx.Delete(&u).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}
