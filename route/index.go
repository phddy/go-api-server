package route

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/gin-gonic/gin"
	"github.com/phddy/go-api-server/common"
	"github.com/phddy/go-api-server/middleware"
	"github.com/phddy/go-api-server/model"
	"github.com/phddy/go-api-server/route/user"
	"net/http"
)

func Router(r *gin.Engine) {
	router := r.Group("/")
	router.Use(middleware.JwtHandler)
	user.Router(router)
	r.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	r.POST("/login", login)
	r.GET("/logout", middleware.JwtHandler, logout)
	r.POST("/join", join)
}

var o = orm.NewOrm()

func join(c *gin.Context) {
	var user *model.User
	c.BindJSON(&user)

	user.EncryptPassword()
	if _, err := o.Insert(user); err != nil {
		c.Error(err)
		return
	}

	token, _ := common.CreateToken(user)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func login(c *gin.Context) {
	var user *model.User
	c.BindJSON(&user)

	user.EncryptPassword()

	if err := o.Read(user, "email", "password"); err != nil {
		c.Error(&common.UnauthorizedError{
			Message: fmt.Sprintf("email(%s) user not found or wrong password", user.Email),
		})
		return
	}

	token, _ := common.CreateToken(user)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func logout(c *gin.Context) {
	authorization := c.Request.Header["Authorization"][0]
	if err := common.RemoveToken(authorization); err != nil {
		c.Error(err)
		c.Abort()
	}
}
