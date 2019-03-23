package service

import (
	"go-helm-rest/config"
	"go-helm-rest/model"

	"github.com/globalsign/mgo"
	"go.mongodb.org/mongo-driver/bson"
)

type Mongo struct{}

func (m *Mongo) InsertStorage(storage *model.Storage) {
	session, err := mgo.Dial(config.MongoUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	if err := session.DB("StorageManageSystem").C("Storage").Insert(storage); err != nil {
		panic(err)
	}
}

func (m *Mongo) ListStorage() *[]model.Storage {
	session, err := mgo.Dial(config.MongoUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	var storageList []model.Storage
	if err := session.DB("StorageManageSystem").C("Storage").Find(nil).All(&storageList); err != nil {
		panic(err)
	}
	return &storageList
}

func (m *Mongo) GetStorage(releaseName string) *model.Storage {
	session, err := mgo.Dial(config.MongoUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	var storage model.Storage
	if err := session.DB("StorageManageSystem").C("Storage").Find(bson.M{"releaseName": releaseName}).One(&storage); err != nil {
		panic(err)
	}
	return &storage
}

func (m *Mongo) UpdateStorage(storage *model.Storage) {
	session, err := mgo.Dial(config.MongoUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	if err := session.DB("StorageManageSystem").C("Storage").Update(bson.M{"releaseName": storage.ReleaseName}, storage); err != nil {
		panic(err)
	}
}

func (m *Mongo) DeleteStorage(releaseName string) {
	session, err := mgo.Dial(config.MongoUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	if err := session.DB("StorageManageSystem").C("Storage").Remove(bson.M{"releaseName": releaseName}); err != nil {
		panic(err)
	}
}
