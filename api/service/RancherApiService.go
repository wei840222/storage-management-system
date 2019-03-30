package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"storage-management-system/config"
	"storage-management-system/model"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

type RancherApiService struct {
	config *config.Config
}

func NewRancherApiService(c *config.Config) *RancherApiService {
	return &RancherApiService{c}
}

func (r *RancherApiService) doApiRequest(method string, url string) []byte {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", r.config.RancherApiToken))
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return body
}

func (r *RancherApiService) GetPodStatus(storage *model.Storage) {
	body := r.doApiRequest("GET", fmt.Sprintf("%s/%s/%s:%s", r.config.RancherApiUrl, "pods", r.config.DeployNamespace, storage.GetResourceName("Pod")))
	storage.Status = gjson.Get(string(body), "state").String()
	log.Printf("GetPodStatus Status:%s", storage.Status)
}

func (r *RancherApiService) GetServiceEndpoint(storage *model.Storage) {
	body := r.doApiRequest("GET", fmt.Sprintf("%s/%s/%s:%s", r.config.RancherApiUrl, "services", r.config.DeployNamespace, storage.GetResourceName("Service")))
	storage.Endpoint = make(map[string]interface{})
	storage.Endpoint["host"] = gjson.Get(string(body), "publicEndpoints.0.addresses.0").String()
	storage.Endpoint["port"] = gjson.Get(string(body), "publicEndpoints.0.port").Int()
	log.Printf("GetServiceEndpoint Host:%s, Port:%d", storage.Endpoint["host"], storage.Endpoint["port"])
}

func (r *RancherApiService) GetPVCStatus(storage *model.Storage) {
	body := r.doApiRequest("GET", fmt.Sprintf("%s/%s/%s:%s", r.config.RancherApiUrl, "persistentVolumeClaims", r.config.DeployNamespace, storage.GetResourceName("PersistentVolumeClaim")))
	storage.PersistentVolumeClaim = make(map[string]interface{})
	storage.PersistentVolumeClaim["id"] = gjson.Get(string(body), "volumeId").String()
	volumeCapacityStr := strings.Split(gjson.Get(string(body), "status.capacity.storage").String(), "Gi")[0]
	volumeCapacity, err := strconv.ParseInt(volumeCapacityStr, 10, 64)
	if err != nil {
		log.Print(err)
		volumeCapacity = 0
	}
	storage.PersistentVolumeClaim["capacity"] = volumeCapacity * 1024 * 1024 * 1024
	body = r.doApiRequest("POST", fmt.Sprintf("%s/%s?action=snapshotList", r.config.LonghornApiUrl, storage.PersistentVolumeClaim["id"]))
	storage.PersistentVolumeClaim["size"] = gjson.Get(string(body), "data.0.size").Int()
	log.Printf("GetPVCStatus ID:%s, Capacity:%d, Size:%d", storage.PersistentVolumeClaim["id"], storage.PersistentVolumeClaim["capacity"], storage.PersistentVolumeClaim["size"])
}
