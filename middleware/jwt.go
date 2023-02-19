package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/phddy/go-api-server/common"
)

func JwtHandler(c *gin.Context) {
	authorization := c.Request.Header["Authorization"][0]
	if err := common.ValidateToken(authorization); err != nil {
		c.Error(err)
		c.Abort()
	} else {
		c.Next()
	}
}
