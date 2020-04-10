package adminAuth

import (
	"net/http"
	"strings"

	tcode "github.com/Biubiubiuuuu/warehouse/server/common/tips/code"
	"github.com/Biubiubiuuuu/warehouse/server/common/tips/msg"
	"github.com/Biubiubiuuuu/warehouse/server/entity"
	"github.com/Biubiubiuuuu/warehouse/server/models"
	"github.com/gin-gonic/gin"
)

// 根据token查询对应账号信息验证当前登录用户或者外部调用 是否有操作curd的权限
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := tcode.SUCCESS
		token := c.Query("token")
		if token == "" {
			authToken := c.GetHeader("Authorization")
			if authToken == "" {
				code = tcode.AUTH_NOT_BEARER
			}
			token = strings.TrimSpace(authToken)
		}
		admin := models.Admin{Token: token}
		user := models.User{Token: token}
		if !admin.CheckAdministrator() && !user.QueryUser() {
			code = tcode.NOT_ADMINISTRATOR
		}
		if code != tcode.SUCCESS {
			message := msg.GetMsg(code)
			responseJson := entity.ResponseJson(false, nil, message)
			c.AbortWithStatusJSON(http.StatusUnauthorized, responseJson)
			return
		}
		c.Next()
	}
}
