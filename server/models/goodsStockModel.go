package models

import "github.com/Biubiubiuuuu/warehouse/server/dbs/mysql"

// 商品库存
type GoodsStock struct {
	Model
	QuantityStock int64 `json:"quantity_stock"` // 库存数量
	QuantitySold  int64 `json:"quantity_sold"`  // 已售数量
	QuantityGifts int64 `json:"quantity_gifts"` // 已赠送数量
	QuantityTotal int64 `json:"quantity_total"` // 总数量
	GoodsTypeID   int64 `json:"goods_type_id"`  // 商品种类ID
}

// 创建商品库存
func (g *GoodsStock) AddGoodsStock() error {
	db := mysql.GetDB()
	return db.Create(&g).Error
}

// 修改商品库存信息
func (g *GoodsStock) UpdateGoodsStock(args map[string]interface{}) error {
	db := mysql.GetDB()
	return db.Model(&g).Update(args).Error
}

// 查看商品库存详情
func (g *GoodsStock) QueryByGoodsStockID() error {
	db := mysql.GetDB()
	return db.First(&g, g.ID).Error
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
