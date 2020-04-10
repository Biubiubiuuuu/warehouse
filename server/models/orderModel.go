package models

import (
	"github.com/Biubiubiuuuu/warehouse/server/dbs/mysql"
)

// 商品订单
type Order struct {
	Model
	UserID          int64            `json:"user_id"`                                               // 用户ID
	OrderNumber     string           `json:"order_number"`                                          // 订单编号
	OrderPrince     float64          `json:"order_prince"`                                          // 成交总金额
	OrderStatus     string           `gorm:"size:2;default:1" json:"order_status"`                  // 订单状态 1.已付款 2.未付款
	OrderGoodsInfos []OrderGoodsInfo `gorm:"many2many:order_goods_infos;" json:"order_goods_infos"` // 订单商品信息
	OrderUserInfo   OrderUserInfo    `json:"order_user_info"`                                       // 订单用户信息
	OrderUserInfoID int64            `json:"-"`
}

// 创建商品订单
func (g *Order) AddOrder() error {
	db := mysql.GetDB()
	tx := db.Begin()
	if err := tx.Create(&g).Error; err != nil {
		tx.Rollback()
		return err
	}
	for _, info := range g.OrderGoodsInfos {
		goodsStock := GoodsStock{}
		goodsStock.GoodsTypeID = info.GoodsTypeID
		if _, err := goodsStock.QueryByID(); err != nil {
			tx.Rollback()
			return err
		}
		quantityStock := goodsStock.QuantityStock - info.GoodsQty
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

// 查看商品订单详情
func (g *Order) QueryByGoodsOrderID() error {
	db := mysql.GetDB()
	return db.First(&g, g.ID).Error
}

// 查看商品订单（分页查询）
func QueryOrderByLimitOffset(pageSize int, page int) (orders []Order) {
	db := mysql.GetDB()
	db.Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&orders)
	return
}

// 总记录数
func QueryOrderCount() (count int) {
	db := mysql.GetDB()
	db.Model(&Order{}).Count(&count)
	return count
}
