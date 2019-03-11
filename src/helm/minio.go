package helm

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

// CreateMinio new minio storage with given accessKey, secretKey, size
func CreateMinio(accessKey string, secretKey string, size string) (map[string]interface{}, error) {
	cmd := exec.Command("bash", "-c", "helm json install --set accessKey="+accessKey+",secretKey="+secretKey+",persistence.size="+size+",service.type=NodePort,resources.limits.memory=512Mi --namespace storage-manage-system-dev stable/minio")
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("CreateMinio failed, ", err)
		return nil, err
	}
	var m map[string]interface{}
	err = json.Unmarshal([]byte(stdout), &m)
	if err != nil {
		fmt.Println("Unmarshal failed, ", err)
		return nil, err
	}
	log.Printf("CreateMinio releaseName: %s", m["releaseName"])
	return m, err
}
