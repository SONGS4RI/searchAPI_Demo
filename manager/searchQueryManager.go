package manager

import (
	elasticSearch "github.com/olivere/elastic/v7"
)

/*
	"query": {
	    "match_all": {}
	  }
*/
func MatchAllQuery() (query *elasticSearch.MatchAllQuery) {
	query = elasticSearch.NewMatchAllQuery()
	return
}

/*
	"query": {
	    "bool": {
	      "should": [
	        {
	          "match": {
	            "movieNm": "바보"
	          }
	        },
	        {
	          "match": {
	            "movieNmEn": "바보"
	          }
	        }
	      ]
	    }
	  }
*/
func BoolQuery(search string) (query *elasticSearch.BoolQuery) {
	query = elasticSearch.NewBoolQuery().Should(
		elasticSearch.NewMatchQuery("movieNm", search),
		elasticSearch.NewMatchQuery("movieNmEn", search),
	)
	return
}
