package rest

import (
	"fmt"
	"github.com/Vladimir77715/fibonacciservice/redis/fibonaccicashed"
	"github.com/Vladimir77715/fibonacciservice/rest/server/fibonacciservice"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	port = 8080 //"The def rest server port"
)

func InitGin(cs *fibonaccicashed.CashedService) *http.Server {
	if p, err := strconv.Atoi(os.Getenv("REST_PORT")); err == nil {
		log.Printf("REST_PORT is: %v", p)
		port = p
	}
	router := gin.Default()
	fibonacciservice.RegisterFibonacciService(router.Group("/"), cs)
	var http2Server = &http2.Server{}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: h2c.NewHandler(router, http2Server),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal("GIN ListenAndServe: ", err.Error())
		}
	}()
	return srv
}
