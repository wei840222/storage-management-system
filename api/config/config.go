package config

import "os"

var TypeChartMap map[string]string
var DeployNameSpace string
var RancherApiUrl string
var LonghornApiUrl string
var WorkloadApiPath string
var ServiceApiPath string
var PVCApiPath string
var ApiToken string
var MongoUrl string

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func init() {
	DeployNameSpace = "storage-manage-system-dev"
	RancherApiUrl = "https://acl.csie.ntut.edu.tw:3000/v3/project/c-6rfmm:p-vtj8t"
	LonghornApiUrl = "https://acl.csie.ntut.edu.tw:3000/k8s/clusters/c-6rfmm/api/v1/namespaces/longhorn-system/services/http:longhorn-frontend:80/proxy/v1/volumes"
	WorkloadApiPath = "workloads"
	ServiceApiPath = "dnsRecords"
	PVCApiPath = "persistentVolumeClaims"
	ApiToken = "token-9s52k:z5m5cwwxft7blkszf8nw2j5th9mlxgjzlmwswk5pgl2bgv87sdvcwc"
	TypeChartMap = make(map[string]string)
	TypeChartMap["minio"] = "stable/minio"
	MongoUrl = getEnv("MONGO_URL", "mongodb://localhost:27017")
}
