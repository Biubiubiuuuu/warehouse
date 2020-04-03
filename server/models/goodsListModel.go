package models

import "github.com/Biubiubiuuuu/warehouse/server/dbs/mysql"

// 商品库存列表
type GoodsList struct {
	Model
	Stock         int64 `json:"stock"`          // 库存数量
	QuantitySold  int64 `json:"quantity_sold"`  // 已售数量
	QuantityGifts int64 `json:"quantity_gifts"` // 已赠送数量
	QuantityTotal int64 `json:"quantity_total"` // 总数量
	GoodsTypeID   int64 `json:"goods_type_id"`  // 商品种类ID
}

// 创建商品清单
func (g *GoodsList) AddGoodsList() error {
	db := mysql.GetDB()
	return db.Create(&g).Error
}

// 修改商品库存信息
func (g *GoodsList) UpdateGoodsList(args map[string]interface{}) error {
	db := mysql.GetDB()
	if err := db.Model(&g).Update(args).Error; err != nil {
		return err
	}
	return nil
}

// 查看商品库存（分页查询）
func QueryGoodsListByLimitOffset(pageSize int, page int) (goodsLists []GoodsList) {
	db := mysql.GetDB()
	db.Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&goodsLists)
	return
}

// 总记录数
func QueryGoodsListCount() (count int) {
	db := mysql.GetDB()
	db.Model(&GoodsList{}).Count(&count)
	return count
}
