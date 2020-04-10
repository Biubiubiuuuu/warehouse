package models

// 订单商品详情
type OrderGoodsInfo struct {
	Model
	GoodsName   string  `json:"goods_name"`                          // 商品名称
	GoodsSpecs  string  `gorm:"size:2;default:1" json:"goods_specs"` // 商品规格 1.盒 2.瓶 3.支
	GoodsPrince float64 `json:"goods_prince"`                        // 商品单价
	GoodsImage  string  `json:"goods_image"`                         // 商品图片
	GoodsQty    int64   `json:"goods_qty"`                           // 商品购买数量
	GoodsTypeID int64   `json:"goods_type_id"`                       // 商品ID
}
