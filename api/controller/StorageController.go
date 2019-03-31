package controller

import (
	"encoding/json"
	"net/http"
	"storage-management-system/model"
	"storage-management-system/service"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

type StorageController struct {
	helmService       *service.HelmService
	mongoService      *service.MongoService
	rancherApiService *service.RancherApiService
}

func NewStorageController(h *service.HelmService, m *service.MongoService, r *service.RancherApiService) *StorageController {
	return &StorageController{h, m, r}
}

func (s *StorageController) returnJSON(c *gin.Context, code int, parm interface{}) {
	parmJSONStr := "data"
	if code >= 400 {
		parmJSONStr = "error"
	}
	c.JSON(code, gin.H{
		"code":      code,
		"message":   http.StatusText(code),
		parmJSONStr: parm,
	})
}

func (s *StorageController) CreateStorage(c *gin.Context) {
	var storage model.Storage
	if err := c.ShouldBindJSON(&storage); err != nil {
		s.returnJSON(c, http.StatusBadRequest, err.Error())
		return
	}
	storage.ID = bson.NewObjectId()
	releaseName, resource, err := s.helmService.CreateStorage(&storage)
	if err != nil {
		s.returnJSON(c, http.StatusBadRequest, err.Error())
		return
	}
	storage.ReleaseName = releaseName
	if err := json.Unmarshal(resource, &storage.Resources); err != nil {
		panic(err)
	}
	s.rancherApiService.GetPodStatus(&storage)
	s.rancherApiService.GetServiceEndpoint(&storage)
	s.mongoService.InsertStorage(&storage)
	s.returnJSON(c, http.StatusCreated, storage)
}

func (s *StorageController) ListStorage(c *gin.Context) {
	storageList := s.mongoService.ListStorage()
	s.returnJSON(c, http.StatusOK, storageList)
}

func (s *StorageController) GetStorage(c *gin.Context) {
	releaseName := c.Param("releaseName")
	storage := s.mongoService.GetStorage(releaseName)
	s.returnJSON(c, http.StatusOK, storage)
}

func (s *StorageController) DeleteStorage(c *gin.Context) {
	releaseName := c.Param("releaseName")
	storage := s.mongoService.GetStorage(releaseName)
	storage.Status = "deleting"
	s.mongoService.UpdateStorage(storage)
	s.helmService.DeleteStorage(releaseName)
	s.mongoService.DeleteStorage(releaseName)
	s.returnJSON(c, http.StatusOK, nil)
}
