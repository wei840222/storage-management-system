package service

import (
	"fmt"
	"go-helm-rest/config"
	"go-helm-rest/model"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

type RancherApi struct{}

func (r *RancherApi) doApiRequest(method string, url string) []byte {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+config.ApiToken)
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

func (r *RancherApi) GetWorkloadStatus(storage *model.Storage) {
	body := r.doApiRequest("GET", fmt.Sprintf("%s/%s/deployment:%s:%s", config.RancherApiUrl, config.WorkloadApiPath, config.DeployNameSpace, storage.GetResourceName("v1beta2/Deployment")))
	storage.Status = gjson.Get(string(body), "state").String()
	log.Printf("GetWorkloadStatus Status:%s", storage.Status)
}

func (r *RancherApi) GetServiceEndpoint(storage *model.Storage) {
	body := r.doApiRequest("GET", fmt.Sprintf("%s/%s/%s:%s", config.RancherApiUrl, config.ServiceApiPath, config.DeployNameSpace, storage.GetResourceName("v1/Service")))
	storage.Endpoint = make(map[string]interface{})
	storage.Endpoint["host"] = gjson.Get(string(body), "publicEndpoints.0.addresses.0").String()
	storage.Endpoint["port"] = gjson.Get(string(body), "publicEndpoints.0.port").Int()
	log.Printf("GetServiceEndpoint Host:%s, Port:%d", storage.Endpoint["host"], storage.Endpoint["port"])
}

func (r *RancherApi) GetPVCStatus(storage *model.Storage) {
	body := r.doApiRequest("GET", fmt.Sprintf("%s/%s/%s:%s", config.RancherApiUrl, config.PVCApiPath, config.DeployNameSpace, storage.GetResourceName("v1/PersistentVolumeClaim")))
	storage.PersistentVolumeClaim = make(map[string]interface{})
	storage.PersistentVolumeClaim["id"] = gjson.Get(string(body), "volumeId").String()
	volumeCapacityStr := strings.Split(gjson.Get(string(body), "status.capacity.storage").String(), "Gi")[0]
	volumeCapacity, err := strconv.ParseInt(volumeCapacityStr, 10, 64)
	if err != nil {
		panic(err)
	}
	storage.PersistentVolumeClaim["capacity"] = volumeCapacity * 1024 * 1024 * 1024
	body = r.doApiRequest("POST", fmt.Sprintf("%s/%s?action=snapshotList", config.LonghornApiUrl, storage.PersistentVolumeClaim["id"]))
	storage.PersistentVolumeClaim["size"] = gjson.Get(string(body), "data.0.size").Int()
	log.Printf("GetPVCStatus ID:%s, Capacity:%d, Size:%d", storage.PersistentVolumeClaim["id"], storage.PersistentVolumeClaim["capacity"], storage.PersistentVolumeClaim["size"])
}
