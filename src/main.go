package main

import (
	"go-helm-rest/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/storage", controller.CreateStorage)
	r.GET("/storage", controller.ListStorage)
	r.GET("/storage/:releaseName", controller.GetStorage)
	r.DELETE("/storage/:releaseName", controller.DeleteStorage)
	r.Run(":8080")
}
