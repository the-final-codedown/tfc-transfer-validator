package cap_reader

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type CapReader struct {
	collection *mongo.Collection
}
type CapReaderInterface interface {
	GetCap(id int64) (int64, error)
}

var client *mongo.Client

func InitializeReader(mongo_adresse string) *CapReader {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongo_adresse))
	if err != nil {
		log.Panic(err)
	}
	capCollection := client.Database("tfc").Collection("cap")
	return &CapReader{collection: capCollection}
}
func DisconnectReader() {
	defer client.Disconnect(context.TODO())

}
func (repository *CapReader) GetCap(id int64) (int32, error) {
	account_filter := bson.D{{"accountID", id}}
	result := struct {
		AccountID int64
		Value     int32
	}{}
	err := repository.collection.FindOne(context.Background(), account_filter).Decode(&result)
	return result.Value, err

}
