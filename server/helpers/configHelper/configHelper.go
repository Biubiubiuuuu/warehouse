package configHelper

import (
	"log"

	"github.com/go-ini/ini"
)

var (
	Cfg                *ini.File
	RunMode            string
	HTTPPort           string
	JwtSecret          string
	DBMysqlType        string
	DBMysqlName        string
	DBMysqlUser        string
	DBMysqlPassword    string
	DBMysqlHost        string
	DBMysqlTablePrefix string
	JwtName            string
	Version            string
	LogDir             string
	LogName            string
	LogLevel           string
	MaxIdleConns       string
	MaxOpenConns       string
)

// init config
func init() {
	var err error
	Cfg, err = ini.Load("./config/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'config/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
	LoadMysql()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	HTTPPort = sec.Key("HTTP_PORT").MustString("8060")
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	JwtName = sec.Key("JWT_NAME").MustString("token")
	Version = sec.Key("VERSION").MustString("1.0")
	LogDir = sec.Key("LOGDIR").MustString("")
	LogName = sec.Key("LOGFILENAME").MustString("lottery.log")
	LogLevel = sec.Key("LOGLEVEL").MustString("info")
}

func LoadMysql() {
	sec, err := Cfg.GetSection("mysql")
	if err != nil {
		log.Fatalf("Fail to get section 'mysql': %v", err)
	}
	DBMysqlType = sec.Key("TYPE").MustString("mysql")
	DBMysqlName = sec.Key("NAME").MustString("test")
	DBMysqlUser = sec.Key("USER").MustString("root")
	DBMysqlPassword = sec.Key("PASSWORD").MustString("")
	DBMysqlHost = sec.Key("HOST").MustString("127.0.0.1:3306")
	DBMysqlTablePrefix = sec.Key("TABLE_PREFIX").MustString("")
	MaxIdleConns = sec.Key("MAXIDLECONNS").MustString("")
	MaxOpenConns = sec.Key("MAXOPENCONNS").MustString("")
}
