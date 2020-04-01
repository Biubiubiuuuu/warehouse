package models

import (
	"time"

	"github.com/Biubiubiuuuu/warehouse/server/dbs/mysql"
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
	db.LogMode(true)
	// init username:Admin,password:123456
	//if CheckAdminUsernameExist("Admin") != nil {
	//	adminM := Admin{Username: "Admin", Password: MD5Helper.EncryptMD5To32Bit("123456"), Administrator: "Y"}
	//AddAdmin(adminM)
	//}
}
