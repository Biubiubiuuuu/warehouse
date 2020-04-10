package models

// 订单用户详情
type OrderUserInfo struct {
	Model
	Provice     string `json:"provice"`       // 省
	City        string `json:"city"`          // 城市
	ShopAddress string `json:"shop_address1"` // 门店详细地址
	Tel         string `json:"tel"`           // 联系电话
}
