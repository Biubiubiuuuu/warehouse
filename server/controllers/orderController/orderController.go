package orderController

import (
	"net/http"
	"strconv"
	"strings"

	tcode "github.com/Biubiubiuuuu/warehouse/server/common/tips/code"
	"github.com/Biubiubiuuuu/warehouse/server/common/tips/msg"
	"github.com/Biubiubiuuuu/warehouse/server/entity"
	"github.com/Biubiubiuuuu/warehouse/server/services/orderService"
	"github.com/gin-gonic/gin"
)

// @Summary 添加订单
// @tags 用户
// @Accept  application/json
// @Produce  json
// @Param body body entity.AddOrder true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/user/order/addOrder [POST]
// @Security ApiKeyAuth
func AddOrder(c *gin.Context) {
	responseData := entity.ResponseData{}
	request := entity.AddOrder{}
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
			responseData = orderService.AddOrder(token, request)
		}
	}
	c.JSON(http.StatusOK, responseData)
}

// @Summary 查询订单详情
// @tags 用户
// @Accept  application/json
// @Produce  json
// @Param id query string true "订单ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/user/order/queryByGoodsOrderID [GET]
// @Security ApiKeyAuth
func QueryByGoodsOrderID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.DefaultQuery("id", "0"), 10, 64)
	responseData := orderService.QueryByGoodsOrderID(id)
	c.JSON(http.StatusOK, responseData)
}

// @Summary 查询用户订单
// @tags 用户
// @Accept  application/json
// @Produce  json
// @Param pageSize query string false "页大小"
// @Param page query string false "页数"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/user/order/queryByOrderUserID [GET]
// @Security ApiKeyAuth
func QueryByOrderUserID(c *gin.Context) {
	responseData := entity.ResponseData{}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "100"))
	token := c.Query("token")
	if token == "" {
		authToken := c.GetHeader("Authorization")
		if authToken == "" {
			responseData.Message = msg.GetMsg(tcode.AUTH_NOT_BEARER)
		}
		token = strings.TrimSpace(authToken)
	}
	if responseData.Message == "" {
		responseData = orderService.QueryByOrderUserID(token, pageSize, page)
	}
	c.JSON(http.StatusOK, responseData)
}

// @Summary 分页查询订单(默认前100条) 并返回总记录数
// @tags 管理员
// @Accept application/x-www-form-urlencoded
// @Produce  json
// @Param pageSize query string false "页大小"
// @Param page query string false "页数"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/order/queryOrderByLimitOffset [GET]
// @Security ApiKeyAuth
func QueryOrderByLimitOffset(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "100"))
	responseData := orderService.QueryOrderByLimitOffset(pageSize, page)
	c.JSON(http.StatusOK, responseData)
}
