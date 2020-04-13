package orderService

import (
	"fmt"
	"strconv"
	"time"

	tcode "github.com/Biubiubiuuuu/warehouse/server/common/tips/code"
	"github.com/Biubiubiuuuu/warehouse/server/common/tips/msg"
	"github.com/Biubiubiuuuu/warehouse/server/entity"
	"github.com/Biubiubiuuuu/warehouse/server/models"
)

// 创建订单
func AddOrder(token string, request entity.AddOrder) (responseData entity.ResponseData) {
	user := models.User{Token: token}
	if !user.QueryUser() {
		responseData.Message = msg.GetMsg(tcode.ADD_ERROR) + "，用户不存在"
		return
	}
	order := models.Order{
		UserID: user.ID,
		OrderUserDetail: models.OrderUserDetail{
			Provice:     request.UserDetails.Provice,
			City:        request.UserDetails.City,
			ShopAddress: request.UserDetails.ShopAddress,
			Tel:         request.UserDetails.Tel,
		},
	}
	var arr []models.OrderGoodsDetail
	for _, item := range request.GoodsDetails {
		g := models.GoodsType{}
		g.ID = item.GoodsTypeID
		if err := g.QueryByGoodsTypeID(); err != nil {
			responseData.Message = msg.GetMsg(tcode.ADD_ERROR) + ",该商品不存在"
			return
		}
		goodsStock := models.GoodsStock{}
		goodsStock.GoodsTypeID = item.GoodsTypeID
		stocks, errStock := goodsStock.QueryByID()
		if errStock != nil {
			responseData.Message = msg.GetMsg(tcode.ADD_ERROR) + ",该商品库存不存在，待添加"
			return
		}
		if stocks.QuantityStock <= 0 || stocks.QuantityStock < item.GoodsQty {
			responseData.Message = msg.GetMsg(tcode.ADD_ERROR) + ",该商品库存不足"
			return
		}
		goodsDetail := models.OrderGoodsDetail{
			GoodsName:   g.GoodsName,
			GoodsSpecs:  g.GoodsSpecs,
			GoodsPrince: g.GoodsPrince,
			GoodsImage:  g.GoodsImage,
			GoodsQty:    item.GoodsQty,
			GoodsTypeID: item.GoodsTypeID,
		}
		// 订单总金额
		str := strconv.FormatInt(item.GoodsQty, 10)
		float, _ := strconv.ParseFloat(str, 64)
		order.OrderPrince += g.GoodsPrince * float
		arr = append(arr, goodsDetail)
	}
	order.OrderGoodsDetails = arr
	// 订单编号 直接取时间戳
	orderNumber := strconv.Itoa(time.Now().Nanosecond())
	order.OrderNumber = orderNumber
	// 订单状态 1.已付款 2.未付款
	order.OrderStatus = "2"
	fmt.Println(order)
	if err := order.AddOrder(); err != nil {
		responseData.Message = "下单失败"
		return
	}
	responseData.Status = true
	responseData.Message = "下单成功"
	return
}

// 查询订单详情
func QueryByGoodsOrderID(id int64) (responseData entity.ResponseData) {
	order := models.Order{}
	order.ID = id
	if err := order.QueryByGoodsOrderID(); err != nil {
		responseData.Message = msg.GetMsg(tcode.QUERY_ERROR)
		return
	}
	data := make(map[string]interface{})
	data["orderInfo"] = order
	responseData.Data = data
	responseData.Status = true
	responseData.Message = msg.GetMsg(tcode.QUERY_SUCCESS)
	return
}

// 查询用户订单
func QueryByOrderUserID(token string, pageSize int, page int) (responseData entity.ResponseData) {
	user := models.User{Token: token}
	if !user.QueryUser() {
		responseData.Message = msg.GetMsg(tcode.ADD_ERROR) + "，用户不存在"
		return
	}
	order := models.Order{}
	order.UserID = user.ID
	orders := order.QueryByOrderUserID(pageSize, page)
	if len(orders) == 0 {
		responseData.Message = msg.GetMsg(tcode.NOTMORE)
	}
	data := make(map[string]interface{})
	data["orders"] = orders
	responseData.Data = data
	responseData.Status = true
	responseData.Message = msg.GetMsg(tcode.QUERY_SUCCESS)
	return
}

// 查询订单（分页） admin
func QueryOrderByLimitOffset(pageSize int, page int) (responseData entity.ResponseData) {
	orders := models.QueryOrderByLimitOffset(pageSize, page)
	responseData.Message = msg.GetMsg(tcode.QUERY_SUCCESS)
	if len(orders) == 0 {
		responseData.Message = msg.GetMsg(tcode.NOTMORE)
	}
	count := models.QueryOrderCount()
	data := make(map[string]interface{})
	data["orders"] = orders
	data["count"] = count
	responseData.Data = data
	responseData.Status = true
	return
}
