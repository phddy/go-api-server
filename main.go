package main

import (
	"github.com/gin-gonic/gin"
	"github.com/phddy/go-api-server/middleware"
	"github.com/phddy/go-api-server/route"
)

func main() {
	r := gin.New()
	r.Use(middleware.ErrorHandler)
	//r.Use(common.Logger())
	route.Router(r)
	r.Run("0.0.0.0:3000")
}
