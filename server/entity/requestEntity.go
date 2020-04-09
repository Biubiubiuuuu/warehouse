package entity

// 管理员登录
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 管理员注册
type Register struct {
	Username      string `json:"username"`      // 用户名
	Password      string `json:"password"`      // 密码
	Administrator string `json:"administrator"` // 超级管理员 Y | N
}

// 管理员修改密码
type UpdatePass struct {
	OldPassword string `json:"oldPassword"` // 旧密码
	Username    string `json:"username"`    // 用户名
	NewPassword string `json:"newPassword"` // 密码
}

// 删除ids
type DeleteIds struct {
	Ids []int64 `json:"ids"` // ids
}

// 添加商品种类
type AddGoodsType struct {
	GoodsName        string  `json:"goods_name"`                            // 商品名称
	GoodsSpecs       string  `json:"goods_specs" enums:"1,2,3" default:"1"` // 商品规格 1.盒 2.瓶 3.支
	GoodsUnitPrince  float64 `json:"goods_unitprince"`                      // 商品成本价
	GoodsPrince      float64 `json:"goods_prince"`                          // 商品销售价
	GoodsImage       string  `json:"goods_image"`                           // 商品图片
	GoodsBatchNumber string  `json:"goods_batch_number"`                    // 生产批号
	GoodsDate        string  `json:"goods_date"`                            // 生产日期
	GoodsState       string  `json:"goods_state" enums:"1,2,3" default:"2"` // 商品状态 1.下架  2.在售
	GoodsCreateAdmin string  `json:"goods_create_admin"`                    // 创建人
}

// 修改商品种类
type UpdateGoodsType struct {
	GoodsID          int64   `json:"goods_id"`                              // 商品ID
	GoodsName        string  `json:"goods_name"`                            // 商品名称
	GoodsSpecs       string  `json:"goods_specs" enums:"1,2,3" default:"1"` // 商品规格 1.盒 2.瓶 3.支
	GoodsUnitPrince  float64 `json:"goods_unitprince"`                      // 商品成本价
	GoodsPrince      float64 `json:"goods_prince"`                          // 商品销售价
	GoodsImage       string  `json:"goods_image"`                           // 商品图片
	GoodsBatchNumber string  `json:"goods_batch_number"`                    // 生产批号
	GoodsDate        string  `json:"goods_date"`                            // 生产日期
	GoodsState       string  `json:"goods_state" enums:"1,2,3" default:"2"` // 商品状态 1.下架  2.在售
}

// 添加商品库存
type AddGoodsStock struct {
	QuantityTotal int64 `json:"quantity_total"` // 总数量
	GoodsTypeID   int64 `json:"goods_type_id"`  // 商品种类ID
}

// 修改商品库存信息
type UpdateGoodsStock struct {
	GoodsStockID int64 `json:"goods_stock_id"` // 库存ID
	AddQuantity  int64 `json:"add_quantity"`   // 添加库存数量
}

// 用户登录
type UserLogin struct {
	Tel      string `json:"tel"`
	Password string `json:"password"`
}

// 用户注册
type UserRegister struct {
	Tel      string `json:"tel"`      // 用户名
	Password string `json:"password"` // 密码
}

// 用户修改密码
type UserUpdatePass struct {
	OldPassword string `json:"oldPassword"` // 旧密码
	Tel         string `json:"tel"`         // 用户名
	NewPassword string `json:"newPassword"` // 密码
}
