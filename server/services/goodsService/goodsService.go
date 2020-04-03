package goodsService

import (
	tcode "github.com/Biubiubiuuuu/warehouse/server/common/tips/code"
	"github.com/Biubiubiuuuu/warehouse/server/common/tips/msg"
	"github.com/Biubiubiuuuu/warehouse/server/entity"
	"github.com/Biubiubiuuuu/warehouse/server/models"
)

// 添加商品种类
func AddGoodsType(username string, g entity.GoodsType) (responseData entity.ResponseData) {
	if username == "" {
		responseData.Message = msg.GetMsg(tcode.NOTNULL)
		return
	}
	admin := models.Admin{Username: username}
	if !admin.CheckAdministrator() {
		responseData.Message = msg.GetMsg(tcode.NOT_ADMINISTRATOR)
		return
	}
	goodsType := models.GoodsType{
		GoodsName:  g.GoodsName,
		GoodsSpecs: g.GoodsSpecs,
	}
}
