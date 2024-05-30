package router

import (
	"searchAPI/controller"

	"github.com/gin-gonic/gin"
)

func(r * Router) TestRouter(router *gin.RouterGroup) {
	c := controller.NewController()

	router.GET("/test", c.TestController)
}