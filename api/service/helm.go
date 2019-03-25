package service

import (
	"errors"
	"go-helm-rest/config"
	"go-helm-rest/model"
	"log"
	"os/exec"
	"strings"

	"github.com/tidwall/gjson"
)

type Helm struct{}

func (h *Helm) CreateStorage(storage *model.Storage) (string, []byte, error) {
	cmd := "helm json install --set "
	for key := range storage.Config {
		cmd += key + "=" + storage.Config[key] + ","
	}
	cmd = strings.TrimSuffix(cmd, ",") + " --namespace " + config.DeployNameSpace + " " + config.TypeChartMap[storage.Type]
	log.Printf("CreateStorage cmd: %s", cmd)
	stdout, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		panic(err)
	}
	if strings.Contains(string(stdout), "Error") {
		return "", nil, errors.New(string(stdout))
	}
	releaseName := gjson.Get(string(stdout), "releaseName").String()
	resource := gjson.Get(string(stdout), "resources").String()
	log.Printf("CreateStorage releaseInfo: %s", stdout)
	return releaseName, []byte(resource), nil
}

func (h *Helm) GetStorage(realseName string) []byte {
	cmd := "helm json status " + realseName
	log.Printf("GetStorage cmd: %s", cmd)
	stdout, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		panic(err)
	}
	log.Printf("GetStorage %s", stdout)
	resource := gjson.Get(string(stdout), "resources").String()
	return []byte(resource)
}

func (h *Helm) DeleteStorage(realseName string) {
	cmd := "helm delete --purge " + realseName
	log.Printf("DeleteStorage cmd: %s", cmd)
	stdout, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		panic(err)
	}
	log.Printf("DeleteStorage %s", stdout)
}
