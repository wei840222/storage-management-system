package main

import (
	"log"

	"go-helm-rest/handler"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {
	router := fasthttprouter.New()
	router.POST("/storage", handler.CreateStorage)

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
