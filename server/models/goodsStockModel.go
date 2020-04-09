package models

import (
	"errors"
	"time"

	"github.com/Biubiubiuuuu/warehouse/server/common/tips/code"
	"github.com/Biubiubiuuuu/warehouse/server/common/tips/msg"
	"github.com/Biubiubiuuuu/warehouse/server/dbs/mysql"
)

// 商品库存
type GoodsStock struct {
	Model
	QuantityStock int64 `json:"quantity_stock"`                        // 库存数量
	QuantitySold  int64 `json:"quantity_sold"`                         // 已售数量
	QuantityGifts int64 `json:"quantity_gifts"`                        // 已赠送数量
	QuantityTotal int64 `json:"quantity_total"`                        // 总数量
	GoodsTypeID   int64 `gorm:"not null;unique;" json:"goods_type_id"` // 商品种类ID
}

// 商品库存详情映射struct
type GoodsStockData struct {
	ID            int64      `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	QuantityStock int64      `json:"quantity_stock"` // 库存数量
	QuantitySold  int64      `json:"quantity_sold"`  // 已售数量
	QuantityGifts int64      `json:"quantity_gifts"` // 已赠送数量
	QuantityTotal int64      `json:"quantity_total"` // 总数量
	GoodsName     string     `json:"goods_name"`     // 商品名称
}

// 创建商品库存
func (g *GoodsStock) AddGoodsStock() error {
	db := mysql.GetDB()
	if db.NewRecord(g.GoodsTypeID) {
		return db.Create(&g).Error
	}
	return errors.New(msg.GetMsg(code.STOCK_EXIST))
}

// 修改商品库存信息
func (g *GoodsStock) UpdateGoodsStock(args map[string]interface{}) error {
	db := mysql.GetDB()
	return db.Model(&g).Updates(args).Error
}

// 查看商品库存详情
func (g *GoodsStock) QueryByGoodsStockID() (goodsStockData GoodsStockData, err error) {
	db := mysql.GetDB()
	err = db.Table("goods_stock").Select("goods_stock.id, goods_stock.created_at, goods_stock.updated_at, goods_stock.deleted_at, goods_stock.quantity_stock, goods_stock.quantity_sold, goods_stock.quantity_total, goods_stock.quantity_gifts, goods_type.goods_name").Joins("left join goods_type on goods_type.id = goods_stock.id").Where("goods_type.id = ?", g.ID).Scan(&goodsStockData).Error
	return goodsStockData, err
}

// 查看商品库存（分页查询）
func QueryGoodsStockByLimitOffset(pageSize int, page int) (goodsStockDatas []GoodsStockData) {
	db := mysql.GetDB()
	db.Table("goods_stock").Select("goods_stock.id, goods_stock.created_at, goods_stock.updated_at, goods_stock.deleted_at, goods_stock.quantity_stock, goods_stock.quantity_sold, goods_stock.quantity_total, goods_stock.quantity_gifts, goods_type.goods_name").Joins("left join goods_type on goods_type.id = goods_stock.id").Limit(pageSize).Offset((page - 1) * pageSize).Order("goods_stock.created_at desc").Scan(&goodsStockDatas)
	return
}

// 总记录数
func QueryGoodsStockCount() (count int) {
	db := mysql.GetDB()
	db.Model(&GoodsStock{}).Count(&count)
	return count
}
