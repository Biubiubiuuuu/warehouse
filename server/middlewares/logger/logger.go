package logger

import (
	"time"

	"github.com/Biubiubiuuuu/warehouse/server/helpers/configHelper"
	"github.com/Biubiubiuuuu/warehouse/server/helpers/fileHelper"
	"github.com/gin-gonic/gin"
)

// 文件路径，以天为目录
var path = configHelper.LogDir + time.Now().Format("20060102") + "/"
var logChannel = make(chan string, 100)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}

func handleLogChannel() {
	if !fileHelper.IsExist(path) {
		fileHelper.CreateDir(path)
	}

}
