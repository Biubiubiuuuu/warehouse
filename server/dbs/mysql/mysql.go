package mysql

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/Biubiubiuuuu/warehouse/server/helpers/configHelper"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlDatabase struct {
	Mysql *gorm.DB
}

var DB *MysqlDatabase
var once sync.Once

// init
func (db *MysqlDatabase) Init() {
	once.Do(func() {
		DB = &MysqlDatabase{
			Mysql: InitDB(),
		}
	})
}

// init and open mysql DB
func InitDB() *gorm.DB {
	var (
		dbType, dbName, user, password, host, tablePrefix string
	)
	dbType = configHelper.DBMysqlType
	dbName = configHelper.DBMysqlName
	user = configHelper.DBMysqlUser
	password = configHelper.DBMysqlPassword
	host = configHelper.DBMysqlHost
	tablePrefix = configHelper.DBMysqlTablePrefix

	connect := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName)
	db, err := gorm.Open(dbType, connect)
	if err != nil {
		log.Fatal(err)
	}
	// set Singular Table
	db.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return fmt.Sprintf("%v%v", tablePrefix, defaultTableName)
	}
	SetDBConns(db)
	return db
}

// Set Max Open Conns
// Set MaxIdle Conns
func SetDBConns(db *gorm.DB) {
	MaxOpenConns, _ := strconv.Atoi(configHelper.MaxOpenConns)
	MaxIdleConns, _ := strconv.Atoi(configHelper.MaxIdleConns)
	db.DB().SetMaxOpenConns(MaxOpenConns)
	db.DB().SetMaxIdleConns(MaxIdleConns)
}

// get mysql DB conn
func GetDB() *gorm.DB {
	return InitDB()
}
