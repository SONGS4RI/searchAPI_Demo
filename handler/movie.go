package handler

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	elasticSearch "github.com/olivere/elastic/v7"

	elasticConn "searchAPI/elasticConn"
	"searchAPI/manager"
	"searchAPI/model"
)

var (
	from int
	size int = 20
)

func EsGetAllMovies(param map[string]string) (result []interface{}, err error) {
	// 인덱스
	index := "movie_search"
	// 검색 시작 값
	page, _ := strconv.Atoi(param["page"])

	from = (page - 1) * size

	query := manager.MatchAllQuery(param)

	movieSource := []string{
		"companies", "companys", "directors", "genreAlt", "movieCd",
		"movieNm", "movieNmEn", "nationAlt", "openDt", "prdtStatNm",
		"prdtYear", "repGenreNm", "repNationNm", "typeNm"}
	fsc := elasticSearch.NewFetchSourceContext(true).Include(movieSource...)

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

	for _, value := range res.Hits.Hits {
		movie := new(model.Movie)
		err := json.Unmarshal(value.Source, &movie)
		if err != nil {
			log.Println("Movie:", err)
		}
		result = append(result, movie)
	}
	return
}
