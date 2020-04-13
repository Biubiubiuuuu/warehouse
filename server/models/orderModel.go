package models

import (
	"github.com/Biubiubiuuuu/warehouse/server/dbs/mysql"
)

// 商品订单
type Order struct {
	Model
	UserID            int64              `json:"user_id"`                                                                 // 用户ID
	OrderNumber       string             `json:"order_number"`                                                            // 订单编号
	OrderPrince       float64            `json:"order_prince"`                                                            // 成交总金额
	OrderStatus       string             `gorm:"size:2;default:2" json:"order_status"`                                    // 订单状态 1.已付款 2.未付款
	OrderGoodsDetails []OrderGoodsDetail `gorm:"foreignkey:OrderID;association_foreignkey:ID" json:"order_goods_details"` // 订单商品信息
	OrderUserDetail   OrderUserDetail    `json:"order_user_details"`                                                      // 订单用户信息
	OrderUserDetailID int64              `json:"order_user_detail_id"`                                                    // 收件人地址ID
}

// 订单商品详情
type OrderGoodsDetail struct {
	ID          int64
	GoodsName   string  `json:"goods_name"`                          // 商品名称
	GoodsSpecs  string  `gorm:"size:2;default:1" json:"goods_specs"` // 商品规格 1.盒 2.瓶 3.支
	GoodsPrince float64 `json:"goods_prince"`                        // 商品单价
	GoodsImage  string  `json:"goods_image"`                         // 商品图片
	GoodsQty    int64   `json:"goods_qty"`                           // 商品购买数量
	GoodsTypeID int64   `json:"goods_type_id"`                       // 商品ID
	OrderID     int64   `gorm:"INDEX" json:"order_id"`               // 订单ID
}

// 订单收件地址详情
type OrderUserDetail struct {
	ID          int64
	Provice     string `json:"provice"`       // 省
	City        string `json:"city"`          // 城市
	ShopAddress string `json:"shop_address1"` // 门店详细地址
	Tel         string `json:"tel"`           // 联系电话
}

// 创建商品订单
func (g *Order) AddOrder() error {
	db := mysql.GetDB()
	tx := db.Begin()
	if err := tx.Create(&g).Error; err != nil {
		tx.Rollback()
		return err
	}
	for _, item := range g.OrderGoodsDetails {
		goodsStock := GoodsStock{}
		goodsStock.GoodsTypeID = item.GoodsTypeID
		var goodsStockData GoodsStockData
		var queryErr error
		if goodsStockData, queryErr = goodsStock.QueryByID(); queryErr != nil {
			tx.Rollback()
			return queryErr
		}
		quantityStock := goodsStockData.QuantityStock - item.GoodsQty
		quantitySold := goodsStockData.QuantitySold + item.GoodsQty
		args := map[string]interface{}{
			"QuantityStock": quantityStock,
			"QuantitySold":  quantitySold,
		}
		if err := goodsStock.UpdateGoodsStock(args); err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

// 付款订单
func (g *Order) UpdateOrderStatus() error {
	db := mysql.GetDB()
	return db.Model(&g).Updates(map[string]interface{}{"order_status": "1"}).Error
}

// 查看商品订单详情
func (g *Order) QueryByGoodsOrderID() error {
	db := mysql.GetDB()
	err := db.First(&g).Model(&g).Related(&g.OrderUserDetail).Related(&g.OrderGoodsDetails).Find(&g).Error
	return err
}

// 查看用户订单
func (g *Order) QueryByOrderUserID(pageSize int, page int) (orders []Order) {
	db := mysql.GetDB()
	db.Where("user_id = ?", g.UserID).Preload("OrderUserDetail").Preload("OrderGoodsDetails").Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&orders)
	return
}

// 查看商品订单（分页查询）
func QueryOrderByLimitOffset(pageSize int, page int) (orders []Order) {
	db := mysql.GetDB()
	db.Preload("OrderUserDetail").Preload("OrderGoodsDetails").Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&orders)
	return
}

// 总记录数
func QueryOrderCount() (count int) {
	db := mysql.GetDB()
	db.Model(&Order{}).Count(&count)
	return count
}
