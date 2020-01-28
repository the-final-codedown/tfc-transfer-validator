package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"math"

	//cap "github.com/the-final-codedown/tfc-cap-updater/proto/tfc/cap/updater"

	"log"
	"net/http"
	"time"
)

type CapReader struct {
	collection *mongo.Collection
}

type CapReaderInterface interface {
	GetCap(id int64) (int64, error)
}

var client *mongo.Client

func InitializeReader(mongoAddress string) *CapReader {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoOptions := options.Client().ApplyURI(mongoAddress)
	mongoOptions.SetMaxPoolSize(10)
	client, err := mongo.Connect(ctx, mongoOptions)
	if err != nil {
		log.Panic(err)
	}
	capCollection := client.Database("tfc").Collection("cap")
	return &CapReader{collection: capCollection}
}

func DisconnectReader() {
	defer client.Disconnect(context.TODO())

}

func (repository *CapReader) GetCap(id string) (int32, error) {
	accountFilter := bson.M{"accountid": id}
	result := struct {
		AccountID string
		Value     int32
		Money     int32
	}{}
	err := repository.collection.FindOne(context.Background(), &accountFilter).Decode(&result)
	if err != nil {
		log.Println("Cap for given id not found")
		return repository.CreateCap(id)
	}
	return result.Value, err
}

func (repository *CapReader) CreateCap(id string) (int32, error) {

	resp, err := http.Get("http://app:8081/accounts/" + id + "/cap")
	if err != nil {
		log.Println("Error in cap GET : ", err)
		return 0, err
	}
	log.Println(resp.Status)
	if resp.Status != "200" {
		return 0, errors.New(resp.Status)
	}
	result := struct {
		Money               int32
		AmountSlidingWindow int32
	}{}
	body, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Error in parsing GET cap request : ", err)
	}

	capStruct := struct {
		AccountID string
		Value     int32
		Money     int32
	}{id,
		int32(math.Min(float64(result.Money), float64(result.AmountSlidingWindow))),
		result.Money}

	log.Println("Response body : ", fmt.Sprint(result.Money), " ", fmt.Sprint(result.AmountSlidingWindow))
	log.Println("Cap object : ", fmt.Sprint(capStruct.AccountID), " ", fmt.Sprint(capStruct.Value), " ", fmt.Sprint(capStruct.Money))
	_, err = repository.collection.InsertOne(context.Background(), capStruct)
	if err != nil {
		log.Println("Error in inserting cap : ", err)
	}

	return capStruct.Value, err
}
