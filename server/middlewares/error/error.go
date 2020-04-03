package error

import (
	"net/http"

	"github.com/Biubiubiuuuu/warehouse/server/common/tips/code"
	"github.com/Biubiubiuuuu/warehouse/server/common/tips/msg"
	"github.com/Biubiubiuuuu/warehouse/server/entity"
	"github.com/gin-gonic/gin"
)

// 404
func NotFound(c *gin.Context) {
	code := code.NOTFOUND
	msg := msg.GetMsg(code)
	response := entity.ResponseJson(false, nil, msg)
	c.JSON(http.StatusNotFound, response)
}
