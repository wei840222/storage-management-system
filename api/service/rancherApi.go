package service

import (
	"go-helm-rest/config"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

type RancherApi struct{}

func (r *RancherApi) GetServiceEndpoint(releaseName string, storageType string) string {
	url := config.RancherApiUrl + "/" + config.ServiceApiPath + "/" + config.DeployNameSpace + ":" + releaseName + "-" + storageType
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+config.ApiToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	// log.Printf("GetServiceEndpoint %s", string(body))
	addresse := gjson.Get(string(body), "publicEndpoints.0.addresses.0").String()
	port := gjson.Get(string(body), "publicEndpoints.0.port").String()
	log.Printf("GetServiceEndpoint http://%s:%s", addresse, port)
	return "http://" + addresse + ":" + port
}

func (r *RancherApi) GetWorkloadStatus(releaseName string, storageType string) string {
	url := config.RancherApiUrl + "/" + config.WorkloadApiPath + "/deployment:" + config.DeployNameSpace + ":" + releaseName + "-" + storageType
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+config.ApiToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	// log.Printf("GetWorkloadStatus %s", string(body))
	status := gjson.Get(string(body), "state").String()
	log.Printf("GetWorkloadStatus %s", status)
	return status
}

func (r *RancherApi) GetPVCStatus(releaseName string, storageType string) (string, int64, int64) {
	url := config.RancherApiUrl + "/" + config.PVCApiPath + "/" + config.DeployNameSpace + ":" + releaseName + "-" + storageType
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+config.ApiToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	// log.Printf("GetPVCStatus %s", string(body))
	volumeID := gjson.Get(string(body), "volumeId").String()
	volumeCapacityStr := strings.Split(gjson.Get(string(body), "status.capacity.storage").String(), "Gi")[0]
	volumeCapacity, err := strconv.ParseInt(volumeCapacityStr, 10, 64)
	if err != nil {
		panic(err)
	}
	volumeCapacity *= 1024 * 1024 * 1024
	url = config.LonghornApiUrl + "/" + volumeID + "?action=snapshotList"
	req, err = http.NewRequest("POST", url, nil)
	req.Header.Add("Authorization", "Bearer "+config.ApiToken)
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ = ioutil.ReadAll(resp.Body)
	// log.Printf("GetVolStatus %s", string(body))
	volumeSize := gjson.Get(string(body), "data.0.size").Int()
	log.Printf("GetPVCStatus %s, %d, %d", volumeID, volumeCapacity, volumeSize)
	return volumeID, volumeCapacity, volumeSize
}
