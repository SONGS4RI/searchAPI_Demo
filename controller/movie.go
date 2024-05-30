package controller

import (
	"net/http"
	"searchAPI/handler"
	"searchAPI/response"

	"github.com/gin-gonic/gin"
)

func (c *Controller) GetMovies(ctx *gin.Context) {
	response := response.Response{}
	query := make(map[string]string)
	query["page"] = ctx.DefaultQuery("page", "1") // 검색 시작 위치 지정

	if result, err := handler.EsGetAllMovies(query); err != nil {
		response.Status = http.StatusInternalServerError
		response.Desc = err.Error()
	} else {
		response.Status = http.StatusOK
		response.Result = result
	}

	ctx.IndentedJSON(http.StatusOK, response)
}
