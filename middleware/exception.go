package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/phddy/go-api-server/common"
	"net/http"
)

var logger = common.GetLogger()

func ErrorHandler(c *gin.Context) {
	c.Next()

	if c.Errors != nil {
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case *common.UnauthorizedError:
				c.AbortWithStatusJSON(http.StatusUnauthorized, e.Json())
			case *common.BaseError:
				c.AbortWithStatusJSON(http.StatusInternalServerError, e.Json())
			default:
				logger.Error(e.Error())
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": e.Error(),
				})
			}
		}
	}
}
