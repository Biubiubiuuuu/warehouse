package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/Biubiubiuuuu/warehouse/server/common/tips/code"
	"github.com/Biubiubiuuuu/warehouse/server/common/tips/msg"
	"github.com/Biubiubiuuuu/warehouse/server/dbs/mysql"
)

// 商品种类
type GoodsType struct {
	Model
	GoodsName           string    `gorm:"not null;unique;size:255" json:"goods_name"` // 商品名称
	GoodsSpecs          string    `gorm:"size:20" json:"goods_specs"`                 // 商品规格 1.盒 2.瓶 3.支
	GoodsUnitPrince     float64   `json:"goods_unitprince"`                           // 商品成本价
	GoodsPrince         float64   `json:"goods_prince"`                               // 商品销售价
	GoodsDiscountPrince float64   `json:"goods_discount_prince"`                      // 商品折扣价
	GoodsImage          string    `json:"goods_image"`                                // 商品图片
	GoodsBatchNumber    string    `gorm:"size:50" json:"goods_batch_number"`          // 生产批号
	GoodsDate           time.Time `json:"goods_date"`                                 // 生产日期
	GoodsState          string    `json:"goods_state"`                                // 商品状态 1.下架  2.在售
}

// 添加商品种类
func (g *GoodsType) AddGoodsType() error {
	db := mysql.GetDB()
	if db.NewRecord(g.GoodsName) {
		return db.Create(&g).Error
	}
	return errors.New(fmt.Sprintf("%v%v", g.GoodsName, msg.GetMsg(code.EXIST)))
}

// 修改商品种类信息
func (g *GoodsType) UpdateGoodsType(args map[string]interface{}) error {
	db := mysql.GetDB()
	if err := db.Model(&g).Update(args).Error; err != nil {
		return err
	}
	return nil
}

// 下架商品
func (g *GoodsType) DeleteGoodsTypes(ids []int64) error {
	db := mysql.GetDB()
	tx := db.Begin()
	for _, id := range ids {
		if id == 0 {
			return errors.New("id is not 0")
		}
		g.ID = id
		if err := tx.Delete(&g).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

// 查看商品种类（分页查询）
func QueryGoodsTypesByLimitOffset(pageSize int, page int) (goodsTypes []GoodsType) {
	db := mysql.GetDB()
	db.Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&goodsTypes)
	return
}

// 总记录数
func QueryGoodsTypesCount() (count int) {
	db := mysql.GetDB()
	db.Model(&GoodsType{}).Count(&count)
	return count
}
