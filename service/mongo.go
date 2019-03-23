package service

import (
	"go-helm-rest/config"
	"go-helm-rest/model"
	"log"

	"github.com/globalsign/mgo"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertStorage(storage *model.Storage) {
	session, err := mgo.Dial(config.MongoUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	if err := session.DB("StorageManageSystem").C("Storage").Insert(storage); err != nil {
		panic(err)
	}
}

func GetStorage(releaseName string) {
	session, err := mgo.Dial(config.MongoUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	storage := &model.Storage{}
	if err := session.DB("StorageManageSystem").C("Storage").Find(bson.M{"releaseName": releaseName}).One(&storage); err != nil {
		panic(err)
	}
	log.Println(storage)
}

func DeleteStorageFromMongo(releaseName string) {
	session, err := mgo.Dial(config.MongoUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	storage := &model.Storage{}
	if err := session.DB("StorageManageSystem").C("Storage").Remove(bson.M{"releaseName": releaseName}); err != nil {
		panic(err)
	}
	log.Println(storage)
}
