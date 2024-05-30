package main

import (
	"log"
	"net/http"
	elasticconn "searchAPI/elasticConn"
	"searchAPI/router"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {
	elasticconn.InitES()

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
		log.Fatal(err)
	}
}

func setUpRouter() http.Handler {
	e := gin.New()
	rt := router.NewRouter()
	rt.TestRouter(e.Group("test")) // localhost:8080/{!HERE!}
	return e
}
