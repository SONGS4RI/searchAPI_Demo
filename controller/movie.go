package controller

import (
	"net/http"
	"searchAPI/handler"
	"searchAPI/response"

	"github.com/gin-gonic/gin"
)

var (
	page   = "page"
	search = "search"
)

func (c *Controller) SearchMovie(ctx *gin.Context) {
	response := response.Response{}
	query := make(map[string]string)
	query[page] = ctx.DefaultQuery(page, "1") // 검색 시작 위치 지정
	query[search] = ctx.Query(search)

	var queryManager func(map[string]string) (result []interface{}, err error)

	if query["search"] == "" { // 검색어 기반 검색이 아닌 경우
		queryManager = handler.EsSearchAllMovies
	} else { // 검색어가 있는 경우
		queryManager = handler.EsSearchNameMovie
	}

	if result, err := queryManager(query); err != nil {
		response.Status = http.StatusInternalServerError
		response.Desc = err.Error()
	} else {
		response.Status = http.StatusOK
		response.Result = result
	}

	ctx.IndentedJSON(http.StatusOK, response)
}
