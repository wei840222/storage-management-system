package model

import (
	"github.com/globalsign/mgo/bson"
)

type Storage struct {
	ID                    bson.ObjectId            `json:"id" bson:"_id,omitempty"`
	ReleaseName           string                   `json:"releaseName" bson:"releaseName"`
	Resources             []map[string]interface{} `json:"resources" bson:"resources"`
	Type                  string                   `json:"type" binding:"required" bson:"type"`
	Config                map[string]string        `json:"config" binding:"required" bson:"config"`
	Endpoint              map[string]interface{}   `json:"endpoint" bson:"endpoint"`
	Status                string                   `json:"status" bson:"status"`
	PersistentVolumeClaim map[string]interface{}   `json:"persistentVolumeClaim" bson:"persistentVolumeClaim"`
}

func (s *Storage) GetResourceName(resourceType string) string {
	var resource map[string]interface{}
	for _, r := range s.Resources {
		if r["name"] == resourceType {
			resource = r
		}
	}
	return resource["resources"].([]interface{})[0].(string)
}
