package controller

import (
	"log"
	"net/http"
	customerror "searchAPI/global/customError"
	"searchAPI/handler"
	"searchAPI/response"

	"github.com/gin-gonic/gin"
)

type MovieController struct {
	Handler handler.MovieHandler
}

func NewMovieController() *MovieController {
	c := MovieController{Handler: &handler.MovieHandlerImpl{}}
	return &c
}

var (
	page   = "page"
	search = "search"
)

// @Tags movies : movies API
// @Summary 영화명 검색 API
// @Description 영화명 검색 API입니다.
// @Accept  json
// @Produce  json
// @Router /movies [get]
// @Param search query string false "영화 제목"
// @Param page query int false "페이지"
// @Success 200 {object} response.Response
func (c *MovieController) SearchMovie(ctx *gin.Context) {
	response := response.Response{}
	query := make(map[string]string)
	query[page] = ctx.DefaultQuery(page, "1") // 검색 시작 위치 지정
	query[search] = ctx.Query(search)

	var queryManager func(handler.MovieHandler, map[string]string) (result []interface{}, cerr customerror.CustomError)

	if query["search"] == "" { // 검색어 기반 검색이 아닌 경우
		queryManager = handler.MovieHandler.EsSearchAllMovies
	} else { // 검색어가 있는 경우
		queryManager = handler.MovieHandler.EsSearchNameMovie
	}

	if result, err := queryManager(c.Handler, query); err.Cerror != nil {
		log.Println(err)
		response.Status = err.Code
		response.Desc = err.Error()
	} else {
		response.Status = http.StatusOK
		response.Result = result
	}

	ctx.IndentedJSON(http.StatusOK, response)
}
