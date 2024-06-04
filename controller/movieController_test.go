package controller_test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"searchAPI/controller"
	customerror "searchAPI/global/customError"
	"searchAPI/handler"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockHandler struct{}

func (m *mockHandler) EsSearchAllMovies(query map[string]string) ([]interface{}, customerror.CustomError) {
	// 가짜 데이터 반환
	return []interface{}{"Movie1", "Movie2"}, customerror.CustomError{}
}

func (m *mockHandler) EsSearchNameMovie(query map[string]string) ([]interface{}, customerror.CustomError) {
	if query["search"] == "fail" {
		return nil, customerror.CustomError{Code: 404, Cerror: customerror.ErrNotFound}
	}
	return []interface{}{"Movie1", "Movie2"}, customerror.CustomError{}
}

func Test_SearchMovie(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var mock handler.MovieHandler = &mockHandler{}

	c := controller.MovieController{Handler: mock}

	pathNQuery := []string{"/movies", "/movies?search='movie'", "/movies?search='movie'&page=2", "/movies?search=fail"}

	for _, target := range pathNQuery {
		log.Println(target, "테스트 시작")
		req, _ := http.NewRequest(http.MethodGet, target, nil)

		// gin Context 생성
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req
		// 요청을 수행할 핸들러 메서드 호출
		c.SearchMovie(ctx)

		// 응답 확인
		assert.Equal(t, http.StatusOK, w.Code)

		// 예상된 JSON 결과
		var expected string
		if target == "/movies?search=fail" {
			expected = `{"desc":"NOT FOUND", "result":null, "status":404}`
		} else {
			expected = `{"desc":"", "result":["Movie1", "Movie2"], "status":200}`
		}

		// 응답 JSON 결과
		var expectedJSON map[string]interface{}
		var actualJSON map[string]interface{}

		// JSON 파싱
		err := json.Unmarshal([]byte(expected), &expectedJSON)
		assert.NoError(t, err) // err 가 nil 인지 확인, JSON 디코딩 과정에서 오류 발생시 테스트 실패

		err = json.Unmarshal(w.Body.Bytes(), &actualJSON)
		assert.NoError(t, err)

		// JSON 객체 비교
		assert.Equal(t, expectedJSON, actualJSON)
	}
}
