package goodsTypeController

import (
	"net/http"
	"strconv"

	tcode "github.com/Biubiubiuuuu/warehouse/server/common/tips/code"
	"github.com/Biubiubiuuuu/warehouse/server/common/tips/msg"
	"github.com/Biubiubiuuuu/warehouse/server/entity"
	"github.com/Biubiubiuuuu/warehouse/server/services/goodsService"
	"github.com/gin-gonic/gin"
)

// @Summary 添加商品种类
// @tags 管理员
// @Accept  application/json
// @Produce  json
// @Param username query string true "username"
// @Param body body entity.AddGoodsType true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/addGoodsType [POST]
// @Security ApiKeyAuth
func AddGoodsType(c *gin.Context) {
	request := entity.AddGoodsType{}
	responseData := entity.ResponseData{}
	if c.ShouldBindJSON(&request) != nil {
		responseData.Message = msg.GetMsg(tcode.PARAMETER_ERROR)
	} else {
		responseData = goodsService.AddGoodsType(request)
	}
	c.JSON(http.StatusOK, responseData)
}

// @Summary 修改商品种类信息
// @tags 管理员
// @Accept  application/json
// @Produce  json
// @Param body body entity.UpdateGoodsType true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/updateGoodsType [PUT]
// @Security ApiKeyAuth
func UpdateGoodsType(c *gin.Context) {
	request := entity.UpdateGoodsType{}
	responseData := entity.ResponseData{}
	if c.ShouldBindJSON(&request) != nil {
		responseData.Message = msg.GetMsg(tcode.PARAMETER_ERROR)
	} else {
		responseData = goodsService.UpdateGoodsType(request)
	}
	c.JSON(http.StatusOK, responseData)
}

// @Summary 查看商品种类详情
// @tags 管理员
// @Accept  application/json
// @Produce  json
// @Param id query string true "商品种类ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/queryByGoodsTypeID [GET]
// @Security ApiKeyAuth
func QueryByGoodsTypeID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.DefaultQuery("id", "0"), 10, 64)
	responseData := goodsService.QueryByGoodsTypeID(id)
	c.JSON(http.StatusOK, responseData)
}

// @Summary 下架商品
// @tags 管理员
// @Accept  application/json
// @Produce  json
// @Param id query string true "id"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/deleteGoodsType [DELETE]
// @Security ApiKeyAuth
func DeleteGoodsType(c *gin.Context) {
	request := entity.DeleteIds{}
	id, _ := strconv.ParseInt(c.DefaultQuery("id", "0"), 10, 64)
	ids := append(request.Ids, id)
	responseData := goodsService.DeleteGoodsTypes(ids)
	c.JSON(http.StatusOK, responseData)
}

// @Summary 批量下架商品
// @tags 管理员
// @Accept  application/json
// @Produce  json
// @Param body body entity.DeleteIds true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/deleteGoodsTypes [DELETE]
// @Security ApiKeyAuth
func DeleteGoodsTypes(c *gin.Context) {
	responseData := entity.ResponseData{}
	request := entity.DeleteIds{}
	// 参数校验
	if c.ShouldBindJSON(&request) != nil {
		responseData.Message = msg.GetMsg(tcode.PARAMETER_ERROR)
	} else {
		responseData = goodsService.DeleteGoodsTypes(request.Ids)
	}
	c.JSON(http.StatusOK, responseData)
}

// @Summary 分页查询商品种类(默认前100条) 并返回总记录数
// @tags 管理员
// @Accept application/x-www-form-urlencoded
// @Produce  json
// @Param pageSize query string false "页大小"
// @Param page query string false "页数"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/queryGoodsTypesByLimitOffset [GET]
// @Security ApiKeyAuth
func QueryGoodsTypesByLimitOffset(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "100"))
	responseData := goodsService.QueryGoodsTypesByLimitOffset(pageSize, page)
	c.JSON(http.StatusOK, responseData)
}
