package handler

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	elasticSearch "github.com/olivere/elastic/v7"

	elasticConn "searchAPI/elasticConn"
	customerror "searchAPI/global/customError"
	"searchAPI/manager"
	"searchAPI/model"
)

var (
	from  int
	size  int    = 20
	index string = "movie_search" // 인덱스
	defaultMovieSource = []string{
		"companies", "companys", "directors", "genreAlt", "movieCd",
		"movieNm", "movieNmEn", "nationAlt", "openDt", "prdtStatNm",
		"prdtYear", "repGenreNm", "repNationNm", "typeNm"}
)

/*
	MatchAllQuery를 사용하는 페이징 검색
*/
func EsSearchAllMovies(param map[string]string) (result []interface{}, cerr customerror.CustomError) {
	// 검색 시작 값
	page, _ := strconv.Atoi(param["page"])

	from = (page - 1) * size

	query := manager.MatchAllQuery()

	fsc := elasticSearch.NewFetchSourceContext(true).Include(defaultMovieSource...)

	client := elasticConn.EsClient
	esSearch := client.Search().
		Query(query). // specify the query
		FetchSourceContext(fsc).
		Index(index). // search in index
		From(from).   // -1 처리(기존 1)
		Size(size).   // take documents 0-9
		Pretty(true). // pretty print request and response JSON
		TrackTotalHits(true).
		Timeout("5s")

	res, err := esSearch.Do(context.Background())

	if err != nil {
		log.Println(err)
		return
	}

	if !isPageAvailable(from, res, &cerr) {
		return
	}

	handleSearchResult(&result, res)

	return
}

/*
	BoolQuery를 사용하는 페이징 검색
*/
func EsSearchNameMovie(param map[string]string) (result []interface{}, cerr customerror.CustomError) {
	// 검색 시작 값
	page, _ := strconv.Atoi(param["page"])

	from = (page - 1) * size

	query := manager.BoolQuery(param["search"])

	fsc := elasticSearch.NewFetchSourceContext(true).Include(defaultMovieSource...)

	client := elasticConn.EsClient
	esSearch := client.Search().
		Query(query). // specify the query
		FetchSourceContext(fsc).
		Index(index). // search in index
		From(from).   // -1 처리(기존 1)
		Size(size).   // take documents 0-9
		Pretty(true). // pretty print request and response JSON
		TrackTotalHits(true).
		Timeout("5s")

	res, err := esSearch.Do(context.Background())

	if err != nil {
		log.Println(err)
		return
	}

	if !isPageAvailable(from, res, &cerr) {
		return
	}

	handleSearchResult(&result, res)

	return
}

/*
	요청 페이지에 검색 결과가 존재하는지 체크하는 함수.
*/
func isPageAvailable(from int, res *elasticSearch.SearchResult, cerr *customerror.CustomError) bool {
	if res.Hits.TotalHits.Value < int64(from) {
		*cerr = *customerror.ErrNotFound
		return false
	}
	return true
}

/*
	요청한 영화 검색 결과 담아주는 함수.
*/
func handleSearchResult(result *[]interface{}, res *elasticSearch.SearchResult) {
	for _, value := range res.Hits.Hits {
		movie := new(model.Movie)
		err := json.Unmarshal(value.Source, &movie)
		if err != nil {
			log.Println("ERROR: Movie:", err)
		}
		*result = append(*result, movie)
	}

}
