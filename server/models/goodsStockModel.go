package models

import (
	"errors"

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
	GoodsStock
	GoodsName string `json:"goods_name"` // 商品名称
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
	return db.Model(&g).Update(args).Error
}

// 查看商品库存详情
func QueryByGoodsStockID(id int64) (result GoodsStockData, err error) {
	db := mysql.GetDB()
	err = db.Table("goods_stock").Select("goods_stock.id, goods_stock.created_at, goods_stock.updated_at, goods_stock.deleted_at, goods_stock.quantity_stock, goods_stock.quantity_sold, goods_stock.quantity_total, goods_type.goods_name").Joins("left join goods_type on goods_type.id = goods_stock.id").Where("goods_type.id = ?", id).Scan(&result).Error
	return result, err
}

// 查看商品库存（分页查询）
func QueryGoodsStockByLimitOffset(pageSize int, page int) (GoodsStocks []GoodsStock) {
	db := mysql.GetDB()
	db.Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&GoodsStocks)
	return
}

// 总记录数
func QueryGoodsStockCount() (count int) {
	db := mysql.GetDB()
	db.Model(&GoodsStock{}).Count(&count)
	return count
}
