package model

import "github.com/globalsign/mgo/bson"

type Storage struct {
	ID          bson.ObjectId     `json:"id" bson:"_id,omitempty"`
	ReleaseName string            `json:"releaseName" bson:"releaseName"`
	Resource    interface{}       `json:"resource" bson:"resource"`
	Type        string            `json:"type" binding:"required" bson:"type"`
	Config      map[string]string `json:"config" binding:"required" bson:"config"`
	Endpoint    string            `json:"endpoint" bson:"endpoint"`
	Status      string            `json:"status" bson:"status"`
}
