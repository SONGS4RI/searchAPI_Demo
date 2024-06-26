package router

import (
	"searchAPI/controller"

	"github.com/gin-gonic/gin"
)

func (r *Router) MovieRouter(router *gin.RouterGroup) {
	c := controller.NewMovieController()

	router.GET("", c.SearchMovie)
}
