package main

import (
	"go-helm-rest/config"
	"go-helm-rest/controller"
	"go-helm-rest/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.New()
	helmService := service.NewHelmService(config)
	rancherApiService := service.NewRancherApiService(config)
	mongoService := service.NewMongoService(config)
	defer mongoService.CloseSession()
	controller := controller.NewStorageController(helmService, mongoService, rancherApiService)
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/storage", controller.CreateStorage)
	r.GET("/storage", controller.ListStorage)
	r.GET("/storage/:releaseName", controller.GetStorage)
	r.DELETE("/storage/:releaseName", controller.DeleteStorage)
	r.Run(":8080")
}
