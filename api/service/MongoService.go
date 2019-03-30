package service

import (
	"go-helm-rest/config"
	"go-helm-rest/model"

	"github.com/globalsign/mgo"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoService struct {
	session    *mgo.Session
	collection *mgo.Collection
}

func NewMongoService(c *config.Config) *MongoService {
	session, err := mgo.Dial(c.MongoUrl)
	if err != nil {
		panic(err)
	}
	collection := session.DB("StorageManageSystem").C("Storage")
	return &MongoService{session, collection}
}

func (m *MongoService) CloseSession() {
	m.session.Close()
}

func (m *MongoService) InsertStorage(storage *model.Storage) {
	if err := m.collection.Insert(storage); err != nil {
		panic(err)
	}
}

func (m *MongoService) ListStorage() *[]model.Storage {
	var storageList []model.Storage
	if err := m.collection.Find(nil).All(&storageList); err != nil {
		panic(err)
	}
	return &storageList
}

func (m *MongoService) GetStorage(releaseName string) *model.Storage {
	var storage model.Storage
	if err := m.collection.Find(bson.M{"releaseName": releaseName}).One(&storage); err != nil {
		panic(err)
	}
	return &storage
}

func (m *MongoService) UpdateStorage(storage *model.Storage) {
	if err := m.collection.Update(bson.M{"releaseName": storage.ReleaseName}, storage); err != nil {
		panic(err)
	}
}

func (m *MongoService) DeleteStorage(releaseName string) {
	if err := m.collection.Remove(bson.M{"releaseName": releaseName}); err != nil {
		panic(err)
	}
}
