package models

import (
	"time"

	"github.com/Biubiubiuuuu/warehouse/server/dbs/mysql"
	"github.com/Biubiubiuuuu/warehouse/server/helpers/MD5Helper"
	"github.com/google/uuid"
)

// base model by gorm.Model
type Model struct {
	ID        int64      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	UUID      uuid.UUID  `json:"uuid"`
}

// init mysql DB and auto migrate models
func Init() {
	mysql.DB.Init()
	db := mysql.GetDB()
	db.AutoMigrate(&Admin{})
	// 添加默认管理员 username:Admin,password:123456
	a := Admin{Username: "admin", Password: MD5Helper.EncryptMD5To32Bit("123456"), Administrator: "Y"}
	if !a.QueryByUsername() {
		a.Register()
	}
}
