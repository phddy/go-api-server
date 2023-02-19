package user

import "github.com/gin-gonic/gin"

func cookie(c *gin.Context) {
	print(c)
}

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user",
	})
}

func Router(r *gin.RouterGroup) {
	router := r.Group("/user")
	router.Use(cookie)
	router.GET("", index)
}
