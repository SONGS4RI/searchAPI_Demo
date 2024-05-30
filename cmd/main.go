package main

import (
	"fmt"
	"log"
	"net/http"
	elasticConn "searchAPI/elasticConn"
	"searchAPI/router"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {
	elasticConn.InitEs()

	portNumber := "8080"
	timeout := 5

	apiServer := &http.Server{
		Addr:           ":" + portNumber,
		Handler:        setUpRouter(),
		ReadTimeout:    time.Duration(timeout) * time.Second,
		WriteTimeout:   time.Duration(timeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	g.Go(func() error {
		return apiServer.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal("SERVER ERROR:", err)
	}
}

func setUpRouter() http.Handler {
	e := gin.New()

	e.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		cache := "nocache"
		if param.Keys != nil {
			cache = fmt.Sprintf("%s", param.Keys["cache"])
		}

		// ex) [Thu Nov 19 14:44:34 KST 2020] 566.0013ms "GET /product/search?target=total&keyword=%EB%85%B8%ED%8A%B8%EB%B6%81&field=prodType,prod_id,shop_id,link_prod_id,prod_name,maker_name,brand_name,imageUrl,simple_desc,add_desc,cm_desc,total_price,min_price,pc_price,mobile_price,cate_c1,cate_c2,cate_c3,cate_c4,writeCnt,score,input_d,make_d,price_d,post_q,bundle,group_seq,shop,shot,priceCompareYN,savePlueQ,unit,totalCapacity,standardCapacity,priceType,initialPrice,priceTypeName,dic_code,actionTag_code,selectYN,descriptionListSeq,imageVersion,cmpnyWriteScoreSum,cmpnyWriteCnt,video_id,optionMarkPosition,optionMarkUrl&maker=2137&avSeq=183089%5E72236%7C39485%5E6297%7C6298&packageYN=Y&sort=9&start=1&limit=40&service=1&keywordType=1&previousKeyword=ssd&volume=vm HTTP/1.1" 200 "" "PostmanRuntime/7.26.5"  cache
		return fmt.Sprintf("%s - [%s] %s \"%s %s %s\" %d \"%s\" \"%s\" %s %s \n",
			param.ClientIP,
			param.TimeStamp.Format(time.UnixDate),
			param.Latency, // request에서 response 까지 걸린 시간
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Request.Referer(),
			param.Request.UserAgent(),
			param.ErrorMessage,
			cache,
		)
	}), gin.Recovery())

	rt := router.NewRouter()
	rt.MovieRouter(e.Group("movies")) // localhost:8080/{!HERE!}
	return e
}
