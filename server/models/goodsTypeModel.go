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
	GoodsName           string    `gorm:"not null;size:50" json:"goods_name"`   // 商品名称
	GoodsSpecs          string    `gorm:"size:2;default:1" json:"goods_specs"`  // 商品规格 1.盒 2.瓶 3.支
	GoodsUnitPrince     float64   `json:"goods_unitprince"`                     // 商品成本价
	GoodsPrince         float64   `json:"goods_prince"`                         // 商品销售价
	GoodsDiscountPrince float64   `json:"goods_discount_prince"`                // 商品折扣价
	GoodsImage          string    `json:"goods_image"`                          // 商品图片
	GoodsBatchNumber    string    `gorm:"size:50" json:"goods_batch_number"`    // 生产批号
	GoodsDate           time.Time `json:"goods_date"`                           // 生产日期
	GoodsState          string    `gorm:"size:2;default:2" json:"goods_state"`  // 商品状态 1.下架  2.在售
	GoodsCreateAdmin    string    `gorm:"not null;" json:"goods_create_aAdmin"` // 创建人
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
	return db.Model(&g).Update(args).Error
}

// 查看商品详情
func (g *GoodsType) QueryByGoodsTypeID() error {
	db := mysql.GetDB()
	return db.First(&g, g.ID).Error
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

// 查询所有商品种类（支持模糊查询）
func (g *GoodsType) QueryAllGoodsTypes() (goodsTypes []GoodsType) {
	db := mysql.GetDB()
	db.Select("ID, GoodsName").Where("GoodsName LIKE ?", "%"+g.GoodsName+"%").Find(&goodsTypes)
	return
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
