package goodsTypeController

import (
	"net/http"
	"strconv"
	"strings"

	tcode "github.com/Biubiubiuuuu/warehouse/server/common/tips/code"
	"github.com/Biubiubiuuuu/warehouse/server/common/tips/msg"
	"github.com/Biubiubiuuuu/warehouse/server/entity"
	"github.com/Biubiubiuuuu/warehouse/server/helpers/configHelper"
	"github.com/Biubiubiuuuu/warehouse/server/helpers/fileHelper"
	"github.com/Biubiubiuuuu/warehouse/server/services/goodsService"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary 添加商品种类
// @tags 商品种类
// @Accept  multipart/form-data
// @Produce  json
// @Param goods_name formData string true "商品名称"
// @Param goods_specs formData string false "商品规格 1.盒 2.瓶 3.支"
// @Param goods_unitprince formData float false "商品成本价"
// @Param goods_prince formData float false "商品销售价"
// @Param goods_image[] formData file false "商品图片"
// @Param goods_batch_number formData string false "生产批号"
// @Param goods_date formData string false "生产日期"
// @Param goods_state formData string false "商品状态 1.下架  2.在售"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/goodsType/addGoodsType [POST]
// @Security ApiKeyAuth
func AddGoodsType(c *gin.Context) {
	responseData := entity.ResponseData{}
	token := c.Query("token")
	if token == "" {
		authToken := c.GetHeader("Authorization")
		if authToken == "" {
			responseData.Message = msg.GetMsg(tcode.AUTH_NOT_BEARER)
		}
		token = strings.TrimSpace(authToken)
	}
	goods_name := c.PostForm("goods_name")
	goods_specs := c.PostForm("goods_specs")
	goods_unitprince := c.PostForm("goods_unitprince")
	floatgoods_unitprince, _ := strconv.ParseFloat(goods_unitprince, 64)
	goods_prince := c.PostForm("goods_prince")
	floatgoods_prince, _ := strconv.ParseFloat(goods_prince, 64)
	goods_batch_number := c.PostForm("goods_batch_number")
	goods_date := c.PostForm("goods_date")
	goods_state := c.PostForm("goods_state")
	request := entity.AddGoodsType{
		GoodsName:        goods_name,
		GoodsSpecs:       goods_specs,
		GoodsUnitPrince:  floatgoods_unitprince,
		GoodsPrince:      floatgoods_prince,
		GoodsBatchNumber: goods_batch_number,
		GoodsDate:        goods_date,
		GoodsState:       goods_state,
	}
	// 获取主机头
	r := c.Request
	host := r.Host
	// 多商品图片
	var imgs []string
	form, _ := c.MultipartForm()
	files := form.File["goods_image[]"]
	for _, file := range files {
		// 文件名 避免重复取uuid
		var filename string
		uuid, _ := uuid.NewUUID()
		arr := strings.Split(file.Filename, ".")
		if strings.EqualFold(arr[1], "png") {
			filename = uuid.String() + ".png"
		} else {
			filename = uuid.String() + ".jpg"
		}
		pathFile := configHelper.ImageDir
		if !fileHelper.IsExist(pathFile) {
			fileHelper.CreateDir(pathFile)
		}
		pathFile = pathFile + filename

		if err := c.SaveUploadedFile(file, pathFile); err == nil {
			imgs = append(imgs, host+"/"+pathFile)
		}
	}
	if responseData.Message == "" {
		responseData = goodsService.AddGoodsType(token, imgs, request)
	}
	c.JSON(http.StatusOK, responseData)
}

// @Summary 修改商品种类信息
// @tags 商品种类
// @Accept  multipart/form-data
// @Produce  json
// @Param goods_id formData int true "商品ID"
// @Param goods_name formData string true "商品名称"
// @Param goods_specs formData string false "商品规格 1.盒 2.瓶 3.支"
// @Param goods_unitprince formData float false "商品成本价"
// @Param goods_prince formData float false "商品销售价"
// @Param goods_image[] formData file false "商品图片"
// @Param goods_batch_number formData string false "生产批号"
// @Param goods_date formData string false "生产日期"
// @Param goods_state formData string false "商品状态 1.下架  2.在售"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/goodsType/updateGoodsType [PUT]
// @Security ApiKeyAuth
func UpdateGoodsType(c *gin.Context) {
	responseData := entity.ResponseData{}
	token := c.Query("token")
	if token == "" {
		authToken := c.GetHeader("Authorization")
		if authToken == "" {
			responseData.Message = msg.GetMsg(tcode.AUTH_NOT_BEARER)
		}
		token = strings.TrimSpace(authToken)
	}
	goods_id := c.PostForm("goods_id")
	intgoods_id, _ := strconv.ParseInt(goods_id, 10, 64)
	goods_name := c.PostForm("goods_name")
	goods_specs := c.PostForm("goods_specs")
	goods_unitprince := c.PostForm("goods_unitprince")
	floatgoods_unitprince, _ := strconv.ParseFloat(goods_unitprince, 64)
	goods_prince := c.PostForm("goods_prince")
	floatgoods_prince, _ := strconv.ParseFloat(goods_prince, 64)
	goods_batch_number := c.PostForm("goods_batch_number")
	goods_date := c.PostForm("goods_date")
	goods_state := c.PostForm("goods_state")
	request := entity.UpdateGoodsType{
		GoodsID:          intgoods_id,
		GoodsName:        goods_name,
		GoodsSpecs:       goods_specs,
		GoodsUnitPrince:  floatgoods_unitprince,
		GoodsPrince:      floatgoods_prince,
		GoodsBatchNumber: goods_batch_number,
		GoodsDate:        goods_date,
		GoodsState:       goods_state,
	}
	// 获取主机头
	r := c.Request
	host := r.Host
	// 多商品图片
	var imgs []string
	form, _ := c.MultipartForm()
	files := form.File["goods_image[]"]
	for _, file := range files {
		// 文件名 避免重复取uuid
		var filename string
		uuid, _ := uuid.NewUUID()
		arr := strings.Split(file.Filename, ".")
		if strings.EqualFold(arr[1], "png") {
			filename = uuid.String() + ".png"
		} else {
			filename = uuid.String() + ".jpg"
		}
		pathFile := configHelper.ImageDir
		if !fileHelper.IsExist(pathFile) {
			fileHelper.CreateDir(pathFile)
		}
		pathFile = pathFile + filename

		if err := c.SaveUploadedFile(file, pathFile); err == nil {
			imgs = append(imgs, host+"/"+pathFile)
		}
	}
	if responseData.Message == "" {
		responseData = goodsService.UpdateGoodsType(token, imgs, request)
	}
	c.JSON(http.StatusOK, responseData)
}

// @Summary 查看商品种类详情
// @tags 商品种类
// @Accept  application/json
// @Produce  json
// @Param id query string true "商品种类ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/goodsType/queryByGoodsTypeID [GET]
// @Security ApiKeyAuth
func QueryByGoodsTypeID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.DefaultQuery("id", "0"), 10, 64)
	responseData := goodsService.QueryByGoodsTypeID(id)
	c.JSON(http.StatusOK, responseData)
}

