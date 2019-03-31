package model

import (
	"strings"

	"github.com/globalsign/mgo/bson"
)

type Storage struct {
	ID                    bson.ObjectId            `json:"id" bson:"_id,omitempty"`
	ReleaseName           string                   `json:"releaseName" bson:"releaseName"`
	ChartName             string                   `json:"chartName" binding:"required" bson:"chartName"`
	Resources             []map[string]interface{} `json:"resources" bson:"resources"`
	Config                map[string]string        `json:"config" binding:"required" bson:"config"`
	Endpoint              map[string]interface{}   `json:"endpoint" bson:"endpoint"`
	Status                string                   `json:"status" bson:"status"`
	PersistentVolumeClaim map[string]interface{}   `json:"persistentVolumeClaim" bson:"persistentVolumeClaim"`
	PrometheusURL         map[string]string        `json:"prometheusUrl" bson:"prometheusUrl"`
}

func (s *Storage) GetResourceName(resourceType string) string {
	var resource map[string]interface{}
	for _, r := range s.Resources {
		if strings.Contains(r["name"].(string), resourceType) {
			resource = r
		}
	}
	return resource["resources"].([]interface{})[0].(string)
}
