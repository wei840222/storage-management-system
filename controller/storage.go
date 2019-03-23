package controller

import (
	"go-helm-rest/model"
	"go-helm-rest/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

func CreateStorage(c *gin.Context) {
	var storage model.Storage
	if err := c.ShouldBindJSON(&storage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   http.StatusText(http.StatusBadRequest),
			"message": err.Error(),
		})
		return
	}
	storage.ID = bson.NewObjectId()
	storage.ReleaseName = service.CreateStorage(&storage)
	storage.Status = service.GetWorkloadStatus(storage.ReleaseName, storage.Type)
	storage.Endpoint = service.GetServiceEndpoint(storage.ReleaseName, storage.Type)
	service.InsertStorage(&storage)
	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "Create Storage Success",
		"data":    storage,
	})

}

func DeleteStorage(c *gin.Context) {
	releaseName := c.Param("releaseName")
	service.DeleteStorage(releaseName)
	service.DeleteStorageFromMongo(releaseName)
	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "Delete Storage Success",
	})
}
