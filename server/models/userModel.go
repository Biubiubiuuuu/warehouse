package models

type User struct {
	Model
	Tel          string `json:"tel"`           // 手机号
	Name         string `json:"name"`          // 姓名
	ShopAddress1 string `json:"shop_address1"` // 门店地址1
	ShopAddress2 string `json:"shop_address2"` // 门店地址2
	ShopAddress3 string `json:"shop_address3"` // 门店地址3
}
