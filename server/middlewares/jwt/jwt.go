package jwt

import (
	"net/http"
	"strings"
	"time"

	tcode "github.com/Biubiubiuuuu/warehouse/server/common/tips/code"
	"github.com/Biubiubiuuuu/warehouse/server/common/tips/msg"
	"github.com/Biubiubiuuuu/warehouse/server/entity"
	"github.com/Biubiubiuuuu/warehouse/server/helpers/jwtHelper"
	"github.com/gin-gonic/gin"
)

// JWT middleware validation
// param query url "token"
// param header "Authorization"
func JWT() gin.HandlerFunc {
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
		claims, err := jwtHelper.ParseToken(token)
		if err != nil {
			code = tcode.TOKEN_AUTH_ERROR
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = tcode.TOKEN_TIMEOUT
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
