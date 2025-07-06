package jwt

import (
	"github.com/endymion/go-base/task-04/common/model/response"
	"github.com/endymion/go-base/task-04/common/util"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken := c.GetHeader("x-jwt-token")
		if jwtToken == "" {
			response.NoAuth(c)
			c.Abort()
			return
		}
		claim, err := util.ParseToken(jwtToken)
		if err == nil {
			// 把claim 的 username 放入 c中
			c.Set("username", claim.Username)
			c.Next()
			return
		}
		if !util.IsTokenExpired(err) {
			response.NoAuth(c)
			c.Abort()
			return
		}
		token, err := util.RefreshToken(jwtToken)
		if err != nil {
			response.NoAuth(c)
			c.Abort()
			return
		}
		c.Set("username", claim.Username)
		c.Header("x-jwt-token", token)
		c.Next()
	}
}
