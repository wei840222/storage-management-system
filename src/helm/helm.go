package helm

import (
	"log"
	"os/exec"
)

// CreateMinio new minio storage with given accessKey, secretKey, size
func CreateMinio(accessKey string, secretKey string, size string) (string, error) {
	cmd := exec.Command("bash", "-c", "helm install --set accessKey="+accessKey+",secretKey="+secretKey+",persistence.size="+size+",service.type=NodePort,resources.limits.memory=512Mi --namespace storage-manage-system-dev stable/minio")
	stdout, err := cmd.CombinedOutput()
	log.Printf("CreateMinio\n%s", string(stdout))
	return string(stdout), err
}
