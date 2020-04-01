package main

import (
	"github.com/Biubiubiuuuu/warehouse/server/helpers/configHelper"
	"github.com/Biubiubiuuuu/warehouse/server/models"
	"github.com/Biubiubiuuuu/warehouse/server/routers"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// DB初始化、数据模型自动迁移
	models.Init()
	// 初始化
	router := routers.Init()
	router.Run(configHelper.HTTPPort)
}
