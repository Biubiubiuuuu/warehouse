package adminController

import (
	"net/http"

	tcode "github.com/Biubiubiuuuu/warehouse/server/common/tips/code"
	"github.com/Biubiubiuuuu/warehouse/server/common/tips/msg"
	"github.com/Biubiubiuuuu/warehouse/server/entity"
	"github.com/Biubiubiuuuu/warehouse/server/services/adminService"
	"github.com/gin-gonic/gin"
)

// @Summary 管理员登录
// @tags 管理员
// @Accept  application/json
// @Produce  json
// @Param body body entity.Login true "body"
// @Success 200 {string} json "{"status": true, "message": "SUCCESS"} {"status": false, "message": "用户名或密码不能为空"}  {"status": false, "message": "ERROR"}"
// @Router /api/v1/admin/Login [POST]
func LoginAdmin(c *gin.Context) {
	request := entity.Login{}
	responseData := entity.ResponseData{}
	if c.ShouldBindJSON(&request) != nil {
		responseData.Message = msg.GetMsg(tcode.PARAMETER_ERROR)
	} else {
		responseData = adminService.LoginAdmin(request, c.ClientIP())
	}
	c.JSON(http.StatusOK, responseData)
}

// @Summary 添加管理员
// @tags 管理员
// @Accept  application/json
// @Produce  json
// @Param body body entity.Register true "body"
// @Success 200 {string} json "{"status": true, "message": "SUCCESS"} {"status": false, "message": "用户名或密码不能为空"}  {"status": false, "message": "ERROR"}"
// @Router /api/v1/admin/addAdmin [POST]
// @Security ApiKeyAuth
func AddAdmin(c *gin.Context) {
	responseData := entity.ResponseData{}
	request := entity.Register{}
	// 参数校验
	if c.ShouldBindJSON(&request) != nil {
		responseData.Message = msg.GetMsg(tcode.PARAMETER_ERROR)
	} else {
		responseData = adminService.AddAdmin(request)
	}
	c.JSON(http.StatusOK, responseData)
}
