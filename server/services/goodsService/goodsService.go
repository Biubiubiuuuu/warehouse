package goodsService

import (
	"fmt"
	"time"

	"github.com/Biubiubiuuuu/warehouse/server/common/tips/code"
	tcode "github.com/Biubiubiuuuu/warehouse/server/common/tips/code"
	"github.com/Biubiubiuuuu/warehouse/server/common/tips/msg"
	"github.com/Biubiubiuuuu/warehouse/server/entity"
	"github.com/Biubiubiuuuu/warehouse/server/models"
)

/*==========================================商品种类curd==================================*/

// 添加商品种类
func AddGoodsType(g entity.AddGoodsType) (responseData entity.ResponseData) {
	admin := models.Admin{
		Username: g.GoodsCreateAdmin,
	}
	if !admin.QueryAdminByUsername() {
		responseData.Message = msg.GetMsg(tcode.ADD_ERROR) + "，管理员不存在"
		return
	}
	goodsDate, _ := time.ParseInLocation("2006-01-02", g.GoodsDate, time.Local)
	goodsType := models.GoodsType{
		GoodsName:        g.GoodsName,
		GoodsSpecs:       g.GoodsSpecs,
		GoodsUnitPrince:  g.GoodsUnitPrince,
		GoodsPrince:      g.GoodsPrince,
		GoodsImage:       g.GoodsImage,
		GoodsBatchNumber: g.GoodsBatchNumber,
		GoodsDate:        goodsDate,
		GoodsCreateAdmin: g.GoodsCreateAdmin,
	}
	if err := goodsType.AddGoodsType(); err != nil {
		responseData.Message = msg.GetMsg(tcode.ADD_ERROR) + fmt.Sprintf(",%v%v", g.GoodsName, msg.GetMsg(code.EXIST))
		return
	}
	responseData.Status = true
	responseData.Message = msg.GetMsg(tcode.ADD_SUCCESS)
	return
}

// 修改商品种类信息
func UpdateGoodsType(g entity.UpdateGoodsType) (responseData entity.ResponseData) {
	goodsType := models.GoodsType{}
	goodsType.ID = g.GoodsID
	if err := goodsType.QueryByGoodsTypeID(); err != nil {
		responseData.Message = msg.GetMsg(tcode.UPDATE_ERROR) + ",商品ID不存在"
		return
	}
	goodsDate, _ := time.ParseInLocation("2006-01-02", g.GoodsDate, time.Local)
	args := map[string]interface{}{
		"GoodsName":        g.GoodsName,
		"GoodsSpecs":       g.GoodsSpecs,
		"GoodsUnitPrince":  g.GoodsUnitPrince,
		"GoodsPrince ":     g.GoodsPrince,
		"GoodsImage":       g.GoodsImage,
		"GoodsBatchNumber": g.GoodsBatchNumber,
		"GoodsDate":        goodsDate,
		"GoodsState":       g.GoodsState,
	}
	if err := goodsType.UpdateGoodsType(args); err != nil {
		responseData.Message = msg.GetMsg(tcode.UPDATE_ERROR)
		return
	}
	responseData.Status = true
	responseData.Message = msg.GetMsg(tcode.UPDATE_SUCCESS)
	return
}

// 查询商品种类详情
func QueryByGoodsTypeID(id int64) (responseData entity.ResponseData) {
	goodsType := models.GoodsType{}
	goodsType.ID = id
	if err := goodsType.QueryByGoodsTypeID(); err != nil {
		responseData.Message = msg.GetMsg(tcode.QUERY_ERROR)
		return
	}
	data := make(map[string]interface{})
	data["goodsType"] = goodsType
	responseData.Data = data
	responseData.Status = true
	responseData.Message = msg.GetMsg(tcode.QUERY_SUCCESS)
	return
}

// 下架商品
func DeleteGoodsTypes(ids []int64) (responseData entity.ResponseData) {
	if len(ids) == 0 {
		responseData.Message = msg.GetMsg(tcode.NOTNULL)
		return
	}
	goodsType := models.GoodsType{}
	if err := goodsType.DeleteGoodsTypes(ids); err != nil {
		responseData.Message = msg.GetMsg(tcode.DELETE_ERROR)
		return
	}
	responseData.Status = true
	responseData.Message = msg.GetMsg(tcode.DELETE_SUCCESS)
	return
}

// 查看商品种类（分页查询）
func QueryGoodsTypesByLimitOffset(pageSize int, page int) (responseData entity.ResponseData) {
	goodsTypes := models.QueryGoodsTypesByLimitOffset(pageSize, page)
	responseData.Message = msg.GetMsg(tcode.QUERY_SUCCESS)
	if len(goodsTypes) == 0 {
		responseData.Message = msg.GetMsg(tcode.NOTMORE)
	}
	count := models.QueryGoodsTypesCount()
	data := make(map[string]interface{})
	data["goodsTypes"] = goodsTypes
	data["count"] = count
	responseData.Data = data
	responseData.Status = true
	return
}

