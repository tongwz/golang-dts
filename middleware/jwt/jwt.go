package jwt

import (
	"golang-dts/pkg/e"
	"golang-dts/pkg/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// jwt gin中间件 在gin访问时就会访问这个中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)

			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": "token验证失败",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
