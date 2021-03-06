package service

import (
	"storage-management-system/config"
	"storage-management-system/model"

	"errors"
	"log"
	"os/exec"
	"strings"

	"github.com/tidwall/gjson"
)

type HelmService struct {
	config *config.Config
}

func NewHelmService(c *config.Config) *HelmService {
	return &HelmService{c}
}

func (h *HelmService) CreateStorage(storage *model.Storage) (string, []byte, error) {
	cmd := "helm json install --set "
	for key := range storage.Config {
		cmd += key + "=" + storage.Config[key] + ","
	}
	cmd = strings.TrimSuffix(cmd, ",") + " --namespace " + h.config.DeployNamespace + " " + storage.ChartName
	log.Printf("CreateStorage cmd: %s", cmd)
	stdout, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	log.Printf("CreateStorage %s", stdout)
	if err != nil {
		log.Panic(err)
	}
	if strings.Contains(string(stdout), "Error") {
		return "", nil, errors.New(string(stdout))
	}
	releaseName := gjson.Get(string(stdout), "releaseName").String()
	resource := gjson.Get(string(stdout), "resources").String()
	return releaseName, []byte(resource), nil
}

func (h *HelmService) GetStorage(realseName string) []byte {
	cmd := "helm json status " + realseName
	log.Printf("GetStorage cmd: %s", cmd)
	stdout, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	log.Printf("GetStorage %s", stdout)
	if err != nil {
		log.Panic(err)
	}
	resource := gjson.Get(string(stdout), "resources").String()
	return []byte(resource)
}

func (h *HelmService) DeleteStorage(realseName string) {
	cmd := "helm delete --purge " + realseName
	log.Printf("DeleteStorage cmd: %s", cmd)
	stdout, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	log.Printf("DeleteStorage %s", stdout)
	if err != nil {
		log.Panic(err)
	}
}
