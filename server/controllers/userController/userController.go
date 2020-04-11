package userController

import (
	"net/http"
	"strconv"
	"strings"

	tcode "github.com/Biubiubiuuuu/warehouse/server/common/tips/code"
	"github.com/Biubiubiuuuu/warehouse/server/common/tips/msg"
	"github.com/Biubiubiuuuu/warehouse/server/entity"
	"github.com/Biubiubiuuuu/warehouse/server/services/userService"
	"github.com/gin-gonic/gin"
)

// @Summary 用户登录
// @tags 用户
// @Accept  application/json
// @Produce  json
// @Param body body entity.UserLogin true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/login [POST]
func LoginUser(c *gin.Context) {
	request := entity.UserLogin{}
	responseData := entity.ResponseData{}
	if c.ShouldBindJSON(&request) != nil {
		responseData.Message = msg.GetMsg(tcode.PARAMETER_ERROR)
	} else {
		responseData = userService.LoginUser(request, c.ClientIP())
	}
	c.JSON(http.StatusOK, responseData)
}

// @Summary 用户注册
// @tags 用户
// @Accept  application/json
// @Produce  json
// @Param body body entity.UserRegister true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/register [POST]
func RegisterUser(c *gin.Context) {
	responseData := entity.ResponseData{}
	request := entity.UserRegister{}
	// 参数校验
	if c.ShouldBindJSON(&request) != nil {
		responseData.Message = msg.GetMsg(tcode.PARAMETER_ERROR)
	} else {
		responseData = userService.RegisterUser(request)
	}
	c.JSON(http.StatusOK, responseData)
}

// @Summary 修改用户密码
// @tags 用户
// @Accept  application/json
// @Produce  json
// @Param body body entity.UserUpdatePass true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/updatePass [PUT]
// @Security ApiKeyAuth
func UpdateUserPass(c *gin.Context) {
	responseData := entity.ResponseData{}
	request := entity.UserUpdatePass{}
	// 参数校验
	if c.ShouldBindJSON(&request) != nil {
		responseData.Message = msg.GetMsg(tcode.PARAMETER_ERROR)
	} else {
		token := c.Query("token")
		if token == "" {
			authToken := c.GetHeader("Authorization")
			if authToken == "" {
				responseData.Message = msg.GetMsg(tcode.AUTH_NOT_BEARER)
			}
			token = strings.TrimSpace(authToken)
		}
		if responseData.Message == "" {
			responseData = userService.UpdateUserPass(token, request)
		}
	}
	c.JSON(http.StatusOK, responseData)
}

// @Summary 添加用户地址信息
// @tags 用户
// @Accept  application/json
// @Produce  json
// @Param body body entity.AddUserInfo true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/addUserInfo [POST]
// @Security ApiKeyAuth
func AddUserInfo(c *gin.Context) {
	responseData := entity.ResponseData{}
	request := entity.AddUserInfo{}
	// 参数校验
	if c.ShouldBindJSON(&request) != nil {
		responseData.Message = msg.GetMsg(tcode.PARAMETER_ERROR)
	} else {
		token := c.Query("token")
		if token == "" {
			authToken := c.GetHeader("Authorization")
			if authToken == "" {
				responseData.Message = msg.GetMsg(tcode.AUTH_NOT_BEARER)
			}
			token = strings.TrimSpace(authToken)
		}
		if responseData.Message == "" {
			responseData = userService.AddUserInfo(token, request)
		}
	}
	c.JSON(http.StatusOK, responseData)
}

// @Summary 查询用户所有地址信息
// @tags 用户
// @Accept  application/json
// @Produce  json
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/queryUserInfoByUserID [GET]
// @Security ApiKeyAuth
func QueryUserInfoByUserID(c *gin.Context) {
	responseData := entity.ResponseData{}
	token := c.Query("token")
	if token == "" {
		authToken := c.GetHeader("Authorization")
		if authToken == "" {
			responseData.Message = msg.GetMsg(tcode.AUTH_NOT_BEARER)
		}
		token = strings.TrimSpace(authToken)
	}
	if responseData.Message == "" {
		responseData = userService.QueryUserInfoByUserID(token)
	}
	c.JSON(http.StatusOK, responseData)
}

// @Summary 查询用户地址信息详情
// @tags 用户
// @Accept  application/json
// @Produce  json
// @Param id query string true "用户地址ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/queryUserInfoByID [GET]
// @Security ApiKeyAuth
func QueryUserInfoByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.DefaultQuery("id", "0"), 10, 64)
	responseData := userService.QueryUserInfoByID(id)
	c.JSON(http.StatusOK, responseData)
}

// @Summary 删除户地址信息
// @tags 用户
// @Accept  application/json
// @Produce  json
// @Param id query string true "用户地址ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/deleteUserInfo [DELETE]
// @Security ApiKeyAuth
func DeleteUserInfo(c *gin.Context) {
	request := entity.DeleteIds{}
	id, _ := strconv.ParseInt(c.DefaultQuery("id", "0"), 10, 64)
	ids := append(request.Ids, id)
	responseData := userService.DeleteUserInfo(ids)
	c.JSON(http.StatusOK, responseData)
}

// @Summary 批量删除户地址信息
// @tags 用户
// @Accept  application/json
// @Produce  json
// @Param body body entity.DeleteIds true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/deleteUserInfos [DELETE]
// @Security ApiKeyAuth
func DeleteUserInfos(c *gin.Context) {
	responseData := entity.ResponseData{}
	request := entity.DeleteIds{}
	// 参数校验
	if c.ShouldBindJSON(&request) != nil {
		responseData.Message = msg.GetMsg(tcode.PARAMETER_ERROR)
	} else {
		responseData = userService.DeleteUserInfo(request.Ids)
	}
	c.JSON(http.StatusOK, responseData)
}

// @Summary 分页查询用户(默认前100条) 并返回总记录数
// @tags 管理员
// @Accept application/x-www-form-urlencoded
// @Produce  json
// @Param pageSize query string false "页大小"
// @Param page query string false "页数"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/users/queryUserByLimitOffset [GET]
// @Security ApiKeyAuth
func QueryUserByLimitOffset(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "100"))
	responseData := userService.QueryUserByLimitOffset(pageSize, page)
	c.JSON(http.StatusOK, responseData)
}
