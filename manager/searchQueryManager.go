package manager

import (
	elasticSearch "github.com/olivere/elastic/v7"
)

func MatchAllQuery(param map[string]string) (query *elasticSearch.MatchAllQuery) {
	query = elasticSearch.NewMatchAllQuery()
	return
}