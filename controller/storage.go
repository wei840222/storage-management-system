package controller

import (
	"go-helm-rest/model"
	"go-helm-rest/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStorage(c *gin.Context) {
	var storage model.Storage
	if err := c.ShouldBindJSON(&storage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"error":   http.StatusText(http.StatusBadRequest),
			"message": err.Error(),
		})
		return
	}
	storage.ID = service.CreateStorage(&storage)
	storage.Status = service.GetWorkloadStatus(storage.ID, storage.Type)
	storage.Endpoint = service.GetServiceEndpoint(storage.ID, storage.Type)
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Create Storage Success",
		"data":    storage,
	})

}
