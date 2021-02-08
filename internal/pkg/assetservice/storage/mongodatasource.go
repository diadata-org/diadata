package assetstorage

import "github.com/diadata-org/diadata/pkg/dia"

type MongoDataSource struct {
}

func NewMongoDataSource() *MongoDataSource {

	return &MongoDataSource{}

}

func (ds *MongoDataSource) Save() {

}

func (ds *MongoDataSource) Get() *dia.Asset {
	return &dia.Asset{}

}
