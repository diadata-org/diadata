package database

import (
	"context"
	"github.com/diadata-org/diadata/pkg/dia"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDatabase struct {
	URI        string
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongo(URI string) (*MongoDatabase, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	md := &MongoDatabase{client: client}
	if err != nil {
		return md, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return md, err
	}

	mod := mongo.IndexModel{
		Keys: bson.M{
			"symbol": 1,
			"name":   1,
		},
		// create UniqueIndex option
		Options: options.Index().SetUnique(true),
	}
	collection := md.client.Database("diadata").Collection("asset")

	md.client.Database("diadata").Collection("asset").Indexes().CreateOne(context.Background(), mod)
	md.collection = collection
	//defer client.Disconnect(ctx)
	return md, nil
}

func (md *MongoDatabase) Save(asset dia.Asset) error {
	_, err := md.collection.InsertOne(context.TODO(), asset)
	if err != nil {
		return err
	}
	return nil
}

func (md *MongoDatabase) GetPage(pageNumber int64) (assets []dia.Asset, err error) {
 	var limit int64
	limit = 100
	skip := limit * pageNumber
 	cursor, err := md.collection.Find(context.Background(), bson.M{}, &options.FindOptions{Limit: &limit, Skip: &skip})

	if err != nil {
		return
 	}
	err = cursor.All(context.Background(),&assets)
	return
}

func (md *MongoDatabase) Count() (int64, error) {
	return md.collection.EstimatedDocumentCount(context.Background())
}

func (md *MongoDatabase) GetByName(name string) (dia.Asset, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var result dia.Asset

	err := md.collection.FindOne(ctx, bson.D{{"symbol", name}}).Decode(&result)
	if err != nil {
		return dia.Asset{}, err
	}

	return result, err

}
