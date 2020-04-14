package routers

import (
	"github.com/Biubiubiuuuu/warehouse/server/controllers/adminController"
	"github.com/Biubiubiuuuu/warehouse/server/controllers/goodsStockController"
	"github.com/Biubiubiuuuu/warehouse/server/controllers/goodsTypeController"
	"github.com/Biubiubiuuuu/warehouse/server/controllers/orderController"
	"github.com/Biubiubiuuuu/warehouse/server/controllers/userController"
	"github.com/Biubiubiuuuu/warehouse/server/docs"
	"github.com/Biubiubiuuuu/warehouse/server/helpers/configHelper"
	"github.com/Biubiubiuuuu/warehouse/server/middlewares/adminAuth"
	"github.com/Biubiubiuuuu/warehouse/server/middlewares/cross"
	err "github.com/Biubiubiuuuu/warehouse/server/middlewares/error"
	"github.com/Biubiubiuuuu/warehouse/server/middlewares/jwt"
	"github.com/Biubiubiuuuu/warehouse/server/middlewares/logger"
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
	// 静态资源路径 /static 开头 或者 取自定义配置
	//router.Static(configHelper.Static, "."+configHelper.Static)
	router.Static("/static", "./static")
	//允许跨域请求
	router.Use(cross.Cors())
	//记录日志
	router.Use(logger.Logger())
	InitAdmin(router)
	InitUser(router)
	//gin swaager
	router.GET("/swagger/*any", ginswagger.WrapHandler(swaggerFiles.Handler))
	//404
	router.NoRoute(err.NotFound)
	return router
}

// init admin
func InitAdmin(router *gin.Engine) {
	// 管理员路由分组
	apiAdmin := router.Group("/api/v1/admin")
	// 管理员 get post update delete...
	apiAdmin.POST("login", adminController.LoginAdmin)
	// 管理员 需要登录授权并验证token
	apiAdmin.Use(jwt.JWT())
	{
		apiAdmin.POST("updateAdminPass", adminController.UpdateAdminPass)

		// 用户curd
		apiUser := apiAdmin.Group("/users")
		apiUser.GET("queryUserByLimitOffset", userController.QueryUserByLimitOffset)

		// 商品curd
		apiGoodsType := apiAdmin.Group("/goodsType")
		apiGoodsType.GET("queryByGoodsTypeID", goodsTypeController.QueryByGoodsTypeID)
		apiGoodsType.GET("queryGoodsTypesByLimitOffset", goodsTypeController.QueryGoodsTypesByLimitOffset)
		apiGoodsType.GET("queryAllGoods", goodsTypeController.QueryAllGoods)

		// 库存curd
		apigoodsStock := apiAdmin.Group("/goodsStock")
		apigoodsStock.GET("queryGoodsStocksByLimitOffset", goodsStockController.QueryGoodsStocksByLimitOffset)
		apigoodsStock.GET("queryByGoodsStockID", goodsStockController.QueryByGoodsStockID)

		// 商品订单curd
		apiOrder := apiAdmin.Group("/order")
		apiOrder.GET("queryOrderByLimitOffset", orderController.QueryOrderByLimitOffset)
	}
	// 管理员 需要管理权限Administrator为Y才能操作
	apiAdmin.Use(jwt.JWT(), adminAuth.AdminAuth())
	{
		// 管理员curd
		apiAdmin.GET("queryAdmins", adminController.QueryAdmins)
		apiAdmin.POST("addAdmin", adminController.AddAdmin)
		apiAdmin.DELETE("deleteAdmin", adminController.DeleteAdmin)
		apiAdmin.DELETE("deleteAdmins", adminController.DeleteAdmins)

		// 商品种类curd
		apiGoodsType := apiAdmin.Group("/goodsType")
		apiGoodsType.POST("addGoodsType", goodsTypeController.AddGoodsType)
		apiGoodsType.PUT("updateGoodsType", goodsTypeController.UpdateGoodsType)
		apiGoodsType.DELETE("deleteGoodsType", goodsTypeController.DeleteGoodsType)
		apiGoodsType.DELETE("deleteGoodsTypes", goodsTypeController.DeleteGoodsTypes)

		// 商品库存curd
		apigoodsStock := apiAdmin.Group("/goodsStock")
		apigoodsStock.POST("addGoodsStock", goodsStockController.AddGoodsStock)
		apigoodsStock.PUT("updateGoodsStock", goodsStockController.UpdateGoodsStock)
	}
}

// init user
func InitUser(router *gin.Engine) {
	//用户路由分组
	apiUser := router.Group("/api/v1")
	apiUser.POST("login", userController.LoginUser)
	apiUser.POST("register", userController.RegisterUser)
	apiUser.Use(jwt.JWT())
	{
		apiUser.PUT("updatePass", userController.UpdateUserPass)
		apiUser.POST("addUserInfo", userController.AddUserInfo)
		apiUser.GET("queryUserInfoByUserID", userController.QueryUserInfoByUserID)
		apiUser.GET("queryUserInfoByID", userController.QueryUserInfoByID)
		apiUser.DELETE("deleteUserInfo", userController.DeleteUserInfo)
		apiUser.DELETE("deleteUserInfos", userController.DeleteUserInfos)

		// 订单
		apiUserOrder := apiUser.Group("/user/order")
		apiUserOrder.POST("addOrder", orderController.AddOrder)
		apiUserOrder.GET("queryByGoodsOrderID", orderController.QueryByGoodsOrderID)
		apiUserOrder.GET("queryByOrderUserID", orderController.QueryByOrderUserID)
	}
}
