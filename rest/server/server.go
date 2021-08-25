package main

import (
	"fibonacciservice/rest/server/fibonacciservice"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func initGIN() *http.Server {
	router := gin.Default()
	fibonacciservice.RegisterFibonacciService(router.Group("/"))
	var http2Server = &http2.Server{}
	srv := &http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(router, http2Server),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal("GIN ListenAndServe: ", err.Error())
		}
	}()
	return srv
}

func main() {
	_ = initGIN()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown all...")
}
