package models

import "github.com/Biubiubiuuuu/warehouse/server/dbs/mysql"

// 商品订单
type GoodsOrder struct {
	Model
	GoodsOrderPrince     float64 `json:"goods_order_prince"`                         // 成交单价
	GoodsOrderQuantity   int64   `json:"goods_order_quantity"`                       // 成交数量
	GoodsOrderRemarks    int64   `json:"goods_order_remarks"`                        // 备注信息 例如：买20送2
	GoodsOrderStatus     string  `gorm:"size:2;default:1" json:"goods_order_status"` // 订单状态 1.已付款 2.未付款
	GoodsOrderCustomerID int64   `json:"goods_order_customer_id"`                    // 买家ID
	GoodsTypeID          int64   `json:"goods_type_id"`                              // 商品种类ID
}

// 创建商品订单
func (g *GoodsOrder) AddGoodsOrder() error {
	db := mysql.GetDB()
	return db.Create(&g).Error
}

// 查看商品订单详情
func (g *GoodsOrder) QueryByGoodsOrderID() error {
	db := mysql.GetDB()
	return db.First(&g, g.ID).Error
}

// 查看商品订单（分页查询）
func QueryGoodsOrderByLimitOffset(pageSize int, page int) (goodsOrders []GoodsOrder) {
	db := mysql.GetDB()
	db.Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&goodsOrders)
	return
}

// 总记录数
func QueryGoodsOrderCount() (count int) {
	db := mysql.GetDB()
	db.Model(&GoodsOrder{}).Count(&count)
	return count
}
