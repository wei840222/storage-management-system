package service

import (
	"encoding/json"
	"log"

	"github.com/robfig/cron"
)

type CronService struct {
	cron              *cron.Cron
	helmService       *HelmService
	mongoService      *MongoService
	rancherApiService *RancherApiService
}

func NewCronService(h *HelmService, m *MongoService, r *RancherApiService) *CronService {
	c := cron.New()
	return &CronService{c, h, m, r}
}

func (c *CronService) Strat() {
	c.cron.AddFunc("*/10 * * * * ?", func() {
		storageList := c.mongoService.ListStorage()
		for _, storage := range *storageList {
			c.rancherApiService.GetPodStatus(&storage)
			c.rancherApiService.GetServiceEndpoint(&storage)
			c.rancherApiService.GetPVCStatus(&storage)
			if err := json.Unmarshal(c.helmService.GetStorage(storage.ReleaseName), &storage.Resources); err != nil {
				panic(err)
			}
			if currStorage := c.mongoService.GetStorage(storage.ReleaseName); currStorage.Status == "deleting" {
				continue
			}
			c.mongoService.UpdateStorage(&storage)
		}
		log.Print("Update all storage status")
	})
	c.cron.Start()
}

func (c *CronService) Stop() {
	c.cron.Stop()
}
