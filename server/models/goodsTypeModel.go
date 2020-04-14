package models

import (
	"errors"
	"time"

	"github.com/Biubiubiuuuu/warehouse/server/dbs/mysql"
)

// 商品种类
type GoodsType struct {
	Model
	GoodsName           string       `gorm:"not null;unique;size:50" json:"goods_name"`                            // 商品名称
	GoodsSpecs          string       `gorm:"size:2;default:1" json:"goods_specs"`                                  // 商品规格 1.盒 2.瓶 3.支
	GoodsUnitPrince     float64      `json:"goods_unitprince"`                                                     // 商品成本价
	GoodsPrince         float64      `json:"goods_prince"`                                                         // 商品销售价
	GoodsDiscountPrince float64      `json:"goods_discount_prince"`                                                // 商品折扣价
	GoodsImages         []GoodsImage `gorm:"foreignkey:GoodsTypeID;association_foreignkey:ID" json:"goods_images"` // 商品图片
	GoodsBatchNumber    string       `gorm:"size:50" json:"goods_batch_number"`                                    // 生产批号
	GoodsDate           time.Time    `json:"goods_date"`                                                           // 生产日期
	GoodsState          string       `gorm:"size:2;default:2" json:"goods_state"`                                  // 商品状态 1.下架  2.在售
	GoodsCreateAdmin    string       `gorm:"not null;" json:"goods_create_aAdmin"`                                 // 创建人
}

type GoodsImage struct {
	ID              int64
	GoodsImageFiles string `json:"goods_image_files"` // 商品图片路径
	GoodsTypeID     int64  `json:"goods_type_id"`     // 商品种类ID
}

type GoodsTypeData struct {
	ID        int64  `json:"id"`
	GoodsName string `json:"goods_name"` // 商品名称
}

// 添加商品种类
func (g *GoodsType) AddGoodsType() error {
	db := mysql.GetDB()
	return db.Create(&g).Error
}

// 修改商品种类信息
func (g *GoodsType) UpdateGoodsType(goodsImage []GoodsImage, args map[string]interface{}) error {
	db := mysql.GetDB()
	db.Model(&g).Association("GoodsImages").Replace(goodsImage)
	return db.Model(&g).Update(args).Error
}

// 查看商品详情
// param ID
// return GoodsType,error
func (g *GoodsType) QueryByGoodsTypeID() error {
	db := mysql.GetDB()
	return db.First(&g).Model(&g).Related(&g.GoodsImages).Find(&g).Error
}

// 下架商品
// param ID
// return error
func (g *GoodsType) DeleteGoodsTypes(ids []int64) error {
	db := mysql.GetDB()
	tx := db.Begin()
	for _, id := range ids {
		if id == 0 {
			return errors.New("id is not 0")
		}
		g.ID = id
		tx.Model(&g).Update("goods_state", "1")
		tx.Model(&g).Association("GoodsImages").Clear()
		if err := tx.Delete(&g).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

// 查询商品种类ID和商品名（支持模糊查询）
// param GoodsName string 商品名
// return []GoodsTypeData
func (g *GoodsType) QueryAllGoods() (goodsTypeDatas []GoodsTypeData) {
	db := mysql.GetDB()
	db.Table("goods_type").Select("id, goods_name").Where("goods_name LIKE ?", "%"+g.GoodsName+"%").Scan(&goodsTypeDatas)
	return
}

// 查看商品种类（分页查询）
// param pageSize int
// param page int
func QueryGoodsTypesByLimitOffset(pageSize int, page int) (goodsTypes []GoodsType) {
	db := mysql.GetDB()
	db.Preload("GoodsImages").Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&goodsTypes)
	return
}

// 商品总记录数
func QueryGoodsTypesCount() (count int) {
	db := mysql.GetDB()
	db.Model(&GoodsType{}).Count(&count)
	return count
}
