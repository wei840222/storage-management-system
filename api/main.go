package main

import (
	"storage-management-system/config"
	"storage-management-system/controller"
	"storage-management-system/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func main() {
	c := dig.New()
	c.Provide(config.New)
	c.Provide(service.NewHelmService)
	c.Provide(service.NewRancherApiService)
	c.Provide(service.NewMongoService)
	c.Provide(service.NewCronService)
	c.Provide(controller.NewStorageController)
	if err := c.Invoke(func(mongoService *service.MongoService, cronService *service.CronService, storageController *controller.StorageController) {
		defer mongoService.CloseSession()
		cronService.Strat()
		defer cronService.Stop()
		r := gin.Default()
		r.Use(cors.Default())
		r.Use(static.Serve("/", static.LocalFile("/app/dist", false)))
		r.POST("/storage", storageController.CreateStorage)
		r.GET("/storage", storageController.ListStorage)
		r.GET("/storage/:releaseName", storageController.GetStorage)
		r.DELETE("/storage/:releaseName", storageController.DeleteStorage)
		r.Run(":8080")
	}); err != nil {
		panic(err)
	}
}
