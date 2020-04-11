package models

import (
	"fmt"

	"github.com/Biubiubiuuuu/warehouse/server/dbs/mysql"
)

// 商品订单
type Order struct {
	Model
	UserID          int64            `json:"user_id"`                                              // 用户ID
	OrderNumber     string           `json:"order_number"`                                         // 订单编号
	OrderPrince     float64          `json:"order_prince"`                                         // 成交总金额
	OrderStatus     string           `gorm:"size:2;default:2" json:"order_status"`                 // 订单状态 1.已付款 2.未付款
	OrderGoodsInfos []OrderGoodsInfo `gorm:"many2many:order_goods_infos" json:"order_goods_infos"` // 订单商品信息
	OrderUserInfo   OrderUserInfo    `gorm:"foreignkey:OrderUserInfoID" json:"order_user_info"`    // 订单用户信息
	OrderUserInfoID int64            `json:"-"`
}

// 订单商品详情
type OrderGoodsInfo struct {
	ID          int64
	GoodsName   string  `json:"goods_name"`                          // 商品名称
	GoodsSpecs  string  `gorm:"size:2;default:1" json:"goods_specs"` // 商品规格 1.盒 2.瓶 3.支
	GoodsPrince float64 `json:"goods_prince"`                        // 商品单价
	GoodsImage  string  `json:"goods_image"`                         // 商品图片
	GoodsQty    int64   `json:"goods_qty"`                           // 商品购买数量
	GoodsTypeID int64   `json:"goods_type_id"`                       // 商品ID
	OrderId     int64   `gorm:"INDEX" json:"order_id"`               // 订单ID
}

// 订单用户详情
type OrderUserInfo struct {
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
	for _, item := range g.OrderGoodsInfos {
		goodsStock := GoodsStock{}
		goodsStock.GoodsTypeID = item.GoodsTypeID
		var goodsStockData GoodsStockData
		var queryErr error
		if goodsStockData, queryErr = goodsStock.QueryByID(); queryErr != nil {
			tx.Rollback()
			return queryErr
		}
		quantityStock := goodsStockData.QuantityStock - item.GoodsQty
		args := map[string]interface{}{
			"QuantityStock": quantityStock,
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
	var orderInfo Order
	db.Where("id = ?", g.ID).First(&orderInfo)
	err := db.Model(&orderInfo).Related(&orderInfo.OrderUserInfo).Related(&orderInfo.OrderGoodsInfos).Find(&orderInfo).Error
	return err
}

// 查看用户订单
func (g *Order) QueryByOrderUserID(pageSize int, page int) (orders []Order) {
	var list []Order
	db := mysql.GetDB()
	err := db.Preload("OrderGoodsInfos").Preload("OrderUserInfo").Where("user_id = ?", g.UserID).Limit(pageSize).Offset((page - 1) * pageSize).Find(&list).Error
	fmt.Println(err.Error())
	return list
}

// 查看商品订单（分页查询）
func QueryOrderByLimitOffset(pageSize int, page int) (orders []Order) {
	var list []Order
	db := mysql.GetDB()
	db.Preload("OrderUserInfo").Preload("OrderGoodsInfo").Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&list)
	return list
}

// 总记录数
func QueryOrderCount() (count int) {
	db := mysql.GetDB()
	db.Model(&Order{}).Count(&count)
	return count
}
