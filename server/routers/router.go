package routers

import (
	"github.com/Biubiubiuuuu/warehouse/server/docs"
	"github.com/Biubiubiuuuu/warehouse/server/helpers/configHelper"
	"github.com/Biubiubiuuuu/warehouse/server/middlewares/cross"
	"github.com/gin-gonic/gin"
	ginswagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// init router
func Init() *gin.Engine {
	//swagger接口文档
	docs.SwaggerInfo.Title = "开放接口"
	docs.SwaggerInfo.Description = ""
	docs.SwaggerInfo.Version = configHelper.Version
	//设置模式，设置模式要放在调用Default()函数之前
	if configHelper.RunMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	//允许跨域请求
	router.Use(cross.Cors())
	InitAdmin(router)
	InitUser(router)
	//gin swaager
	router.GET("/swagger/*any", ginswagger.WrapHandler(swaggerFiles.Handler))
	//404
	//router.NoRoute(err.NotFound)
	return router
}

// init admin
func InitAdmin(router *gin.Engine) {
	//管理员路由分组
	//apiAdmin := router.Group("/api/v1/admin")
}

// init user
func InitUser(router *gin.Engine) {
	//用户路由分组
	//api := router.Group("/api/v1")
}
