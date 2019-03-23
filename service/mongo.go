package service

import (
	"go-helm-rest/config"
	"go-helm-rest/model"
	"log"

	"github.com/globalsign/mgo"
)

func Insert2Mongo(storage *model.Storage) {
	session, err := mgo.Dial(config.MongoUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	if err := session.DB("test").C("Storage").Insert(storage); err != nil {
		log.Fatal(err)
	}
}