// 查询商品种类ID和商品名（支持模糊查询）
func QueryAllGoods(goodsName string) (responseData entity.ResponseData) {
	goodsType := models.GoodsType{}
	goodsType.GoodsName = goodsName
	goods := goodsType.QueryAllGoods()
	if len(goods) == 0 {
		responseData.Message = msg.GetMsg(tcode.NOTMORE)
	}
	data := make(map[string]interface{})
	data["goodsTypes"] = goods
	responseData.Data = data
	responseData.Status = true
	responseData.Message = msg.GetMsg(tcode.QUERY_SUCCESS)
	return
}

/*==========================================end==========================================*/

/*==========================================商品库存curd==================================*/

// 创建商品库存
func AddGoodsStock(g entity.AddGoodsStock) (responseData entity.ResponseData) {
	goodsStock := models.GoodsStock{
		QuantityTotal: g.QuantityTotal,
		GoodsTypeID:   g.GoodsTypeID,
	}
	goodsType := models.GoodsType{}
	goodsType.ID = g.GoodsTypeID
	if err := goodsType.QueryByGoodsTypeID(); err != nil {
		responseData.Message = msg.GetMsg(tcode.QUERY_ERROR) + "，该商品种类不存在，请先添加商品种类"
		return
	}
	if err := goodsStock.AddGoodsStock(); err != nil {
		responseData.Message = err.Error()
		return
	}
	responseData.Status = true
	responseData.Message = msg.GetMsg(tcode.ADD_SUCCESS)
	return
}

// 增加商品库存总数量
func UpdateGoodsStock(g entity.UpdateGoodsStock) (responseData entity.ResponseData) {
	goodsStock := models.GoodsStock{}
	goodsStock.ID = g.GoodsStockID
	var data models.GoodsStockData
	var errQueryByGoodsStockID error
	if data, errQueryByGoodsStockID = goodsStock.QueryByGoodsStockID(); errQueryByGoodsStockID != nil {
		responseData.Message = msg.GetMsg(tcode.UPDATE_ERROR) + "，该商品库存不存在，请先添加商品库存信息再修改"
		return
	}
	quantityTotal := data.QuantityTotal + g.AddQuantity
	quantityStock := data.QuantityStock + g.AddQuantity
	args := map[string]interface{}{
		"QuantityTotal": quantityTotal,
		"QuantityStock": quantityStock,
	}
	if err := goodsStock.UpdateGoodsStock(args); err != nil {
		responseData.Message = msg.GetMsg(tcode.UPDATE_ERROR)
		return
	}
	responseData.Status = true
	responseData.Message = msg.GetMsg(tcode.UPDATE_SUCCESS)
	return
}

// 查询商品种类详情
func QueryByGoodsStockID(id int64) (responseData entity.ResponseData) {
	goodsStock := models.GoodsStock{}
	goodsStock.ID = id
	goodsStockData, err := goodsStock.QueryByGoodsStockID()
	if err != nil {
		responseData.Message = msg.GetMsg(tcode.QUERY_ERROR) + "，该商品库存不存在"
		return
	}
	data := make(map[string]interface{})
	data["goodsStock"] = goodsStockData
	responseData.Data = data
	responseData.Status = true
	responseData.Message = msg.GetMsg(tcode.QUERY_SUCCESS)
	return
}

// 查看商品库存（分页查询）
func QueryGoodsStocksByLimitOffset(pageSize int, page int) (responseData entity.ResponseData) {
	goodsStocks := models.QueryGoodsStockByLimitOffset(pageSize, page)
	responseData.Message = msg.GetMsg(tcode.QUERY_SUCCESS)
	if len(goodsStocks) == 0 {
		responseData.Message = msg.GetMsg(tcode.NOTMORE)
	}
	count := models.QueryGoodsStockCount()
	data := make(map[string]interface{})
	data["goodsStocks"] = goodsStocks
	data["count"] = count
	responseData.Data = data
	responseData.Status = true
	return
}

/*==========================================end==========================================*/

/*==========================================商品订单curd==================================*/

// 查看商品订单（分页查询）
func QueryGoodsOrderByLimitOffset(pageSize int, page int) (responseData entity.ResponseData) {
	goodsOrders := models.QueryGoodsOrderByLimitOffset(pageSize, page)
	responseData.Message = msg.GetMsg(tcode.QUERY_SUCCESS)
	if len(goodsOrders) == 0 {
		responseData.Message = msg.GetMsg(tcode.NOTMORE)
	}
	count := models.QueryGoodsStockCount()
	data := make(map[string]interface{})
	data["goodsOrders"] = goodsOrders
	data["count"] = count
	responseData.Data = data
	responseData.Status = true
	return
}

/*==========================================end==========================================*/
