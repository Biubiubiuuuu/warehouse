package goodsStockController

import (
	"net/http"
	"strconv"

	tcode "github.com/Biubiubiuuuu/warehouse/server/common/tips/code"
	"github.com/Biubiubiuuuu/warehouse/server/common/tips/msg"
	"github.com/Biubiubiuuuu/warehouse/server/entity"
	"github.com/Biubiubiuuuu/warehouse/server/services/goodsService"
	"github.com/gin-gonic/gin"
)

// @Summary 添加商品库存
// @tags 商品库存
// @Accept  application/json
// @Produce  json
// @Param body body entity.AddGoodsStock true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/goodsStock/addGoodsStock [POST]
// @Security ApiKeyAuth
func AddGoodsStock(c *gin.Context) {
	request := entity.AddGoodsStock{}
	responseData := entity.ResponseData{}
	if c.ShouldBindJSON(&request) != nil {
		responseData.Message = msg.GetMsg(tcode.PARAMETER_ERROR)
	} else {
		responseData = goodsService.AddGoodsStock(request)
	}
	c.JSON(http.StatusOK, responseData)
}

// @Summary 增加商品库存、总数量
// @tags 商品库存
// @Accept  application/json
// @Produce  json
// @Param body body entity.UpdateGoodsStock true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/goodsStock/updateGoodsStock [PUT]
// @Security ApiKeyAuth
func UpdateGoodsStock(c *gin.Context) {
	request := entity.UpdateGoodsStock{}
	responseData := entity.ResponseData{}
	if c.ShouldBindJSON(&request) != nil {
		responseData.Message = msg.GetMsg(tcode.PARAMETER_ERROR)
	} else {
		responseData = goodsService.UpdateGoodsStock(request)
	}
	c.JSON(http.StatusOK, responseData)
}

// @Summary 查看商品库存详情
// @tags 商品库存
// @Accept  application/json
// @Produce  json
// @Param id query string true "商品库存ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/goodsStock/queryByGoodsStockID [GET]
// @Security ApiKeyAuth
func QueryByGoodsStockID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.DefaultQuery("id", "0"), 10, 64)
	responseData := goodsService.QueryByGoodsStockID(id)
	c.JSON(http.StatusOK, responseData)
}

// @Summary 分页查询商品库存(默认前100条) 并返回总记录数
// @tags 商品库存
// @Accept application/x-www-form-urlencoded
// @Produce  json
// @Param pageSize query string false "页大小"
// @Param page query string false "页数"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/goodsStock/queryGoodsStocksByLimitOffset [GET]
// @Security ApiKeyAuth
func QueryGoodsStocksByLimitOffset(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "100"))
	responseData := goodsService.QueryGoodsStocksByLimitOffset(pageSize, page)
	c.JSON(http.StatusOK, responseData)
}
