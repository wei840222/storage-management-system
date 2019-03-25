package model

import "github.com/globalsign/mgo/bson"

type Storage struct {
	ID             bson.ObjectId     `json:"id" bson:"_id,omitempty"`
	ReleaseName    string            `json:"releaseName" bson:"releaseName"`
	Resource       interface{}       `json:"resource" bson:"resource"`
	Type           string            `json:"type" binding:"required" bson:"type"`
	Config         map[string]string `json:"config" binding:"required" bson:"config"`
	Endpoint       string            `json:"endpoint" bson:"endpoint"`
	Status         string            `json:"status" bson:"status"`
	VolumeID       string            `json:"volumeId" bson:"volumeId"`
	VolumeCapacity int64            `json:"volumeCapacity" bson:"volumeCapacity"`
	VolumeSize     int64             `json:"volumeSize" bson:"volumeSize"`
}
