package config

var TypeChartMap map[string]string
var DeployNameSpace string
var RancherApiUrl string
var WorkloadApiPath string
var ServiceApiPath string
var ApiToken string

func init() {
	DeployNameSpace = "storage-manage-system-dev"
	RancherApiUrl = "https://acl.csie.ntut.edu.tw:3000/v3/project/c-6rfmm:p-vtj8t"
	WorkloadApiPath = "workloads"
	ServiceApiPath = "dnsRecords"
	ApiToken = "token-9s52k:z5m5cwwxft7blkszf8nw2j5th9mlxgjzlmwswk5pgl2bgv87sdvcwc"
	TypeChartMap = make(map[string]string)
	TypeChartMap["minio"] = "stable/minio"
}
