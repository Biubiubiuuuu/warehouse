package goodsService

import (
	tcode "github.com/Biubiubiuuuu/warehouse/server/common/tips/code"
	"github.com/Biubiubiuuuu/warehouse/server/common/tips/msg"
	"github.com/Biubiubiuuuu/warehouse/server/entity"
	"github.com/Biubiubiuuuu/warehouse/server/models"
)

// 添加商品种类
func AddGoodsType(g entity.AddGoodsType) (responseData entity.ResponseData) {
	goodsType := models.GoodsType{
		GoodsName:        g.GoodsName,
		GoodsSpecs:       g.GoodsSpecs,
		GoodsUnitPrince:  g.GoodsUnitPrince,
		GoodsPrince:      g.GoodsPrince,
		GoodsImage:       g.GoodsImage,
		GoodsBatchNumber: g.GoodsBatchNumber,
		GoodsDate:        g.GoodsDate,
	}
	if err := goodsType.AddGoodsType(); err != nil {
		responseData.Message = msg.GetMsg(tcode.ADD_ERROR)
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
	args := map[string]interface{}{
		"GoodsName":        g.GoodsName,
		"GoodsSpecs":       g.GoodsSpecs,
		"GoodsUnitPrince":  g.GoodsUnitPrince,
		"GoodsPrince ":     g.GoodsPrince,
		"GoodsImage":       g.GoodsImage,
		"GoodsBatchNumber": g.GoodsBatchNumber,
		"GoodsDate":        g.GoodsDate,
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
	admins := models.QueryGoodsTypesByLimitOffset(pageSize, page)
	responseData.Message = msg.GetMsg(tcode.QUERY_SUCCESS)
	if len(admins) == 0 {
		responseData.Message = msg.GetMsg(tcode.NOTMORE)
	}
	count := models.QueryGoodsTypesCount()
	data := make(map[string]interface{})
	data["goodsTypes"] = admins
	data["count"] = count
	responseData.Data = data
	responseData.Status = true
	return
}
