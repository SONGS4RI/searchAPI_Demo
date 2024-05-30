package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) TestController(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}