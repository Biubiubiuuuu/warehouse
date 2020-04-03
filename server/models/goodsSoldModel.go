package models

import "github.com/Biubiubiuuuu/warehouse/server/dbs/mysql"

// 已售商品清单
type GoodsSold struct {
	Model
	GoodsTypeID       int64   `json:"goods_type_id"`        // 商品种类ID
	GoodsSoldPrince   float64 `json:"goods_sold_prince"`    // 成交单价
	GoodsSoldQuantity int64   `json:"goods_sold_quantity"`  // 成交数量
	GoodsSoldCustomer int64   `json:"goods_sold_ customer"` // 买家ID
	GoodsSoldRemarks  int64   `json:"goods_sold_ remarks"`  // 备注信息 例如：买20送2
}

// 创建商品清单
func (g *GoodsSold) AddGoodsSold() error {
	db := mysql.GetDB()
	return db.Create(&g).Error
}

// 查看商品清单（分页查询）
func QueryGoodsSoldByLimitOffset(pageSize int, page int) (goodsSolds []GoodsSold) {
	db := mysql.GetDB()
	db.Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&goodsSolds)
	return
}

// 总记录数
func QueryGoodsSoldCount() (count int) {
	db := mysql.GetDB()
	db.Model(&GoodsSold{}).Count(&count)
	return count
}
