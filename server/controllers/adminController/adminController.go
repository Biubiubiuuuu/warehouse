package adminController

import (
	"net/http"
	"strconv"

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

// @Summary 删除管理员
// @tags 管理员
// @Accept  application/json
// @Produce  json
// @Param id query string true "id"
// @Param username query string true "username"
// @Success 200 {string} json "{"status": true, "message": "SUCCESS"} {"status": false, "message": "用户名或密码不能为空"}  {"status": false, "message": "ERROR"}"
// @Router /api/v1/admin/deleteAdmin [DELETE]
// @Security ApiKeyAuth
func DeleteAdmin(c *gin.Context) {
	responseData := entity.ResponseData{}
	request := entity.DeleteIds{}
	id, _ := strconv.ParseInt(c.DefaultQuery("id", "0"), 10, 64)
	username := c.DefaultQuery("username", "")
	ids := append(request.Ids, id)
	responseData = adminService.DeleteAdmin(username, ids)
	c.JSON(http.StatusOK, responseData)
}

// @Summary 批量删除管理员
// @tags 管理员
// @Accept  application/json
// @Produce  json
// @Param username query string true "username"
// @Param body body entity.DeleteIds true "body"
// @Success 200 {string} json "{"status": true, "message": "SUCCESS"} {"status": false, "message": "用户名或密码不能为空"}  {"status": false, "message": "ERROR"}"
// @Router /api/v1/admin/DeleteAdmins [DELETE]
// @Security ApiKeyAuth
func DeleteAdmins(c *gin.Context) {
	responseData := entity.ResponseData{}
	request := entity.DeleteIds{}
	username := c.DefaultQuery("username", "")
	// 参数校验
	if c.ShouldBindJSON(&request) != nil {
		responseData.Message = msg.GetMsg(tcode.PARAMETER_ERROR)
	} else {
		responseData = adminService.DeleteAdmin(username, request.Ids)
	}
	c.JSON(http.StatusOK, responseData)
}

// @Summary 修改管理员密码
// @tags 管理员
// @Accept  application/json
// @Produce  json
// @Param body body entity.UpdatePass true "body"
// @Success 200 {string} json "{"status": true, "message": "SUCCESS"} {"status": false, "message": "用户名或密码不能为空"}  {"status": false, "message": "ERROR"}"
// @Router /api/v1/admin/updateAdminPass [UPDATE]
// @Security ApiKeyAuth
func UpdateAdminPass(c *gin.Context) {
	responseData := entity.ResponseData{}
	request := entity.UpdatePass{}
	// 参数校验
	if c.ShouldBindJSON(&request) != nil {
		responseData.Message = msg.GetMsg(tcode.PARAMETER_ERROR)
	} else {
		responseData = adminService.UpdateAdminPass(request)
	}
	c.JSON(http.StatusOK, responseData)
}

// @Summary 分页查询用户(默认前100条) 并返回总记录数
// @tags 管理员
// @Accept application/x-www-form-urlencoded
// @Produce  json
// @Param pageSize query string false "页大小"
// @Param page query string false "页数"
// @Success 200 {string} json "{"status": true, "message": "SUCCESS", "data":{}} {"status": false, "message": "ERROR"}"
// @Router /api/v1/admin/queryAdmins [GET]
// @Security ApiKeyAuth
func QueryAdmins(c *gin.Context) {
	responseData := entity.ResponseData{}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "100"))
	responseData = adminService.QueryByLimitOffset(pageSize, page)
	c.JSON(http.StatusOK, responseData)
}
