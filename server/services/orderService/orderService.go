package orderService

import (
	"strconv"
	"time"

	tcode "github.com/Biubiubiuuuu/warehouse/server/common/tips/code"
	"github.com/Biubiubiuuuu/warehouse/server/common/tips/msg"
	"github.com/Biubiubiuuuu/warehouse/server/entity"
	"github.com/Biubiubiuuuu/warehouse/server/models"
)

// 创建订单
/*
请求参数
先查询商品是否存在  不存在返回error
检查库存是否充足    不充足  提示
更新商品库存（加锁）
订单表添加一条数据
*/
func AddOrder(token string, request entity.AddOrder) (responseData entity.ResponseData) {
	user := models.User{Token: token}
	if !user.QueryUser() {
		responseData.Message = msg.GetMsg(tcode.ADD_ERROR) + "，用户不存在"
		return
	}
	order := models.Order{
		UserID: user.ID,
		OrderUserInfo: models.OrderUserInfo{
			Provice:     request.OrderUserInfo.Provice,
			City:        request.OrderUserInfo.City,
			ShopAddress: request.OrderUserInfo.ShopAddress,
			Tel:         request.OrderUserInfo.Tel,
		},
	}
	var orderGoodsInfos []models.OrderGoodsInfo
	for _, item := range request.OrderGoodsInfo {
		g := models.GoodsType{}
		g.ID = item.GoodsTypeID
		if err := g.QueryByGoodsTypeID(); err != nil {
			responseData.Message = msg.GetMsg(tcode.ADD_ERROR) + ",该商品不存在"
			return
		}
		goodsStock := models.GoodsStock{}
		goodsStock.GoodsTypeID = item.GoodsTypeID
		stocks, errStock := goodsStock.QueryByGoodsStockID()
		if errStock != nil {
			responseData.Message = msg.GetMsg(tcode.ADD_ERROR) + ",该商品库存不存在，待添加"
			return
		}
		if stocks.QuantityStock <= 0 {
			responseData.Message = msg.GetMsg(tcode.ADD_ERROR) + ",该商品库存不足"
			return
		}
		orderGoodsInfo := models.OrderGoodsInfo{
			GoodsName:   g.GoodsName,
			GoodsSpecs:  g.GoodsSpecs,
			GoodsPrince: g.GoodsPrince,
			GoodsImage:  g.GoodsImage,
			GoodsQty:    item.GoodsQty,
		}
		orderGoodsInfos = append(orderGoodsInfos, orderGoodsInfo)
	}
	order.OrderGoodsInfos = orderGoodsInfos

	goodsOrderNumber := strconv.Itoa(time.Now().Nanosecond())
	return
}
