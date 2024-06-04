package controller_test

import (
	"errors"
	"searchAPI/controller"
	customerror "searchAPI/global/customError"
	"testing"

	"github.com/gin-gonic/gin"
)

type mockHandler struct{}

func (m *mockHandler) EsSearchAllMovies(query map[string]string) ([]interface{}, customerror.CustomError) {
	// 가짜 데이터 반환
	return []interface{}{"Movie1", "Movie2"}, customerror.CustomError{}
}

func (m *mockHandler) EsSearchNameMovie(query map[string]string) ([]interface{}, customerror.CustomError) {
	if query["search"] == "fail" {
		return nil, customerror.CustomError{Code: 404, Cerror: errors.New("movie not found")}
	}
	return []interface{}{"SearchMovie1"}, customerror.CustomError{}
}

func Test_SearchMovie(t *testing.T) {
	gin.SetMode(gin.TestMode)

	c := controller.MovieController{}
	mock := &mockHandler{}

}
