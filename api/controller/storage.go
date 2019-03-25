package controller

import (
	"encoding/json"
	"go-helm-rest/model"
	"go-helm-rest/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

var helm *service.Helm
var mongo *service.Mongo
var rancherApi *service.RancherApi

func init() {
	helm = &service.Helm{}
	mongo = &service.Mongo{}
	rancherApi = &service.RancherApi{}
}

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
	releaseName, resource, err := helm.CreateStorage(&storage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   http.StatusText(http.StatusBadRequest),
			"message": err.Error(),
		})
		return
	}
	storage.ReleaseName = releaseName
	if err := json.Unmarshal(resource, &storage.Resources); err != nil {
		panic(err)
	}
	rancherApi.GetWorkloadStatus(&storage)
	rancherApi.GetServiceEndpoint(&storage)
	mongo.InsertStorage(&storage)
	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "Create Storage Success",
		"data":    storage,
	})
}

func ListStorage(c *gin.Context) {
	storageList := mongo.ListStorage()
	for _, storage := range *storageList {
		rancherApi.GetWorkloadStatus(&storage)
		rancherApi.GetServiceEndpoint(&storage)
		rancherApi.GetPVCStatus(&storage)
		if err := json.Unmarshal(helm.GetStorage(storage.ReleaseName), &storage.Resources); err != nil {
			panic(err)
		}
		mongo.UpdateStorage(&storage)
	}
	storageList = mongo.ListStorage()
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "List Storage Success",
		"data":    storageList,
	})
}

func GetStorage(c *gin.Context) {
	releaseName := c.Param("releaseName")
	storage := mongo.GetStorage(releaseName)
	rancherApi.GetWorkloadStatus(storage)
	rancherApi.GetServiceEndpoint(storage)
	rancherApi.GetPVCStatus(storage)
	if err := json.Unmarshal(helm.GetStorage(storage.ReleaseName), &storage.Resources); err != nil {
		panic(err)
	}
	mongo.UpdateStorage(storage)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Get Storage Success",
		"data":    storage,
	})
}

func DeleteStorage(c *gin.Context) {
	releaseName := c.Param("releaseName")
	helm.DeleteStorage(releaseName)
	mongo.DeleteStorage(releaseName)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Delete Storage Success",
	})
}