// @Summary 下架商品
// @tags 商品种类
// @Accept  application/json
// @Produce  json
// @Param id query string true "id"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/goodsType/deleteGoodsType [DELETE]
// @Security ApiKeyAuth
func DeleteGoodsType(c *gin.Context) {
	request := entity.DeleteIds{}
	id, _ := strconv.ParseInt(c.DefaultQuery("id", "0"), 10, 64)
	ids := append(request.Ids, id)
	responseData := goodsService.DeleteGoodsTypes(ids)
	c.JSON(http.StatusOK, responseData)
}

// @Summary 批量下架商品
// @tags 商品种类
// @Accept  application/json
// @Produce  json
// @Param body body entity.DeleteIds true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/goodsType/deleteGoodsTypes [DELETE]
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
// @tags 商品种类
// @Accept application/x-www-form-urlencoded
// @Produce  json
// @Param pageSize query string false "页大小"
// @Param page query string false "页数"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/goodsType/queryGoodsTypesByLimitOffset [GET]
// @Security ApiKeyAuth
func QueryGoodsTypesByLimitOffset(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "100"))
	responseData := goodsService.QueryGoodsTypesByLimitOffset(pageSize, page)
	c.JSON(http.StatusOK, responseData)
}

// @Summary 查询商品种类ID和商品名（支持模糊查询）
// @tags 商品种类
// @Accept application/x-www-form-urlencoded
// @Produce  json
// @Param goods_name query string false "商品名称"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/goodsType/queryAllGoods [GET]
// @Security ApiKeyAuth
func QueryAllGoods(c *gin.Context) {
	goods_name := c.DefaultQuery("goods_name", "")
	responseData := goodsService.QueryAllGoods(goods_name)
	c.JSON(http.StatusOK, responseData)
}
